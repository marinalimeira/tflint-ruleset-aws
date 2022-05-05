// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecretsmanagerSecretInvalidKmsKeyIDRule checks the pattern is valid
type AwsSecretsmanagerSecretInvalidKmsKeyIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsSecretsmanagerSecretInvalidKmsKeyIDRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretInvalidKmsKeyIDRule() *AwsSecretsmanagerSecretInvalidKmsKeyIDRule {
	return &AwsSecretsmanagerSecretInvalidKmsKeyIDRule{
		resourceType:  "aws_secretsmanager_secret",
		attributeName: "kms_key_id",
		max:           2048,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretInvalidKmsKeyIDRule) Name() string {
	return "aws_secretsmanager_secret_invalid_kms_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretInvalidKmsKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretInvalidKmsKeyIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretInvalidKmsKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretInvalidKmsKeyIDRule) Check(runner tflint.Runner) error {
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
					"kms_key_id must be 2048 characters or less",
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
