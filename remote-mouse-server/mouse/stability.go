package mouse

import (
	"math"
	"time"
)

type StabilizationOptions struct {
	DeadZone       int     // Ignore movements smaller than this value (in pixels)
	SmoothingLevel float64 // 0.0-1.0: higher values mean more smoothing
	JiggleFilter   bool    // Enable anti-jiggle filtering
	AntiDrift      bool    // Enable anti-drift compensation
	
	lastX          int
	lastY          int
	lastMoveTime   time.Time
	velocityX      float64
	velocityY      float64
	histories      []PositionHistory
	historyPointer int
	historySize    int
}

type PositionHistory struct {
	X          int
	Y          int
	Time       time.Time
	VelocityX  float64
	VelocityY  float64
}

func DefaultStabilizationOptions() *StabilizationOptions {
	return &StabilizationOptions{
		DeadZone:       2,
		SmoothingLevel: 0.3,
		JiggleFilter:   true,
		AntiDrift:      true,
		historySize:    5,
		histories:      make([]PositionHistory, 5),
		lastMoveTime:   time.Now(),
	}
}

func (s *StabilizationOptions) ProcessMovement(deltaX, deltaY int) (int, int, bool) {
	now := time.Now()
	
	if s.DeadZone > 0 {
		if math.Abs(float64(deltaX)) < float64(s.DeadZone) {
			deltaX = 0
		}
		if math.Abs(float64(deltaY)) < float64(s.DeadZone) {
			deltaY = 0
		}
	}
	
	// If no movement after dead zone, nothing to do
	if deltaX == 0 && deltaY == 0 {
		return 0, 0, false
	}
	
	// Apply anti-jiggle filtering
	if s.JiggleFilter {
		timeElapsed := now.Sub(s.lastMoveTime).Seconds()
		
		s.histories[s.historyPointer] = PositionHistory{
			X:         deltaX,
			Y:         deltaY,
			Time:      now,
			VelocityX: float64(deltaX) / timeElapsed,
			VelocityY: float64(deltaY) / timeElapsed,
		}
		
		s.historyPointer = (s.historyPointer + 1) % s.historySize
		
		sumX, sumY := 0, 0

		for i := 0; i < s.historySize; i++ {
			sumX += s.histories[i].X
			sumY += s.histories[i].Y
		}
		

		absSum := math.Abs(float64(sumX)) + math.Abs(float64(sumY))
		absTotal := 0.0

		for i := 0; i < s.historySize; i++ {
			absTotal += math.Abs(float64(s.histories[i].X))
			absTotal += math.Abs(float64(s.histories[i].Y))
		}
		
		// If we have high movement but low net movement, it's likely jiggle
		if absTotal > 0 && absSum/absTotal < 0.3 && absTotal > float64(s.DeadZone*s.historySize) {
			if math.Abs(float64(sumX)) > math.Abs(float64(sumY)) {
				deltaY = 0
				deltaX = int(float64(sumX) / float64(s.historySize))
			} else {
				deltaX = 0
				deltaY = int(float64(sumY) / float64(s.historySize))
			}
		}
	}
	
	// Apply smoothing (if needed)
	if s.SmoothingLevel > 0 {

		timeDelta := now.Sub(s.lastMoveTime).Seconds()

		if timeDelta > 0 {
			currentVelocityX := float64(deltaX) / timeDelta
			currentVelocityY := float64(deltaY) / timeDelta
			
			s.velocityX = s.velocityX*(s.SmoothingLevel) + currentVelocityX*(1-s.SmoothingLevel)
			s.velocityY = s.velocityY*(s.SmoothingLevel) + currentVelocityY*(1-s.SmoothingLevel)
			
			deltaX = int(s.velocityX * timeDelta)
			deltaY = int(s.velocityY * timeDelta)
		}
	}
	
	if s.AntiDrift {
		timeSinceLastMove := now.Sub(s.lastMoveTime).Milliseconds()

		if timeSinceLastMove > 2000 {  // 2 seconds threshold
			s.lastMoveTime = now
			s.velocityX = 0
			s.velocityY = 0
			
			if math.Abs(float64(deltaX)) <= float64(s.DeadZone*2) && 
			   math.Abs(float64(deltaY)) <= float64(s.DeadZone*2) {
				return 0, 0, false
			}
		}
	}
	
	s.lastMoveTime = now
	s.lastX += deltaX
	s.lastY += deltaY
	
	return deltaX, deltaY, true
} 