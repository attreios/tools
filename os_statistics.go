package tools

import (
	"runtime"
)

type Memory struct {
	Total       uint64
	Used        uint64
	PercentUsed int8
}

func MemoryStats() Memory {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return Memory{Total: uint64(m.Sys/1024/1024), Used: uint64(m.HeapAlloc/1024/1024), PercentUsed: int8(m.HeapAlloc / m.Sys * 100)}
}
