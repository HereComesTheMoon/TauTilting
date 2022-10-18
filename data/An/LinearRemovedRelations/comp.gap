comp := function(n, K)
	local orientation, arrows, KQ, I, Q, i, k, arr;

    orientation := [];
    for i in [1..n-1] do
        Add(orientation, "r");
    od;
    Q := DynkinQuiver("A", n, orientation);

	KQ := PathAlgebra(Rationals, Q);

    I := [];
    arrows := GeneratorsOfAlgebra(KQ);
    # First n are the length 0 paths. n + 1, n + 2 are the paths we don't want to include
    for i in [3+n..2*n-K] do
        arr := arrows[i];
        for k in [1..K-1] do
            arr := arr * arrows[i+k];
        od;
        Add(I, arr);
    od;
    I := Ideal(KQ, I);

    return KQ/I;
end;
