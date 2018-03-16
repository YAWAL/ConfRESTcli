package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"strconv"

	"github.com/YAWAL/ConfRESTcli/entities"
	"github.com/YAWAL/GetMeConfAPI/api"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func getByNameHandler(cc *configClient) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		configType := c.Param("type")
		configName := c.Param("name")
		resultConfig, err := cc.retrieveConfig(configName, configType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"config": resultConfig,
		})
	})
}

func getByTypeHandler(cc *configClient) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		configType := c.Param("type")
		resultConfig, err := cc.retrieveConfigs(&configType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"config": resultConfig,
		})
	})
}

func createConfigHandler(cc *configClient) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		createResult, err := cc.createConfig(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"config": createResult,
		})
	})
}

func deleteConfigHandler(cc *configClient) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		deleteResult, err := cc.deleteConfig(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"config": deleteResult,
		})
	})
}

func updateConfigHandler(cc *configClient) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		updateResult, err := cc.updateConfig(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"config": updateResult,
		})
	})
}

func statusInfoHandler() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusText(http.StatusOK),
		})
	})
}

func (client *configClient) selectType(cType string) (entities.ConfigInterface, error) {
	switch cType {
	case mongoConf:
		return new(entities.Mongodb), nil
	case tempConf:
		return new(entities.Tempconfig), nil
	case tsConf:
		return new(entities.Tsconfig), nil
	default:
		log.Printf("Such config: %v does not exist", cType)
		return nil, errors.New("config does not exist")
	}
}

func (client *configClient) retrieveConfig(configName, configType string) (entities.ConfigInterface, error) {
	config, err := client.grpcClient.GetConfigByName(context.Background(), &api.GetConfigByNameRequest{ConfigName: configName, ConfigType: configType})
	if err != nil {
		log.Printf("Error during retrieving config has occurred: %v", err)
		return nil, err
	}
	switch configType {
	case mongoConf:
		var mongodb entities.Mongodb
		err := json.Unmarshal(config.Config, &mongodb)
		if err != nil {
			log.Printf("Unmarshal mongodb err: %v", err)
			return nil, err
		}
		return mongodb, err
	case tempConf:
		var tempconfig entities.Tempconfig
		err := json.Unmarshal(config.Config, &tempconfig)
		if err != nil {
			log.Printf("Unmarshal tempconfig err: %v", err)
			return nil, err
		}
		return tempconfig, err
	case tsConf:
		var tsconfig entities.Tsconfig
		err := json.Unmarshal(config.Config, &tsconfig)
		if err != nil {
			log.Printf("Unmarshal tsconfig err: %v", err)
			return nil, err
		}
		return tsconfig, err

	}
	log.Printf("Such config: %v does not exist", configType)
	return nil, errors.New("config does not exist")

}

func (client *configClient) retrieveConfigs(configType *string) ([]entities.ConfigInterface, error) {
	stream, err := client.grpcClient.GetConfigsByType(context.Background(), &api.GetConfigsByTypeRequest{ConfigType: *configType})
	if err != nil {
		log.Printf("Error during retrieving stream configs has occurred:%v", err)
		return nil, err
	}
	var resultConfigs []entities.ConfigInterface
	for {
		config, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error during streaming has occurred: %v", err)
			return nil, err
		}
		switch *configType {
		case mongoConf:
			var mongodb entities.Mongodb
			err := json.Unmarshal(config.Config, &mongodb)
			if err != nil {
				log.Printf("Unmarshal mongodb err: %v", err)
				return nil, err
			}
			resultConfigs = append(resultConfigs, mongodb)
		case tempConf:
			var tempconfig entities.Tempconfig
			err := json.Unmarshal(config.Config, &tempconfig)
			if err != nil {
				log.Printf("Unmarshal tempconfig err: %v", err)
				return nil, err
			}
			resultConfigs = append(resultConfigs, tempconfig)
		case tsConf:
			var tsconfig entities.Tsconfig
			err := json.Unmarshal(config.Config, &tsconfig)
			if err != nil {
				log.Printf("Unmarshal tsconfig err: %v", err)
				return nil, err
			}
			resultConfigs = append(resultConfigs, tsconfig)
		default:
			log.Printf("Such config: %v does not exist", *configType)
			return nil, err
		}
	}
	return resultConfigs, nil
}

