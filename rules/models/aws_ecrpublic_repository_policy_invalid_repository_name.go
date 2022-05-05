// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule checks the pattern is valid
type AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule returns new rule with default attributes
func NewAwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule() *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule {
	return &AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule{
		resourceType:  "aws_ecrpublic_repository_policy",
		attributeName: "repository_name",
		max:           205,
		min:           2,
		pattern:       regexp.MustCompile(`^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Name() string {
	return "aws_ecrpublic_repository_policy_invalid_repository_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Check(runner tflint.Runner) error {
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
					"repository_name must be 205 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"repository_name must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
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
