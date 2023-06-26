package simulator

import (
	"math"
	"testing"
)

func TestSimulator_Simulate(t *testing.T) {
	tests := []struct {
		name                 string
		startAngle           float64
		releaseAngle         float64
		wantDistance         float64
		ballDistanceFromAxis float64
	}{
		{
			name:                 "Release and start angle are the same",
			startAngle:           45.0,
			releaseAngle:         45.0,
			ballDistanceFromAxis: 0.165,
			wantDistance:         0,
		},
		{
			name:                 "Release angle is 30 degrees greater than start angle",
			startAngle:           30.0,
			releaseAngle:         60.0,
			ballDistanceFromAxis: 0.165,
			wantDistance:         38,
		},
		// Other test cases
	}

	arm := NewArm(0.2, 0.015, 2700, 2.0, 20.0)
	ball := NewBall(0.015, 7850)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulator := NewSimulator(arm, ball, tt.ballDistanceFromAxis)
			distance := simulator.Simulate(tt.startAngle, tt.releaseAngle)
			if math.Abs(distance-tt.wantDistance) > 1 {
				t.Errorf("want distance %v, got %v", tt.wantDistance, distance)
			}
		})
	}
}
