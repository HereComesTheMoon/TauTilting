LoadPackage("qpa");
LoadPackage("json");


ComputeAlgebrasUpTo := function(folder, n, K)
    local k, nn, folderInner;

    if not IsDirectoryPath(folder) then
        Error(Concatenation("The directory ", folder, " does not exist."));
    fi;

    if not IsInt(K) then
        Error(Concatenation("K = ", String(K), " is not an integer."));
    fi;

    for k in [2..K] do 
        ComputeAlgebrasModK(folder, 1, n, k);
    od;
end;

ComputeAlgebrasModK := function(folder, lower_n, upper_n, K)
    local k, nn, folderInner;

    if not IsDirectoryPath(folder) then
        Error(Concatenation("The directory ", folder, " does not exist."));
    fi;

    if not IsInt(K) then
        Error(Concatenation("K = ", String(K), " is not an integer."));
    fi;

    folderInner := Concatenation(folder, "/Mod", String(K));

    if not IsDirectoryPath(folderInner) then
        Print(Concatenation("The directory ", folder, "/Mod", String(K), " does not exist. Making directory.\n"));
        Exec(Concatenation("mkdir ", folderInner));
    fi;

    for nn in [lower_n..upper_n] do
        ComputeAlgebra(folder, nn, K);
    od;
end;


ComputeAlgebra := function(folder, n, K)
    local A, output, folderInner, result, orbits, proj, data, tau_rigidity_matrix, results, data_algebra, out_stream;

    if not IsDirectoryPath(folder) then
        Error(Concatenation("The directory ", folder, " does not exist."));
    fi;

    # This reads a function comp(n) into memory.
    Read(Concatenation(folder, "/comp.gap"));

    folderInner := Concatenation(folder, "/Mod", String(K));
    if not IsDirectoryPath(folderInner) then
        Error(Concatenation("The directory ", folder, "Mod", String(K), " does not exist.\n"));
    fi;

    output := Filename(Directory(folderInner), Concatenation("alg_", String(n), ".alg"));

    A := comp(n, K);
    SaveAlgebra(A, output, "delete");

	# proj := IndecProjectiveModules(A);
	proj := IndecProjectiveModules(A);
    proj := Concatenation(proj, [SimpleModules(A)[2]]);
    Print(proj);

    # QPA struggles with the transpose of a module for the one vertex, zero arrows quiver
    if NumberOfArrows(QuiverOfPathAlgebra(A)) = 0 then
        if n <> 1 then
            Error("Algebra with zero arrows, and more than one vertex. Disconnected.");
        fi;
        orbits := [proj];
        data := [rec(
                id := 0,
                orbit := 1,
                orbit_position := 0,
                proj_dim := 0,
                inj_dim := 0,
                dim_vector := DimensionVector(proj[1]),
                tau_rigidity_row := [1],
            )];
    else 
        orbits := ComputeOrbits(proj);

        data := ComputeModuleDataFromOrbits(orbits);
    fi;

    data_algebra := rec(
        number_vertices := NumberOfVertices(QuiverOfPathAlgebra(A)),
        number_arrows := NumberOfArrows(QuiverOfPathAlgebra(A)),
        number_orbits := Length(orbits),
        number_modules := Length(data),
        algebra_data := ReadAll(InputTextFile(output)),
        indecomposables := data,
    );


    out_stream := OutputTextFile(Filename(Directory(folderInner), Concatenation("data_", String(n), ".json")), false);

    GapToJsonStream(out_stream, data_algebra);
    CloseStream(out_stream);
end;

# Almost always these are the projective modules.
# If they are not, 
ComputeOrbits := function(orbitReps)
    local orbits, numberOrbits, orbit, orbitRep, m, N;
	numberOrbits := Length(orbitReps);
    orbits := [];

	for m in [1..numberOrbits] do
        orbitRep := orbitReps[m];
		orbit := [orbitRep];

		N := DTr(orbitRep,-1);

		while not IsZero(N) and N <> orbitRep do
			Add(orbit, N);
			N := DTr(N,-1);;
		od;

		Add(orbits, orbit);
	od;

    if not IsDuplicateFreeList(Flat(orbits)) then
        Error("Chosen AR-translate orbits intersect for this choice of modules!");
    fi;

    return orbits;
end;

ComputeModuleDataFromOrbits := function(orbits)
    local id, moduleData, pdim, idim, orbit, orbit_counter, orbit_position, result_record, module, dimv, row, modulesFlattened, N;
	moduleData := [];

    id := 1;

    modulesFlattened := Flat(orbits);

    orbit_counter := 1;
    for orbit in orbits do 
        orbit_position := 0;
        for module in orbit do 
            if module <> modulesFlattened[id] then
                Error("");
            fi;
			pdim := ProjDimensionOfModule(module, 2*Length(orbits));
			idim := InjDimensionOfModule(module, 2*Length(orbits));
            if pdim = false then
                Print(Concatenation("Warning: Could not compute projective dimension of module ", String(module)));
                pdim := -1;
            fi;
            if idim = false then
                Print(Concatenation("Warning: Could not compute injective dimension of module ", String(module)));
                idim := -1;
            fi;

            row := [];

            for N in modulesFlattened do
                if HomOverAlgebra(module, DTr(N,1)) = [ ] then
                    Add(row, 1); # Flipped from old code
                else 
                    Add(row, 0);
                fi;
            od;



            result_record := rec(
                id := id - 1, # Zero-indexing
                orbit := orbit_counter,
                orbit_position := orbit_position,
                proj_dim := pdim,
                inj_dim := idim,
                dim_vector := DimensionVector(module),
                tau_rigidity_row := row,
            );

            Add(moduleData, result_record);

            id := id + 1;
            orbit_position := orbit_position + 1;
        od;
        orbit_counter := orbit_counter + 1;
    od;
            
    return moduleData;
end;


# Adjacency matrix whose (i,j) entry is 1 iff M_i \oplus M_j is tau-rigid, otherwise 0. M_i is the indecomposable module with id i.
# Currently not used!
ComputeTauRigidityMatrix := function(orbits)
    local matrixSize, matrix, k, M, N, modules;

    # modules has to be in the same order as moduleData
    modules := Flat(orbits);

	matrixSize := Length(modules);

	matrix := [];
	k := 0;

	for M in modules do
		Add(matrix, []);
		k := k+1;
		for N in modules do
			if HomOverAlgebra(M, DTr(N,1)) = [ ] then
				Add(matrix[k], 1); # Flipped from old code
			else 
				Add(matrix[k], 0);
			fi;
		od;
	od;

    return matrix;
end;

Print("Needs: A file comp.gap in a folder (eg. ./data/alg_name/comp.gap), which contains a single function comp(n, K), which takes two integers and returns an algebra with n vertices. K generally refers to the highest power of the arrow ideal that is factored out.\n\n");
Print("Call ComputeAlgebrasModK(folder, lower_n, upper_n, K) to compute alg_n.gap, data_n.json for all values of n in the range, and put these files in ./data/alg_name/ModK/.");
