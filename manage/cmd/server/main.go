package main

import "manage/pkg/config"
import "github.com/gin-gonic/gin"

func main() {
	config.GetConfig()
	route := gin.New()
	route.Group("/api/v1")
}
