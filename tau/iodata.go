package tau

import (
    "log"
    "os"
    "encoding/csv"
    "strings"
    "strconv"
    "io"
    "fmt"
    "math"
    "time"
    "errors"
)

// TODO: Read needs to be rewritten to account for new format

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
    //file, err := os.Open(location)
    if err != nil {
        fmt.Printf("Unable to find the file %v. Redirecting to stdout instead.\n", location)
        file = os.Stdout
    }
    defer file.Close()
    //w := fmt.Fprintf(

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
            //row_strings[i] = " " + strconv.Itoa(v)
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


func initialiseTautiltingCsv(folder string) {
    loc := folder + "TAUTILTING.csv"
    _, err := os.Stat(loc)
    if errors.Is(err, os.ErrNotExist) {
        // File does not exist, initialise.
        for i := 0; i < 32; i++ {
            next_file_path := fmt.Sprint(folder, "/alg_", i, "_rigiditymatrix.csv")
            fmt.Println(next_file_path)
            _, err := os.Stat(next_file_path)
            if errors.Is(err, os.ErrNotExist) {
                if i == 31 {
                    log.Fatalf("No alg_i_rigiditymatrix.csv files.\n")
                }
                continue
            }
            // Else, create new TAUTILTING.csv file
            new_tau := make([][]int, 0, i)
            for j := 0; j < i; j++ {
                new_tau = append(new_tau, make([]int, i))
            }
            WriteSliceToCsv(loc, new_tau)
            return
        }
    } else {
        log.Printf("%v already exists.\n", loc)
    }
}

// Given a folder, read in TAUTILTING.csv. Count the rows, then use Get_indecomposables() to read in alg_len_dimv.csv, alg_len_rigiditymatrix.csv, alg_len_modules.csv (from this folder), and compute the next row. Then append the row to the file, and save it. Also save DIAGONALS.csv, and DIAGONALS_PADDED.csv
func ComputeNextRow(folder string, number_threads int, granularity int) {
    tau_out := folder + "/TAUTILTING.csv" // Repeated slash may be turned into a single one. Slash="/" good here.
    //diag_out := folder + "/DIAGONALS.csv"
    _, err := os.Stat(folder)
    if errors.Is(err, os.ErrNotExist) {
        log.Fatalf("Folder %v does not exist.\n", folder)
    }
    _, err = os.Stat(tau_out)
    if errors.Is(err, os.ErrNotExist) {
        initialiseTautiltingCsv(folder)
    }
    // Else, actually compute the next row.
    data := ReadCsvToSlice(tau_out)
    next_modules := Get_indecomposables(folder, len(data))

    t0 := time.Now().Truncate(time.Minute)
    fmt.Printf("\nComputing algebra %v. It is now %v:%v.\n", len(data), t0.Hour(), t0.Minute())
    next_row := Enumerate_tau_tilting_modules(next_modules, number_threads, granularity)

    data = append(data, next_row)
    data = PadJaggedArray(data, len(next_row))
    WriteSliceToCsv(tau_out, data)
    return
}
