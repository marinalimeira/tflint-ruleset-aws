// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule checks the pattern is valid
type AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule() *AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule {
	return &AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule{
		resourceType:  "aws_secretsmanager_secret_rotation",
		attributeName: "rotation_lambda_arn",
		max:           2048,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule) Name() string {
	return "aws_secretsmanager_secret_rotation_invalid_rotation_lambda_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretRotationInvalidRotationLambdaArnRule) Check(runner tflint.Runner) error {
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
					"rotation_lambda_arn must be 2048 characters or less",
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
