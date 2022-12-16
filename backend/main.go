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
	avatarRepo := repository.NewAvatarRepository(conn)
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

	authServ := service.NewAuthService(avatarRepo, userRepo, conf, settingRepo)
	avatarServ := service.NewAvatarService(avatarRepo)
	boxServ := service.NewBoxService(conf, boxRepo, boxMemRepo, boxSecretRepo, canSubRepo, canPubRepo)
	boxMemberServ := service.NewBoxMemberService(userRepo, boxMemRepo)
	controlServ := service.NewControlService(boxRepo, controlRepo, layoutRepo, widgetControlRepo)
	displayDataServ := service.NewDisplayDataService(displayRepo, displayDataRepo)
	displayServ := service.NewDisplayService(boxRepo, displayRepo, layoutRepo, widgetDisplayRepo)
	mqttServ := service.NewMqttAuthService(boxRepo, canSubRepo, canPubRepo, userRepo)
	userServ := service.NewUserService(userRepo)
	widgetCtServ := service.NewWidgetControlService(widgetControlRepo)
	widgetDpServ := service.NewWidgetDisplayService(widgetDisplayRepo)

	// handler
	authHandler := handler.NewAuthHandler(authServ)
	avatarHandler := handler.NewAvatarHandler(avatarServ)
	boxHandler := handler.NewBoxHandler(boxServ)
	boxMemberHandler := handler.NewBoxMemberHandler(boxMemberServ)
	controlHandler := handler.NewControlHandler(controlServ)
	displayDataHandler := handler.NewDisplayDataHandler(displayDataServ)
	displayHandler := handler.NewDisplayHandler(displayServ)
	mqttHandler := handler.NewMqttHandler(mqttServ)
	userHandler := handler.NewUserHandler(userServ)
	widgetCtHandler := handler.NewWidgetControlHandler(widgetCtServ)
	widgetDpHandler := handler.NewWidgetDisplayHandler(widgetDpServ)

	// middle ware
	jwtWare := middleware.NewJWTWare(conf)
	graudRole := middleware.NewGraudRole(boxMemRepo)

	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1"})
	// core

	r.Use(GinMiddleware("http://localhost:3000"))

	// auth
	authGroup := r.Group("auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)
	// avatar
	avatarGroup := r.Group("avatar", jwtWare.JWTWare)
	avatarGroup.PUT("", avatarHandler.Update)

	// box group
	boxGroup := r.Group("boxes")
	boxGroup.Use(jwtWare.JWTWare)
	boxGroup.POST("", boxHandler.Create)
	boxGroup.GET("", boxHandler.FindBoxes)
	boxGroup.GET("/members", boxHandler.Member)
	boxGroup.GET("/:boxId", boxHandler.FindBox)
	boxGroup.PUT("/:boxId", graudRole.CanWrite, boxHandler.Update)
	boxGroup.DELETE("/:boxId", graudRole.CanWrite, boxHandler.Delete)

	// box member
	boxMemberGroup := r.Group("boxes/:boxId/members", jwtWare.JWTWare)
	boxMemberGroup.GET("", boxMemberHandler.BoxMembers)
	boxMemberGroup.POST("", graudRole.CanWrite, boxMemberHandler.AddMember)

	// control
	controlGroup := r.Group("boxes/:boxId/controls", jwtWare.JWTWare)
	controlGroup.POST("", graudRole.CanWrite, controlHandler.Create)
	controlGroup.GET("", controlHandler.FindControls)
	controlGroup.PUT("/:controlId", graudRole.CanWrite, controlHandler.Update)
	controlGroup.DELETE("/:controlId", graudRole.CanWrite, controlHandler.Delete)
	// display data
	displayDataGroup := r.Group("boxes/:boxId/displays/:displayId/data", jwtWare.JWTWare)
	displayDataGroup.GET("", displayDataHandler.Displays)

	// display
	displayGroup := r.Group("boxes/:boxId/displays", jwtWare.JWTWare)
	displayGroup.POST("", graudRole.CanWrite, displayHandler.Create)
	displayGroup.GET("", displayHandler.FindDisplays)
	displayGroup.PUT("/:displayId", graudRole.CanWrite, displayHandler.Update)
	displayGroup.DELETE("/:displayId", graudRole.CanWrite, displayHandler.Delete)

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

	// -----
	r.Static("/asset", "./public/asset")

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
