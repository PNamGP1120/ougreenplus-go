package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewAWSConfig(region string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		panic("❌ Failed to load AWS config: " + err.Error())
	}
	return cfg
}
