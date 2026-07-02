package controllers

import "github.com/gin-gonic/gin"

func TestRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": gin.H{"biryani": "delicious"},
	})
}
