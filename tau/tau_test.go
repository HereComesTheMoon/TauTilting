package tau

import (
    "testing"
    "fmt"
    "runtime"
)


func TestEnumerateLinearAnMod5(t *testing.T) {
    runtime.SetBlockProfileRate(1)
    runtime.SetMutexProfileFraction(1)
    wanted := ReadCsvToSlice("./test/LinearAnMod5/LinearAnMod5_TAU_RESULT.csv")

    loc := "./test/LinearAnMod5/"
    for k := 0; k < 10; k++ {
        fmt.Printf("Counting for %v.\n", loc + "alg_" + fmt.Sprint(k))
        //result := CountSTautiltingModulesStatic(locations[k])
        alg := ReadJsonDataToAlgebra(fmt.Sprintf("%v/data_%d.json", loc, k + 1))
        
        if !alg.SanityCheck() {
            t.Fatal("Modules did not pass sanity check!.")
        }
        result := Enumerate_tau_tilting_modules(alg.Indecomposables, 8, 3)
        
        fmt.Printf("WANTED: %v\n", wanted[k + 1])
        fmt.Printf("   GOT: %v\n", result)
        for i, v := range wanted[k + 1] {
            // v != 0 since once we hit 0, the rest is just padding
            if v != 0 && v != result[i] {
                t.Fatal("Value incorrect.")
            }
        }
    }
}


func TestListLinearAnMod5(t *testing.T) {
    runtime.SetBlockProfileRate(1)
    runtime.SetMutexProfileFraction(1)
    wanted := ReadCsvToSlice("./test/LinearAnMod5/LinearAnMod5_TAU_RESULT.csv")

    loc := "./test/LinearAnMod5/"
    for k := 0; k < 10; k++ {
        fmt.Printf("Counting for %v.\n", loc + "alg_" + fmt.Sprint(k))
        alg := ReadJsonDataToAlgebra(fmt.Sprintf("%v/data_%d.json", loc, k + 1))
        if !alg.SanityCheck() {
            t.Fatal("Modules did not pass sanity check!.")
        }
        modules := alg.Indecomposables

        result_modules := List_tau_tilting_modules(modules, 8, 3)

        result := make([]int, len(result_modules))
        for i, modules := range result_modules {
            result[i] = len(modules)
        }
        
        fmt.Printf("WANTED: %v\n", wanted[k + 1])
        fmt.Printf("   GOT: %v\n", result)
        for i, v := range wanted[k + 1] {
            // v != 0 since once we hit 0, the rest is just padding
            if v != 0 && v != result[i] {
                t.Fatal("Value incorrect.")
            }
        }
    }
}

