// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSchemasRegistryInvalidDescriptionRule checks the pattern is valid
type AwsSchemasRegistryInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsSchemasRegistryInvalidDescriptionRule returns new rule with default attributes
func NewAwsSchemasRegistryInvalidDescriptionRule() *AwsSchemasRegistryInvalidDescriptionRule {
	return &AwsSchemasRegistryInvalidDescriptionRule{
		resourceType:  "aws_schemas_registry",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsSchemasRegistryInvalidDescriptionRule) Name() string {
	return "aws_schemas_registry_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSchemasRegistryInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSchemasRegistryInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSchemasRegistryInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSchemasRegistryInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 256 characters or less",
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
