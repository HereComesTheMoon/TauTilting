package main

import (
    "errors"
	"fmt"
    "log"
    "os"
	"time"

    "github.com/HereComesTheMoon/TauTilting/tau"
)

func main() {
    t0 := time.Now()
    loc := "./data/CommutativeSquare/B/Mod"


    ComputeData(loc + "2", 7, 3, false)
    ComputeData(loc + "3", 7, 3, false)
    ComputeData(loc + "4", 7, 3, false)
    ComputeData(loc + "5", 7, 3, false)

    //ComputeData("./data/Special/SkewedTriangle/Mod2", 8, 3, false)

    delta := time.Since(t0)
    fmt.Printf("Total process took %v seconds.\n", delta)
}

func ComputeData(folder string, number_threads int, granularity int, overwrite bool) error {
    data_output := folder + "/TAUTILTING.csv" 

    _, err := os.Stat(folder)
    if errors.Is(err, os.ErrNotExist) {
        log.Printf("Folder %v does not exist.\n", folder)
        return err
    }

    _, err = os.Stat(data_output)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            fmt.Printf("No computations yet. Try to create file %v.\n", data_output)
            tau.InitialiseResultsFile(folder)
        } else {
            log.Fatalf("Oh no! %v", err) 
        }
    }


    for {
        _, err = compNextRow(folder, number_threads, granularity)
        if errors.Is(err, os.ErrNotExist) {
            fmt.Printf("Done computing in %v.", folder)
            break
        }
        if err != nil {
            panic(1)
        }
    }

    return nil
}

func compNextRow(folder string, number_threads int, granularity int) ([]int, error) {
    data_output := folder + "/TAUTILTING.csv" 

    already_computed_data := tau.ReadCsvToSlice(data_output)

    next_alg_data := fmt.Sprintf("%v/data_%d.json", folder, len(already_computed_data))

    alg, err := tau.ReadJsonDataToAlgebra(next_alg_data)
    if err != nil { return []int{}, err }

    next_modules := alg.Indecomposables

    t0 := time.Now().Truncate(time.Minute)
    fmt.Printf("\nComputing algebra %v in %v. It is now %v:%v.\n", len(already_computed_data), folder, t0.Hour(), t0.Minute())

    next_row := tau.Enumerate_tau_tilting_modules(next_modules, number_threads, granularity)

    already_computed_data = append(already_computed_data, next_row)
    already_computed_data = tau.PadJaggedArray(already_computed_data, len(next_row))

    tau.WriteSliceToCsv(data_output, already_computed_data)
    return next_row, nil
}
