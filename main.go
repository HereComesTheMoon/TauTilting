package main

import (
	"fmt"
	"time"

    "github.com/HereComesTheMoon/TauTilting/tau"
)

func main() {
    t0 := time.Now()
    loc := "./data/2/Mod5/"

    for {
        err := tau.ComputeNextRow(loc, 8, 3)
        if err != nil {
            break
        }
    }

    delta := time.Since(t0)
    fmt.Printf("Total process took %v seconds.\n", delta)
}

