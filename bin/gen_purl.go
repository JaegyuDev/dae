package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/3AM-Developer/dae/database"

	"github.com/3AM-Developer/dae/initializers"
	"github.com/3AM-Developer/dae/models"
)

func generatePreSignedURL() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println("Error generating random bytes:", err)
		return "", err
	}
	token := hex.EncodeToString(bytes)
	log.Println("Generated token:", token)

	otu := models.OneTimeURL{
		Token: token,
		Used:  false,
	}

	// Save the token to the database
	if err := models.CreatePreSignedURL(&otu); err != nil {
		log.Println("Error saving token to the database:", err)
		return "", err
	}

	return fmt.Sprintf("/signup?token=%s", token), nil
}

func init() {
	initializers.LoadEnvVars()
	database.ConnectToDb()
	initializers.SyncDB()
}

func main() {
	// Generate a new token
	url, err := generatePreSignedURL()
	if err != nil {
		log.Fatal(err)
	}
	// Print the pre-signed URL to the console
	fmt.Printf("Generated Pre-Signed URL: %s", url)
}
