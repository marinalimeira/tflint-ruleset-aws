// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudformationStackSetInvalidTemplateURLRule checks the pattern is valid
type AwsCloudformationStackSetInvalidTemplateURLRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudformationStackSetInvalidTemplateURLRule returns new rule with default attributes
func NewAwsCloudformationStackSetInvalidTemplateURLRule() *AwsCloudformationStackSetInvalidTemplateURLRule {
	return &AwsCloudformationStackSetInvalidTemplateURLRule{
		resourceType:  "aws_cloudformation_stack_set",
		attributeName: "template_url",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackSetInvalidTemplateURLRule) Name() string {
	return "aws_cloudformation_stack_set_invalid_template_url"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackSetInvalidTemplateURLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackSetInvalidTemplateURLRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackSetInvalidTemplateURLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackSetInvalidTemplateURLRule) Check(runner tflint.Runner) error {
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
					"template_url must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"template_url must be 1 characters or higher",
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
