package simulator

import (
	"math"
)

// Arm represents a mechanical cylindrical arm connected to the motor.
type Arm struct {
	Angle           float64 // The current angle of the arm
	AngularVel      float64 // The angular velocity of the arm
	Length          float64
	Mass            float64
	maxSpeed        float64
	momentOfInertia float64
	torque          float64
}

func NewArm(length, diameter, density, torque, maxSpeed float64) *Arm {
	mass := calculateMass(length, diameter, density)
	momentOfInertia := calculateMomentOfInertia(length, mass)

	return &Arm{
		Length:          length,
		torque:          torque,
		maxSpeed:        maxSpeed,
		momentOfInertia: momentOfInertia,
		Mass:            mass,
	}
}

// Move advances the movement of the arm given a time increment and an additional moment of inertia.
func (a *Arm) Move(time float64, additionalMomentOfInertia float64) {
	totalMomentOfInertia := a.momentOfInertia + additionalMomentOfInertia
	if a.AngularVel < a.maxSpeed {
		a.AngularVel += a.torque / totalMomentOfInertia * time
	}
	a.Angle = math.Mod(a.Angle+a.AngularVel*time, 2*math.Pi)
}

func calculateMass(length, diameter, density float64) float64 {
	radius := diameter / 2
	volume := math.Pi * math.Pow(radius, 2) * length // Volume in cubic meters
	mass := density * volume                         // Mass in kg
	return mass
}

func calculateMomentOfInertia(length, mass float64) float64 {
	return mass * math.Pow(length, 2)
}
