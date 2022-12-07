package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/suttapak/siot-backend/config"
	"github.com/suttapak/siot-backend/db"
	"github.com/suttapak/siot-backend/external"
	"github.com/suttapak/siot-backend/handler"
	"github.com/suttapak/siot-backend/middleware"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/service"
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
	controlDataRepo := repository.NewControlDataRepository(conn)
	displayRepo := repository.NewDisplayRepository(conn)
	displayDataRepo := repository.NewDisplayDataRepositoryDb(conn)
	layoutRepo := repository.NewLayoutRepository(conn)
	userRepo := repository.NewUserRepositoryDB(conn)
	settingRepo := repository.NewSettingRepository(conn)
	widgetControlRepo := repository.NewWidgetControlRepository(conn)
	widgetDisplayRepo := repository.NewWidgetDisplayRepository(conn)

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

	r.Use(GinMiddleware("http://localhost:3000"))

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
	boxMemberGroup.GET("", boxMemberHandler.BoxMembers)
	boxMemberGroup.POST("", boxMemberHandler.AddMember)

	// control
	controlGroup := r.Group("boxes/:boxId/controls", jwtWare.JWTWare)
	controlGroup.POST("", controlHandler.Create)
	controlGroup.GET("", controlHandler.FindControls)
	// display

	displayGroup := r.Group("boxes/:boxId/displays", jwtWare.JWTWare)
	displayGroup.POST("", displayHandler.Create)
	displayGroup.GET("", displayHandler.FindDisplays)

	mqttGroup := r.Group("mqtt")
	mqttGroup.POST("/auth", mqttHandler.Auth)
	mqttGroup.POST("/acl", mqttHandler.ACLCheck)
	mqttGroup.POST("/admin", mqttHandler.Admin)

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

	mqtt := external.NewMqttClient(conf)
	server := socketio.NewServer(nil)
	wsServ := service.NewWsService(mqtt, boxRepo, controlRepo, displayRepo)
	wsHandler := handler.NewWsHandler(wsServ, server)
	mqttMachine := external.NewMQTTMachine(mqtt, server, canSubRepo, controlRepo, controlDataRepo, displayRepo, displayDataRepo)
	go mqttMachine.MQTTMachine()

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		return nil
	})

	server.OnEvent("", "subscript", wsHandler.Subscript)
	server.OnEvent("", "publish", wsHandler.Publish)
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	// run server
	err := r.Run(fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		panic(err)
	}

}

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}
