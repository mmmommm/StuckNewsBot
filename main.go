package main

import (
    "goPractice/handler"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.POST("/api/v1/contact", handler.SendMail)
    r.Run()
}