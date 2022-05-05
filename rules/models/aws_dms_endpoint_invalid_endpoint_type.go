// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDmsEndpointInvalidEndpointTypeRule checks the pattern is valid
type AwsDmsEndpointInvalidEndpointTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDmsEndpointInvalidEndpointTypeRule returns new rule with default attributes
func NewAwsDmsEndpointInvalidEndpointTypeRule() *AwsDmsEndpointInvalidEndpointTypeRule {
	return &AwsDmsEndpointInvalidEndpointTypeRule{
		resourceType:  "aws_dms_endpoint",
		attributeName: "endpoint_type",
		enum: []string{
			"source",
			"target",
		},
	}
}

// Name returns the rule name
func (r *AwsDmsEndpointInvalidEndpointTypeRule) Name() string {
	return "aws_dms_endpoint_invalid_endpoint_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDmsEndpointInvalidEndpointTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDmsEndpointInvalidEndpointTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDmsEndpointInvalidEndpointTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDmsEndpointInvalidEndpointTypeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as endpoint_type`, truncateLongMessage(val)),
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
