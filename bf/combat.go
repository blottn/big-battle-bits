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
	image.Point{0, 0}, // TODO decide if this is right
	image.Point{1, 0},
	//	image.Point{-1, -1},
	image.Point{0, -1},
	//	image.Point{1, -1},
}

func StepCombat(bg *BattleGround, orders Orders, teamColors TeamColors) error {
	newArmies := map[image.Point]ArmyControl{}
	for point, control := range bg.armies {
		team := control.GetWinner()
		if team == nil {
			existingAc, ok := newArmies[point]
			if !ok {
				newArmies[point] = control
			} else {
				existingAc.Add(control)
				newArmies[point] = existingAc
			}
			// doesn't exert other influence therefore we just continue
			continue
		}
		// find correct team TODO make this not stupid...
		name, err := teamColors.findName(*team)
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
			if _, ok := newArmies[target]; !ok {
				newArmies[target] = ArmyControl{}
			}
			ang := math.Atan2(float64(square.Y), float64(square.X))
			newArmies[target].AddInfluence(*team, p.Priority(ang))
		}
	}
	// Remove oceans
	for p, _ := range newArmies {
		if bg.IsOcean(p) {
			delete(newArmies, p)
		}
	}

	for p, nc := range newArmies {
		// Check for residual control
		if ac, ok := bg.armies[p]; ok {
			ac.Add(nc)
			bg.armies[p] = ac
		} else {
			bg.armies[p] = nc
		}
	}
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
