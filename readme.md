# Launcher Simulator

Launcher Simulator is an application written in Go that simulates the physics of a launcher throwing a ball. The
application focuses on the calculation of physical parameters such as mass, moment of inertia, velocity, etc.

## Assumptions

The simulator assumes that the launching arm is made from 6061 aluminum alloy and the ball is crafted from 1018 steel.
The simulator does not consider air resistance in its calculations. As such, it provides a simplified representation of
the launcher's physics in a vacuum.

The arm length is measured at 200mm, and the arm diameter is measured at 15mm. The ball diameter is measured at 15mm.

## Prerequisites

Ensure you have Go installed on your system. This application is developed using Go version 1.19, but it should work
with more recent versions as well.

## Building the Application

In the root directory of the application, execute the following command:

```bash
go build -o launcher ./cmd/launcher-sim
```

This will build the application and create an executable named `launcher` in the current directory.

## Usage

Provide the application with the necessary parameters like arm length, arm diameter, arm density, ball diameter, ball
density, etc. The application will simulate the physics of the launcher and provide the calculated parameters.

Example:

```bash
./launcher -torque=2 -speed=20 -start_angle=0 -release_angle=45
```

## Testing

Run the tests using the following command:

```bash
go test ./...
```

## Future Enhancements

### Air Resistance

In the future, the Launcher Simulator could include air resistance in the calculation for a more realistic simulation.

### Adjustable Material Properties

To allow for more versatility, the simulator could be upgraded to enable the specification of different materials and
measurements for the arm and ball at runtime.

