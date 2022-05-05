// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLbTargetGroupInvalidProtocolRule checks the pattern is valid
type AwsLbTargetGroupInvalidProtocolRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLbTargetGroupInvalidProtocolRule returns new rule with default attributes
func NewAwsLbTargetGroupInvalidProtocolRule() *AwsLbTargetGroupInvalidProtocolRule {
	return &AwsLbTargetGroupInvalidProtocolRule{
		resourceType:  "aws_lb_target_group",
		attributeName: "protocol",
		enum: []string{
			"HTTP",
			"HTTPS",
			"TCP",
			"TLS",
			"UDP",
			"TCP_UDP",
			"GENEVE",
		},
	}
}

// Name returns the rule name
func (r *AwsLbTargetGroupInvalidProtocolRule) Name() string {
	return "aws_lb_target_group_invalid_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLbTargetGroupInvalidProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLbTargetGroupInvalidProtocolRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLbTargetGroupInvalidProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLbTargetGroupInvalidProtocolRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as protocol`, truncateLongMessage(val)),
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
