package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

// TestUnnecessaryStmt rule.
func TestUnnecessaryStmt(t *testing.T) {
	testRule(t, "unnecessary_stmt", &rule.UnnecessaryStmtRule{})
}
