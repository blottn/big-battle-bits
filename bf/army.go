package bf

import (
    "encoding/json"
	"image"
	"image/color"
)

var (
	Black  = color.RGBA{0, 0, 0, 255}
	Pink   = color.RGBA{244, 3, 252, 255}
	Orange = color.RGBA{244, 3, 252, 255}
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
    bytes, err :=  json.Marshal(team)
    if err != nil {
        return err.Error()
    }
    return string(bytes)
}

func (team *Team) Set(in string) error {
    return json.Unmarshal([]byte(in), team)
}


