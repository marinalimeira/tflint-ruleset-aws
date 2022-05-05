// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWafv2WebACLAssociationInvalidResourceArnRule checks the pattern is valid
type AwsWafv2WebACLAssociationInvalidResourceArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWafv2WebACLAssociationInvalidResourceArnRule returns new rule with default attributes
func NewAwsWafv2WebACLAssociationInvalidResourceArnRule() *AwsWafv2WebACLAssociationInvalidResourceArnRule {
	return &AwsWafv2WebACLAssociationInvalidResourceArnRule{
		resourceType:  "aws_wafv2_web_acl_association",
		attributeName: "resource_arn",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsWafv2WebACLAssociationInvalidResourceArnRule) Name() string {
	return "aws_wafv2_web_acl_association_invalid_resource_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafv2WebACLAssociationInvalidResourceArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafv2WebACLAssociationInvalidResourceArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafv2WebACLAssociationInvalidResourceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafv2WebACLAssociationInvalidResourceArnRule) Check(runner tflint.Runner) error {
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
					"resource_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resource_arn must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
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
