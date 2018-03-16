package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getByNameHandler(cc *configClient) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		resultConfig, err := cc.retrieveConfig(c)
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
		resultConfig, err := cc.retrieveConfigs(c)
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
