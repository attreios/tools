package tools

import (
	"testing"
)

func TestMemoryStats(t *testing.T) {
	m:=MemoryStats()
	t.Logf("Total: %v",m.Total)
	t.Logf("Used: %v",m.Used)
	t.Logf("Percent: %v",m.PercentUsed)
}
