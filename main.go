package main

import (
	"fmt"
	"time"

    "github.com/HereComesTheMoon/TauTilting/tau"
)

func main() {
    t0 := time.Now()
    loc := "./data/7/Mod5/"
    tau.ComputeNextRow(loc, 8, 3)

    delta := time.Since(t0)
    fmt.Printf("Total process took %v seconds.\n", delta)
}

