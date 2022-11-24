package app

import (
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"kingclover.com/api/pkg/config"
	"kingclover.com/api/pkg/database"
)

func Bootstrap() {
	database.InitializeDB()

	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	c := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8000"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}

	r.Use(cors.New(c))

	SetupRoutes(r)

	r.Run(config.ServerPort)
}
