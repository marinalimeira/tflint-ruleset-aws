// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverRuleInvalidResolverEndpointIDRule checks the pattern is valid
type AwsRoute53ResolverRuleInvalidResolverEndpointIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53ResolverRuleInvalidResolverEndpointIDRule returns new rule with default attributes
func NewAwsRoute53ResolverRuleInvalidResolverEndpointIDRule() *AwsRoute53ResolverRuleInvalidResolverEndpointIDRule {
	return &AwsRoute53ResolverRuleInvalidResolverEndpointIDRule{
		resourceType:  "aws_route53_resolver_rule",
		attributeName: "resolver_endpoint_id",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverRuleInvalidResolverEndpointIDRule) Name() string {
	return "aws_route53_resolver_rule_invalid_resolver_endpoint_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverRuleInvalidResolverEndpointIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverRuleInvalidResolverEndpointIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverRuleInvalidResolverEndpointIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverRuleInvalidResolverEndpointIDRule) Check(runner tflint.Runner) error {
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
					"resolver_endpoint_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resolver_endpoint_id must be 1 characters or higher",
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
