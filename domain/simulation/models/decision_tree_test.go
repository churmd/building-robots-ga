package models_test

import (
	"context"
	"testing"

	"github.com/churmd/building-robots-ga/domain/simulation/models"
)

func TestTurnRight(t *testing.T) {
	testCases := []struct {
		name              string
		initialDirection  models.Direction
		expectedDirection models.Direction
	}{
		{
			name:              "Turning right from North",
			initialDirection:  models.North,
			expectedDirection: models.East,
		},
		{
			name:              "Turning right from East",
			initialDirection:  models.East,
			expectedDirection: models.South,
		},
		{
			name:              "Turning right from South",
			initialDirection:  models.South,
			expectedDirection: models.West,
		},
		{
			name:              "Turning right from West",
			initialDirection:  models.West,
			expectedDirection: models.North,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			robot := models.Robot{
				Direction: tc.initialDirection,
			}
			world := models.World{
				Robots: []models.Robot{robot},
			}

			updatedWorld, err := models.TurnRight(context.Background(), world, 0)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if updatedWorld.Robots[0].Direction != tc.expectedDirection {
				t.Errorf("expected direction %v, got %v", tc.expectedDirection, updatedWorld.Robots[0].Direction)
			}
		})
	}

	t.Run("when direction is unknown", func(t *testing.T) {
		robot := models.Robot{
			Direction: models.Direction(-1),
		}
		world := models.World{
			Robots: []models.Robot{robot},
		}

		_, err := models.TurnRight(context.Background(), world, 0)

		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})
}
