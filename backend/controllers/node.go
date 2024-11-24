package controllers

import (
    "net/http"
    "vpn-backend/models"
    "vpn-backend/services"

    "github.com/gin-gonic/gin"
)

// 获取节点列表
func GetNodes(c *gin.Context) {
    nodes := services.FetchAllNodes()
    c.JSON(http.StatusOK, gin.H{"nodes": nodes})
}

// 添加节点
func AddNode(c *gin.Context) {
    var node models.Node
    if err := c.ShouldBindJSON(&node); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    services.AddNode(node)
    c.JSON(http.StatusOK, gin.H{"message": "Node added successfully"})
}

// 更新节点
func UpdateNode(c *gin.Context) {
    id := c.Param("id")
    var node models.Node
    if err := c.ShouldBindJSON(&node); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    services.UpdateNode(id, node)
    c.JSON(http.StatusOK, gin.H{"message": "Node updated successfully"})
}

// 删除节点
func DeleteNode(c *gin.Context) {
    id := c.Param("id")
    services.DeleteNode(id)
    c.JSON(http.StatusOK, gin.H{"message": "Node deleted successfully"})
}
