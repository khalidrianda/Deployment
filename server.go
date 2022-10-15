package main

import (
	"github.com/ALTA-BE12-KhalidRianda/Deployment/config"
	"github.com/ALTA-BE12-KhalidRianda/Deployment/features/user/delivery"
	"github.com/ALTA-BE12-KhalidRianda/Deployment/features/user/repository"
	"github.com/ALTA-BE12-KhalidRianda/Deployment/features/user/services"
	"github.com/ALTA-BE12-KhalidRianda/Deployment/utils/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)
	uRepo := repository.New(db)
	uService := services.New(uRepo)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	delivery.New(e, uService)

	e.Logger.Fatal(e.Start(":8000"))
}
