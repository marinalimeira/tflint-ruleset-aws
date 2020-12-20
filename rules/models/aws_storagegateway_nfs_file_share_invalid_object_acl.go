// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayNfsFileShareInvalidObjectACLRule checks the pattern is valid
type AwsStoragegatewayNfsFileShareInvalidObjectACLRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsStoragegatewayNfsFileShareInvalidObjectACLRule returns new rule with default attributes
func NewAwsStoragegatewayNfsFileShareInvalidObjectACLRule() *AwsStoragegatewayNfsFileShareInvalidObjectACLRule {
	return &AwsStoragegatewayNfsFileShareInvalidObjectACLRule{
		resourceType:  "aws_storagegateway_nfs_file_share",
		attributeName: "object_acl",
		enum: []string{
			"private",
			"public-read",
			"public-read-write",
			"authenticated-read",
			"bucket-owner-read",
			"bucket-owner-full-control",
			"aws-exec-read",
		},
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Name() string {
	return "aws_storagegateway_nfs_file_share_invalid_object_acl"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as object_acl`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}