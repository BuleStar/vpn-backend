# **VPN 项目开发文档**

## **项目背景与目标**
目标是开发一个支持 **Clash** 和 **Shadowrocket** 订阅的VPN服务，用户通过Web界面管理VPN节点，实现订阅链接的生成和分发。支持以下功能：
- VPN 节点的增删查改。
- 生成 Clash 和 Shadowrocket 格式的订阅链接。
- 提供用户友好的Web管理界面。
- 使用高效和可扩展的技术栈，便于维护和扩展。

---

## **技术栈**
- **后端**: Go + Gin
- **前端**: React.js + Ant Design
- **数据库**: PostgreSQL（存储持久化数据） + Redis（加速缓存请求）
- **部署**: Docker 容器化，支持快速上线。

---

## **项目结构**
```plaintext
vpn-project/
├── backend/             # 后端代码
│   ├── main.go          # 程序入口
│   ├── config/          # 配置文件相关
│   ├── controllers/     # 业务逻辑控制器
│   ├── models/          # 数据库模型
│   ├── routes/          # 路由定义
│   ├── services/        # 逻辑服务层
├── frontend/            # 前端代码
│   ├── public/          # 静态资源
│   ├── src/
│   │   ├── components/  # 公共组件
│   │   ├── pages/       # 页面
│   │   ├── api/         # 封装API请求
│   │   ├── App.js       # React 主入口
├── docker-compose.yml   # Docker 容器编排
├── README.md            # 项目说明
```

---

## **功能详细说明**

### 1. **后端功能**
#### 1.1 节点管理
提供对VPN节点的增删查改功能。

- **数据模型**
```sql
CREATE TABLE nodes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    address VARCHAR(255),
    port INTEGER,
    type VARCHAR(50),       -- 节点类型 (e.g., Vmess, Shadowsocks)
    config TEXT             -- 节点配置
);
```

- **API 路由**
```go
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
```

- **节点控制器**
```go
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
```

#### 1.2 订阅生成
根据用户Key生成Clash或Shadowrocket订阅格式。

- **示例实现**
```go
package controllers

import (
    "net/http"
    "vpn-backend/services"

    "github.com/gin-gonic/gin"
)

// 生成订阅链接
func GenerateSubscription(c *gin.Context) {
    userKey := c.Param("userKey")
    subscription := services.GenerateSubscription(userKey)
    c.String(http.StatusOK, subscription)
}
```

- **服务实现**
```go
package services

import (
    "vpn-backend/models"
    "strings"
)

func GenerateSubscription(userKey string) string {
    nodes := FetchAllNodes() // 获取所有节点
    var subscriptionBuilder strings.Builder

    for _, node := range nodes {
        // 根据Clash或Shadowrocket格式生成订阅内容
        subscriptionBuilder.WriteString(node.ToClashFormat())
    }

    return subscriptionBuilder.String()
}
```

---

### 2. **前端功能**
#### 2.1 节点管理页面
- **功能描述**
  - 展示所有节点。
  - 支持添加、编辑和删除节点。
  
- **页面代码**
```jsx
import React, { useState, useEffect } from "react";
import { Table, Button, Modal, Form, Input } from "antd";
import axios from "axios";

const NodeManagement = () => {
  const [nodes, setNodes] = useState([]);
  const [visible, setVisible] = useState(false);
  const [form] = Form.useForm();

  useEffect(() => {
    fetchNodes();
  }, []);

  const fetchNodes = async () => {
    const res = await axios.get("/api/nodes");
    setNodes(res.data.nodes);
  };

  const handleAddNode = async (values) => {
    await axios.post("/api/nodes", values);
    fetchNodes();
    setVisible(false);
  };

  return (
    <div>
      <Button type="primary" onClick={() => setVisible(true)}>Add Node</Button>
      <Table
        dataSource={nodes}
        columns={[
          { title: "Name", dataIndex: "name" },
          { title: "Address", dataIndex: "address" },
        ]}
      />
      <Modal
        visible={visible}
        title="Add Node"
        onCancel={() => setVisible(false)}
        onOk={() => form.submit()}
      >
        <Form form={form} onFinish={handleAddNode}>
          <Form.Item name="name" label="Name">
            <Input />
          </Form.Item>
          <Form.Item name="address" label="Address">
            <Input />
          </Form.Item>
          <Form.Item name="port" label="Port">
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default NodeManagement;
```

---

### 3. **部署与运行**
#### 3.1 Docker Compose 配置
```yaml
version: '3.9'
services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"

  frontend:
    build: ./frontend
    ports:
      - "3000:3000"

  db:
    image: postgres:latest
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: vpn
    ports:
      - "5432:5432"

  redis:
    image: redis:latest
    ports:
      - "6379:6379"
```

#### 3.2 启动项目
```bash
docker-compose up --build
```

---

### **后续改进**
1. 支持用户认证和权限管理。
2. 添加订阅的访问日志功能。
3. 优化订阅生成的性能，适配更多格式。

---

### **导入其他 Clash 数据订阅**
#### 1. **导入订阅 API**
- **API 路由**
```go
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
```

- **导入订阅控制器**
```go
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
```

- **导入订阅服务**
```go
package services

import (
    "encoding/json"
    "errors"
    "net/http"
    "vpn-backend/models"
    "strings"
)

// ImportSubscription imports data subscriptions from external sources
func ImportSubscription(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return errors.New("failed to fetch subscription data")
    }

    var nodes []models.Node
    if err := json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
        return err
    }

    for _, node := range nodes {
        if err := AddNode(node); err != nil {
            return err
        }
    }

    return nil
}
```

---

如需补充细节或进一步改进，请提出具体需求！
