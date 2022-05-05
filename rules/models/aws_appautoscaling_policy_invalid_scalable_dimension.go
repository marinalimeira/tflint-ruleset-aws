// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppautoscalingPolicyInvalidScalableDimensionRule checks the pattern is valid
type AwsAppautoscalingPolicyInvalidScalableDimensionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppautoscalingPolicyInvalidScalableDimensionRule returns new rule with default attributes
func NewAwsAppautoscalingPolicyInvalidScalableDimensionRule() *AwsAppautoscalingPolicyInvalidScalableDimensionRule {
	return &AwsAppautoscalingPolicyInvalidScalableDimensionRule{
		resourceType:  "aws_appautoscaling_policy",
		attributeName: "scalable_dimension",
		enum: []string{
			"ecs:service:DesiredCount",
			"ec2:spot-fleet-request:TargetCapacity",
			"elasticmapreduce:instancegroup:InstanceCount",
			"appstream:fleet:DesiredCapacity",
			"dynamodb:table:ReadCapacityUnits",
			"dynamodb:table:WriteCapacityUnits",
			"dynamodb:index:ReadCapacityUnits",
			"dynamodb:index:WriteCapacityUnits",
			"rds:cluster:ReadReplicaCount",
			"sagemaker:variant:DesiredInstanceCount",
			"custom-resource:ResourceType:Property",
			"comprehend:document-classifier-endpoint:DesiredInferenceUnits",
			"comprehend:entity-recognizer-endpoint:DesiredInferenceUnits",
			"lambda:function:ProvisionedConcurrency",
			"cassandra:table:ReadCapacityUnits",
			"cassandra:table:WriteCapacityUnits",
			"kafka:broker-storage:VolumeSize",
			"elasticache:replication-group:NodeGroups",
			"elasticache:replication-group:Replicas",
			"neptune:cluster:ReadReplicaCount",
		},
	}
}

// Name returns the rule name
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Name() string {
	return "aws_appautoscaling_policy_invalid_scalable_dimension"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppautoscalingPolicyInvalidScalableDimensionRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as scalable_dimension`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
