package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// GetCPUUsage fetches CPU utilization
func GetCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}

// GetMemoryUsage fetches RAM usage
func GetMemoryUsage() (uint64, uint64, float64, error) {
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, err
	}
	return memStats.Total, memStats.Used, memStats.UsedPercent, nil
}

// GetDiskUsage fetches disk usage
func GetDiskUsage() (uint64, uint64, float64, error) {
	diskStats, err := disk.Usage("/")
	if err != nil {
		return 0, 0, 0, err
	}
	return diskStats.Total, diskStats.Used, diskStats.UsedPercent, nil
}

// GetNetworkStats fetches network usage
func GetNetworkStats() (uint64, uint64, error) {
	netStats, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, err
	}
	return netStats[0].BytesSent, netStats[0].BytesRecv, nil
}

// PrintStats prints all system stats
func PrintStats() {
	cpuUsage, _ := GetCPUUsage()
	totalMem, usedMem, memUsage, _ := GetMemoryUsage()
	totalDisk, usedDisk, diskUsage, _ := GetDiskUsage()
	sentBytes, receivedBytes, _ := GetNetworkStats()

	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	fmt.Printf("Memory Usage: %.2f%% (%d/%d MB)\n", memUsage, usedMem/1024/1024, totalMem/1024/1024)
	fmt.Printf("Disk Usage: %.2f%% (%d/%d GB)\n", diskUsage, usedDisk/1024/1024/1024, totalDisk/1024/1024/1024)
	fmt.Printf("Network: Sent: %d KB, Received: %d KB\n", sentBytes/1024, receivedBytes/1024)
}
