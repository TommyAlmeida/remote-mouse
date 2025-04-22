package mouse

import (
	"fmt"
	"sync"

	"github.com/tommyalmeida/remote-mouse/mouse/native"
)

// ClickType represents a valid mouse click operation
type ClickType string

// Valid mouse click types
const (
	LeftClick   ClickType = "left"
	RightClick  ClickType = "right"
	DoubleClick ClickType = "double"
)

type MouseState int

const (
	Up   MouseState = iota
	Down
)

// Config holds the mouse controller configuration
type Config struct {
	// SpeedFactor multiplies delta movements (1.0 = normal, 0.5 = slower, 2.0 = faster)
	SpeedFactor float64
	// EnforceBounds prevents mouse from moving outside screen boundaries
	EnforceBounds bool
	// Silent disables logging to stdout
	Silent bool
	
	// Stabilization holds the stabilization options
	Stabilization *StabilizationOptions
	
	// screenWidth and screenHeight cache the screen dimensions
	screenWidth  int
	screenHeight int
	// mutex for thread-safety
	mu sync.RWMutex
}

// ScreenWidth returns the cached screen width
func (c *Config) ScreenWidth() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.screenWidth
}

// ScreenHeight returns the cached screen height
func (c *Config) ScreenHeight() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.screenHeight
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	w, h := native.GetScreenSize()
	return &Config{
		SpeedFactor:   1.0,
		EnforceBounds: true,
		Silent:        false,
		screenWidth:   w,
		screenHeight:  h,
		Stabilization: DefaultStabilizationOptions(),
	}
}

// Controller manages mouse interactions with configurable behavior
type Controller struct {
	config *Config
}

// NewController creates a new mouse controller with the given configuration
func NewController(config *Config) *Controller {
	if config == nil {
		config = DefaultConfig()
	}
	return &Controller{
		config: config,
	}
}

// Move moves the mouse cursor by the given delta amounts,
// applying speed factor and bounds checking according to configuration
func (c *Controller) Move(deltaX, deltaY int) {
	c.config.mu.RLock()
	defer c.config.mu.RUnlock()
	
	// Apply stabilization if enabled
	if c.config.Stabilization != nil {
		stabilizedX, stabilizedY, shouldMove := c.config.Stabilization.ProcessMovement(deltaX, deltaY)
		if !shouldMove {
			return
		}
		deltaX, deltaY = stabilizedX, stabilizedY
	}
	
	// Apply speed factor
	adjustedDeltaX := int(float64(deltaX) * c.config.SpeedFactor)
	adjustedDeltaY := int(float64(deltaY) * c.config.SpeedFactor)
	
	// Get current position
	currentX, currentY := native.GetMousePosition()
	
	// Calculate new position
	newX := currentX + adjustedDeltaX
	newY := currentY + adjustedDeltaY
	
	// Enforce screen boundaries if configured
	if c.config.EnforceBounds {
		if newX < 0 {
			newX = 0
		} else if newX >= c.config.screenWidth {
			newX = c.config.screenWidth - 1
		}
		
		if newY < 0 {
			newY = 0
		} else if newY >= c.config.screenHeight {
			newY = c.config.screenHeight - 1
		}
	}
	
	// Move the mouse
	native.MoveAbsolute(newX, newY)
	
	// Log the movement if not silent
	if !c.config.Silent {
		fmt.Printf("Moved mouse to: %d,%d (delta: %d,%d, adjusted: %d,%d)\n", 
			newX, newY, deltaX, deltaY, adjustedDeltaX, adjustedDeltaY)
	}
}

// SetLeftButton sets the left mouse button state
func (c *Controller) SetLeftButton(state MouseState) {
	c.config.mu.RLock()
	defer c.config.mu.RUnlock()
	
	if state == Down {
		native.LeftDown()
		if !c.config.Silent {
			fmt.Println("Left mouse button down")
		}
	} else {
		native.LeftUp()
		if !c.config.Silent {
			fmt.Println("Left mouse button up")
		}
	}
}

// SetRightButton sets the right mouse button state
func (c *Controller) SetRightButton(state MouseState) {
	c.config.mu.RLock()
	defer c.config.mu.RUnlock()
	
	if state == Down {
		native.RightDown()
		if !c.config.Silent {
			fmt.Println("Right mouse button down")
		}
	} else {
		native.RightUp()
		if !c.config.Silent {
			fmt.Println("Right mouse button up")
		}
	}
}

// Click performs a mouse click of the specified type
func (c *Controller) Click(clickType ClickType) error {
	c.config.mu.RLock()
	defer c.config.mu.RUnlock()
	
	switch clickType {
	case LeftClick:
		native.LeftClick()
		if !c.config.Silent {
			fmt.Println("Left mouse click")
		}
	case RightClick:
		native.RightClick()
		if !c.config.Silent {
			fmt.Println("Right mouse click")
		}
	case DoubleClick:
		native.DoubleClick()
		if !c.config.Silent {
			fmt.Println("Double mouse click")
		}
	default:
		err := fmt.Errorf("unknown click type: %s", clickType)
		if !c.config.Silent {
			fmt.Println(err)
		}
		return err
	}
	
	return nil
}

// UpdateConfig updates the controller's configuration
func (c *Controller) UpdateConfig(config *Config) {
	c.config.mu.Lock()
	defer c.config.mu.Unlock()
	
	if config != nil {
		// Only update fields that are explicitly set in the input config
		if config.SpeedFactor != 0 {
			c.config.SpeedFactor = config.SpeedFactor
		}
		
		// For boolean fields, we need to check if they were provided in the input
		// For now, we just apply them as is, assuming a non-nil config means these fields were set
		c.config.EnforceBounds = config.EnforceBounds
		c.config.Silent = config.Silent
		
		if config.Stabilization != nil {
			c.config.Stabilization = config.Stabilization
		}
		
		if c.config.screenWidth == 0 || c.config.screenHeight == 0 {
			c.config.screenWidth, c.config.screenHeight = native.GetScreenSize()
		}
	}
}

// UpdateStabilization updates the stabilization options for the controller
func (c *Controller) UpdateStabilization(options *StabilizationOptions) {
	c.config.mu.Lock()
	defer c.config.mu.Unlock()
	
	c.config.Stabilization = options
}

// For backward compatibility with existing code
var defaultController *Controller

// lazily initialize the default controller
func getDefaultController() *Controller {
	if defaultController == nil {
		defaultController = NewController(DefaultConfig())
	}
	return defaultController
}

// Move uses the default controller to move the mouse
func Move(deltaX, deltaY int) {
	getDefaultController().Move(deltaX, deltaY)
}

// Click uses the default controller to perform a click
func Click(clickType string) {
	getDefaultController().Click(ClickType(clickType))
}

// SetLeftButton uses the default controller to set the left mouse button state
func SetLeftButton(state MouseState) {
	getDefaultController().SetLeftButton(state)
}

// SetRightButton uses the default controller to set the right mouse button state
func SetRightButton(state MouseState) {
	getDefaultController().SetRightButton(state)
}

// UpdateStabilization uses the default controller to update the stabilization options
func UpdateStabilization(options *StabilizationOptions) {
	getDefaultController().UpdateStabilization(options)
} 