// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppconfigDeploymentStrategyInvalidReplicateToRule checks the pattern is valid
type AwsAppconfigDeploymentStrategyInvalidReplicateToRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppconfigDeploymentStrategyInvalidReplicateToRule returns new rule with default attributes
func NewAwsAppconfigDeploymentStrategyInvalidReplicateToRule() *AwsAppconfigDeploymentStrategyInvalidReplicateToRule {
	return &AwsAppconfigDeploymentStrategyInvalidReplicateToRule{
		resourceType:  "aws_appconfig_deployment_strategy",
		attributeName: "replicate_to",
		enum: []string{
			"NONE",
			"SSM_DOCUMENT",
		},
	}
}

// Name returns the rule name
func (r *AwsAppconfigDeploymentStrategyInvalidReplicateToRule) Name() string {
	return "aws_appconfig_deployment_strategy_invalid_replicate_to"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppconfigDeploymentStrategyInvalidReplicateToRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppconfigDeploymentStrategyInvalidReplicateToRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppconfigDeploymentStrategyInvalidReplicateToRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppconfigDeploymentStrategyInvalidReplicateToRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as replicate_to`, truncateLongMessage(val)),
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
