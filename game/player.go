package game

import (
	"big-battle-bits/bf"
	"image"
	"image/color"
)

type PlayerConfigs map[string]PlayerConfig

func (pcs *PlayerConfigs) getOrders() bf.Orders {
	orders := bf.Orders{}
	for name, config := range *pcs {
		orders[name] = *config.Priority
	}
	return orders
}

func (pcs *PlayerConfigs) getTeamColors() bf.TeamColors {
	tcs := bf.TeamColors{}
	for name, config := range *pcs {
		tcs[name] = bf.Team{*config.Color}
	}
	return tcs
}

type PlayerConfig struct {
	Priority *bf.Angle
	Start    *image.Point
	Color    *color.RGBA
}

func NewDefaultPlayerConfig() PlayerConfig {
	return PlayerConfig{
		&bf.Angle{0.0},
		&image.Point{0, 0},
		&bf.Black,
	}
}

func (pc *PlayerConfig) Merge(pc2 PlayerConfig) {
	if pc.Priority == nil {
		pc.Priority = pc2.Priority
	}
	if pc.Start == nil {
		pc.Start = pc2.Start
	}
	if pc.Color == nil {
		pc.Color = pc2.Color
	}
}

func (p *PlayerConfig) GetTeam() bf.Team {
	return bf.Team{
		*p.Color,
	}
}
