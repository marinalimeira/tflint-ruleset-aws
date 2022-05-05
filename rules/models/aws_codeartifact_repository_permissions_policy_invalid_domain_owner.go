// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule checks the pattern is valid
type AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule returns new rule with default attributes
func NewAwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule() *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule {
	return &AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule{
		resourceType:  "aws_codeartifact_repository_permissions_policy",
		attributeName: "domain_owner",
		max:           12,
		min:           12,
		pattern:       regexp.MustCompile(`^[0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule) Name() string {
	return "aws_codeartifact_repository_permissions_policy_invalid_domain_owner"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainOwnerRule) Check(runner tflint.Runner) error {
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
					"domain_owner must be 12 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"domain_owner must be 12 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]{12}$`),
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