func (client *configClient) createConfig(c *gin.Context) (*api.Responce, error) {
	configType := c.Param("type")
	switch configType {
	case mongoConf:
		domain := c.Request.PostFormValue("domain")
		mongodb, err := strconv.ParseBool(c.Request.PostFormValue("mongodb"))
		if err != nil {
			return nil, err
		}
		host := c.Request.PostFormValue("host")
		port := c.Request.PostFormValue("port")
		c := entities.Mongodb{Domain: domain, Mongodb: mongodb, Host: host, Port: port}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		result, err := client.grpcClient.CreateConfig(context.Background(), &api.Config{ConfigType: configType, Config: bytes})
		if err != nil {
			return nil, err
		}
		return result, err

	case tempConf:
		restApiRoot := c.Request.PostFormValue("restApiRoot")
		host := c.Request.PostFormValue("host")
		port := c.Request.PostFormValue("port")
		remoting := c.Request.PostFormValue("remoting")
		legasyExplorer, err := strconv.ParseBool(c.Request.PostFormValue("legasyExplorer"))
		if err != nil {
			return nil, err
		}
		c := entities.Tempconfig{RestApiRoot: restApiRoot, Host: host, Port: port, Remoting: remoting, LegasyExplorer: legasyExplorer}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		result, err := client.grpcClient.CreateConfig(context.Background(), &api.Config{ConfigType: configType, Config: bytes})
		if err != nil {
			return nil, err
		}
		return result, err
	case tsConf:
		module := c.Request.PostFormValue("module")
		target := c.Request.PostFormValue("target")
		sourceMap, err := strconv.ParseBool(c.Request.PostFormValue("sourceMap"))
		if err != nil {
			return nil, err
		}
		excluding, err := strconv.Atoi(c.Request.PostFormValue("excluding"))
		if err != nil {
			return nil, err
		}

		c := entities.Tsconfig{Module: module, Target: target, SourceMap: sourceMap, Excluding: excluding}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		result, err := client.grpcClient.CreateConfig(context.Background(), &api.Config{ConfigType: configType, Config: bytes})
		if err != nil {
			return nil, err
		}
		return result, err
	}
	log.Printf("Such config: %v does not exist", configType)
	return nil, errors.New("config does not exist")
}

func (client *configClient) updateConfig(c *gin.Context) (*api.Responce, error) {
	configType := c.Param("type")
	switch configType {
	case mongoConf:
		domain := c.Request.PostFormValue("domain")
		mongodb, err := strconv.ParseBool(c.Request.PostFormValue("mongodb"))
		if err != nil {
			return nil, err
		}
		host := c.Request.PostFormValue("host")
		port := c.Request.PostFormValue("port")
		c := entities.Mongodb{Domain: domain, Mongodb: mongodb, Host: host, Port: port}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		result, err := client.grpcClient.UpdateConfig(context.Background(), &api.Config{ConfigType: configType, Config: bytes})
		if err != nil {
			return nil, err
		}
		return result, err

	case tempConf:
		restApiRoot := c.Request.PostFormValue("restApiRoot")
		host := c.Request.PostFormValue("host")
		port := c.Request.PostFormValue("port")
		remoting := c.Request.PostFormValue("remoting")
		legasyExplorer, err := strconv.ParseBool(c.Request.PostFormValue("legasyExplorer"))
		if err != nil {
			return nil, err
		}
		c := entities.Tempconfig{RestApiRoot: restApiRoot, Host: host, Port: port, Remoting: remoting, LegasyExplorer: legasyExplorer}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		result, err := client.grpcClient.UpdateConfig(context.Background(), &api.Config{ConfigType: configType, Config: bytes})
		if err != nil {
			return nil, err
		}
		return result, err
	case tsConf:
		module := c.Request.PostFormValue("module")
		target := c.Request.PostFormValue("target")
		sourceMap, err := strconv.ParseBool(c.Request.PostFormValue("sourceMap"))
		if err != nil {
			return nil, err
		}
		excluding, err := strconv.Atoi(c.Request.PostFormValue("excluding"))
		if err != nil {
			return nil, err
		}

		c := entities.Tsconfig{Module: module, Target: target, SourceMap: sourceMap, Excluding: excluding}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		result, err := client.grpcClient.UpdateConfig(context.Background(), &api.Config{ConfigType: configType, Config: bytes})
		if err != nil {
			return nil, err
		}
		return result, err
	}
	log.Printf("Such config: %v does not exist", configType)
	return nil, errors.New("config does not exist")
}

func (client *configClient) deleteConfig(c *gin.Context) (*api.Responce, error) {
	configType := c.Param("type")
	configName := c.Param("name")
	return client.grpcClient.DeleteConfig(context.Background(), &api.DeleteConfigRequest{ConfigName: configName, ConfigType: configType})
}
