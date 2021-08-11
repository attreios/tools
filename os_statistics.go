package tools

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/sys/unix"
)

type Memory struct {
	Total       uint64
	Used        uint64
	PercentUsed int8
}

type CPU struct {
	Cores uint8
}

func MemoryStats() Memory {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return Memory{
		Total:       uint64(m.Sys / 1024 / 1024),
		Used:        uint64(m.HeapAlloc / 1024 / 1024),
		PercentUsed: int8(m.HeapAlloc / m.Sys * 100),
	}
}

func CPUStats() CPU {
	return CPU{
		Cores: uint8(runtime.NumCPU()),
	}
}

func DiskStats() uint64{
	var stat unix.Statfs_t

	wd, err := os.Getwd()

	unix.Statfs(wd, &stat)

	// Available blocks * size per block = available space in bytes
	return stat.Bavail * uint64(stat.Bsize)
}