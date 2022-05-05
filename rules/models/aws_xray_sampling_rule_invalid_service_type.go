// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsXraySamplingRuleInvalidServiceTypeRule checks the pattern is valid
type AwsXraySamplingRuleInvalidServiceTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsXraySamplingRuleInvalidServiceTypeRule returns new rule with default attributes
func NewAwsXraySamplingRuleInvalidServiceTypeRule() *AwsXraySamplingRuleInvalidServiceTypeRule {
	return &AwsXraySamplingRuleInvalidServiceTypeRule{
		resourceType:  "aws_xray_sampling_rule",
		attributeName: "service_type",
		max:           64,
	}
}

// Name returns the rule name
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Name() string {
	return "aws_xray_sampling_rule_invalid_service_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Check(runner tflint.Runner) error {
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
					"service_type must be 64 characters or less",
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
