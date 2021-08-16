package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

const (
	baseURL      = "http://localhost:8080"
	productURL   = "/products"
	cartURL      = "/carts"
	addToCartURL = cartURL + "/items"
	orderURL     = "/orders"
)

type BuyerID int

func (id BuyerID) String() string {
	return fmt.Sprintf("%d", id)
}

var (
	productID   int
	cartIDs     = make(map[BuyerID]int)
	cartItemIDs = make(map[BuyerID]int)
	buyerIDs    = []BuyerID{90, 11, 15, 111, 221, 1311, 14511, 23451, 55112, 4111}
	wg          = new(sync.WaitGroup)
)

func main() {
	createProduct()
	getCardID()
	addToCart()
	order()
	wg.Wait()
}

func createProduct() {
	productData := map[string]interface{}{
		"name":      "Product Test 4",
		"seller_id": 10,
		"price":     10000,
		"active":    true,
		"stock":     10,
		"weight":    1000,
	}
	client := &http.Client{}
	jsonString, _ := json.Marshal(productData)
	jsonBody := bytes.NewReader(jsonString)
	req, _ := http.NewRequest(http.MethodPost, baseURL+productURL, jsonBody)
	resp, _ := client.Do(req)
	if resp.StatusCode == 201 {
		response := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&response)
		productID = getIDFromResponse(response)
	}
}

func getCardID() {
	for _, buyerID := range buyerIDs {
		client := &http.Client{}
		r, _ := http.NewRequest(http.MethodGet, baseURL+cartURL+"?user_id="+buyerID.String(), nil)
		resp, _ := client.Do(r)
		if resp.StatusCode == 200 {
			response := make(map[string]interface{})
			json.NewDecoder(resp.Body).Decode(&response)
			cartIDfloat := response["data"].(map[string]interface{})["id"].(float64)
			cartIDs[buyerID] = int(cartIDfloat)
		}
	}
}

func addToCart() {
	for _, buyerID := range buyerIDs {
		addToCartData := map[string]interface{}{
			"cart_id":    cartIDs[buyerID],
			"product_id": productID,
			"quantity":   rand.Intn(2) + 1, // +1 preventif if get 0 value
		}
		client := &http.Client{}
		jsonString, _ := json.Marshal(addToCartData)
		jsonBody := bytes.NewReader(jsonString)
		req, _ := http.NewRequest(http.MethodPost, baseURL+addToCartURL, jsonBody)
		resp, _ := client.Do(req)
		if resp.StatusCode == 201 {
			response := make(map[string]interface{})
			json.NewDecoder(resp.Body).Decode(&response)
			cartItemIDs[buyerID] = getIDFromResponse(response)
		}
	}
}

func getIDFromResponse(response map[string]interface{}) int {
	floatID := response["data"].(map[string]interface{})["id"].(float64)
	return int(floatID)
}

func order() {
	for _, buyerID := range buyerIDs {
		wg.Add(1)
		go func(buyerID BuyerID, wg *sync.WaitGroup) {
			defer wg.Done()
			orderData := map[string]interface{}{
				"buyer_id":      buyerID,
				"cart_item_ids": []int{cartItemIDs[buyerID]},
			}
			client := &http.Client{}
			jsonString, _ := json.Marshal(orderData)
			jsonBody := bytes.NewReader(jsonString)
			log.Printf("buyer %d start checkout\n", buyerID)
			req, _ := http.NewRequest(http.MethodPost, baseURL+orderURL, jsonBody)
			resp, _ := client.Do(req)
			if resp.StatusCode == 201 {
				log.Printf("buyer id %d got success make the order with response code %d\n", buyerID, resp.StatusCode)
			} else {
				response := make(map[string]interface{})
				json.NewDecoder(resp.Body).Decode(&response)
				reason := response["error"].(string)
				log.Printf("buyer id %d fail make the order because %s with response code %d\n", buyerID, reason, resp.StatusCode)
			}
		}(buyerID, wg)
	}
}
