// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogProductInvalidTypeRule checks the pattern is valid
type AwsServicecatalogProductInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	enum          []string
}

// NewAwsServicecatalogProductInvalidTypeRule returns new rule with default attributes
func NewAwsServicecatalogProductInvalidTypeRule() *AwsServicecatalogProductInvalidTypeRule {
	return &AwsServicecatalogProductInvalidTypeRule{
		resourceType:  "aws_servicecatalog_product",
		attributeName: "type",
		max:           8191,
		enum: []string{
			"CLOUD_FORMATION_TEMPLATE",
			"MARKETPLACE",
		},
	}
}

// Name returns the rule name
func (r *AwsServicecatalogProductInvalidTypeRule) Name() string {
	return "aws_servicecatalog_product_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogProductInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogProductInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogProductInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogProductInvalidTypeRule) Check(runner tflint.Runner) error {
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
					"type must be 8191 characters or less",
					attribute.Expr.Range(),
				)
			}
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
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
