from analysis import check_and_read_diags_padded, gen_function_F, coeffs_of_generating_function
from analysis import genf_from_rec, gen_function_F2
from analysis import find_best_recurrence
from tabulate import tabulate
import sympy as sp

data_folder = "../data/"


class Solution:
    def __init__(self, location: str, hypo: dict[int, str], offset: dict[int, int], n: dict[int, int], m: int):
        self.location: str = data_folder + location + "/"
        self.hypo: dict[int, str] = hypo
        self.offset: dict[int, int] = offset
        self.n: dict[int, int] = n
        self.m: int = m

    def check_specific_M(self, M: int, verbose: bool = False) -> bool:
        """Check whether a hypothesis holds for a specific value of M. If there are mistakes, it will complain."""
        diags = check_and_read_diags_padded(self.location + f"/Mod{M}/")
        result = True

        n = self.n[M] + 1
        for k in range(n):
            offset = max(self.offset[M] - (M-1)*k, 0)
            assert offset >= 0
            genf = self.hypo[M].format(F=gen_function_F(M, k))
            prediction = coeffs_of_generating_function(genf, n)
            diag = diags[k]

            assert len(prediction) >= n # This should never fail
            if len(diag) != n: # This might happen if n is badly chosen.
                print(f"Bad n! n is probably too high. Maybe too low. {M=}, {k=}, {self.n[M]=}, {len(diag)-1=}")
                assert False

            given = diag[offset:n]
            wanted = prediction[offset:]
            if given != wanted:
                result = False
            if given != wanted or verbose:
                output = tabulate(
                    [
                        [f"{M=}"] + list(range(n)),
                        [f"{k=}"] + ["X" if offset <= i < n else "" for i in range(n)],
                        ["GIVEN"] + diag,
                        ["WANTED"] + prediction,
                        ["WRONG?"] + ["?" if a != b else "" for a, b in zip(diag, prediction)]
                    ]
                )
                print(output)
        return result

    def check_all(self, verbose: bool = False):
        assert len(self.hypo) == len(self.offset) == len(self.n)
        assert self.hypo.keys() == self.offset.keys() == self.n.keys()

        if all(self.check_specific_M(M, verbose) for M in self.hypo):
            print(f"Sequence of algebras {self.location} passed.")

    def find_solution(self, M: int) -> str:
        from sympy.abc import x
        k = 5
        diags = check_and_read_diags_padded(self.location + f"/Mod{M}/")

        last = "0"

        # Upper bound, at which point. Better bound is around self.n[M] / M.
        # Expect process to terminate way before, though.
        for k in range(self.n[M]):
            # Find generating function with quotient F_{M,k}, which fits diags[k]
            _pk = genf_from_rec(gen_function_F2(M, k), diags[k])

            # Remove the extraneous x**k stuff.
            _pk = "(" + str(sp.sympify(_pk + f" *x**(-{max(self.m, k)})").simplify()) + ")"

            # FORMAT NICELY:
            pk = _pk[1:-1].split()
            pk = [a + " " + b for a, b in zip((["+"] + pk)[::2], pk[::2])]
            pk = reversed(pk)
            pk = " ".join(pk)[2:]
            # The formatting has a bug if the polynomial with the highest degree has a negative coefficient. In this case it will print "+ - coefficient X^degree". This is technically correct,
            # but it requires manually fixing when used together with print_solution_table
            # Try this:
            pk = pk.replace("+ -", "- ")
            # Check that formatted is same thing:
            assert sp.Eq(sp.sympify(_pk), sp.sympify(pk))
            # pk2 = reversed(pk[1:-1].split(sep=" "))

            print(pk)

            # Check if the polynomial we got is the same as the previous one
            # Then we found a solution, in all likelihood
            f = sp.Poly(f"{pk} - ({last})", x)
            if f.degree() < 0:
                # If this assert goes wrong, then the string formatting might be wrong
                assert last == pk
                print(f"Found possible solution for {self.location}Mod{M}/ at {k=}.")
                print()
                print(pk)
                print()
                return pk
            else:
                last = pk

        print("No solution found.")
        return ""

    def print_solution_table(self):
        """Note: This was used purely for my thesis, and is horrible code. The resulting table  is not guaranteed to be perfect, especially the last column may need to be adjusted."""
        print("\\begin{table}[!htbp]")
        print("\t\\caption{Results: XXX}")
        print("\t\\centering")
        print("\\begin{tabular}{|c|p{10.1cm}|r|}")
        print("\t\\hline")
        print("\tM & Generator & \\text{Verified for \\quad} \\\\")
        print("\t\\midrule")
        for M, sol in self.hypo.items():
            p = sol[:-3].replace("x", "X").replace("**", "^").replace("*", "").replace("^", "^{").replace(" +", "} +").replace(" -", "} -").replace(")", "})").replace("(", "({")
            row = f"\t{M} & $F_{{{M},k}}(X){p}$ & "
            if self.offset[M] == 0:
                row += f"$0 \\leq n \\leq {self.n[M]}$ \\\\"
            else:
                row += f"${self.offset[M]} - {M-1}k \\leq n \\leq {self.n[M]}$ \\\\"
            print(row)
            print("\t\\hline")
        print("\\end{tabular}")
        print("\\end{table}")



sol0 = Solution(
    location="0",
    hypo={
        2: "()*{F}",
        3: "()*{F}",
        4: "()*{F}",
        5: "()*{F}",
    },
    offset={
        2: 6,
        3: 8,
        4: 10,
        5: 12,
    },
    n={
        2: 24,
        3: 23,
        4: 22,
        5: 21,
    },
    m=0,
)






# def factor(sol: Solution):
    # for k, pol in sol.hypo.items():
        # pol = pol.replace("*{F}", "")
        # p = sp.poly_from_expr(sp.sympify(pol))
        # print(p)
        # print(sp.polys.factor(pol))


# def common_zeroes(sol: Solution):
    # from itertools import combinations
    # for p, q in combinations(sol.hypo.values(), 2):
         # p = p.replace("*{F}", "")
         # q = q.replace("*{F}", "")
         # print(sp.polys.resultant(p, q))


# for p in sol.hypo.values():
    #     for q in sol.hypo.values():
    #         if p == q:
    #             continue
    #         p = p.replace("*{F}", "")
    #         q = q.replace("*{F}", "")
    #         print(sp.polys.resultant(p, q))
    #

# sols = [sol2, sol4, sol5, sol6, sol7, sol10, sol11]

if __name__ == '__main__':
    # sol6.print_solution_table()
    # sol5.find_solution(2)
    # sol7.check_specific_M(7)
    # sol2.find_solution(2)
    diags = check_and_read_diags_padded("../data/2/Mod2/")
    find_best_recurrence(diags[0])
