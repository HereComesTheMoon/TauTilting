package tau

import (
    "testing"
    "fmt"
    "runtime"
)


//func BenchmarkStaticBasic(b *testing.B) {
    //runtime.SetBlockProfileRate(1)
    //runtime.SetMutexProfileFraction(1)
    //loc := "../Data/2/Mod2/alg_"
    //for i := 0; i < b.N; i++ {
        //fmt.Println("Testing:")
        //CountSTautiltingModulesStatic(loc + "20")
    //}
//}

//func BenchmarkStaticMod4(b *testing.B) {
    //runtime.SetBlockProfileRate(1)
    //runtime.SetMutexProfileFraction(1)
    //loc := "../Data/2/Mod4/alg_"
    //for i := 0; i < b.N; i++ {
        //fmt.Println("Testing:")
        //CountSTautiltingModulesStatic(loc + "17")
    //}
//}

//func TestStaticLinearAnMod5(t *testing.T) {
    //runtime.SetBlockProfileRate(1)
    //runtime.SetMutexProfileFraction(1)
    //wanted := iodata.ReadCsvToSlice("../TestData/LinearAnMod5/LinearAnMod5_TAU_RESULT.csv")

    //loc := "../TestData/LinearAnMod5/alg_"
    //locations := [10]string{
        //loc + "1",
        //loc + "2",
        //loc + "3",
        //loc + "4",
        //loc + "5",
        //loc + "6",
        //loc + "7",
        //loc + "8",
        //loc + "9",
        //loc + "10",
    //}
    //for k := 0; k < len(locations); k++ {
        //result := CountSTautiltingModulesStatic(locations[k])
        //for i, v := range wanted[k + 1] {
            //// v != 0 since once we hit 0, the rest is just padding
            //if v != 0 && v != result[i] {
                //fmt.Printf("Counting for %v.\n", locations[k])
                //fmt.Printf("WANTED: %v\n", wanted[k + 1])
                //fmt.Printf("   GOT: %v\n", result)
                //t.Fatal("Value incorrect.")
            //}
        //}
    //}
//}


func TestEnumerateLinearAnMod5(t *testing.T) {
    runtime.SetBlockProfileRate(1)
    runtime.SetMutexProfileFraction(1)
    wanted := ReadCsvToSlice("./test/LinearAnMod5/LinearAnMod5_TAU_RESULT.csv")

    loc := "./test/LinearAnMod5/"
    for k := 0; k < 10; k++ {
        fmt.Printf("Counting for %v.\n", loc + "alg_" + fmt.Sprint(k))
        //result := CountSTautiltingModulesStatic(locations[k])
        modules := ReadJsonDataToAlgebra(fmt.Sprintf("%v/data_%d.json", loc, k + 1)).Indecomposables
        result := Enumerate_tau_tilting_modules(modules, 8, 3)
        
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
        //result := CountSTautiltingModulesStatic(locations[k])
        modules := ReadJsonDataToAlgebra(fmt.Sprintf("%v/data_%d.json", loc, k + 1)).Indecomposables
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

