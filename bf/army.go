package bf

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
)

var (
	Black  = color.RGBA{0, 0, 0, 255}
	Pink   = color.RGBA{244, 3, 252, 255}
	Orange = color.RGBA{244, 3, 252, 255}
	Green  = color.RGBA{0, 255, 0, 255}
	Lilac  = color.RGBA{185, 66, 245, 255}
)

type Armies map[image.Point]Team

func (a Armies) overlay(b Armies) {
	for p, team := range b {
		a[p] = team
	}
}

type Team struct {
	Color color.RGBA
}

// flag.Value interface
func (team *Team) String() string {
	bytes, err := json.Marshal(team)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (team *Team) Set(in string) error {
	return json.Unmarshal([]byte(in), team)
}

func (t1 Team) Equals(t2 Team) bool {
	r1, g1, b1, a1 := t1.Color.RGBA()
	r2, g2, b2, a2 := t2.Color.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

type TeamColors map[string]Team

func (teamColors TeamColors) findName(t Team) (string, error) {
	for name, team := range teamColors {
		if team.Equals(t) {
			return name, nil
		}
	}
	return "", fmt.Errorf("Team %v is running rogue with no owning player", t)
}

// Represents the amount various armies control a point
type ArmyControl map[Team]float64

func (ac ArmyControl) GetWinner() *Team {
	highest := 0.5 // need at least 0.5 to be considered in control
	var winner *Team
	for team, val := range ac {
		if val > highest {
			winner = &team
			highest = val
		}
	}
	return winner
}

func (ac ArmyControl) AddInfluence(t Team, influence float64) {
	existing, ok := ac[t]
	if !ok {
		existing = 0.0
	}
	ac[t] = existing + influence
}

func TotalControl(t Team) ArmyControl {
	return ArmyControl(map[Team]float64{t: 1.0})
}

func (ac ArmyControl) String() string {
	s := ""
	for t, v := range ac {
		s += fmt.Sprintf("%v : %f\n", t, v)
	}
	return s
}

func (ac ArmyControl) Add(ac2 ArmyControl) {
	for team, v := range ac2 {
		ac.AddInfluence(team, v)
	}
}
