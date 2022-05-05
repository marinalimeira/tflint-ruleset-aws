// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule checks the pattern is valid
type AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsoadminAccountAssignmentInvalidPermissionSetArnRule returns new rule with default attributes
func NewAwsSsoadminAccountAssignmentInvalidPermissionSetArnRule() *AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule {
	return &AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule{
		resourceType:  "aws_ssoadmin_account_assignment",
		attributeName: "permission_set_arn",
		max:           1224,
		min:           10,
		pattern:       regexp.MustCompile(`^arn:aws:sso:::permissionSet/(sso)?ins-[a-zA-Z0-9-.]{16}/ps-[a-zA-Z0-9-./]{16}$`),
	}
}

// Name returns the rule name
func (r *AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule) Name() string {
	return "aws_ssoadmin_account_assignment_invalid_permission_set_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsoadminAccountAssignmentInvalidPermissionSetArnRule) Check(runner tflint.Runner) error {
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
					"permission_set_arn must be 1224 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"permission_set_arn must be 10 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws:sso:::permissionSet/(sso)?ins-[a-zA-Z0-9-.]{16}/ps-[a-zA-Z0-9-./]{16}$`),
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
