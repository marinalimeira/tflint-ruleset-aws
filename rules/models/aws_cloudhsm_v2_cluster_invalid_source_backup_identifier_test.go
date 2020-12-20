// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCloudhsmV2ClusterInvalidSourceBackupIdentifierRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudhsm_v2_cluster" "foo" {
	source_backup_identifier = "rtq2dwi2gq6"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCloudhsmV2ClusterInvalidSourceBackupIdentifierRule(),
					Message: `"rtq2dwi2gq6" does not match valid pattern ^backup-[2-7a-zA-Z]{11,16}$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudhsm_v2_cluster" "foo" {
	source_backup_identifier = "backup-rtq2dwi2gq6"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCloudhsmV2ClusterInvalidSourceBackupIdentifierRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}