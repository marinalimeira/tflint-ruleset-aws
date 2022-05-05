// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigConfigRuleInvalidDescriptionRule checks the pattern is valid
type AwsConfigConfigRuleInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsConfigConfigRuleInvalidDescriptionRule returns new rule with default attributes
func NewAwsConfigConfigRuleInvalidDescriptionRule() *AwsConfigConfigRuleInvalidDescriptionRule {
	return &AwsConfigConfigRuleInvalidDescriptionRule{
		resourceType:  "aws_config_config_rule",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsConfigConfigRuleInvalidDescriptionRule) Name() string {
	return "aws_config_config_rule_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigConfigRuleInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigConfigRuleInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigConfigRuleInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigConfigRuleInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 256 characters or less",
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
