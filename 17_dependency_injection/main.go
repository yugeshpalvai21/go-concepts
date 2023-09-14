package main

import (
	"context"
	"log"
	"os"

	"go.uber.org/fx"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{Logger: log.New(os.Stdout, "[MyService] ", log.LstdFlags)}
}

type MyService struct {
	Logger *Logger
}

func NewMyService(logger *Logger) *MyService {
	return &MyService{Logger: logger}
}

func main() {
	app := fx.New(
		fx.Provide(NewLogger),
		fx.Provide(NewMyService),
		fx.Invoke(StartApplication),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Error starting the application: %v", err)
	}
	defer app.Stop(context.Background())
}

func StartApplication(service *MyService) {
	service.Logger.Print("Application started...")
}
