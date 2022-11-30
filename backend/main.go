package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/suttapak/siot-backend/config"
	"github.com/suttapak/siot-backend/db"
	"github.com/suttapak/siot-backend/handler"
	"github.com/suttapak/siot-backend/middleware"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/logs"
)

func main() {
	conf := config.Default()

	conn := db.GetPostgresInstance(conf, true)
	// repository
	boxMemRepo := repository.NewBoxMemberRepository(conn)
	boxRepo := repository.NewBoxRepository(conn)
	boxSecretRepo := repository.NewBoxSecretRepository(conn)
	canSubRepo := repository.NewCanSubRepository(conn)
	canPubRepo := repository.NewCanPubRepository(conn)
	controlRepo := repository.NewControlRepository(conn)
	displayRepo := repository.NewDisplayRepository(conn)
	layoutRepo := repository.NewLayoutRepository(conn)
	userRepo := repository.NewUserRepositoryDB(conn)
	settingRepo := repository.NewSettingRepository(conn)
	roleRepo := repository.NewRoleRepository(conn)
	widgetControlRepo := repository.NewWidgetControlRepository(conn)
	widgetDisplayRepo := repository.NewWidgetDisplayRepository(conn)

	_ = roleRepo

	// service || use-case

	authServ := service.NewAuthService(userRepo, conf, settingRepo)
	boxServ := service.NewBoxService(conf, boxRepo, boxMemRepo, boxSecretRepo, canSubRepo, canPubRepo)
	boxMemberServ := service.NewBoxMemberService(userRepo, boxMemRepo)
	controlServ := service.NewControlService(boxRepo, controlRepo, layoutRepo, widgetControlRepo)
	displayServ := service.NewDisplayService(boxRepo, displayRepo, layoutRepo, widgetDisplayRepo)
	mqttServ := service.NewMqttAuthService(boxRepo, canSubRepo, canPubRepo, userRepo)
	userServ := service.NewUserService(userRepo)
	widgetCtServ := service.NewWidgetControlService(widgetControlRepo)
	widgetDpServ := service.NewWidgetDisplayService(widgetDisplayRepo)

	// handler
	authHandler := handler.NewAuthHandler(authServ)
	boxHandler := handler.NewBoxHandler(boxServ)
	boxMemberHandler := handler.NewBoxMemberHandler(boxMemberServ)
	controlHandler := handler.NewControlHandler(controlServ)
	displayHandler := handler.NewDisplayHandler(displayServ)
	mqttHandler := handler.NewMqttHandler(mqttServ)
	userHandler := handler.NewUserHandler(userServ)
	widgetCtHandler := handler.NewWidgetControlHandler(widgetCtServ)
	widgetDpHandler := handler.NewWidgetDisplayHandler(widgetDpServ)

	jwtWare := middleware.NewJWTWare(conf)

	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1"})
	// core

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", "Access-Control-Allow-Origin")

	r.Use(cors.New(config))

	// auth
	authGroup := r.Group("auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)
	// box group
	boxGroup := r.Group("boxes")
	boxGroup.Use(jwtWare.JWTWare)
	boxGroup.POST("", boxHandler.Create)
	boxGroup.GET("", boxHandler.FindBoxes)
	boxGroup.GET("/:boxId", boxHandler.FindBox)

	// box member
	boxMemberGroup := r.Group("boxes/:boxId/members", jwtWare.JWTWare)
	{
		boxMemberGroup.GET("", boxMemberHandler.BoxMembers)
		boxMemberGroup.POST("", boxMemberHandler.AddMember)
	}

	// control
	controlGroup := r.Group("boxes/:boxId/controls", jwtWare.JWTWare)
	controlGroup.POST("", controlHandler.Create)
	controlGroup.GET("", controlHandler.FindControls)
	// display

	displayGroup := r.Group("boxes/:boxId/displays", jwtWare.JWTWare)
	displayGroup.POST("", displayHandler.Create)
	displayGroup.GET("", displayHandler.FindDisplays)

	mqttGroup := r.Group("mqtt")
	{
		mqttGroup.POST("/auth", mqttHandler.Auth)
		mqttGroup.POST("/acl", mqttHandler.ACLCheck)
		mqttGroup.POST("/admin", mqttHandler.Admin)
	}

	// user group
	userGroup := r.Group("user")
	userGroup.Use(jwtWare.JWTWare)
	userGroup.GET("", userHandler.FindUser)
	userGroup.GET("/:userId", userHandler.FindUserById)

	// widget display group
	widgetDpGroup := r.Group("widgets/displays")
	widgetDpGroup.GET("", widgetDpHandler.Widgets)
	widgetDpGroup.GET("/:widgetId", widgetDpHandler.Widget)
	widgetDpGroup.POST("", widgetDpHandler.Create)

	widgetCtGroup := r.Group("widgets/controls")
	widgetCtGroup.GET("", widgetCtHandler.Widgets)
	widgetCtGroup.GET("/:widgetId", widgetCtHandler.Widget)
	widgetCtGroup.POST("", widgetCtHandler.Create)

	client := mqttFunc()

	socketServer := testSocketIo(r, *client)
	// toodosfjas

	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer socketServer.Close()

	r.GET("/socket.io/*any", gin.WrapH(socketServer))
	r.POST("/socket.io/*any", gin.WrapH(socketServer))
	err := r.Run(fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		panic(err)
	}

	// run server

}

type Message struct {
	Topic string  `json:"topic"`
	Data  float64 `json:"data"`
	BoxId string  `json:"boxId"`
}

type Mq struct {
	Msg string `json:"msg"`
}

func testSocketIo(router *gin.Engine, client mqtt.Client) *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "subscribe", func(s socketio.Conn, msg interface{}) {

		m, err := json.Marshal(msg)
		if err != nil {
			s.Emit("suttapak", err)
			return
		}
		message := Message{}
		if err := json.Unmarshal(m, &message); err != nil {
			s.Emit("suttapak", err)
			return
		}

		client.Subscribe("betamanga-0384d656/test", 1, func(c mqtt.Client, m mqtt.Message) {

			mq := Mq{}
			json.Unmarshal(m.Payload(), &mq)

			s.Emit(message.BoxId, mq)
		})

		s.Emit(message.BoxId, message)
	})

	server.OnEvent("/", "publish", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	return server

}

func mqttFunc() *mqtt.Client {

	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetUsername("fd3782ff-a373-4146-8534-b85f78a78aa7").SetPassword("spider09")
	opts.SetClientID(getRandStr())
	opts.SetCleanSession(true)
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		panic(err)
	})
	client := mqtt.NewClient(opts)

	if err := client.Connect().Error(); err != nil {
		logs.Error(err)
		panic(err)
	}

	// go func() {
	// 	i := 0
	// 	for {
	// 		type User struct {
	// 			Email    string `json:"email"`
	// 			Password string `json:"password"`
	// 		}

	// 		u := User{"suttapak.matee@gmail.com", "password"}
	// 		b, _ := json.Marshal(u)

	// 		t := client.Publish("betamanga-0384d656/test", 1, true, b)
	// 		i++

	// 		<-t.Done()
	// 		logs.Debug("done")

	// 		if t.Error() != nil {
	// 			logs.Error(t.Error())
	// 		}
	// 		time.Sleep(3 * time.Second)

	// 	}

	// }()
	return &client

}

func getRandStr() string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	length := 6
	randStr := make([]byte, length)
	rand.Read(randStr)
	for i, b := range randStr {
		randStr[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(randStr)
}
