// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaPermissionInvalidQualifierRule checks the pattern is valid
type AwsLambdaPermissionInvalidQualifierRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidQualifierRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidQualifierRule() *AwsLambdaPermissionInvalidQualifierRule {
	return &AwsLambdaPermissionInvalidQualifierRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "qualifier",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^(|[a-zA-Z0-9$_-]+)$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidQualifierRule) Name() string {
	return "aws_lambda_permission_invalid_qualifier"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidQualifierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaPermissionInvalidQualifierRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidQualifierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidQualifierRule) Check(runner tflint.Runner) error {
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
					"qualifier must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"qualifier must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(|[a-zA-Z0-9$_-]+)$`),
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
