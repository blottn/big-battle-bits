package main

import "big-battle-bits/bf"
import "flag"
import "image"


func main() {
    team := bf.Team{bf.Pink}
    flag.Var(&team, "-team", "json representation of the color to add")
    x := flag.Int("-x", 250, "x coordinate")
    y := flag.Int("-y", 100, "y coordinate")
    adder := func(bg *bf.BattleGround) (error) {
        (*bg).Add(team, image.Point{*x,*y})
        return nil
    }
    bf.Mutate("bg-1.png", "mutated.png", adder)
    //mutator := func(bg BattleGround) (BattleGround, error) {
    //    bg.Add()
    //}
}
