package main

import (
	"fmt"
	//"runtime"
	//"strconv"
	"time"

    "github.com/HereComesTheMoon/TauTilting/tau"
)

func main() {
    t0 := time.Now()

    //folder := "./Data/TEST/5/"
    //_ = folder;
    //tau.ComputeNextRow(folder, 8, 3)
    //folder := "./Data/7/Mod5/"
    //modules := tau.Get_indecomposables(folder, 3)
    //modules.SanityCheck()

    //results := tau.List_tau_tilting_modules(modules, 3, 3)
    //fmt.Printf("Results: %v\n", results)
    //trialRun2(1)
    //trialRun(1)
    //loc := "./Data/14/Mod2/"
    //data := make([][]int, 0) // iodata.ReadCsvToSlice(loc + "TAUTILTING.csv")
    //data = iodata.ReadCsvToSlice(loc + "TAUTILTING.csv")
    //iodata.WriteDiagonalsToCsv(loc + "DIAGONALS.csv", data)

    //for k := 0; k < 30; k++ {
        //tau.AddRowCountSTautiltingModulesStatic(loc)
        //data = iodata.ReadCsvToSlice(loc + "TAUTILTING.csv")
        //iodata.WriteDiagonalsToCsv(loc + "DIAGONALS.csv", data)
    //}

    alg := tau.ReadJsonDataToAlgebra("./tau/test/LinearAnMod5/data_5.json")
    fmt.Printf("%+v", alg)

    if !alg.SanityCheck() { panic(1) }

    println("HELLO")

    res := tau.Enumerate_tau_tilting_modules(alg.Indecomposables, 7, 3)

    fmt.Printf("%v", res)


    //tau.ReadJson2()


    delta := time.Since(t0)
    fmt.Printf("Total process took %v seconds.\n", delta)
}

func timeTrack(t0 time.Time, name string) {
    delta := time.Since(t0)
    fmt.Printf("Function %s had TOTAL TIME: %s\n", name, delta)
}

