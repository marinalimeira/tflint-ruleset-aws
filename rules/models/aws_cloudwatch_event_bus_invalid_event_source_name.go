// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventBusInvalidEventSourceNameRule checks the pattern is valid
type AwsCloudwatchEventBusInvalidEventSourceNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventBusInvalidEventSourceNameRule returns new rule with default attributes
func NewAwsCloudwatchEventBusInvalidEventSourceNameRule() *AwsCloudwatchEventBusInvalidEventSourceNameRule {
	return &AwsCloudwatchEventBusInvalidEventSourceNameRule{
		resourceType:  "aws_cloudwatch_event_bus",
		attributeName: "event_source_name",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^aws\.partner(/[\.\-_A-Za-z0-9]+){2,}$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventBusInvalidEventSourceNameRule) Name() string {
	return "aws_cloudwatch_event_bus_invalid_event_source_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventBusInvalidEventSourceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventBusInvalidEventSourceNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventBusInvalidEventSourceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventBusInvalidEventSourceNameRule) Check(runner tflint.Runner) error {
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
					"event_source_name must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"event_source_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^aws\.partner(/[\.\-_A-Za-z0-9]+){2,}$`),
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
