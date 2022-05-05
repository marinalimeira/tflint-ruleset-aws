// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventRuleInvalidRoleArnRule checks the pattern is valid
type AwsCloudwatchEventRuleInvalidRoleArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchEventRuleInvalidRoleArnRule returns new rule with default attributes
func NewAwsCloudwatchEventRuleInvalidRoleArnRule() *AwsCloudwatchEventRuleInvalidRoleArnRule {
	return &AwsCloudwatchEventRuleInvalidRoleArnRule{
		resourceType:  "aws_cloudwatch_event_rule",
		attributeName: "role_arn",
		max:           1600,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Name() string {
	return "aws_cloudwatch_event_rule_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventRuleInvalidRoleArnRule) Check(runner tflint.Runner) error {
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
					"role_arn must be 1600 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 1 characters or higher",
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
