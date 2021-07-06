package main

import (
    "flag"
    "fmt"
    "image"
    "os"
    "big-battle-bits/bf"
)

func main() {
    team := bf.Team{bf.Pink}
    flag.Var(&team, "team", "json representation of the color to add")
    input := flag.String("input", "", "Input map file to start from")
    output := flag.String("output", "", "File to write map to afterwards")
    x := flag.Int("x", 100, "x coordinate")
    y := flag.Int("y", 100, "y coordinate")
    flag.Parse()
    if *input == "" {
        fmt.Println("Input file is required")
        flag.Usage()
        os.Exit(1)
    }
    if *output == "" {
        fmt.Println("Output file is required")
        flag.Usage()
        os.Exit(1)
    }
    b, err := bf.ReadFromFile(*input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    b.Add(team, image.Point{*x, *y})
    err = bf.WriteTo(b, *output)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}
