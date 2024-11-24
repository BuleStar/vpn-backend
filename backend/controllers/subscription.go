package controllers

import (
    "net/http"
    "vpn-backend/services"

    "github.com/gin-gonic/gin"
)

// 导入订阅
func ImportSubscription(c *gin.Context) {
    var importData struct {
        URL string `json:"url" binding:"required"`
    }

    if err := c.ShouldBindJSON(&importData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.ImportSubscription(importData.URL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Subscription imported successfully"})
}
