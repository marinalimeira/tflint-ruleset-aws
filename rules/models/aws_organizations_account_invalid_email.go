// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsOrganizationsAccountInvalidEmailRule checks the pattern is valid
type AwsOrganizationsAccountInvalidEmailRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsAccountInvalidEmailRule returns new rule with default attributes
func NewAwsOrganizationsAccountInvalidEmailRule() *AwsOrganizationsAccountInvalidEmailRule {
	return &AwsOrganizationsAccountInvalidEmailRule{
		resourceType:  "aws_organizations_account",
		attributeName: "email",
		max:           64,
		min:           6,
		pattern:       regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsAccountInvalidEmailRule) Name() string {
	return "aws_organizations_account_invalid_email"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsAccountInvalidEmailRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsAccountInvalidEmailRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsAccountInvalidEmailRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsAccountInvalidEmailRule) Check(runner tflint.Runner) error {
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
					"email must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"email must be 6 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\s@]+@[^\s@]+\.[^\s@]+$`),
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
