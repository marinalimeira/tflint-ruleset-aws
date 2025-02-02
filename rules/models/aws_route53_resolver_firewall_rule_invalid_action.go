// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverFirewallRuleInvalidActionRule checks the pattern is valid
type AwsRoute53ResolverFirewallRuleInvalidActionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsRoute53ResolverFirewallRuleInvalidActionRule returns new rule with default attributes
func NewAwsRoute53ResolverFirewallRuleInvalidActionRule() *AwsRoute53ResolverFirewallRuleInvalidActionRule {
	return &AwsRoute53ResolverFirewallRuleInvalidActionRule{
		resourceType:  "aws_route53_resolver_firewall_rule",
		attributeName: "action",
		enum: []string{
			"ALLOW",
			"BLOCK",
			"ALERT",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Name() string {
	return "aws_route53_resolver_firewall_rule_invalid_action"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as action`, truncateLongMessage(val)),
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
