package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func getCredentials() aws.Config {
	// Check if Environment variables exist, error out if they are blank
	keyEnv := os.Getenv("AWS_ACCESS_KEY_ID")
	secretEnv := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if keyEnv == "" || secretEnv == "" {
		log.Fatalf("Error: Check Env Vars AWS_ACCESS_KEY_ID: %s AWS_SECRET_ACCESS_KEY: %s", keyEnv, secretEnv)
	}

	// Set Config and return it
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	return cfg
}

func main() {

	//TODO: Accept Input

	//TODO: Get Credentials

	//TODO: Create Bucket

	//TODO: List Buckets

	//TODO: Delete Bucket

}
