// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDirectoryServiceDirectoryInvalidSizeRule checks the pattern is valid
type AwsDirectoryServiceDirectoryInvalidSizeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDirectoryServiceDirectoryInvalidSizeRule returns new rule with default attributes
func NewAwsDirectoryServiceDirectoryInvalidSizeRule() *AwsDirectoryServiceDirectoryInvalidSizeRule {
	return &AwsDirectoryServiceDirectoryInvalidSizeRule{
		resourceType:  "aws_directory_service_directory",
		attributeName: "size",
		enum: []string{
			"Small",
			"Large",
		},
	}
}

// Name returns the rule name
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Name() string {
	return "aws_directory_service_directory_invalid_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDirectoryServiceDirectoryInvalidSizeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as size`, truncateLongMessage(val)),
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
