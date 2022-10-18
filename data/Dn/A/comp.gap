comp := function(n, K)
	local orientation, Q, i;

    if n = 1 then
        orientation := [];
        Q := DynkinQuiver("A", n, orientation);
    elif n = 2 then
        orientation := [];
        Q := Quiver(2, []);
    elif n = 3 then
        orientation := ["r", "l"];
        Q := DynkinQuiver("A", n, orientation);
    else
        orientation := [];
        for i in [2..n] do
            Add(orientation, "r");
        od;
        Q := DynkinQuiver("D", n, orientation);
    fi;

	return TruncatedPathAlgebra(Rationals, Q, K);
end;
