// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayGatewayInvalidMediumChangerTypeRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidMediumChangerTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayGatewayInvalidMediumChangerTypeRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidMediumChangerTypeRule() *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule {
	return &AwsStoragegatewayGatewayInvalidMediumChangerTypeRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "medium_changer_type",
		max:           50,
		min:           2,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Name() string {
	return "aws_storagegateway_gateway_invalid_medium_changer_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Check(runner tflint.Runner) error {
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
					"medium_changer_type must be 50 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"medium_changer_type must be 2 characters or higher",
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
