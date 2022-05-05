// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMServerCertificateInvalidCertificateChainRule checks the pattern is valid
type AwsIAMServerCertificateInvalidCertificateChainRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMServerCertificateInvalidCertificateChainRule returns new rule with default attributes
func NewAwsIAMServerCertificateInvalidCertificateChainRule() *AwsIAMServerCertificateInvalidCertificateChainRule {
	return &AwsIAMServerCertificateInvalidCertificateChainRule{
		resourceType:  "aws_iam_server_certificate",
		attributeName: "certificate_chain",
		max:           2097152,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMServerCertificateInvalidCertificateChainRule) Name() string {
	return "aws_iam_server_certificate_invalid_certificate_chain"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMServerCertificateInvalidCertificateChainRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMServerCertificateInvalidCertificateChainRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMServerCertificateInvalidCertificateChainRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMServerCertificateInvalidCertificateChainRule) Check(runner tflint.Runner) error {
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
					"certificate_chain must be 2097152 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"certificate_chain must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
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
