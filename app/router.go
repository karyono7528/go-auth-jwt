package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/karyono7528/go-auth-jwt/controller/users"
)

func mapUrls() {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.POST("/api/v1/register", users.Register)
	router.POST("/api/v1/login", users.Login)
	router.GET("/api/v1/user", users.Get)
	router.GET("/api/v1/logout", users.Logout)
}
