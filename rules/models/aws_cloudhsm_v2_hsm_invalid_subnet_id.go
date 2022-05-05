// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudhsmV2HsmInvalidSubnetIDRule checks the pattern is valid
type AwsCloudhsmV2HsmInvalidSubnetIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudhsmV2HsmInvalidSubnetIDRule returns new rule with default attributes
func NewAwsCloudhsmV2HsmInvalidSubnetIDRule() *AwsCloudhsmV2HsmInvalidSubnetIDRule {
	return &AwsCloudhsmV2HsmInvalidSubnetIDRule{
		resourceType:  "aws_cloudhsm_v2_hsm",
		attributeName: "subnet_id",
		pattern:       regexp.MustCompile(`^subnet-[0-9a-fA-F]{8,17}$`),
	}
}

// Name returns the rule name
func (r *AwsCloudhsmV2HsmInvalidSubnetIDRule) Name() string {
	return "aws_cloudhsm_v2_hsm_invalid_subnet_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudhsmV2HsmInvalidSubnetIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudhsmV2HsmInvalidSubnetIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudhsmV2HsmInvalidSubnetIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudhsmV2HsmInvalidSubnetIDRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^subnet-[0-9a-fA-F]{8,17}$`),
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
