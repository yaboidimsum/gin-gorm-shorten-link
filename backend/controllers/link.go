package controllers

import (
	"go-url-shortener/config"
	"go-url-shortener/models"
	"go-url-shortener/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortLink( c *gin.Context){

	//Konsep Data Transfer Object (likely vulneraibility)
	/*
	Model (models.Link): Adalah cerminan Tabel Database.
	Input (var input): Adalah cerminan Formulir Website.
	
	*/
	var input struct{
		OriginalURL string `json:"original_url" binding:"required"`
		Title string `json:"title" binding:"required"`
	}

	// Input validation JSON
	// Bedaya shouldBind dan Bind adalah, kalau shouldBind kamu bisa custom error message, kalau bindJson saja dia akan mengembalikan error tanpa custom error message
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"missing required inputs"})
		return
	}

	// Make link model
	link := models.Link{
		OriginalURL: input.OriginalURL,
		Title: input.Title,
		ShortCode: utils.GenerateCode(6),
	}

	err = config.DB.Create(&link).Error;

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to create (500)"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"link is shortened",
		"short_url":"http://localhost:8000/" + link.ShortCode,
		"data": link,
	})
}

func RedirectToOriginal(c *gin.Context){
	shortCode := c.Param("code")

	var link models.Link

	err:=config.DB.Where("short_code = ?", shortCode).First(&link).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":"link may be expired or does not exist",
		})
		return
	}

	c.Redirect(http.StatusFound, link.OriginalURL)
}