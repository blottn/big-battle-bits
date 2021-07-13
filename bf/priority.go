package bf

import (
	"encoding/json"
	"math"
)

type Orders map[string]Angle

// flag.Value interface
func (orders Orders) String() string {
	bytes, err := json.Marshal(orders)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (orders Orders) Set(in string) error {
	return json.Unmarshal([]byte(in), &orders)
}

type Angle struct {
	V float64
}

func (a Angle) toRad() float64 {
	return (a.V / 360) * math.Pi * 2.0
}

func (a1 Angle) Priority(a2 float64) float64 {
	diff := math.Abs(a1.toRad() - a2)
	weight := math.Cos(diff)
	if weight < 0 {
		return 0
	}
	return weight
}

// TODO allow rendering of player angle as an image
