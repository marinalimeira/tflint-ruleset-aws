// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderImagePipelineInvalidDescriptionRule checks the pattern is valid
type AwsImagebuilderImagePipelineInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsImagebuilderImagePipelineInvalidDescriptionRule returns new rule with default attributes
func NewAwsImagebuilderImagePipelineInvalidDescriptionRule() *AwsImagebuilderImagePipelineInvalidDescriptionRule {
	return &AwsImagebuilderImagePipelineInvalidDescriptionRule{
		resourceType:  "aws_imagebuilder_image_pipeline",
		attributeName: "description",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsImagebuilderImagePipelineInvalidDescriptionRule) Name() string {
	return "aws_imagebuilder_image_pipeline_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderImagePipelineInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderImagePipelineInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderImagePipelineInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderImagePipelineInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"description must be 1 characters or higher",
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
