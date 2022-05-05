// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventAPIDestinationInvalidDescriptionRule checks the pattern is valid
type AwsCloudwatchEventAPIDestinationInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventAPIDestinationInvalidDescriptionRule returns new rule with default attributes
func NewAwsCloudwatchEventAPIDestinationInvalidDescriptionRule() *AwsCloudwatchEventAPIDestinationInvalidDescriptionRule {
	return &AwsCloudwatchEventAPIDestinationInvalidDescriptionRule{
		resourceType:  "aws_cloudwatch_event_api_destination",
		attributeName: "description",
		max:           512,
		pattern:       regexp.MustCompile(`^.*$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventAPIDestinationInvalidDescriptionRule) Name() string {
	return "aws_cloudwatch_event_api_destination_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventAPIDestinationInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventAPIDestinationInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventAPIDestinationInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventAPIDestinationInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*$`),
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
