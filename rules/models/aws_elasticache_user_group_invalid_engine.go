// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElastiCacheUserGroupInvalidEngineRule checks the pattern is valid
type AwsElastiCacheUserGroupInvalidEngineRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsElastiCacheUserGroupInvalidEngineRule returns new rule with default attributes
func NewAwsElastiCacheUserGroupInvalidEngineRule() *AwsElastiCacheUserGroupInvalidEngineRule {
	return &AwsElastiCacheUserGroupInvalidEngineRule{
		resourceType:  "aws_elasticache_user_group",
		attributeName: "engine",
		pattern:       regexp.MustCompile(`^[a-zA-Z]*$`),
	}
}

// Name returns the rule name
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Name() string {
	return "aws_elasticache_user_group_invalid_engine"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastiCacheUserGroupInvalidEngineRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z]*$`),
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
