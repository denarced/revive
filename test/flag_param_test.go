package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

func TestFlagParam(t *testing.T) {
	testRule(t, "flag_param", &rule.FlagParamRule{})
}
