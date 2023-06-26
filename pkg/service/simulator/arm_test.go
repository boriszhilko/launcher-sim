package simulator

import (
	"math"
	"testing"
)

func TestNewArm(t *testing.T) {
	// GIVEN
	length := 2.0
	diameter := 0.05
	density := 7850.0 // density of steel in kg/m3
	torque := 1.0
	maxSpeed := 1.0
	wantMass := 30.82
	wantMoment := 123.30

	// WHEN
	arm := NewArm(length, diameter, density, torque, maxSpeed)

	//THEN
	if math.Abs(arm.Mass-wantMass) > 0.01 {
		t.Errorf("want mass %v, got %v", wantMass, arm.Mass)
	}
	if math.Abs(arm.momentOfInertia-wantMoment) > 0.01 {
		t.Errorf("want moment of inertia %v, got %v", wantMoment, arm.momentOfInertia)
	}
}

func TestArm_Move(t *testing.T) {
	tests := []struct {
		name                      string
		length                    float64
		diameter                  float64
		density                   float64
		torque                    float64
		maxSpeed                  float64
		moveTime                  float64
		additionalMomentOfInertia float64
		wantAngle                 float64
		wantVelocity              float64
	}{
		{
			name:                      "Move should correctly change angle and angular velocity",
			length:                    2.0,
			diameter:                  0.05,
			density:                   7850,
			torque:                    1.0,
			maxSpeed:                  1.0,
			moveTime:                  0.5,
			additionalMomentOfInertia: 0.5,
			wantAngle:                 0.00202,
			wantVelocity:              0.00405,
		},
		// Other test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arm := NewArm(tt.length, tt.diameter, tt.density, tt.torque, tt.maxSpeed)
			arm.Move(tt.moveTime, tt.additionalMomentOfInertia)
			if math.Abs(arm.Angle-tt.wantAngle) > 0.00001 {
				t.Errorf("want angle %v, got %v", tt.wantAngle, arm.Angle)
			}
			if math.Abs(arm.AngularVel-tt.wantVelocity) > 0.0001 {
				t.Errorf("want angular velocity %v, got %v", tt.wantVelocity, arm.AngularVel)
			}
		})
	}
}
