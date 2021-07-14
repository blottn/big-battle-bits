package bf

import (
	"fmt"
	"image"
	"math"
)

var influencedSquares = []image.Point{
	//  image.Point{-1, 1},
	image.Point{0, 1},
	//	image.Point{1, 1},
	image.Point{-1, 0},
	//	image.Point{0, 0}, // TODO decide if this is right
	image.Point{1, 0},
	//	image.Point{-1, -1},
	image.Point{0, -1},
	//	image.Point{1, -1},
}

func StepCombat(bg *BattleGround, orders Orders, teamColors TeamColors) error {
	weights := map[image.Point]map[Team]float64{}
	for point, team := range bg.armies {
		// find correct team
		name, err := teamColors.findName(team)
		if err != nil {
			return err
		}
		p, ok := orders[name]
		if !ok {
			return fmt.Errorf("Missing orders for team %v", team)
		}
		for _, square := range influencedSquares {
			target := point.Add(square)

			// make sure we are already tracking
			if _, ok := weights[target]; !ok {
				weights[target] = map[Team]float64{}
			}
			// make sure we are already tracking that team
			if _, ok := weights[target][team]; !ok {
				weights[target][team] = 0
			}
			weights[target][team] += p.Priority(math.Atan2(float64(square.Y), float64(square.X)))

		}
	}
	newArmies := Armies{}
	for v, teamWeights := range weights {
		winner := getHighestTeam(teamWeights)
		if winner == nil {
			continue
		}
		newArmies[v] = *winner
	}

	for p, _ := range newArmies {
		if bg.IsOcean(p) {
			delete(newArmies, p)
		}
	}
	bg.armies.overlay(newArmies)
	return nil
}

func getHighestTeam(weights map[Team]float64) *Team {
	if weights == nil {
		return nil
	}
	var winner Team
	var winningWeight = -1.0

	for team, val := range weights {
		if val > winningWeight {
			winner = team
			winningWeight = val
		}
	}

	// Disallow freebies
	if winningWeight == 0 {
		return nil
	}
	return &winner
}
