// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGuarddutyIpsetInvalidLocationRule checks the pattern is valid
type AwsGuarddutyIpsetInvalidLocationRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGuarddutyIpsetInvalidLocationRule returns new rule with default attributes
func NewAwsGuarddutyIpsetInvalidLocationRule() *AwsGuarddutyIpsetInvalidLocationRule {
	return &AwsGuarddutyIpsetInvalidLocationRule{
		resourceType:  "aws_guardduty_ipset",
		attributeName: "location",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyIpsetInvalidLocationRule) Name() string {
	return "aws_guardduty_ipset_invalid_location"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyIpsetInvalidLocationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyIpsetInvalidLocationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyIpsetInvalidLocationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyIpsetInvalidLocationRule) Check(runner tflint.Runner) error {
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
					"location must be 300 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"location must be 1 characters or higher",
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
