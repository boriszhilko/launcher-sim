package simulator

import "math"

// Ball represents a physical ball used in the simulation.
type Ball struct {
	Diameter  float64
	Density   float64
	Mass      float64
	Velocity  float64
	LaunchVel float64
}

func NewBall(diameter, density float64) *Ball {
	mass := calculateBallMass(diameter, density)

	ball := &Ball{
		Diameter: diameter,
		Density:  density,
		Mass:     mass,
	}
	return ball
}

func calculateBallMass(diameter, density float64) float64 {
	radius := diameter / 2
	volume := 4 / 3 * math.Pi * math.Pow(radius, 3) // Volume in cubic meters
	return density * volume                         // Mass in kg
}
