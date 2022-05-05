// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMServerCertificateInvalidPathRule checks the pattern is valid
type AwsIAMServerCertificateInvalidPathRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMServerCertificateInvalidPathRule returns new rule with default attributes
func NewAwsIAMServerCertificateInvalidPathRule() *AwsIAMServerCertificateInvalidPathRule {
	return &AwsIAMServerCertificateInvalidPathRule{
		resourceType:  "aws_iam_server_certificate",
		attributeName: "path",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^(\x{002F})|(\x{002F}[\x{0021}-\x{007F}]+\x{002F})$`),
	}
}

// Name returns the rule name
func (r *AwsIAMServerCertificateInvalidPathRule) Name() string {
	return "aws_iam_server_certificate_invalid_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMServerCertificateInvalidPathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMServerCertificateInvalidPathRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMServerCertificateInvalidPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMServerCertificateInvalidPathRule) Check(runner tflint.Runner) error {
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
					"path must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"path must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(\x{002F})|(\x{002F}[\x{0021}-\x{007F}]+\x{002F})$`),
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
