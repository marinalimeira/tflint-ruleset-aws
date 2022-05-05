// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDirectoryServiceDirectoryInvalidNameRule checks the pattern is valid
type AwsDirectoryServiceDirectoryInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsDirectoryServiceDirectoryInvalidNameRule returns new rule with default attributes
func NewAwsDirectoryServiceDirectoryInvalidNameRule() *AwsDirectoryServiceDirectoryInvalidNameRule {
	return &AwsDirectoryServiceDirectoryInvalidNameRule{
		resourceType:  "aws_directory_service_directory",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+$`),
	}
}

// Name returns the rule name
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Name() string {
	return "aws_directory_service_directory_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+$`),
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
