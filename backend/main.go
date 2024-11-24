package main

import (
    "vpn-backend/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := routes.SetupRouter()
    r.Run(":8080")
}
