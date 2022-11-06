# TauTilting
## Enumeration of sequences of support 𝜏-tilting modules with prescribed support-rank.
This here is research work that will only make sense to the few people in the world who've both studied 𝜏-tilting theory, and have a decent understanding of GAP/QPA, Go, and SymPy. We're computing sequences enumerating the number of support 𝜏-tilting modules with prescribed support defect over sequences of algebras, and fitting generating functions to the resulting data.

If this work looks relevant to your own or there are any questions, please contact the author. 

# How To
There needs to be a folder "MyResultFolder". This folder needs to contain a "comp.gap" file, containing a single function definition. This function is called "comp". It takes two positive integer arguments, n and K, and returns an algebra object.
n denotes the number of vertices.
K denotes the power of the arrow ideal taken as quotient, usually.
Example:

   ```gap
    comp := function(n, K)
        local A, Q, orientation;
        orientation = [];
        for i in [1..n-1] do
            Add(orientation, "r");
        od;
        Q := DynkinQuiver("A", n, orientation);
        A := TruncatedPathAlgebra(Rationals, Q, K);

        return A;
    end;
   ```


Open GAP in this folder, and call the following:

   ```gap
    Read("compute_data.gap");

    ComputeAlgebrasModK("data/MyResultFolder", n_start, n_end, K);
   ```

The code will compute all algebras, and indecomposable modules (along with whether the direct sum of any two modules is \tau-rigid or not) for the algebras given by `comp(n, K)`, where `n` ranges from `n_start` to `n_end`, inclusive. The data will be stored as `alg_n.gap` and as `data_n.json`, in the .JSON format. The `data_n.json` contains the `alg_n.gap` data, and is human readable.


To use this data to generate a table enumerating the number of support \tau-tilting modules with prescribed support rank for a given sequence of algebras, call
```go
ComputeData("data/MyResultFolder", number_threads, granularity, false)
```
in `main.go`, and run the code via `go run .`

The resulting data can then be further analyzed using the tools/scripts in the `analysis/` folder. Creating a `Solution` object with a given set of data and a hypothesis attached to it, allows checking whether the enumerated values can be approximated using a family of generating functions or not.






