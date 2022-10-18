comp := function(n, K)
	local arrows, KQ, I, Q, i, arr, k;

    if n <= 3 then
        Q := Quiver(1, []);
    else
        arrows := [[1,2],[1,3],[2,4],[3,4]];
        for i in [5..n] do
            Add(arrows, [i-1,i]);
        od;
        Q := Quiver(n, arrows);
    fi;

	KQ := PathAlgebra(Rationals, Q);

    if n <= 3 then 
        I := Ideal(KQ, []);
    else
        I := [KQ.a1 * KQ.a3 - KQ.a2 * KQ.a4];
        arrows := GeneratorsOfAlgebra(KQ);
        if n >= 5 then 
            Add(I, KQ.a4 * KQ.a5);
            Add(I, KQ.a3 * KQ.a5);
        fi;
        for i in [1..n-3-K] do
            # Arrows gains one element per n, for vertex n. Hence the shift by n.
            arr := arrows[n+4+i];
            for k in [1..K-1] do
                arr := arr * arrows[n+4+i+k];
            od;
            # WITH relations between the commutative square and the arrow a_5 here
            Add(I, arr);
        od;
        I := Ideal(KQ, I);
    fi;

    return KQ/I;
end;
