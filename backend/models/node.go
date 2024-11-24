package models

type Node struct {
    ID      int    `json:"id" gorm:"primary_key"`
    Name    string `json:"name"`
    Address string `json:"address"`
    Port    int    `json:"port"`
    Type    string `json:"type"`
    Config  string `json:"config"`
}

func (n *Node) ToClashFormat() string {
    // Implement the logic to convert the node to Clash format
    return ""
}
