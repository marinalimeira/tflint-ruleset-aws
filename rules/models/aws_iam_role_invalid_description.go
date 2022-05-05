// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMRoleInvalidDescriptionRule checks the pattern is valid
type AwsIAMRoleInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsIAMRoleInvalidDescriptionRule returns new rule with default attributes
func NewAwsIAMRoleInvalidDescriptionRule() *AwsIAMRoleInvalidDescriptionRule {
	return &AwsIAMRoleInvalidDescriptionRule{
		resourceType:  "aws_iam_role",
		attributeName: "description",
		max:           1000,
		pattern:       regexp.MustCompile(`^[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}]*$`),
	}
}

// Name returns the rule name
func (r *AwsIAMRoleInvalidDescriptionRule) Name() string {
	return "aws_iam_role_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMRoleInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMRoleInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMRoleInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMRoleInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 1000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}]*$`),
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
