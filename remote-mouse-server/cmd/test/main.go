package main

import (
	"fmt"
	"time"

	"github.com/tommyalmeida/remote-mouse/mouse"
)

func main() {
	fmt.Println("Testing native mouse implementation...")
	
	// Get screen size
	config := mouse.DefaultConfig()
	fmt.Printf("Screen size: %dx%d\n", config.ScreenWidth(), config.ScreenHeight())
	
	// Create controller with default config
	ctrl := mouse.NewController(nil)
	
	// Wait 2 seconds before starting
	fmt.Println("Starting in 2 seconds...")
	time.Sleep(2 * time.Second)
	
	// Move the mouse in a square pattern
	fmt.Println("Moving mouse...")
	ctrl.Move(100, 0)
	time.Sleep(500 * time.Millisecond)
	
	ctrl.Move(0, 100)
	time.Sleep(500 * time.Millisecond)
	
	ctrl.Move(-100, 0)
	time.Sleep(500 * time.Millisecond)
	
	ctrl.Move(0, -100)
	time.Sleep(500 * time.Millisecond)
	
	// Test clicks
	fmt.Println("Testing left click...")
	ctrl.Click(mouse.LeftClick)
	time.Sleep(1 * time.Second)
	
	fmt.Println("Testing right click...")
	ctrl.Click(mouse.RightClick)
	time.Sleep(1 * time.Second)
	
	fmt.Println("Testing double click...")
	ctrl.Click(mouse.DoubleClick)
	time.Sleep(1 * time.Second)
	
	fmt.Println("Test completed!")
} 