// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule checks the pattern is valid
type AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule returns new rule with default attributes
func NewAwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule() *AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule {
	return &AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule{
		resourceType:  "aws_fsx_ontap_storage_virtual_machine",
		attributeName: "root_volume_security_style",
		enum: []string{
			"UNIX",
			"NTFS",
			"MIXED",
		},
	}
}

// Name returns the rule name
func (r *AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule) Name() string {
	return "aws_fsx_ontap_storage_virtual_machine_invalid_root_volume_security_style"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxOntapStorageVirtualMachineInvalidRootVolumeSecurityStyleRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as root_volume_security_style`, truncateLongMessage(val)),
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
