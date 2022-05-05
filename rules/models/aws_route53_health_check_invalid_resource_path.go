// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53HealthCheckInvalidResourcePathRule checks the pattern is valid
type AwsRoute53HealthCheckInvalidResourcePathRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53HealthCheckInvalidResourcePathRule returns new rule with default attributes
func NewAwsRoute53HealthCheckInvalidResourcePathRule() *AwsRoute53HealthCheckInvalidResourcePathRule {
	return &AwsRoute53HealthCheckInvalidResourcePathRule{
		resourceType:  "aws_route53_health_check",
		attributeName: "resource_path",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsRoute53HealthCheckInvalidResourcePathRule) Name() string {
	return "aws_route53_health_check_invalid_resource_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53HealthCheckInvalidResourcePathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53HealthCheckInvalidResourcePathRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53HealthCheckInvalidResourcePathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53HealthCheckInvalidResourcePathRule) Check(runner tflint.Runner) error {
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
					"resource_path must be 255 characters or less",
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
