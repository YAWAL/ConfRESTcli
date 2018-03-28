package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"strconv"
	"time"

	"github.com/YAWAL/GetMeConfAPI/api"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	mongoConf = "mongodb"
	tsConf    = "tsconfig"
	tempConf  = "tempconfig"

	defaultClientPort     = "8080"
	defaultServiceHost    = "localhost"
	defaultServicePort    = "3000"
	defaultServicePortStr = "1000"
)

type configClient struct {
	grpcClient      api.ConfigServiceClient
	contextDeadline time.Duration
}

func main() {

	port := os.Getenv("CLIENT_PORT")
	if port == "" {
		port = defaultClientPort
	}
	serviceHost := os.Getenv("SERVICE_HOST")
	if port == "" {
		serviceHost = defaultServiceHost
	}
	servicePort := os.Getenv("SERVICE_PORT")
	if servicePort == "" {
		servicePort = defaultServicePort
	}
	contDeadlineStr := os.Getenv("CONTEXT_DEADLINE_MILLISECONDS")
	if contDeadlineStr == "" {
		contDeadlineStr = defaultServicePortStr
	}
	contDeadline, err := strconv.Atoi(contDeadlineStr)
	if err != nil {
		log.Fatalf("can not convert context deadline value to int: %v", err)
	}

	address := fmt.Sprintf("%s:%s", serviceHost, servicePort)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dialContext error has occurred: %v", err)
	}
	conn.GetState()
	log.Printf("State: %v", conn.GetState())
	defer conn.Close()

	grpcClient := api.NewConfigServiceClient(conn)
	cc := configClient{grpcClient: grpcClient, contextDeadline: time.Duration(contDeadline)}

	log.Printf("Processing client...")

	//http server
	router := gin.Default()
	router.GET("/getConfig/:type/:name", getByNameHandler(&cc))

	router.GET("/getConfig/:type", getByTypeHandler(&cc))

	router.POST("/createConfig/:type", createConfigHandler(&cc))

	router.DELETE("/deleteConfig/:type/:name", deleteConfigHandler(&cc))

	router.PUT("/updateConfig/:type", updateConfigHandler(&cc))

	router.GET("/info", statusInfoHandler())

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	defer srv.Shutdown(context.Background())
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("filed to run server: %v", err)
	}
}
