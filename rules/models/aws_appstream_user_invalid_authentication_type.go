// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamUserInvalidAuthenticationTypeRule checks the pattern is valid
type AwsAppstreamUserInvalidAuthenticationTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppstreamUserInvalidAuthenticationTypeRule returns new rule with default attributes
func NewAwsAppstreamUserInvalidAuthenticationTypeRule() *AwsAppstreamUserInvalidAuthenticationTypeRule {
	return &AwsAppstreamUserInvalidAuthenticationTypeRule{
		resourceType:  "aws_appstream_user",
		attributeName: "authentication_type",
		enum: []string{
			"API",
			"SAML",
			"USERPOOL",
			"AWS_AD",
		},
	}
}

// Name returns the rule name
func (r *AwsAppstreamUserInvalidAuthenticationTypeRule) Name() string {
	return "aws_appstream_user_invalid_authentication_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamUserInvalidAuthenticationTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamUserInvalidAuthenticationTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamUserInvalidAuthenticationTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamUserInvalidAuthenticationTypeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as authentication_type`, truncateLongMessage(val)),
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
