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
