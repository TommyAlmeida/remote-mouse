package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/tommyalmeida/remote-mouse/mouse"
)

type WebSocketConfig struct {
	MouseConfig *mouse.Config
	Verbose bool
}

func DefaultWebSocketConfig() *WebSocketConfig {
	return &WebSocketConfig{
		MouseConfig: mouse.DefaultConfig(),
		Verbose:     true,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	// activeConnections keeps track of all active WebSocket connections
	activeConnections int
	// connectionMutex protects the activeConnections counter
	connectionMutex sync.Mutex
)

func GetActiveConnectionCount() int {
	connectionMutex.Lock()
	
	defer connectionMutex.Unlock()

	return activeConnections
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	handler := NewWebSocketHandler(DefaultWebSocketConfig())
	handler.ServeHTTP(w, r)
}

type WebSocketHandler struct {
	config       *WebSocketConfig
	mouseCtrl    *mouse.Controller
}

func NewWebSocketHandler(config *WebSocketConfig) *WebSocketHandler {
	if config == nil {
		config = DefaultWebSocketConfig()
	}
	
	return &WebSocketHandler{
		config:    config,
		mouseCtrl: mouse.NewController(config.MouseConfig),
	}
}

func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	defer conn.Close()

	connectionMutex.Lock()
	activeConnections++
	connectionMutex.Unlock()
	
	defer func() {
		connectionMutex.Lock()
		activeConnections--
		connectionMutex.Unlock()
	}()

	if h.config.Verbose {
		fmt.Printf("New connection established from %s (active: %d)\n", 
			r.RemoteAddr, GetActiveConnectionCount())
	}

	for {
		// Read message from client
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, 
				websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("WebSocket error: %v\n", err)
			}
			break
		}
		
		messageStr := string(message)
		
		// Handle click commands
		if strings.HasPrefix(messageStr, "click:") {
			clickType := strings.TrimPrefix(messageStr, "click:")
			h.mouseCtrl.Click(mouse.ClickType(clickType))
			continue
		}
		
		// Handle left button commands
		if strings.HasPrefix(messageStr, "leftbutton:") {
			state := strings.TrimPrefix(messageStr, "leftbutton:")
			if state == "down" {
				h.mouseCtrl.SetLeftButton(mouse.Down)
			} else if state == "up" {
				h.mouseCtrl.SetLeftButton(mouse.Up)
			}
			continue
		}
		
		// Handle right button commands
		if strings.HasPrefix(messageStr, "rightbutton:") {
			state := strings.TrimPrefix(messageStr, "rightbutton:")
			if state == "down" {
				h.mouseCtrl.SetRightButton(mouse.Down)
			} else if state == "up" {
				h.mouseCtrl.SetRightButton(mouse.Up)
			}
			continue
		}
		
		// Handle configuration commands
		if strings.HasPrefix(messageStr, "config:") {
			h.handleConfigCommand(strings.TrimPrefix(messageStr, "config:"))
			continue
		}
		
		// Handle stabilization commands
		if strings.HasPrefix(messageStr, "stabilize:") {
			h.handleStabilizationCommand(strings.TrimPrefix(messageStr, "stabilize:"))
			continue
		}
		
		// Handle movement commands
		coords := strings.Split(messageStr, ",")
		if len(coords) != 2 {
			if h.config.Verbose {
				fmt.Println("Invalid message format. Expected 'deltaX,deltaY', 'click:type', 'leftbutton:state', 'rightbutton:state', 'config:...' or 'stabilize:...'")
			}
			continue
		}

		deltaX, err := strconv.Atoi(coords[0])
		if err != nil {
			if h.config.Verbose {
				fmt.Println("Invalid x delta coordinate:", err)
			}
			continue
		}

		deltaY, err := strconv.Atoi(coords[1])
		if err != nil {
			if h.config.Verbose {
				fmt.Println("Invalid y delta coordinate:", err)
			}
			continue
		}

		h.mouseCtrl.Move(deltaX, deltaY)
	}
	
	if h.config.Verbose {
		fmt.Printf("Connection closed from %s (active: %d)\n", 
			r.RemoteAddr, GetActiveConnectionCount())
	}
}

