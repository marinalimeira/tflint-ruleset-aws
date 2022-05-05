// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecretsmanagerSecretPolicyInvalidPolicyRule checks the pattern is valid
type AwsSecretsmanagerSecretPolicyInvalidPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSecretsmanagerSecretPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretPolicyInvalidPolicyRule() *AwsSecretsmanagerSecretPolicyInvalidPolicyRule {
	return &AwsSecretsmanagerSecretPolicyInvalidPolicyRule{
		resourceType:  "aws_secretsmanager_secret_policy",
		attributeName: "policy",
		max:           20480,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Name() string {
	return "aws_secretsmanager_secret_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretPolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
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
					"policy must be 20480 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy must be 1 characters or higher",
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
