package simulator

import (
	"math"
)

const (
	tolerance = 0.2 * math.Pi / 180 // 1 degrees in radians
	g         = 9.81                // Acceleration due to gravity
)

type Simulator struct {
	arm                  *Arm
	ball                 *Ball
	ballDistanceFromAxis float64
	ballMomentOfInertia  float64
}

func NewSimulator(arm *Arm, ball *Ball, ballDistanceFromAxis float64) *Simulator {
	ballMomentOfInertia := ball.Mass * math.Pow(ballDistanceFromAxis, 2)

	return &Simulator{
		arm:                  arm,
		ball:                 ball,
		ballDistanceFromAxis: ballDistanceFromAxis,
		ballMomentOfInertia:  ballMomentOfInertia,
	}
}

// Simulate simulates the movements of the arm and the ball from a given start angle to a release angle.
// It returns the horizontal distance covered by the ball after the launch.
func (s *Simulator) Simulate(startAngle, releaseAngle float64) float64 {
	s.setInitialAngle(startAngle)
	s.accelerateToReleaseAngle(releaseAngle)
	s.launchBall()

	// Calculate and return the horizontal distance covered
	return s.calculateHorizontalDistance()
}

func (s *Simulator) setInitialAngle(startAngle float64) {
	s.arm.Angle = startAngle * math.Pi / 180
}

func (s *Simulator) accelerateToReleaseAngle(releaseAngle float64) {
	const dt = 0.01
	for math.Abs(s.arm.Angle-releaseAngle*math.Pi/180) > tolerance {
		s.arm.Move(dt, s.ballMomentOfInertia)
	}
}

func (s *Simulator) launchBall() {
	s.ball.Velocity = s.ballDistanceFromAxis / s.arm.Length * s.arm.AngularVel
	s.ball.LaunchVel = s.ball.Velocity
	s.arm.AngularVel = 0
}

func (s *Simulator) calculateHorizontalDistance() float64 {
	referenceAngle := math.Pi/2 - s.arm.Angle
	Vx := math.Abs(s.ball.LaunchVel * math.Cos(referenceAngle))
	Vy := math.Abs(s.ball.LaunchVel * math.Sin(referenceAngle))

	fallTime := s.calculateFallTime()
	totalFlightTime := s.calculateTotalFlightTime(Vy, fallTime)

	return Vx * totalFlightTime
}

func (s *Simulator) calculateFallTime() float64 {
	var fallDistance float64
	if math.Abs(s.arm.Angle-math.Pi/2) < tolerance {
		fallDistance = s.ballDistanceFromAxis
	} else {
		fallDistance = math.Abs(s.ballDistanceFromAxis * math.Sin(s.arm.Angle))
	}

	return math.Sqrt(2 * fallDistance / g)
}

func (s *Simulator) calculateTotalFlightTime(Vy, fallTime float64) float64 {
	T := 2 * Vy / g
	if s.arm.Angle <= math.Pi/2 {
		T += fallTime
	} else {
		T = fallTime
	}

	return T
}
