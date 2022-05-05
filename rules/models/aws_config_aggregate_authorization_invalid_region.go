// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigAggregateAuthorizationInvalidRegionRule checks the pattern is valid
type AwsConfigAggregateAuthorizationInvalidRegionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigAggregateAuthorizationInvalidRegionRule returns new rule with default attributes
func NewAwsConfigAggregateAuthorizationInvalidRegionRule() *AwsConfigAggregateAuthorizationInvalidRegionRule {
	return &AwsConfigAggregateAuthorizationInvalidRegionRule{
		resourceType:  "aws_config_aggregate_authorization",
		attributeName: "region",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Name() string {
	return "aws_config_aggregate_authorization_invalid_region"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigAggregateAuthorizationInvalidRegionRule) Check(runner tflint.Runner) error {
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
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"region must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"region must be 1 characters or higher",
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
