package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"strconv"

	"time"

	"github.com/YAWAL/ConfRESTcli/entitie"
	"github.com/YAWAL/GetMeConfAPI/api"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (client *configClient) selectType(cType string) (entitie.ConfigInterface, error) {
	switch cType {
	case mongoConf:
		return new(entitie.Mongodb), nil
	case tempConf:
		return new(entitie.Tempconfig), nil
	case tsConf:
		return new(entitie.Tsconfig), nil
	default:
		log.Printf("Such config: %v does not exist", cType)
		return nil, errors.New("config does not exist")
	}
}

func (client *configClient) retrieveConfig(c *gin.Context) (entitie.ConfigInterface, error) {
	configType := c.Param("type")
	configName := c.Param("name")
	ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
	defer cancel()
	config, err := client.grpcClient.GetConfigByName(ctx, &api.GetConfigByNameRequest{ConfigName: configName, ConfigType: configType})
	if err != nil {
		log.Printf("Error during retrieving config has occurred: %v", err)
		return nil, err
	}
	switch configType {
	case mongoConf:
		var mongodb entitie.Mongodb
		err := json.Unmarshal(config.Config, &mongodb)
		if err != nil {
			log.Printf("Unmarshal mongodb err: %v", err)
			return nil, err
		}
		return mongodb, err
	case tempConf:
		var tempconfig entitie.Tempconfig
		err := json.Unmarshal(config.Config, &tempconfig)
		if err != nil {
			log.Printf("Unmarshal tempconfig err: %v", err)
			return nil, err
		}
		return tempconfig, err
	case tsConf:
		var tsconfig entitie.Tsconfig
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

func (client *configClient) retrieveConfigs(c *gin.Context) ([]entitie.ConfigInterface, error) {
	configType := c.Param("type")
	ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
	defer cancel()
	stream, err := client.grpcClient.GetConfigsByType(ctx, &api.GetConfigsByTypeRequest{ConfigType: configType})
	if err != nil {
		log.Printf("Error during retrieving stream configs has occurred:%v", err)
		return nil, err
	}
	var resultConfigs []entitie.ConfigInterface
	for {
		config, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error during streaming has occurred: %v", err)
			return nil, err
		}
		switch configType {
		case mongoConf:
			var mongodb entitie.Mongodb
			err := json.Unmarshal(config.Config, &mongodb)
			if err != nil {
				log.Printf("Unmarshal mongodb err: %v", err)
				return nil, err
			}
			resultConfigs = append(resultConfigs, mongodb)
		case tempConf:
			var tempconfig entitie.Tempconfig
			err := json.Unmarshal(config.Config, &tempconfig)
			if err != nil {
				log.Printf("Unmarshal tempconfig err: %v", err)
				return nil, err
			}
			resultConfigs = append(resultConfigs, tempconfig)
		case tsConf:
			var tsconfig entitie.Tsconfig
			err := json.Unmarshal(config.Config, &tsconfig)
			if err != nil {
				log.Printf("Unmarshal tsconfig err: %v", err)
				return nil, err
			}
			resultConfigs = append(resultConfigs, tsconfig)
		default:
			log.Printf("Such config: %v does not exist", configType)
			return nil, errors.New("config does not exist")
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
		c := entitie.Mongodb{Domain: domain, Mongodb: mongodb, Host: host, Port: port}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
		defer cancel()
		result, err := client.grpcClient.CreateConfig(ctx, &api.Config{ConfigType: configType, Config: bytes})
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
		c := entitie.Tempconfig{RestApiRoot: restApiRoot, Host: host, Port: port, Remoting: remoting, LegasyExplorer: legasyExplorer}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
		defer cancel()
		result, err := client.grpcClient.CreateConfig(ctx, &api.Config{ConfigType: configType, Config: bytes})
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
		c := entitie.Tsconfig{Module: module, Target: target, SourceMap: sourceMap, Excluding: excluding}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
		defer cancel()
		result, err := client.grpcClient.CreateConfig(ctx, &api.Config{ConfigType: configType, Config: bytes})
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
		c := entitie.Mongodb{Domain: domain, Mongodb: mongodb, Host: host, Port: port}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
		defer cancel()
		result, err := client.grpcClient.UpdateConfig(ctx, &api.Config{ConfigType: configType, Config: bytes})
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
		c := entitie.Tempconfig{RestApiRoot: restApiRoot, Host: host, Port: port, Remoting: remoting, LegasyExplorer: legasyExplorer}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
		defer cancel()
		result, err := client.grpcClient.UpdateConfig(ctx, &api.Config{ConfigType: configType, Config: bytes})
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

		c := entitie.Tsconfig{Module: module, Target: target, SourceMap: sourceMap, Excluding: excluding}
		bytes, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
		defer cancel()
		result, err := client.grpcClient.UpdateConfig(ctx, &api.Config{ConfigType: configType, Config: bytes})
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
	ctx, cancel := context.WithTimeout(context.Background(), client.contextDeadline*time.Millisecond)
	defer cancel()
	return client.grpcClient.DeleteConfig(ctx, &api.DeleteConfigRequest{ConfigName: configName, ConfigType: configType})
}
