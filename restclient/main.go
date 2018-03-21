package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/YAWAL/ConfRESTcli/api"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	mongoConf = "mongodb"
	tsConf    = "tsconfig"
	tempConf  = "tempconfig"

	defaultClientPort  = "8080"
	defaultServiceHost = "localhost"
	defaultServicePort = "3000"
)

type configClient struct {
	grpcClient api.ConfigServiceClient
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

	address := fmt.Sprintf("%s:%s", serviceHost, servicePort)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dialContext error has occurred: %v", err)
	}
	conn.GetState()
	log.Printf("State: %v", conn.GetState())
	defer conn.Close()

	grpcClient := api.NewConfigServiceClient("configservice", microclient.DefaultClient)
	cc := configClient{grpcClient: grpcClient}

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
