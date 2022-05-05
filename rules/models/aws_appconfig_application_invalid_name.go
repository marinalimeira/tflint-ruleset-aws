// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppconfigApplicationInvalidNameRule checks the pattern is valid
type AwsAppconfigApplicationInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppconfigApplicationInvalidNameRule returns new rule with default attributes
func NewAwsAppconfigApplicationInvalidNameRule() *AwsAppconfigApplicationInvalidNameRule {
	return &AwsAppconfigApplicationInvalidNameRule{
		resourceType:  "aws_appconfig_application",
		attributeName: "name",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppconfigApplicationInvalidNameRule) Name() string {
	return "aws_appconfig_application_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppconfigApplicationInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppconfigApplicationInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppconfigApplicationInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppconfigApplicationInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
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
