ComputeAlgebrasUpTo := function(folder, n, K)
    local k, nn, folderInner;

    if not IsDirectoryPath(folder) then
        Error(Concatenation("The directory ", folder, " does not exist."));
    fi;

    if not IsInt(K) then
        Error(Concatenation("K = ", String(K), " is not an integer."));
    fi;
    for k in [2..K] do 
        folderInner := Concatenation(folder, "/Mod", String(k));

        if not IsDirectoryPath(folderInner) then
            Print(Concatenation("The directory ", folder, "/Mod", String(k), " does not exist. Making directory.\n"));
            Exec(Concatenation("mkdir ", folderInner));
        fi;

        for nn in [1..n] do
            ComputeAlgebra(folder, nn, k);
        od;
    od;
end;


ComputeAlgebra := function(folder, n, K)
    local A, output, folderInner;

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
end;


