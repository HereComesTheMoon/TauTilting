comp := function(n, K)
	local orientation, Q, i, A;

    orientation := [[n-2, n]];
    for i in [1..n-1] do
        Add(orientation, [i,i+1]);
    od;

    Print(orientation);
    Q := Quiver(n, orientation);

    A := TruncatedPathAlgebra(Rationals, Q, 2);

    return A;

end;
