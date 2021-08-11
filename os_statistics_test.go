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

func TestCPUStats(t *testing.T){
	c:=CPUStats()
	t.Logf("Cores: %v", c.Cores)
}

func TestDiskSpace(t *testing.T){
	s:=DiskSpace()
	t.Logf("space: %s", s)
}
