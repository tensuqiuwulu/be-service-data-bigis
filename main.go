package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tensuqiuwulu/be-service-data-bigis/config"
	"github.com/tensuqiuwulu/be-service-data-bigis/controller"
	"github.com/tensuqiuwulu/be-service-data-bigis/exceptions"
	"github.com/tensuqiuwulu/be-service-data-bigis/repository"
	"github.com/tensuqiuwulu/be-service-data-bigis/routes"
	"github.com/tensuqiuwulu/be-service-data-bigis/service"
	"github.com/tensuqiuwulu/be-service-data-bigis/utilities"
)

func main() {
	appConfig := config.GetConfig()

	// Database
	DBConn := repository.NewDatabaseConnection(&appConfig.Database)

	// Timezone
	location, err := time.LoadLocation(appConfig.Timezone.Timezone)
	time.Local = location
	fmt.Println("Location:", location, err)

	// Server App
	fmt.Println("Server App : ", string(appConfig.Application.Server))

	// Logger
	logrusLogger := utilities.NewLogger(appConfig.Log)

	// Validator
	validate := validator.New()

	e := echo.New()

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      nil,
		ErrorMessage: "Request Timeout",
		Timeout:      15 * time.Second,
	}))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = exceptions.ErrorHandler
	e.Use(middleware.RequestID())

	// Repository
	bigisRepository := repository.NewBigisRepository(&appConfig.Database)

	// Service
	bigisService := service.NewBigisService(
		DBConn,
		validate,
		logrusLogger,
		bigisRepository,
	)

	bigisController := controller.NewBigisController(
		logrusLogger,
		bigisService,
	)

	// Route
	routes.BigisRoute(e, bigisController)

	// Careful shutdown
	go func() {
		if err := e.Start(":" + strconv.Itoa(int(appConfig.Webserver.Port))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// mysql database
	repository.Close(DBConn)
	fmt.Println("Echo was successful shutdown.")

}
