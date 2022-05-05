// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAPIGatewayIntegrationInvalidConnectionTypeRule checks the pattern is valid
type AwsAPIGatewayIntegrationInvalidConnectionTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayIntegrationInvalidConnectionTypeRule returns new rule with default attributes
func NewAwsAPIGatewayIntegrationInvalidConnectionTypeRule() *AwsAPIGatewayIntegrationInvalidConnectionTypeRule {
	return &AwsAPIGatewayIntegrationInvalidConnectionTypeRule{
		resourceType:  "aws_api_gateway_integration",
		attributeName: "connection_type",
		enum: []string{
			"INTERNET",
			"VPC_LINK",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Name() string {
	return "aws_api_gateway_integration_invalid_connection_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as connection_type`, truncateLongMessage(val)),
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
