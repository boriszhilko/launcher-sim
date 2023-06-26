package main

import (
	"flag"
	"fmt"
	"launcher-sim/pkg/service/simulator"
)

const (
	armLength            = 0.2   // 200mm
	armDiameter          = 0.015 // 15mm
	armDensity           = 2700  // kg/m^3 (6061 aluminum)
	ballDiameter         = 0.015 // 15mm
	ballDensity          = 7850  // kg/m^3 (1018 steel)
	ballDistanceFromAxis = 0.165 // 165mm
)

func main() {
	torque := flag.Float64("torque", 2, "the torque of the motor in NM")
	maxSpeed := flag.Float64("speed", 20, "the maximum speed of the motor in radians/second")
	releaseAngle := flag.Float64("release_angle", 45, "the angle at which the ball is released from the arm")
	startAngle := flag.Float64("start_angle", 0, "the starting angle of the arm")
	flag.Parse()

	arm := simulator.NewArm(armLength, armDiameter, armDensity, *torque, *maxSpeed)
	ball := simulator.NewBall(ballDiameter, ballDensity)

	sim := simulator.NewSimulator(arm, ball, ballDistanceFromAxis)
	res := sim.Simulate(*startAngle, *releaseAngle)

	fmt.Printf("The ball traveled %f meters.\n", res)
}
