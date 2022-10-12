comp := function(n, K)
    local A, Q;
    Q := DynkinQuiver("A", 3, ["r", "r"]);
    A := TruncatedPathAlgebra(Rationals, Q, K);

    return A;
end;

