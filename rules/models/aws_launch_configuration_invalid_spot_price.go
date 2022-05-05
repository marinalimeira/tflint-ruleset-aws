// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLaunchConfigurationInvalidSpotPriceRule checks the pattern is valid
type AwsLaunchConfigurationInvalidSpotPriceRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsLaunchConfigurationInvalidSpotPriceRule returns new rule with default attributes
func NewAwsLaunchConfigurationInvalidSpotPriceRule() *AwsLaunchConfigurationInvalidSpotPriceRule {
	return &AwsLaunchConfigurationInvalidSpotPriceRule{
		resourceType:  "aws_launch_configuration",
		attributeName: "spot_price",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsLaunchConfigurationInvalidSpotPriceRule) Name() string {
	return "aws_launch_configuration_invalid_spot_price"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLaunchConfigurationInvalidSpotPriceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLaunchConfigurationInvalidSpotPriceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLaunchConfigurationInvalidSpotPriceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLaunchConfigurationInvalidSpotPriceRule) Check(runner tflint.Runner) error {
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
					"spot_price must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"spot_price must be 1 characters or higher",
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
