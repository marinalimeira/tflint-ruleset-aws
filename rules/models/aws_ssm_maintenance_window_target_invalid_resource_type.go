// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule checks the pattern is valid
type AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmMaintenanceWindowTargetInvalidResourceTypeRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowTargetInvalidResourceTypeRule() *AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule {
	return &AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule{
		resourceType:  "aws_ssm_maintenance_window_target",
		attributeName: "resource_type",
		enum: []string{
			"INSTANCE",
			"RESOURCE_GROUP",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule) Name() string {
	return "aws_ssm_maintenance_window_target_invalid_resource_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowTargetInvalidResourceTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as resource_type`, truncateLongMessage(val)),
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
