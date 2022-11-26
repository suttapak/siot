package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/config"
	"github.com/suttapak/siot-backend/db"
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
	userServ := service.NewUserService(userRepo)
	widgetCtServ := service.NewWidgetControlService(widgetControlRepo)
	widgetDpServ := service.NewWidgetDisplayService(widgetDisplayRepo)

	// handler
	authHandler := handler.NewAuthHandler(authServ)
	boxHandler := handler.NewBoxHandler(boxServ)
	boxMemberHandler := handler.NewBoxMemberHandler(boxMemberServ)
	controlHandler := handler.NewControlHandler(controlServ)
	displayHandler := handler.NewDisplayHandler(displayServ)
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
	boxMemberGroup := r.Group("boxes/:boxId/members")
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

	// run server
	err := r.Run(fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		panic(err)
	}

}

// TODO create control service : create find all
// TODO create display service : create find all

// TODO create widget control adn display
