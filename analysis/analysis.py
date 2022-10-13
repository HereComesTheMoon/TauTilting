import functools
import typing
from collections import Counter

from sympy.abc import n, x
import sympy as sp
import csv

@functools.cache
def catalan(n: int) -> int:
    if n <= 1:
        return 1
    return sum(catalan(n - 1 - k) * catalan(k) for k in range(n))


@functools.cache
def gen_function_Q(K: int) -> str:
    summies = [str(catalan(i - 1)) + f"*x**{i}" for i in range(1, K + 1)]
    return "(1 - " + " - ".join(summies) + ")"


@functools.cache
def gen_function_F(K: int, k: int) -> str:
    # TODO: Check if this checks out
    gen_Q = gen_function_Q(K)
    return f"(x**{k}*{gen_Q}**({-k - 1}))"


@functools.cache
def gen_function_F2(K: int, k: int) -> str:
    # TODO: Check if this checks out
    gen_Q = gen_function_Q(K)
    return f"({gen_Q}**({-k - 1}))"


def genf_from_rec(quotient: str, starting_vals: typing.Sequence) -> str:
    """Given a sequence of values, and a quotient Q, this finds a polynomial P, such that P/Q generates the starting values."""
    goal = " + ".join([str(a) + f"*x**{k}" for k, a in enumerate(starting_vals, 0)])
    goal = sp.sympify(goal)

    coeffs = [f"a{k}" for k in range(len(starting_vals))]
    pol = " + ".join(f"a{k}*x**{k}" for k in range(len(starting_vals)))
    pol_result = " + ".join("{" + f"a{k}" + "}" + f"*x**{k}" for k in range(len(starting_vals)))
    pol = sp.sympify(pol)


    # = coeffs_of_generating_function(sp.sympify(quotient), 5*len(starting_vals))
    quotient_pol = sp.sympify(quotient).series(x, n=len(starting_vals) + 1)

    comparison = quotient_pol * pol
    comparison = comparison.series(x, n=len(coeffs))
    comparison = comparison.removeO()
    # print(comparison)

    # print(f"{goal} === {comparison}")

    sol = sp.solvers.solvers.solve_undetermined_coeffs(sp.Eq(goal, comparison), coeffs, x)
    sol = {str(k): v for k, v in sol.items()}
    result = pol_result.format(**sol)
    result = sp.sympify(result)
    return "(" + str(result) + ")"


def coeffs_of_generating_function(formula: str, degree: int, trim: bool = False) -> list[int]:
    """Compute coefficients of generating function."""
    symp = sp.sympify(formula)
    p = symp.series(n=degree + 3)
    coeffs = [p.coeff(x, k) for k in range(degree)]
    return coeffs


def find_best_recurrence(seq: list[int]) -> tuple[int, ...]:
    recurrences = []
    for k in range(len(seq)):
        next = seq[k:]
        depth = len(next)
        encoded_sequence = sp.SeqPer(next, (n, 0, depth))
        seq_result = encoded_sequence.find_linear_recurrence(depth, d=(depth // 2 - 1))
        if seq_result:
            recurrences.append(tuple(seq_result))
    count = Counter(recurrences)
    best = count.most_common(1)[0][0]
    assert 2 * count.most_common(1)[0][1] > count.total()
    print(f"{best=}")
    if len(count) > 1:
        print(f"Total counted: {count.total()}, Second best={count.most_common(2)[1]}")
    return best


def check_and_read_diags(location: str):
    taus = read_tautilting(location + "/TAUTILTING.csv")
    # diags = read_diags(location + "/DIAGONALS.csv")
    # for a, b in zip(iterate_diags(taus), diags):
        # length = min(len(a), len(b))
        # assert a[:length] == b[:length]
    return [row for row in iterate_diags(taus)]


def check_and_read_diags_padded(location: str) -> list[list[int]]:
    diags = check_and_read_diags(location)
    padded = [
        [0] * k + diag for k, diag in enumerate(diags)
    ]
    for row in padded:
        assert len(padded[0]) == len(row)

    return padded

def read_tautilting(location: str):
    with open(location) as f:
        reader = csv.reader(f)
        table = [
            list(map(int, row)) for row in reader
        ]
    return table


def iterate_diags(table: list[list[int]]) -> typing.Generator[list[int], None, None]:
    for row in table:
        assert len(row) == len(table[0])
    for k in range(len(table[0]) - 1, 0, -1):
        diag = [table[i][k + i] for i in range(min(len(table), len(table[0]) - k))]
        if any(diag):
            yield diag
    for k in range(len(table)):
        diag = [table[k + i][i] for i in range(min(len(table[0]), len(table) - k))]
        if any(diag):
            yield diag