func (h *WebSocketHandler) handleConfigCommand(configCmd string) {
	parts := strings.Split(configCmd, "=")
	if len(parts) != 2 {
		if h.config.Verbose {
			fmt.Println("Invalid config command format. Expected 'key=value'")
		}
		return
	}
	
	key := parts[0]
	value := parts[1]
	
	switch key {
	case "speed":
		if speed, err := strconv.ParseFloat(value, 64); err == nil {
			newConfig := &mouse.Config{}
			newConfig.SpeedFactor = speed
			
			h.mouseCtrl.UpdateConfig(newConfig)
			if h.config.Verbose {
				fmt.Printf("Mouse speed set to %.2f\n", speed)
			}
		}
	case "bounds":
		if bounds, err := strconv.ParseBool(value); err == nil {
			newConfig := &mouse.Config{}
			
			newConfig.EnforceBounds = bounds
			
			h.mouseCtrl.UpdateConfig(newConfig)
			if h.config.Verbose {
				fmt.Printf("Enforce bounds set to %v\n", bounds)
			}
		}
	case "silent":
		if silent, err := strconv.ParseBool(value); err == nil {
			newConfig := &mouse.Config{}
			
			newConfig.Silent = silent
			
			h.mouseCtrl.UpdateConfig(newConfig)
			if h.config.Verbose {
				fmt.Printf("Silent mode set to %v\n", silent)
			}
		}
	default:
		if h.config.Verbose {
			fmt.Printf("Unknown config key: %s\n", key)
		}
	}
}

// handleStabilizationCommand processes stabilization commands from the client
func (h *WebSocketHandler) handleStabilizationCommand(cmd string) {
	parts := strings.Split(cmd, "=")
	if len(parts) != 2 {
		if h.config.Verbose {
			fmt.Println("Invalid stabilization command format. Expected 'key=value'")
		}
		return
	}
	
	key := parts[0]
	value := parts[1]
	
	stabOptions := h.config.MouseConfig.Stabilization
	if stabOptions == nil {
		stabOptions = mouse.DefaultStabilizationOptions()
	}
	
	switch key {
	case "deadzone":
		if val, err := strconv.Atoi(value); err == nil {
			stabOptions.DeadZone = val
			if h.config.Verbose {
				fmt.Printf("Dead zone set to %d\n", val)
			}
		}
	case "smoothing":
		if val, err := strconv.ParseFloat(value, 64); err == nil {
			stabOptions.SmoothingLevel = val
			if h.config.Verbose {
				fmt.Printf("Smoothing level set to %.2f\n", val)
			}
		}
	case "jiggle":
		if val, err := strconv.ParseBool(value); err == nil {
			stabOptions.JiggleFilter = val
			if h.config.Verbose {
				fmt.Printf("Jiggle filter set to %v\n", val)
			}
		}
	case "drift":
		if val, err := strconv.ParseBool(value); err == nil {
			stabOptions.AntiDrift = val
			if h.config.Verbose {
				fmt.Printf("Anti-drift set to %v\n", val)
			}
		}
	case "enable":
		if val, err := strconv.ParseBool(value); err == nil {
			if val {
				h.mouseCtrl.UpdateStabilization(stabOptions)
				if h.config.Verbose {
					fmt.Println("Stabilization enabled")
				}
			} else {
				h.mouseCtrl.UpdateStabilization(nil)
				if h.config.Verbose {
					fmt.Println("Stabilization disabled")
				}
			}
		}
	default:
		if h.config.Verbose {
			fmt.Printf("Unknown stabilization key: %s\n", key)
		}
	}
} 