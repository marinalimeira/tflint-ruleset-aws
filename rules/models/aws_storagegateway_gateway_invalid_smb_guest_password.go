// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayGatewayInvalidSmbGuestPasswordRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidSmbGuestPasswordRule() *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule {
	return &AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "smb_guest_password",
		max:           512,
		min:           6,
		pattern:       regexp.MustCompile(`^[ -~]+$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Name() string {
	return "aws_storagegateway_gateway_invalid_smb_guest_password"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Check(runner tflint.Runner) error {
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
					"smb_guest_password must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"smb_guest_password must be 6 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`smb_guest_password does not match valid pattern ^[ -~]+$`,
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
