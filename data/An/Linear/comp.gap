comp := function(n, K)
#Computes linearly oriented truncated A_n and saves it as alg_n.gap. KQ/KQ+^7
	local orientation, Q, i;

    orientation := [];
    for i in [1..n-1] do
        Add(orientation, "r");
    od;

	Q := DynkinQuiver("A", n, orientation);

    return TruncatedPathAlgebra(Rationals, Q, K);
end;
