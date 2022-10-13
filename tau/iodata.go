package tau

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

/// Read input data.json for an algebra
func ReadJsonDataToAlgebra(location string) Algebra {
    f, err := os.Open(location)

    if err != nil {
        log.Fatal(err)
    }

    dec := json.NewDecoder(f)

    alg := Algebra{}

    dec.Decode(&alg)

    return alg
}


// Reads Csv file with ints to slice. Warning: Doesn't handle rows ending with comma or comma+whitespace
func ReadCsvToSlice(location string) [][]int {
    f, err := os.Open(location)

    if err != nil {
        log.Fatal(err)
    }

    arr := make([][]int, 0)
    r := csv.NewReader(f)

    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        row := make([]int, len(record))
        
        for i, v := range(record) {
            v := strings.TrimSpace(v)
            value, err := strconv.Atoi(strings.TrimSpace(v))
            if err != nil {
                log.Fatal(err)
            }
            row[i] = value
        }
        arr = append(arr, row)
    }
    return arr
}

// Given some 2-slice of ints, write conveniently formatted human readable .csv. Delimiter = ','. Add whitespaces so that all columns have the same width.
func WriteSliceToCsv(location string, data [][]int) {
    file, err := os.Create(location)
    if err != nil {
        fmt.Printf("Unable to find the file %v. Redirecting to stdout instead.\n", location)
        file = os.Stdout
    }
    defer file.Close()

    // Find biggest number, then take log10 + 1 to get amount of necessary padding
    max := 1
    for _, row := range data {
        for _, val := range row {
            if val > max {
                max = val
            }
        }
    }
    max = int(math.Log10(float64(max))) + 1 // All columns evenly-spaced, wide enough to fit all given integers.


    for _, row := range data {
        row_strings := make([]string, len(row))
        for i, v := range row {
            row_strings[i] = fmt.Sprintf("%*d", max, v)
        }
        _, err := fmt.Fprint(file, strings.Join(row_strings, ",") + "\n")
        if err != nil {
            log.Fatalln("error writing record to csv:", err)
        }
    }
}


// pad a jagged slice of ints with zeros until each row fulfills len(row) = length. 
func PadJaggedArray(arr [][]int, length int) [][]int {
    for _, v := range arr {
        if len(v) > length { 
            log.Println("Some rows are longer than the given padding size.") 
            break
        }

    }
    
    for i := range arr {
        arr[i] = append(arr[i], make([]int, length - len(arr[i]))...)
    }

    return arr
}


func initialiseResultsFile(folder string) {
    i := 0
    for ; i < 31; i++ {
        next_file_path := fmt.Sprint(folder, "/data_", i, ".json")
        fmt.Print(next_file_path)
        _, err := os.Stat(next_file_path)

        if errors.Is(err, os.ErrNotExist) { continue }
        if err != nil { log.Fatalf("Oh no! %w", err) }

        break
    }
    if i == 31 {
        log.Fatalf("No data files data_1.json ... data_30.json. Wrong folder?")
    }

    // Else, create new TAUTILTING.csv file
    new_tau := make([][]int, 0, i)
    for j := 0; j < i; j++ {
        new_tau = append(new_tau, make([]int, i))
    }
    WriteSliceToCsv(folder + "/TAUTILTING.csv", new_tau)
}

// Given a folder, read in TAUTILTING.csv. Count the rows, then use ReadJsonDataToAlgebra() to read in data_k.json (from the same folder), and compute the next row. Then append the row to the file, and save it.
func ComputeNextRow(folder string, number_threads int, granularity int) {
    tau_out := folder + "/TAUTILTING.csv" 

    _, err := os.Stat(folder)
    if errors.Is(err, os.ErrNotExist) {
        log.Fatalf("Folder %v does not exist.\n", folder)
    }
    _, err = os.Stat(tau_out)
    if errors.Is(err, os.ErrNotExist) {
        fmt.Printf("No computations yet. Try to create file %v.", tau_out)
        initialiseResultsFile(folder)
    }

    // Else, actually compute the next row.
    already_computed_data := ReadCsvToSlice(tau_out)
    next_modules := ReadJsonDataToAlgebra(fmt.Sprintf("%v/data_%d.json", folder, len(already_computed_data))).Indecomposables

    t0 := time.Now().Truncate(time.Minute)
    fmt.Printf("\nComputing algebra %v. It is now %v:%v.\n", len(already_computed_data), t0.Hour(), t0.Minute())
    next_row := Enumerate_tau_tilting_modules(next_modules, number_threads, granularity)

    already_computed_data = append(already_computed_data, next_row)
    already_computed_data = PadJaggedArray(already_computed_data, len(next_row))
    WriteSliceToCsv(tau_out, already_computed_data)
    return
}
