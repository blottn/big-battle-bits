package bf

import (
	"fmt"
	"math"
)

var influencedSquares = []Vector{
	Vector{-1, 1},
	Vector{0, 1},
	Vector{1, 1},
	Vector{-1, 0},
	Vector{0, 0}, // TODO decide if this is right
	Vector{1, 0},
	Vector{-1, -1},
	Vector{0, -1},
	Vector{1, -1},
}

func StepCombat(bg *BattleGround, orders Orders, teamColors TeamColors) error {
	weights := map[Vector]map[Team]float64{}
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
			ang := math.Atan2(square.y, square.x)
			pVal, err := getPriority(p, ang)
			if err != nil {
				return err
			}
			target := Vector{square.x + float64(point.X), square.y + float64(point.Y)}
			// make sure there's something at that vector
			if _, ok := weights[target]; !ok {
				weights[target] = map[Team]float64{}
			}
			if _, ok := weights[target][team]; !ok {
				weights[target][team] = 0
			}
			weights[target][team] += pVal
		}
	}
	newArmies := Armies{}
	for v, teamWeights := range weights {
		winner := getHighestTeam(teamWeights)
		if winner == nil {
			continue
		}
		newArmies[v.toPoint()] = *winner
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
