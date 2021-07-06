package bf

import (
    "fmt"
    "image"
    "math"
)
// Commands
type Prioritiser interface {
    // Priority returns a value from 0->1 for each angle
    Priority(angle float64) float64
}

func getPriority(p Prioritiser, angle float64) (float64, error) {
    priority := p.Priority(angle)
    if priority < 0 || priority > 1 {
        return -1, fmt.Errorf("Invalid prioritisation at angle %v", angle)
    }
    return priority, nil
}

type Vector struct {
    x, y float64
}


func NewVector(x, y float64) Vector {
    return Vector{x,y}
}

func newVector(angle float64) Vector {
    return Vector{math.Cos(angle), math.Sin(angle)}
}

func (v1 Vector) dot(v2 Vector) float64 {
    return v1.x * v2.x + v1.y * v2.y
}

func (v Vector) magnitude() float64 {
    return math.Sqrt(v.x*v.x + v.y * v.y)
}

func (v1 Vector) scalarProjection(v2 Vector) float64 {
    return v1.dot(v2) / v2.magnitude()
}

// TODO fix this, it returns values even when 180degrees out of phase
func (v1 Vector) Priority(angle float64) float64 {
    // return scalarprojection (scaled)
    scalar := newVector(angle).scalarProjection(v1)
    if scalar <= 0 {
        return 0
    }
    return scalar / v1.magnitude()
}

func (v1 Vector) toPoint() image.Point {
    return image.Point{int(v1.x), int(v1.y)}
}


// TODO make this an image.Image
