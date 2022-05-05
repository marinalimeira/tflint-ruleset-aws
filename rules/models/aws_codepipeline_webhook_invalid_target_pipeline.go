// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodepipelineWebhookInvalidTargetPipelineRule checks the pattern is valid
type AwsCodepipelineWebhookInvalidTargetPipelineRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodepipelineWebhookInvalidTargetPipelineRule returns new rule with default attributes
func NewAwsCodepipelineWebhookInvalidTargetPipelineRule() *AwsCodepipelineWebhookInvalidTargetPipelineRule {
	return &AwsCodepipelineWebhookInvalidTargetPipelineRule{
		resourceType:  "aws_codepipeline_webhook",
		attributeName: "target_pipeline",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[A-Za-z0-9.@\-_]+$`),
	}
}

// Name returns the rule name
func (r *AwsCodepipelineWebhookInvalidTargetPipelineRule) Name() string {
	return "aws_codepipeline_webhook_invalid_target_pipeline"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodepipelineWebhookInvalidTargetPipelineRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodepipelineWebhookInvalidTargetPipelineRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodepipelineWebhookInvalidTargetPipelineRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodepipelineWebhookInvalidTargetPipelineRule) Check(runner tflint.Runner) error {
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
					"target_pipeline must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"target_pipeline must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9.@\-_]+$`),
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
