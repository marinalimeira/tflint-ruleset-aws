// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule checks the pattern is valid
type AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaLayerVersionPermissionInvalidOrganizationIDRule returns new rule with default attributes
func NewAwsLambdaLayerVersionPermissionInvalidOrganizationIDRule() *AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule {
	return &AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule{
		resourceType:  "aws_lambda_layer_version_permission",
		attributeName: "organization_id",
		max:           34,
		pattern:       regexp.MustCompile(`^o-[a-z0-9]{10,32}$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule) Name() string {
	return "aws_lambda_layer_version_permission_invalid_organization_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaLayerVersionPermissionInvalidOrganizationIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"organization_id must be 34 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^o-[a-z0-9]{10,32}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}