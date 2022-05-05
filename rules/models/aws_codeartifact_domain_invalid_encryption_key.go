// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactDomainInvalidEncryptionKeyRule checks the pattern is valid
type AwsCodeartifactDomainInvalidEncryptionKeyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactDomainInvalidEncryptionKeyRule returns new rule with default attributes
func NewAwsCodeartifactDomainInvalidEncryptionKeyRule() *AwsCodeartifactDomainInvalidEncryptionKeyRule {
	return &AwsCodeartifactDomainInvalidEncryptionKeyRule{
		resourceType:  "aws_codeartifact_domain",
		attributeName: "encryption_key",
		max:           1011,
		min:           1,
		pattern:       regexp.MustCompile(`^\S+$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Name() string {
	return "aws_codeartifact_domain_invalid_encryption_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactDomainInvalidEncryptionKeyRule) Check(runner tflint.Runner) error {
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
					"encryption_key must be 1011 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"encryption_key must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\S+$`),
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
