// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule checks the pattern is valid
type AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsDelegatedAdministratorInvalidAccountIDRule returns new rule with default attributes
func NewAwsOrganizationsDelegatedAdministratorInvalidAccountIDRule() *AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule {
	return &AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule{
		resourceType:  "aws_organizations_delegated_administrator",
		attributeName: "account_id",
		max:           12,
		pattern:       regexp.MustCompile(`^\d{12}$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule) Name() string {
	return "aws_organizations_delegated_administrator_invalid_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsDelegatedAdministratorInvalidAccountIDRule) Check(runner tflint.Runner) error {
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
					"account_id must be 12 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\d{12}$`),
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
