// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayGatewayInvalidGatewayTimezoneRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidGatewayTimezoneRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayGatewayInvalidGatewayTimezoneRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidGatewayTimezoneRule() *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule {
	return &AwsStoragegatewayGatewayInvalidGatewayTimezoneRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "gateway_timezone",
		max:           10,
		min:           3,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Name() string {
	return "aws_storagegateway_gateway_invalid_gateway_timezone"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidGatewayTimezoneRule) Check(runner tflint.Runner) error {
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
					"gateway_timezone must be 10 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"gateway_timezone must be 3 characters or higher",
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
