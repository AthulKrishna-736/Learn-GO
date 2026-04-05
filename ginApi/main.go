package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gte=0"`
}

type PersonForm struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

type User struct {
	Name  string `json:"name" binding:"required"`
	Age   uint   `json:"age" binding:"required"`
	Role  string `json:"role" binding:"omitempty"`
	Title string `json:"title" binding:"omitempty"`
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "route tested"})
	})

	router.POST("/person", func(c *gin.Context) {
		var req Person
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": http.StatusBadRequest})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Received data successfully", "data": req})
	})

	router.POST("/person-form", func(c *gin.Context) {
		var req PersonForm

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully received form data", "data": req})
	})

	router.PATCH("/person", func(c *gin.Context) {
		var req Person
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully updated data", "data": req})
	})

	router.DELETE("/person", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, nil)
	})

	router.GET("/person/:id", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Query("name")

		c.JSON(http.StatusOK, gin.H{"message": "Request with query params received successfully", "data": []string{id, name}})
	})

	router.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err})
			return
		}

		c.JSON(http.StatusCreated, User{user.Name, user.Age, "admin", "sample user profile title"})
	})

	_ = router.Run(":8080")
}
