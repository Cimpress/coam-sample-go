package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	sub                = "adfs|cbaldauf@cimpress.com"
	resourceType       = "merchants"
	resourceIdentifier = "vistaprint"
)

type OAuthTokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func main() {
	// Load environment variables from .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		log.Fatal("Cannot find .env file")
	}

	// Read env variables
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("CLIENT_ID or CLIENT_SECRET missing!")
	}

	// Request access token for client_id
	body := map[string]interface{}{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"audience":      "https://api.cimpress.io/",
		"grant_type":    "client_credentials",
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Retrieving IAM access token for client ID %s\n", clientID)

	resp, err := http.Post("https://cimpress.auth0.com/oauth/token", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	t := OAuthTokenResp{}
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	log.Printf("Access Token: %s\n", t.AccessToken)

	// Request IAM permisisons
	url := fmt.Sprintf("https://api.cimpress.io/auth/iam/v0/user-permissions/%s/%s/%s", sub, resourceType, resourceIdentifier)
	authHeader := fmt.Sprintf("Bearer %s", t.AccessToken)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", authHeader)

	log.Printf("Retrieving IAM permissions for %s on %s %s", sub, resourceType, resourceIdentifier)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("===== RESPONSE =====")
	log.Println(string(respBody))
}
