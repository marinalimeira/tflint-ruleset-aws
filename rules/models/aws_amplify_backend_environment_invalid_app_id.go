// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyBackendEnvironmentInvalidAppIDRule checks the pattern is valid
type AwsAmplifyBackendEnvironmentInvalidAppIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAmplifyBackendEnvironmentInvalidAppIDRule returns new rule with default attributes
func NewAwsAmplifyBackendEnvironmentInvalidAppIDRule() *AwsAmplifyBackendEnvironmentInvalidAppIDRule {
	return &AwsAmplifyBackendEnvironmentInvalidAppIDRule{
		resourceType:  "aws_amplify_backend_environment",
		attributeName: "app_id",
		max:           20,
		min:           1,
		pattern:       regexp.MustCompile(`^d[a-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsAmplifyBackendEnvironmentInvalidAppIDRule) Name() string {
	return "aws_amplify_backend_environment_invalid_app_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyBackendEnvironmentInvalidAppIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyBackendEnvironmentInvalidAppIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyBackendEnvironmentInvalidAppIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyBackendEnvironmentInvalidAppIDRule) Check(runner tflint.Runner) error {
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
					"app_id must be 20 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"app_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^d[a-z0-9]+$`),
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
