comp := function(n, K)
#Computes almost linearly oriented truncated A_n and saves it as alg_n.gap. KQ/KQ+^2
	local orientation, Q, i, A, outputFile;

    if n < 5 then
        orientation := [];
        Q := DynkinQuiver("A", 1, orientation);

        A := TruncatedPathAlgebra(Rationals, Q, 2);

        outputFile := Filename(DirectoryCurrent(), Concatenation("alg_", String(n), ".gap"));

        # "keep", "delete", "query". Recompute the algebras?
        SaveAlgebra(A, outputFile, "delete");
    else
        orientation := [[1,3],[2,3],[n-2, n-1],[n-2,n]];
        for i in [3..n-3] do
            Add(orientation, [i,i+1]);
        od;

        Q := Quiver(n, orientation);

        A := TruncatedPathAlgebra(Rationals, Q, 2);

        return A;

    fi;
end;
