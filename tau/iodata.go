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



//func ReadJson(location string) {
    //f, err := os.Open(location)

    //if err != nil {
        //log.Fatal(err)
    //}

    //dec := json.NewDecoder(f)

    //chunk := dataBlock{}
    ////alg := Algebra{}

    //t, err := dec.Token()
    //if err != nil { log.Fatal(err) }
    //if t != json.Delim('{') { log.Fatalf("Got: '%v' of type %T. Expected: '{'", t, t)}

    //t, err = dec.Token()
    //if err != nil { log.Fatal(err) }
    //if t != "algebra" { log.Fatalf("Got: '%v' of type %T. Expected: 'algebra'", t, t)}

    //dec.Decode(&chunk.Algebra)
    ////fmt.Printf("%+v\n", chunk.Algebra)

    //t, err = dec.Token()
    //if err != nil { log.Fatal(err) }
    //if t != "modules" { log.Fatalf("Got: '%v' of type %T. Expected: 'modules'", t, t)}

    ////chunk.Modules = make([]Indecs, 0, chunk.Algebra.Number_modules)

    ////for i := 0; i < chunk.Algebra.Number_modules; i++ {
        
    ////}
    ////dec.Decode(&chunk.Modules)
    ////fmt.Printf("%+v\n", chunk.Modules)

    ////for i := 0; i < 2; i++ {
        ////t, err := dec.Token()
        ////if err != nil { log.Fatal(err) }
        ////fmt.Printf("%T: %v\n", t, t)
    ////}


    ////var alg_data Algebra
    ////err = dec.Decode(&alg_data)

    ////if err != nil {
        ////log.Fatal(err)
    ////}

    
    ////fmt.Printf("%v", alg_data)
//}

//func ReadJson2() {
    //const rawTest = `{"algebra_data" : "IsQuotientOfPathAlgebra\nVertices:\n1, 2, 3\nArrows:\na_1:1->2, a_2:2->3\nField:\nRationals\nRelations:\n(1)*a_1*a_2\n","number_arrows" : 2,"number_modules" : 5,"number_orbits" : 3,"number_vertices" : 3}`

    //dec := json.NewDecoder(strings.NewReader(rawTest))

    ////record := new(map[string]any);

    //record := Algebra{};

    //dec.Decode(&record);
    //dec.Decode(&record);
    //dec.Decode(&record);
    //dec.Decode(&record);
    //dec.Decode(&record);

    ////for k, v := range *record {
        ////fmt.Printf("%T — %T\n", k, v)
        ////fmt.Printf("%v — %v\n", k, v)
    ////}

    //fmt.Print(record)
    
//}

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
    tau_out := folder + "/TAUTILTING.csv" 
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
    already_computed_data := ReadCsvToSlice(tau_out)
    next_modules := ReadJsonDataToAlgebra(fmt.Sprintf("%v/alg_%d.json", folder, len(already_computed_data))).Indecomposables

    t0 := time.Now().Truncate(time.Minute)
    fmt.Printf("\nComputing algebra %v. It is now %v:%v.\n", len(already_computed_data), t0.Hour(), t0.Minute())
    next_row := Enumerate_tau_tilting_modules(next_modules, number_threads, granularity)

    already_computed_data = append(already_computed_data, next_row)
    already_computed_data = PadJaggedArray(already_computed_data, len(next_row))
    WriteSliceToCsv(tau_out, already_computed_data)
    return
}
