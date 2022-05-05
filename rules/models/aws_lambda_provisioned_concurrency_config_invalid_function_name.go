// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule checks the pattern is valid
type AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule returns new rule with default attributes
func NewAwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule() *AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule {
	return &AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule{
		resourceType:  "aws_lambda_provisioned_concurrency_config",
		attributeName: "function_name",
		max:           140,
		min:           1,
		pattern:       regexp.MustCompile(`^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule) Name() string {
	return "aws_lambda_provisioned_concurrency_config_invalid_function_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaProvisionedConcurrencyConfigInvalidFunctionNameRule) Check(runner tflint.Runner) error {
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
					"function_name must be 140 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"function_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?$`),
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
