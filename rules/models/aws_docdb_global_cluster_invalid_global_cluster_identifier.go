// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule checks the pattern is valid
type AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule returns new rule with default attributes
func NewAwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule() *AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule {
	return &AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule{
		resourceType:  "aws_docdb_global_cluster",
		attributeName: "global_cluster_identifier",
		max:           255,
		min:           1,
		pattern:       regexp.MustCompile(`^[A-Za-z][0-9A-Za-z-:._]*$`),
	}
}

// Name returns the rule name
func (r *AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule) Name() string {
	return "aws_docdb_global_cluster_invalid_global_cluster_identifier"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDocDBGlobalClusterInvalidGlobalClusterIdentifierRule) Check(runner tflint.Runner) error {
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
					"global_cluster_identifier must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"global_cluster_identifier must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z][0-9A-Za-z-:._]*$`),
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
