// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferAccessInvalidHomeDirectoryRule checks the pattern is valid
type AwsTransferAccessInvalidHomeDirectoryRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsTransferAccessInvalidHomeDirectoryRule returns new rule with default attributes
func NewAwsTransferAccessInvalidHomeDirectoryRule() *AwsTransferAccessInvalidHomeDirectoryRule {
	return &AwsTransferAccessInvalidHomeDirectoryRule{
		resourceType:  "aws_transfer_access",
		attributeName: "home_directory",
		max:           1024,
		pattern:       regexp.MustCompile(`^$|/.*`),
	}
}

// Name returns the rule name
func (r *AwsTransferAccessInvalidHomeDirectoryRule) Name() string {
	return "aws_transfer_access_invalid_home_directory"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferAccessInvalidHomeDirectoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferAccessInvalidHomeDirectoryRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferAccessInvalidHomeDirectoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferAccessInvalidHomeDirectoryRule) Check(runner tflint.Runner) error {
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
					"home_directory must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^$|/.*`),
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
