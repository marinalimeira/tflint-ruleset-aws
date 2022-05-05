// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightUserInvalidNamespaceRule checks the pattern is valid
type AwsQuicksightUserInvalidNamespaceRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsQuicksightUserInvalidNamespaceRule returns new rule with default attributes
func NewAwsQuicksightUserInvalidNamespaceRule() *AwsQuicksightUserInvalidNamespaceRule {
	return &AwsQuicksightUserInvalidNamespaceRule{
		resourceType:  "aws_quicksight_user",
		attributeName: "namespace",
		max:           64,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9._-]*$`),
	}
}

// Name returns the rule name
func (r *AwsQuicksightUserInvalidNamespaceRule) Name() string {
	return "aws_quicksight_user_invalid_namespace"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightUserInvalidNamespaceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightUserInvalidNamespaceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightUserInvalidNamespaceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightUserInvalidNamespaceRule) Check(runner tflint.Runner) error {
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
					"namespace must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9._-]*$`),
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
