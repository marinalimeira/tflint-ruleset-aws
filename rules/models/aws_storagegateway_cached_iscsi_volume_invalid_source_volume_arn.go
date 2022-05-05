// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule checks the pattern is valid
type AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule returns new rule with default attributes
func NewAwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule() *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule {
	return &AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule{
		resourceType:  "aws_storagegateway_cached_iscsi_volume",
		attributeName: "source_volume_arn",
		max:           500,
		min:           50,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Name() string {
	return "aws_storagegateway_cached_iscsi_volume_invalid_source_volume_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule) Check(runner tflint.Runner) error {
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
					"source_volume_arn must be 500 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"source_volume_arn must be 50 characters or higher",
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
