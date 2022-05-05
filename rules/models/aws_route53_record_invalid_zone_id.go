// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53RecordInvalidZoneIDRule checks the pattern is valid
type AwsRoute53RecordInvalidZoneIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53RecordInvalidZoneIDRule returns new rule with default attributes
func NewAwsRoute53RecordInvalidZoneIDRule() *AwsRoute53RecordInvalidZoneIDRule {
	return &AwsRoute53RecordInvalidZoneIDRule{
		resourceType:  "aws_route53_record",
		attributeName: "zone_id",
		max:           32,
	}
}

// Name returns the rule name
func (r *AwsRoute53RecordInvalidZoneIDRule) Name() string {
	return "aws_route53_record_invalid_zone_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53RecordInvalidZoneIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53RecordInvalidZoneIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53RecordInvalidZoneIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53RecordInvalidZoneIDRule) Check(runner tflint.Runner) error {
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
					"zone_id must be 32 characters or less",
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
