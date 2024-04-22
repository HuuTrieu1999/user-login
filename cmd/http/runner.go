package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	http2 "login/internal/controller/http"
	"login/internal/controller/http/server/http"
	"login/internal/core/config"
	"login/internal/core/service/user"
	"login/internal/infra/repository"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		os.Exit(1)
	}
	instance := gin.New()
	instance.Use(gin.Logger())
	instance.Use(gin.Recovery())
	clientOptions := options.Client().ApplyURI(viper.GetString("mongodb.uri"))
	client, _ := mongo.Connect(context.Background(), clientOptions)
	collection := client.Database(viper.GetString("mongodb.db-name")).Collection(viper.GetString("mongodb.collection"))

	// Initialize repository
	userRepo := repository.NewUserRepo(collection)

	// Create the UserService
	userService := user.NewUserService(userRepo, viper.GetString("secret-key"))

	// Create the UserController
	userController := http2.NewUserController(instance, userService)

	// Initialize the routes for UserController
	userController.InitRouter()

	// Create the HTTP server
	httpServer := http.NewHttpServer(
		instance,
		config.HttpServerConfig{
			Port: viper.GetUint("server.http.port"),
		},
	)

	// Start the HTTP server
	httpServer.Start()
	defer func(httpServer http.HttpServer) {
		err := httpServer.Close()
		if err != nil {
			log.Printf("failed to close http server %v", err)
		}
	}(httpServer)

	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}
