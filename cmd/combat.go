package main

import (
    "big-battle-bits/bf"
    "flag"
    "fmt"
    "os"
)

func main() {
    var orders bf.Orders
    iter := flag.Int("iter", 100, "Number of iterations to run")
    input := flag.String("input", "", "Input map file to start from")
    output := flag.String("output", "", "File to write map to afterwards")
    flag.Var(&orders, "orders", "Orders per team")
    flag.Parse()

    if *input == "" {
        flag.Usage()
        fmt.Println("Input file can't be empty")
        os.Exit(1)
    }

    if *output == "" {
        flag.Usage()
        fmt.Println("Output file can't be empty")
        os.Exit(1)
    }
    bg, err := bf.ReadFromFile(*input)
    if err != nil {
        fmt.Println(err.Error())
    }
    for i := 0; i < *iter; i++ {
        err = bf.StepCombat(bg, orders)
        if err !=nil {
            fmt.Println(err.Error())
            return
        }
    }
    bf.WriteTo(bg, *output)
}
