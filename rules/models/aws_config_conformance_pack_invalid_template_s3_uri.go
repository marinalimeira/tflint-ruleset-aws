// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigConformancePackInvalidTemplateS3URIRule checks the pattern is valid
type AwsConfigConformancePackInvalidTemplateS3URIRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsConfigConformancePackInvalidTemplateS3URIRule returns new rule with default attributes
func NewAwsConfigConformancePackInvalidTemplateS3URIRule() *AwsConfigConformancePackInvalidTemplateS3URIRule {
	return &AwsConfigConformancePackInvalidTemplateS3URIRule{
		resourceType:  "aws_config_conformance_pack",
		attributeName: "template_s3_uri",
		max:           1024,
		min:           1,
		pattern:       regexp.MustCompile(`^s3://.*$`),
	}
}

// Name returns the rule name
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Name() string {
	return "aws_config_conformance_pack_invalid_template_s3_uri"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigConformancePackInvalidTemplateS3URIRule) Check(runner tflint.Runner) error {
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
					"template_s3_uri must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"template_s3_uri must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^s3://.*$`),
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
