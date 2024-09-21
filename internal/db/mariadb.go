package db

import (
    "fmt"
    "github.com/alirezadp10/chat/internal/configs"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "sync"
)

var db *gorm.DB
var dbOnce sync.Once

func Connection() *gorm.DB {
    dbOnce.Do(func() {
        newLogger := logger.New(
            log.New(log.Writer(), "\r\n", log.LstdFlags),
            logger.Config{
                LogLevel: logger.Silent,
            },
        )

        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            configs.Mariadb()["user"],
            configs.Mariadb()["password"],
            configs.Mariadb()["host"],
            configs.Mariadb()["port"],
            configs.Mariadb()["database"],
        )

        var err error
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
            Logger: newLogger,
        })
        if err != nil {
            log.Fatal("Failed to connect to the database:", err)
        }
        log.Println("Database connection established")
    })

    return db
}
