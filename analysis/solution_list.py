from solutions import Solution

DATA_FOLDER = "../data/"

SolAnAlmostLinear = Solution(
    location= DATA_FOLDER + "An/AlmostLinear/",
    hypo={
        2: "(1 + 2*x**3)*{F}",
        3: "(1 + 5*x**4)*{F}",
        4: "(1 + 14*x**5)*{F}",
        5: "(1 + 42*x**6)*{F}",
    },
    offset={
        2: 0,
        3: 0,
        4: 0,
        5: 0,
    },
    n={
        2: 25,
        3: 21, # XXX: Changed this on the 13.10.22 from 21 to 20
        4: 20,
        5: 20,
    },
    m=0,
)

SolAlnAlmostLinearB = Solution(
    location= DATA_FOLDER + "An/AlmostLinearB/",
    hypo={
        2: "(1 + 4*x**3 + 5*x**4 - 2*x**5)*{F}",
        3: "(1 + 10*x**4 + 14*x**5 - 5*x**6 - 10*x**7)*{F}",
        4: "(1 + 28*x**5 + 42*x**6 - 14*x**7 - 28*x**8 - 70*x**9)*{F}",
        5: "(1 + 84*x**6 + 132*x**7 - 42*x**8 - 84*x**9 - 210*x**10 - 588*x**11)*{F}",
    },
    offset={
        2: 0*4,
        3: 0*5,
        4: 0*6,
        5: 7,
    },
    n={
        2: 24,
        3: 23,
        4: 17,
        5: 18,
    },
    m=0,
)

SolAnLinear = Solution(
    location= DATA_FOLDER + "An/Linear/",
    hypo={
        2: "(1)*{F}",
        3: "(1)*{F}",
        4: "(1)*{F}",
        5: "(1)*{F}",
        6: "(1)*{F}",
        7: "(1)*{F}",
    },
    offset={
        2: 0,
        3: 0,
        4: 0,
        5: 0,
        6: 0,
        7: 0,
    },
    n={
        2: 24,
        3: 22,
        4: 20,
        5: 20,
        6: 20,
        7: 20,
    },
    m=0,
)

SolAnRemovedRelation = Solution(
    location= DATA_FOLDER + "An/LinearRemovedRelation/",
    hypo={
        2: "(1 + 2*x**3)*{F}",
        3: "(1 + 5*x**4)*{F}",
        4: "(1 + 14*x**5)*{F}",
        5: "(1 + 42*x**6)*{F}",
    },
    offset={
        2: 0,
        3: 0,
        4: 0,
        5: 0,
    },
    n={
        2: 20,
        3: 20, # XXX: Changed this on the 13.10.22 from 21 to 20
        4: 20,
        5: 20,
    },
    m=0,
)

SolAnRemovedRelations = Solution(
    location= DATA_FOLDER + "An/LinearRemovedRelations/",
    hypo={
        2: "(1 + 4*x**3 + 5*x**4 - 2*x**5)*{F}",
        3: "(1 + 10*x**4 + 14*x**5 - 5*x**6 - 10*x**7)*{F}",
        4: "(1 + 28*x**5 + 42*x**6 - 14*x**7 - 28*x**8 - 70*x**9)*{F}",
        5: "(1 + 84*x**6 + 132*x**7 - 42*x**8 - 84*x**9 - 210*x**10 - 588*x**11)*{F}",
    },
    offset={
        2: 0*4,
        3: 0*5,
        4: 0*6,
        5: 7,
    },
    n={
        2: 24,
        3: 23,
        4: 17,
        5: 18,
    },
    m=0,
)


SolDn = Solution(
    location= DATA_FOLDER + "Dn/A/",
    hypo={
        2: "(1 + 2*x**3 - x**4)*{F}",
        3: "(1 + 2*x**3 + 9*x**4 - 4*x**5 - 4*x**6)*{F}",
        4: "(1 + 2*x**3 + 9*x**4 + 36*x**5 - 14*x**6 - 20*x**7 - 25*x**8)*{F}",
    },
    offset={
        2: 0,
        3: 0,
        4: 0,
    },
    n={
        2: 24,
        3: 21,
        4: 19,
    },
    m=0,
)


SolDnE = Solution(
    location= DATA_FOLDER + "Dn/Euclidean/",
    hypo={
        2: "(1 + 4*x**3 - 2*x**4 + 4*x**6 - 4*x**7 + x**8)*{F}",
    },
    offset={
        2: 0,
    },
    n={
        2: 28,
    },
    m=5,
)

SolCommutativeSquare = Solution(
    location= DATA_FOLDER + "CommutativeSquare/A/",
    hypo={
        2: "(1 + x**2 + 8*x**3 + 13*x**4 + 8*x**5 - 9*x**6 + 4*x**7)*{F}",
        3: "(1 + x**2 + 2*x**3 + 23*x**4 + 48*x**5 + 23*x**6 - 94*x**7 - 58*x**8 + 48*x**9 + 40*x**10)*{F}",
        4: "(1 + x**2 + 2*x**3 + 8*x**4 + 76*x**5 + 177*x**6 + 70*x**7 - 394*x**8 - 692*x**9 - 523*x**10 + 542*x**11 + 685*x**12 + 700*x**13)*{F}",
        5: "(1 + x**2 + 2*x**3 + 8*x**4 + 34*x**5 + 261*x**6 + 658*x**7 + 222*x**8 - 1602*x**9 - 3155*x**10 - 6374*x**11 - 5503*x**12 + 6258*x**13 + 9744*x**14 + 14504*x**15 + 16464*x**16)*{F}",
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
    m=4,
)

SolCommutativeSquareB = Solution(
    location= DATA_FOLDER + "CommutativeSquare/B/",
    hypo={
        2: "(1 + x**2 + 4*x**3 + 3*x**4 - 2*x**5 + x**6)*{F}",
        3: "(1 + x**2 - 2*x**3 + 3*x**4 - 2*x**5 + 9*x**6 + 2*x**7 - 4*x**8 - 8*x**9)*{F}",
        4: "(1 + x**2 - 2*x**3 - 12*x**4 - 2*x**5 + 9*x**6 + 52*x**7 + 76*x**8 - 28*x**9 - 85*x**10 - 150*x**11 - 125*x**12)*{F}",
        5: "(1 + x**2 - 2*x**3 - 12*x**4 - 44*x**5 + 9*x**6 + 52*x**7 + 216*x**8 + 406*x**9 + 447*x**10 - 458*x**11 - 1161*x**12 - 2226*x**13 - 2940*x**14 - 2744*x**15)*{F}",
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
    m=4,
)


SolAnAlmostLinear.check_all()
