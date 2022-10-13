package tau

import (
	"fmt"
	//"graph/iodata"
	"log"
)

// Refactorisation:
// Main functions:
// 1. enumerateTauForAlgebra([][]int rigidityMatrix, [][]int dimensionVectors, int numberThreads, int granularity)
//Have a few checks that rigidityMatrix and dimensionVectors match at all
//2. Same as above, but given some file location, append to the corresponding output file
//3. Actually return a slice containing all tau-tilting modules
//For this might as well implement an 'indecomposable' struct, containing the same information as the python and gap scripts, ie. id, dimension vector, orbit, orbit_representative

// Clique object
//type Clique struct {
//vertices []int
//cnbrs []int
//dimv uint
//}

type Indecomposable struct {
    id int // id, for bookkeeping purposes
    dimv []int // dimension vector
    orbit int // Each orbit corresponds to a projective module, say P(orbit)
    orbit_pos int // Orbit position. Now the indecomposable is isomorphic to \tau^{-orbit_pos} P(orbit)
    pdim int // projective dimension
    idim int // injective dimension
    tau_rigidity []int // tau_rigidity[j] equals 0 if \Hom_A(M, \tau N) = 0, where M is the current indecomposable, and N is the indecomposable with id j. // Check if this is the case, might be flipped.
}


type AllIndecomposables []Indecomposable



func (mods AllIndecomposables) SanityCheck() bool {
    // Check that each id is unique
    //unique_ids := make(map[int]struct{})
    unique_ids := map[int]struct{}{}
    for _, m := range mods {
        unique_ids[m.id] = struct{}{}
    }
    all_ids_unique := len(unique_ids) == len(mods)

    // Check that each (orbit, orbit_pos) is unique
    unique_orbits := map[[2]int]struct{}{}
    for _, m := range mods {
        o := [2]int{m.orbit, m.orbit_pos}
        unique_orbits[o] = struct{}{}
    }
    all_orbits_unique := len(unique_orbits) == len(mods)

    // Check that each dimv has the same length
    num_vertices := len(mods[0].dimv)
    all_same_number_of_vertices := true
    for _, m := range mods {
        if len(m.dimv) != num_vertices {
            all_same_number_of_vertices = false
            break
        }
    }

    // Check that the number of entries of tau_rigidity equals the number of modules
    number_tau_rigidity_relations_equals_number_modules := true
    for _, m := range mods {
        if len(m.tau_rigidity) != len(mods) {
            number_tau_rigidity_relations_equals_number_modules = false
            break
        }
    }

    // Warn if dimv isn't unique
    // Can I use a set/map of slices for this? Should not be possible, since slices are pointers. What's the idiomatic way of doing this?
    // There is no idiomatic way. Convert to strings.
    dimvs := map[string]struct{}{}
    for _, m := range mods {
        dimvs[fmt.Sprint(m.dimv)] = struct{}{}
    }
    all_dimvs_unique := len(dimvs) == len(mods)


    // Warn if number of simples, projectives, injectives doesn't match
    number_projectives := 0
    number_injectives := 0
    number_simples := 0

    sum := func(dimv []int) int {
        res := 0
        for _, v := range dimv {
            res += v
        }
        return res
    }
    for _, m := range mods {
        if m.pdim == 0 {
            number_projectives++
        }
        if m.idim == 0 {
            number_injectives++
        }
        if sum(m.dimv) == 1 {
            number_simples++
        }
    }
    number_projectives_matches := number_projectives == num_vertices
    number_injectives_matches := number_injectives == num_vertices
    number_simples_matches := number_simples == num_vertices

    // Warn if projective isn't orbit_pos == 0
    projectives_at_right_orbit_pos := true
    for _, m := range mods {
        if m.pdim == 0 && m.orbit_pos != 0 {
            projectives_at_right_orbit_pos = false
            break
        }
    }


    checks := []bool{
        all_ids_unique,
        all_orbits_unique,
        all_same_number_of_vertices,
        number_tau_rigidity_relations_equals_number_modules,
        number_projectives_matches,
        number_injectives_matches,
        number_simples_matches,
        projectives_at_right_orbit_pos,
        all_dimvs_unique,
    }
    checks_names := []string{
        "all_ids_unique",
        "all_orbits_unique",
        "all_same_number_of_vertices",
        "number_tau_rigidity_relations_equals_number_modules",
        "number_projectives_matches",
        "number_injectives_matches",
        "number_simples_matches",
        "projectives_at_right_orbit_pos",
        "all_dimvs_unique",
    }

    passed := true
    for i, v := range checks {
        if !v {
            log.Printf("%v violated!\n", checks_names[i])
            passed = false
        }
    }

    return passed
}


func Get_indecomposables(folder string, number_vertices int) AllIndecomposables{
    loc := folder + "/alg_" + fmt.Sprint(number_vertices)

    dimvs := ReadCsvToSlice(loc + "_dimv.csv")
    module_data := ReadCsvToSlice(loc + "_modules.csv")
    tau_rigidity_matrix := ReadCsvToSlice(loc + "_rigiditymatrix.csv")

    if len(dimvs) != len(module_data) {
        log.Fatalf("%v != %v\n", len(dimvs), len(module_data))
    }

    modules := make(AllIndecomposables, 0, len(dimvs))

    for i := range dimvs {
        indec := Indecomposable{
            id: i,
            dimv: dimvs[i],
            orbit: module_data[i][0],
            orbit_pos: module_data[i][1],
            pdim: module_data[i][2],
            idim: module_data[i][3],
            tau_rigidity: tau_rigidity_matrix[i],
        }
        modules = append(modules, indec)
    }

    return modules
}
