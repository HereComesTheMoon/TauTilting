comp := function(n, K)
#Computes almost linearly oriented truncated A_n and saves it as alg_n.gap. KQ/KQ+^K
	local orientation, Q, i;

    if n = 1 then
        orientation := [];
    else
        orientation := ["l"];
        for i in [1..n-2] do
            Add(orientation, "r");
        od;
    fi;

	Q := DynkinQuiver("A", n, orientation);

	return TruncatedPathAlgebra(Rationals, Q, K);

end;

