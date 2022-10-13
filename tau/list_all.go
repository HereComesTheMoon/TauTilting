package tau

import (
	"fmt"
    "sync"
    "time"
    "log"
    "math/bits"
)

/// Clique object. Represents a basic tau-rigid module. 
type cliqueFull struct {
    vertices []int // Slice containing the IDs of the indecomposable modules whose direct sum represents a specific basic tau-rigid module
    cnbrs []int // Slice containing the IDs of all indecomposable modules X, such that Clique \oplus X is tau-rigid
    dimv uint // Stores at which vertices the dimension vector is zero/nonzero, in particular, the support rank.
}


/// Generates a slice containing all support tau-tilting modules over some algebra
func List_tau_tilting_modules(modules []Indecomposable, number_threads int, granularity int) [][][]int {
    number_modules := len(modules)
    number_vertices := len(modules[0].Dim_vector)

    t0 := time.Now()
    // Adjacency matrix for our tau-rigidity graph
    adj := make([][]bool, 0, number_modules)
    for i := 0; i < number_modules; i++ {
        next_row := make([]bool, number_modules)
        adj = append(adj, next_row)

        for j := 0; j < number_modules; j++ {
            // Note the i < j. Enforces an order on our vertices, prevents loops and double counting
            // Otherwise, false if the direct sum of modules[i] and modules[j] is not \tau-rigid
            if modules[i].Tau_rigidity_row[j] == 0 || modules[j].Tau_rigidity_row[i] == 0 {
                adj[i][j] = false
            } else if i < j {
                adj[i][j] = true
            }
        }
    }

    // The m-th bit of dimvs[id] is 1 iff module[id] is non-zero at vertex m
    dimvs := make([]uint, 0, number_modules)

    // new thread-safe WrappedStack, which will be used to store cliques
    ws := newWStack[cliqueFull](number_modules)

    result := make([][][]int, number_vertices + 1)
    result[0] = make([][]int, 1) // To account for the 0-module
    result[1] = make([][]int, 0, number_modules)

    for i, m := range modules {
        if i != m.Id { log.Fatalf("i=%v != m.id=%v\n", i, m.Id) }
        uint_dimv := dimv_to_uint(m.Dim_vector)

        next_module := cliqueFull{
            vertices: []int{m.Id},
            cnbrs: make([]int, 0, number_modules - 1),
            dimv: uint_dimv,
        }

        for id, tau_rigid := range adj[m.Id] {
            if tau_rigid { next_module.cnbrs = append(next_module.cnbrs, id) }
        }

        dimvs = append(dimvs, uint_dimv)
        ws.push(next_module)
        if bits.OnesCount(next_module.dimv) == len(next_module.vertices) {
            result[1] = append(result[1], []int{m.Id})
        }
    }

    delta := time.Since(t0)
    fmt.Printf("Preparation took %v seconds.\n", delta)

    mutex_result := sync.Mutex{}


    // Find cliques:
    fmt.Print("Thread execution times: ")
    wg := sync.WaitGroup{}
    for i := 0; i < number_threads; i++ {
        wg.Add(1)
        go thread_wrapper_list_tau_tilting_modules(ws, dimvs, adj, result, &mutex_result, &wg, granularity)
    }
    wg.Wait()
    fmt.Println()
    return result[:number_vertices + 1]
}

func thread_wrapper_list_tau_tilting_modules(
    ws *wStack[cliqueFull],
    dimvs []uint,
    adj [][]bool,
    result [][][]int,
    mutexResult *sync.Mutex,
    wg *sync.WaitGroup,
    granularity int) {

    t0 := time.Now()
    thread_result := make([][][]int, len(result))

    defer wg.Done()
    defer func() {
        mutexResult.Lock()
        for i := range(thread_result) {
            result[i] = append(result[i], thread_result[i]...)
        }
        mutexResult.Unlock()
        fmt.Printf("%s, ", time.Since(t0).String())
    }()

    for {
        now, err := ws.popFront()
        if err != nil {
            break
        }

        thread_list_tau_tilting_modules(now, dimvs, adj, thread_result, ws, granularity)
    }
}


func thread_list_tau_tilting_modules(now cliqueFull , dimvs []uint, adj [][]bool, thread_result [][][]int, ws *wStack[cliqueFull], granularity int) {
    next := cliqueFull{
        vertices: now.vertices,
        cnbrs: make([]int, len(now.cnbrs)),
        dimv: now.dimv,
    }

    next.vertices = append(next.vertices, -1)

    for _, v := range now.cnbrs {
        next.vertices[len(next.vertices) - 1] = v
        next.dimv = now.dimv | dimvs[v]

        next.cnbrs = next.cnbrs[:len(now.cnbrs)]
        counter := 0
        for _, w := range now.cnbrs {
            if adj[v][w] {
                next.cnbrs[counter] = w
                counter++
            }
        }
        next.cnbrs = next.cnbrs[:counter]

        if bits.OnesCount(next.dimv) == len(next.vertices) {
            thread_result[len(next.vertices)] = append(thread_result[len(next.vertices)], next.vertices)
        }

        if len(next.cnbrs) == 0 {
            continue
        }

        if len(next.vertices) <= granularity {
            // At this point a deep copy is necessary. Multi-threading is fun.
            deep_next := cliqueFull{
                vertices: next.vertices,
                //vertices: make([]int, len(next.vertices)),
                cnbrs: make([]int, len(next.cnbrs)),
                dimv: next.dimv,
            }
            // I feel like copying the vertices here should be necessary, but the program appears to work either way.
            //copy(deep_next.vertices, next.vertices)
            // Copying the common neighbours is definitely necessary.
            copy(deep_next.cnbrs, next.cnbrs)
            ws.push(deep_next)
            continue
        }
        thread_list_tau_tilting_modules(next, dimvs, adj, thread_result, ws, granularity)
    }
}

