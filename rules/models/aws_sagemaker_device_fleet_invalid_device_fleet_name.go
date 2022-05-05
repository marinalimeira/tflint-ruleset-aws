// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule checks the pattern is valid
type AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerDeviceFleetInvalidDeviceFleetNameRule returns new rule with default attributes
func NewAwsSagemakerDeviceFleetInvalidDeviceFleetNameRule() *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule {
	return &AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule{
		resourceType:  "aws_sagemaker_device_fleet",
		attributeName: "device_fleet_name",
		max:           63,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Name() string {
	return "aws_sagemaker_device_fleet_invalid_device_fleet_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Check(runner tflint.Runner) error {
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
					"device_fleet_name must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"device_fleet_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`),
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
