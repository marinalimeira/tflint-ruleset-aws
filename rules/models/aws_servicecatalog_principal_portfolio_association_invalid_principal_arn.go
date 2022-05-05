// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule checks the pattern is valid
type AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule returns new rule with default attributes
func NewAwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule() *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule {
	return &AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule{
		resourceType:  "aws_servicecatalog_principal_portfolio_association",
		attributeName: "principal_arn",
		max:           1000,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule) Name() string {
	return "aws_servicecatalog_principal_portfolio_association_invalid_principal_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalArnRule) Check(runner tflint.Runner) error {
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
					"principal_arn must be 1000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"principal_arn must be 1 characters or higher",
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
