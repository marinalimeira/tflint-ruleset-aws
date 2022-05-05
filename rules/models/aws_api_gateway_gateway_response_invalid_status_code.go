// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAPIGatewayGatewayResponseInvalidStatusCodeRule checks the pattern is valid
type AwsAPIGatewayGatewayResponseInvalidStatusCodeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAPIGatewayGatewayResponseInvalidStatusCodeRule returns new rule with default attributes
func NewAwsAPIGatewayGatewayResponseInvalidStatusCodeRule() *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule {
	return &AwsAPIGatewayGatewayResponseInvalidStatusCodeRule{
		resourceType:  "aws_api_gateway_gateway_response",
		attributeName: "status_code",
		pattern:       regexp.MustCompile(`^[1-5]\d\d$`),
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Name() string {
	return "aws_api_gateway_gateway_response_invalid_status_code"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayGatewayResponseInvalidStatusCodeRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[1-5]\d\d$`),
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
