// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchMetricAlarmInvalidTreatMissingDataRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidTreatMissingDataRule() *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule {
	return &AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "treat_missing_data",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_treat_missing_data"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Check(runner tflint.Runner) error {
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
					"treat_missing_data must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"treat_missing_data must be 1 characters or higher",
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
