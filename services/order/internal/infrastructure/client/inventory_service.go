package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const inventoryServiceAddress = "http://localhost:8071"

type InventoryService struct {
	httpClient *http.Client
}

func NewInventoryService(httpClient *http.Client) *InventoryService {
	return &InventoryService{httpClient: httpClient}
}

func (i *InventoryService) ReserveStock(itemID int64, itemCount int32, orderID int64) error {
	data := struct {
		ItemID    int64 `json:"item_id"`
		ItemCount int32 `json:"item_count"`
		OrderID   int64 `json:"order_id"`
	}{
		ItemID:    itemID,
		ItemCount: itemCount,
		OrderID:   orderID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	url := fmt.Sprintf("%s/%s", inventoryServiceAddress, "stock/reserve")
	contentType := "application/json"
	resp, err := i.httpClient.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("httpClient.Post: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code: 200; received: %d", resp.StatusCode)
	}

	return nil
}

func (i *InventoryService) UnReserveStock(orderID int64) error {
	data := struct {
		OrderID int64 `json:"order_id"`
	}{
		OrderID: orderID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	url := fmt.Sprintf("%s/%s", inventoryServiceAddress, "stock/unreserve")
	contentType := "application/json"
	resp, err := i.httpClient.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("httpClient.Post: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code: 200; received: %d", resp.StatusCode)
	}

	return nil
}
