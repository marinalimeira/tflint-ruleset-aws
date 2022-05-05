// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMAccessKeyInvalidStatusRule checks the pattern is valid
type AwsIAMAccessKeyInvalidStatusRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsIAMAccessKeyInvalidStatusRule returns new rule with default attributes
func NewAwsIAMAccessKeyInvalidStatusRule() *AwsIAMAccessKeyInvalidStatusRule {
	return &AwsIAMAccessKeyInvalidStatusRule{
		resourceType:  "aws_iam_access_key",
		attributeName: "status",
		enum: []string{
			"Active",
			"Inactive",
		},
	}
}

// Name returns the rule name
func (r *AwsIAMAccessKeyInvalidStatusRule) Name() string {
	return "aws_iam_access_key_invalid_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMAccessKeyInvalidStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMAccessKeyInvalidStatusRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMAccessKeyInvalidStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMAccessKeyInvalidStatusRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as status`, truncateLongMessage(val)),
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
