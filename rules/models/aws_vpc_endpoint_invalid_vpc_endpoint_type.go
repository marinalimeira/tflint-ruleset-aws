// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsVpcEndpointInvalidVpcEndpointTypeRule checks the pattern is valid
type AwsVpcEndpointInvalidVpcEndpointTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsVpcEndpointInvalidVpcEndpointTypeRule returns new rule with default attributes
func NewAwsVpcEndpointInvalidVpcEndpointTypeRule() *AwsVpcEndpointInvalidVpcEndpointTypeRule {
	return &AwsVpcEndpointInvalidVpcEndpointTypeRule{
		resourceType:  "aws_vpc_endpoint",
		attributeName: "vpc_endpoint_type",
		enum: []string{
			"Interface",
			"Gateway",
			"GatewayLoadBalancer",
		},
	}
}

// Name returns the rule name
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Name() string {
	return "aws_vpc_endpoint_invalid_vpc_endpoint_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as vpc_endpoint_type`, truncateLongMessage(val)),
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
