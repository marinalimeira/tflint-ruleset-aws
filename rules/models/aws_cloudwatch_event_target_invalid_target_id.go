// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventTargetInvalidTargetIDRule checks the pattern is valid
type AwsCloudwatchEventTargetInvalidTargetIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventTargetInvalidTargetIDRule returns new rule with default attributes
func NewAwsCloudwatchEventTargetInvalidTargetIDRule() *AwsCloudwatchEventTargetInvalidTargetIDRule {
	return &AwsCloudwatchEventTargetInvalidTargetIDRule{
		resourceType:  "aws_cloudwatch_event_target",
		attributeName: "target_id",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\.\-_A-Za-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Name() string {
	return "aws_cloudwatch_event_target_invalid_target_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventTargetInvalidTargetIDRule) Check(runner tflint.Runner) error {
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
					"target_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"target_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\.\-_A-Za-z0-9]+$`),
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
