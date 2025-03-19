package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(0).
		SetBorders(true)

	// Title with emoji
	title := tview.NewTextView().
		SetText("[::b]📊 System Resource Monitor[::-]").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	// Text fields for system stats
	cpuText := tview.NewTextView().SetDynamicColors(true)
	memText := tview.NewTextView().SetDynamicColors(true)
	diskText := tview.NewTextView().SetDynamicColors(true)
	netText := tview.NewTextView().SetDynamicColors(true)

	// Add items to grid
	grid.AddItem(title, 0, 0, 1, 1, 0, 0, false).
		AddItem(cpuText, 1, 0, 1, 1, 0, 0, false).
		AddItem(memText, 2, 0, 1, 1, 0, 0, false).
		AddItem(diskText, 3, 0, 1, 1, 0, 0, false).
		AddItem(netText, 4, 0, 1, 1, 0, 0, false)

	// Function to generate ASCII bar graphs
	generateBarGraph := func(percentage float64, width int) string {
		fullBlocks := int((percentage / 100.0) * float64(width))
		return strings.Repeat("█", fullBlocks) + strings.Repeat("░", width-fullBlocks)
	}

	// Goroutine to update stats
	go func() {
		for {
			cpuUsage, _ := GetCPUUsage()
			_, usedMem, memUsage, _ := GetMemoryUsage()
			_, usedDisk, diskUsage, _ := GetDiskUsage()
			sentBytes, receivedBytes, _ := GetNetworkStats()

			app.QueueUpdateDraw(func() {
				// Use ANSI color codes instead of lipgloss
				cpuText.SetText(fmt.Sprintf("[red]CPU Usage: %.2f%% \n%s[::-]", cpuUsage, generateBarGraph(cpuUsage, 20)))
				memText.SetText(fmt.Sprintf("[green]Memory Usage: %.2f%% (%d MB)\n%s[::-]", memUsage, usedMem/1024/1024, generateBarGraph(memUsage, 20)))
				diskText.SetText(fmt.Sprintf("[blue]Disk Usage: %.2f%% (%d GB)\n%s[::-]", diskUsage, usedDisk/1024/1024/1024, generateBarGraph(diskUsage, 20)))
				netText.SetText(fmt.Sprintf("[yellow]Network: Sent: %d KB, Received: %d KB[::-]", sentBytes/1024, receivedBytes/1024))
			})
			time.Sleep(2 * time.Second)
		}
	}()

	// Run app
	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
