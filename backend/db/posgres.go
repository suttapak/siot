package db

import (
	"fmt"

	"github.com/suttapak/siot-backend/config"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetPostgresInstance(cfg *config.Configs, migrate bool) *gorm.DB {
	//dsn = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		cfg.PG.Username,
		cfg.PG.Password,
		cfg.PG.Host,
		cfg.PG.Port,
		cfg.PG.DB,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// if migrate {
	if db.AutoMigrate(
		&model.User{},
		&model.Avatar{},
		&model.UserSetting{},
		&model.Role{},
		&model.CanPublish{},
		&model.UserNotification{},
		&model.CanSubscribe{},
		&model.Box{},
		&model.BoxMember{},
		&model.BoxSecret{},
		&model.BoxMemberPermission{},
		&model.Layout{},
		&model.WidgetControl{},
		&model.Control{},
		&model.ControlData{},
		&model.WidgetDisplay{},
		&model.Display{},
		&model.DisplayData{},
		&model.UserOnline{},
	) != nil {
		panic("Error when run migrations")
	}
	// if err != nil {
	// 	panic("Error when run migrations")
	// }
	// }
	return db
}
