// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsOpsworksStackInvalidDefaultRootDeviceTypeRule checks the pattern is valid
type AwsOpsworksStackInvalidDefaultRootDeviceTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsOpsworksStackInvalidDefaultRootDeviceTypeRule returns new rule with default attributes
func NewAwsOpsworksStackInvalidDefaultRootDeviceTypeRule() *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule {
	return &AwsOpsworksStackInvalidDefaultRootDeviceTypeRule{
		resourceType:  "aws_opsworks_stack",
		attributeName: "default_root_device_type",
		enum: []string{
			"ebs",
			"instance-store",
		},
	}
}

// Name returns the rule name
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Name() string {
	return "aws_opsworks_stack_invalid_default_root_device_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOpsworksStackInvalidDefaultRootDeviceTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as default_root_device_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}