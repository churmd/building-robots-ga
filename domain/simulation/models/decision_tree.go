package models

import (
	"context"
	"fmt"
)

type RobotDecisionTree func(ctx context.Context, world World, robotIndex int) (World, error)

var turnRightMapper = map[Direction]Direction{
	North: East,
	East:  South,
	South: West,
	West:  North,
}

func TurnRight(ctx context.Context, world World, robotIndex int) (World, error) {
	currentDir := world.Robots[robotIndex].Direction
	updatedDir, ok := turnRightMapper[currentDir]
	if !ok {
		return World{}, fmt.Errorf("turn right received an unexpected direction %v", currentDir)
	}

	world.Robots[robotIndex].Direction = updatedDir
	return world, nil
}
