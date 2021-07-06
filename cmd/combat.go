package main

import (
    "big-battle-bits/bf"
    "flag"
    "fmt"
    "image"
    "math"
)

func main() {
    iter := flag.Int("iter", 100, "Number of iterations to run")
    flag.Parse()

    // DUMMY DATA
    goDown := bf.NewVector(1,-1)
    goUp := bf.NewVector(1,1)
    team1 := bf.Team{bf.Pink}
    team2 := bf.Team{bf.Black}
    p := &map[bf.Team]bf.Prioritiser{
        team1: goUp,
        team2: goDown,
    }
    // END DUMMY DATA

    bg, err := bf.ReadFromFile("start.png")
    if err != nil {
        fmt.Println(err.Error())
    }
    bf.Mutate("start.png", "start.png", addTeam2)
    bg, err := bf.ReadFromFile("start.png")
    if err != nil {
        fmt.Println(err.Error())
    }
    for i := 0; i < *iter; i++ {
        err = bf.StepCombat(bg, *p)
        if err !=nil {
            fmt.Println(err.Error())
            return
        }
        if i > (*iter) / 2 {
            fmt.Println("redirecting")
            goDown = bf.NewVector(0, -1)
            *p = map[bf.Team]bf.Prioritiser{
                team1: goUp,
                team2: goDown,
            }
        }
    }
    bf.WriteTo(bg, "end.png")
}
