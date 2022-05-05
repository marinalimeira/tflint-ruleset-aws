// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationSmbInvalidUserRule checks the pattern is valid
type AwsDatasyncLocationSmbInvalidUserRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationSmbInvalidUserRule returns new rule with default attributes
func NewAwsDatasyncLocationSmbInvalidUserRule() *AwsDatasyncLocationSmbInvalidUserRule {
	return &AwsDatasyncLocationSmbInvalidUserRule{
		resourceType:  "aws_datasync_location_smb",
		attributeName: "user",
		max:           104,
		pattern:       regexp.MustCompile(`^[^\x5B\x5D\\/:;|=,+*?]{1,104}$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationSmbInvalidUserRule) Name() string {
	return "aws_datasync_location_smb_invalid_user"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationSmbInvalidUserRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationSmbInvalidUserRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationSmbInvalidUserRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationSmbInvalidUserRule) Check(runner tflint.Runner) error {
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
					"user must be 104 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\x5B\x5D\\/:;|=,+*?]{1,104}$`),
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
