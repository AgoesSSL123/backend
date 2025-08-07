package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type LoginInput struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

func Login(c *gin.Context) {
  var input LoginInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "Format data salah"})
    return
  }

  // Contoh hardcoded — nanti diganti cek dari DB
  if input.Email == "agus@gmail.com" && input.Password == "admin123" {
    c.JSON(http.StatusOK, gin.H{"message": "Login sukses"})
  } else {
    c.JSON(http.StatusUnauthorized, gin.H{"message": "Email atau password salah"})
  }
}