package computeoptimizer

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
)

var (
	computeoptimizerClient *computeoptimizer.Client
)

type OptimizerInput struct {
	ResourceType types.ResourceType
	ResourceArn  string
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("personal"))

	if err != nil {
		log.Fatal(err)
	}

	if computeoptimizerClient == nil {
		computeoptimizerClient = computeoptimizer.NewFromConfig(cfg)
	}
}

func (c *ComputeOptimizerClient) CreateOptimizer(input OptimizerInput) (string, error) {
	_, err := computeoptimizerClient.PutRecommendationPreferences(context.TODO(), &computeoptimizer.PutRecommendationPreferencesInput{
		ResourceType:                  input.ResourceType,
		EnhancedInfrastructureMetrics: types.EnhancedInfrastructureMetricsInactive,
		Scope: &types.Scope{
			Name:  "ResourceArn",
			Value: aws.String(input.ResourceArn),
		},
	})

	return input.ResourceArn, err
}
