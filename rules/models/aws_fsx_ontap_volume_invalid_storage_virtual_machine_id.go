// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule checks the pattern is valid
type AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule returns new rule with default attributes
func NewAwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule() *AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule {
	return &AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule{
		resourceType:  "aws_fsx_ontap_volume",
		attributeName: "storage_virtual_machine_id",
		max:           21,
		min:           21,
		pattern:       regexp.MustCompile(`^(svm-[0-9a-f]{17,})$`),
	}
}

// Name returns the rule name
func (r *AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule) Name() string {
	return "aws_fsx_ontap_volume_invalid_storage_virtual_machine_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxOntapVolumeInvalidStorageVirtualMachineIDRule) Check(runner tflint.Runner) error {
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
					"storage_virtual_machine_id must be 21 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"storage_virtual_machine_id must be 21 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(svm-[0-9a-f]{17,})$`),
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
