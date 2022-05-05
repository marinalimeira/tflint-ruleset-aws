// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWorklinkFleetInvalidAuditStreamArnRule checks the pattern is valid
type AwsWorklinkFleetInvalidAuditStreamArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsWorklinkFleetInvalidAuditStreamArnRule returns new rule with default attributes
func NewAwsWorklinkFleetInvalidAuditStreamArnRule() *AwsWorklinkFleetInvalidAuditStreamArnRule {
	return &AwsWorklinkFleetInvalidAuditStreamArnRule{
		resourceType:  "aws_worklink_fleet",
		attributeName: "audit_stream_arn",
		pattern:       regexp.MustCompile(`^arn:aws:kinesis:.+:[0-9]{12}:stream/AmazonWorkLink-.*$`),
	}
}

// Name returns the rule name
func (r *AwsWorklinkFleetInvalidAuditStreamArnRule) Name() string {
	return "aws_worklink_fleet_invalid_audit_stream_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorklinkFleetInvalidAuditStreamArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWorklinkFleetInvalidAuditStreamArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWorklinkFleetInvalidAuditStreamArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorklinkFleetInvalidAuditStreamArnRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws:kinesis:.+:[0-9]{12}:stream/AmazonWorkLink-.*$`),
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
