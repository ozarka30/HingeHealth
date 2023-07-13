package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Array and slice don't have a contains function, this one will confirm the Region is a Valid Region
func validRegion(region string) bool {

	VALID_REGIONS := [27]string{"af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-southeast-1",
		"ap-southeast-2", "ap-southeast-3", "ca-central-1", "cn-north-1", "cn-northwest-1", "EU", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1",
		"eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2"}

	for _, v := range VALID_REGIONS {
		if v == region {
			return true
		}
	}

	return false

}

// Verify all listed requirements are set before starting
func checkRequirements() {
	// Check if AWS Creds exist, error out if they are blank
	keyEnv := os.Getenv("AWS_ACCESS_KEY_ID")
	secretEnv := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if keyEnv == "" || secretEnv == "" {
		log.Fatalf("Error: Check Env Vars AWS_ACCESS_KEY_ID: %s AWS_SECRET_ACCESS_KEY: %s", keyEnv, secretEnv)
	}
	// Check that the Region Env Var is set, also that it's valid
	region := os.Getenv("AWS_REGION")
	if !validRegion(region) {
		log.Fatalf("Error: Verify your region is Valid or Exists AWS_REGION: %s", region)
	}

}

func getCreds(ctx context.Context) aws.Config {

	// Set Config and return it
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Error: Unable to load Configuration, %v", err)
	}

	return cfg
}

// Create a Bucket
// TODO: Bucket Name Validation
// TODO: Bucket Exists Validation
func createBucket(ctx context.Context, bucketName string) {
	// Create client
	client := s3.NewFromConfig(getCreds(ctx))
	_, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		log.Fatalf("Error: Unable to create Bucket, %v", err)
	}
}

func main() {
	ctx := context.TODO()

	checkRequirements()

	fmt.Println(cbo)
	//TODO: Accept Input

	//TODO: Create Bucket

	//TODO: List Buckets

	//TODO: Delete Bucket

}
