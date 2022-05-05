// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCurReportDefinitionInvalidS3PrefixRule checks the pattern is valid
type AwsCurReportDefinitionInvalidS3PrefixRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsCurReportDefinitionInvalidS3PrefixRule returns new rule with default attributes
func NewAwsCurReportDefinitionInvalidS3PrefixRule() *AwsCurReportDefinitionInvalidS3PrefixRule {
	return &AwsCurReportDefinitionInvalidS3PrefixRule{
		resourceType:  "aws_cur_report_definition",
		attributeName: "s3_prefix",
		max:           256,
		pattern:       regexp.MustCompile(`^[0-9A-Za-z!\-_.*\'()/]*$`),
	}
}

// Name returns the rule name
func (r *AwsCurReportDefinitionInvalidS3PrefixRule) Name() string {
	return "aws_cur_report_definition_invalid_s3_prefix"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCurReportDefinitionInvalidS3PrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCurReportDefinitionInvalidS3PrefixRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCurReportDefinitionInvalidS3PrefixRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCurReportDefinitionInvalidS3PrefixRule) Check(runner tflint.Runner) error {
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
					"s3_prefix must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9A-Za-z!\-_.*\'()/]*$`),
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
