// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule checks the pattern is valid
type AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule returns new rule with default attributes
func NewAwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule() *AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule {
	return &AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule{
		resourceType:  "aws_fsx_openzfs_file_system",
		attributeName: "daily_automatic_backup_start_time",
		max:           5,
		min:           5,
		pattern:       regexp.MustCompile(`^([01]\d|2[0-3]):?([0-5]\d)$`),
	}
}

// Name returns the rule name
func (r *AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Name() string {
	return "aws_fsx_openzfs_file_system_invalid_daily_automatic_backup_start_time"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxOpenzfsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Check(runner tflint.Runner) error {
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
					"daily_automatic_backup_start_time must be 5 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"daily_automatic_backup_start_time must be 5 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([01]\d|2[0-3]):?([0-5]\d)$`),
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
