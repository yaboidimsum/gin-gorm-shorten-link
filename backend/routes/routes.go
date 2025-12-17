package routes

import (
	"time" // Jangan lupa import ini untuk time.Hour

	"github.com/gin-contrib/cors" // Import library CORS
	"github.com/gin-gonic/gin"
	
	"go-url-shortener/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},		
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},		
		ExposeHeaders:    []string{"Content-Length"},		
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// --- SETUP ROUTING ---
	v1 := r.Group("/v1")
	{
		// Ingat: Frontend nanti POST ke http://localhost:8080/v1/shorten
		v1.POST("/shorten", controllers.CreateShortLink)
	}

	// Endpoint Redirect (Public Access)
	r.GET("/:code", controllers.RedirectToOriginal)
}