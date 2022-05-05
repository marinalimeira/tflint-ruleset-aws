// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyAppInvalidIAMServiceRoleArnRule checks the pattern is valid
type AwsAmplifyAppInvalidIAMServiceRoleArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsAmplifyAppInvalidIAMServiceRoleArnRule returns new rule with default attributes
func NewAwsAmplifyAppInvalidIAMServiceRoleArnRule() *AwsAmplifyAppInvalidIAMServiceRoleArnRule {
	return &AwsAmplifyAppInvalidIAMServiceRoleArnRule{
		resourceType:  "aws_amplify_app",
		attributeName: "iam_service_role_arn",
		max:           1000,
		pattern:       regexp.MustCompile(`^(?s).*$`),
	}
}

// Name returns the rule name
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Name() string {
	return "aws_amplify_app_invalid_iam_service_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Check(runner tflint.Runner) error {
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
					"iam_service_role_arn must be 1000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(?s).*$`),
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
