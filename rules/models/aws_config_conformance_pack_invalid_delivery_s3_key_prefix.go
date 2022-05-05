// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule checks the pattern is valid
type AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule returns new rule with default attributes
func NewAwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule() *AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule {
	return &AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule{
		resourceType:  "aws_config_conformance_pack",
		attributeName: "delivery_s3_key_prefix",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule) Name() string {
	return "aws_config_conformance_pack_invalid_delivery_s3_key_prefix"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigConformancePackInvalidDeliveryS3KeyPrefixRule) Check(runner tflint.Runner) error {
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
					"delivery_s3_key_prefix must be 1024 characters or less",
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
