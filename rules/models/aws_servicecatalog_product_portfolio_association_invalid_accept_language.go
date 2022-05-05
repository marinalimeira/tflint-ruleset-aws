// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule checks the pattern is valid
type AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule returns new rule with default attributes
func NewAwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule() *AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule {
	return &AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule{
		resourceType:  "aws_servicecatalog_product_portfolio_association",
		attributeName: "accept_language",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule) Name() string {
	return "aws_servicecatalog_product_portfolio_association_invalid_accept_language"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogProductPortfolioAssociationInvalidAcceptLanguageRule) Check(runner tflint.Runner) error {
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
					"accept_language must be 100 characters or less",
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
