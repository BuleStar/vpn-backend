package routes

import (
    "github.com/gin-gonic/gin"
    "vpn-backend/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    
    // 节点管理
    r.GET("/api/nodes", controllers.GetNodes)
    r.POST("/api/nodes", controllers.AddNode)
    r.PUT("/api/nodes/:id", controllers.UpdateNode)
    r.DELETE("/api/nodes/:id", controllers.DeleteNode)
    
    // 订阅管理
    r.GET("/api/subscribe/:userKey", controllers.GenerateSubscription)
    
    // 导入订阅
    r.POST("/api/import", controllers.ImportSubscription)
    
    return r
}
