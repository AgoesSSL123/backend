package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/agus/my-hospital-app/config"
)

type RegisterInput struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

func Register(c *gin.Context) {
  var input RegisterInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  _, err := config.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", input.Username, input.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}