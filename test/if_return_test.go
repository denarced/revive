package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

// TestIfReturn rule.
func TestIfReturn(t *testing.T) {
	testRule(t, "if_return", &rule.IfReturnRule{})
}
