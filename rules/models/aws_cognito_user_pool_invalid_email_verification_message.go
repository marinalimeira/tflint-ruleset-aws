// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCognitoUserPoolInvalidEmailVerificationMessageRule checks the pattern is valid
type AwsCognitoUserPoolInvalidEmailVerificationMessageRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule returns new rule with default attributes
func NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule() *AwsCognitoUserPoolInvalidEmailVerificationMessageRule {
	return &AwsCognitoUserPoolInvalidEmailVerificationMessageRule{
		resourceType:  "aws_cognito_user_pool",
		attributeName: "email_verification_message",
		max:           20000,
		min:           6,
		pattern:       regexp.MustCompile(`^[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{####\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Name() string {
	return "aws_cognito_user_pool_invalid_email_verification_message"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Check(runner tflint.Runner) error {
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
					"email_verification_message must be 20000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"email_verification_message must be 6 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{####\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*$`),
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
