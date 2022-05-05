// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule checks the pattern is valid
type AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule returns new rule with default attributes
func NewAwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule() *AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule {
	return &AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule{
		resourceType:  "aws_route53recoverycontrolconfig_safety_rule",
		attributeName: "control_panel_arn",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^[A-Za-z0-9:\/_-]*$`),
	}
}

// Name returns the rule name
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule) Name() string {
	return "aws_route53recoverycontrolconfig_safety_rule_invalid_control_panel_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53recoverycontrolconfigSafetyRuleInvalidControlPanelArnRule) Check(runner tflint.Runner) error {
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
					"control_panel_arn must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"control_panel_arn must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9:\/_-]*$`),
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
