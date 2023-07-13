# Bucket Utility Tool

This tool is used to create and delete buckets in AWS

## Requirements

Set the following environment variables
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_REGION="us-east-1"

These are valid regions
{"af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-southeast-1","ap-southeast-2", "ap-southeast-3", "ca-central-1", "cn-north-1", "cn-northwest-1", "EU", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1","eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2"}

## Running the Tool
Install the project
`go install cloud-provision`

Run the project
`/go/bin/cloud-provision`

You will be asked if you want to create or delete a bucket, make sure you use lowercase

You will then be asked what the name of the bucket is

