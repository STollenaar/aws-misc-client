package computeoptimizer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

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

// CreateOptimizer, creates a recommendation preference for a resource
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

// GetRecommendationPreferences, get the basic preferences of a specific resource type
func (c *ComputeOptimizerClient) GetRecommendationPreferences(input OptimizerInput) ([]types.RecommendationPreferencesDetail, error) {

	out, err := computeoptimizerClient.GetRecommendationPreferences(context.TODO(), &computeoptimizer.GetRecommendationPreferencesInput{
		ResourceType: input.ResourceType,
		Scope: &types.Scope{
			Name:  "ResourceArn",
			Value: aws.String(input.ResourceArn),
		},
	})
	return out.RecommendationPreferencesDetails, err
}

// GetEC2InstanceRecommentdations, get the recommendations of an ec2 instance
func (c *ComputeOptimizerClient) GetEC2InstanceRecommentdations(input OptimizerInput) (types.InstanceRecommendation, error) {
	out, err := computeoptimizerClient.GetEC2InstanceRecommendations(context.TODO(), &computeoptimizer.GetEC2InstanceRecommendationsInput{
		InstanceArns: []string{
			input.ResourceArn,
		},
	})

	if err != nil {
		return types.InstanceRecommendation{}, err
	}
	if len(out.Errors) > 0 {
		return types.InstanceRecommendation{}, flattenErrors(out.Errors)
	}

	return out.InstanceRecommendations[0], nil
}

// GetAutoscalingRecommendations, get the recommendations of an autoscaling group
func (c *ComputeOptimizerClient) GetAutoscalingRecommendations(input OptimizerInput) (types.AutoScalingGroupRecommendation, error) {
	out, err := computeoptimizerClient.GetAutoScalingGroupRecommendations(context.TODO(), &computeoptimizer.GetAutoScalingGroupRecommendationsInput{
		AutoScalingGroupArns: []string{
			input.ResourceArn,
		},
	})

	if err != nil {
		return types.AutoScalingGroupRecommendation{}, err
	}
	if len(out.Errors) > 0 {
		return types.AutoScalingGroupRecommendation{}, flattenErrors(out.Errors)
	}

	return out.AutoScalingGroupRecommendations[0], nil
}

func flattenErrors(rErrors []types.GetRecommendationError) error {
	var messages []string

	for _, e := range rErrors {
		messages = append(messages, fmt.Sprintf(
			"Error Code: %s, Identifier: %s, Message: %s",
			*e.Code,
			*e.Identifier,
			*e.Message,
		))
	}
	return errors.New(strings.Join(messages, "\n"))
}
