# MT19937-64 reference output

`mt19937-64.out.txt` is what the authors' own program prints, unmodified:

    curl -O https://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/VERSIONS/C-LANG/mt19937-64.c
    gcc -O2 -o mtref mt19937-64.c && ./mtref > mt19937-64.out.txt

Its `main` seeds with `init_by_array64({0x12345, 0x23456, 0x34567, 0x45678})` and prints
1000 `genrand64_int64` values followed by 1000 `genrand64_real2` values -- **from one
continuous stream**, so a test has to ask for them in that order.

`mt19937-64-seed.out.txt` covers `init_genrand64`, which that `main` never calls. It is
the same file with `main` replaced by one that calls `init_genrand64(1234)` and prints
1000 `genrand64_int64` values. Only `main` differs; the algorithm is the authors' code.

The published `mt19937-64.out.txt` on that site is currently a 404, which is why these are
regenerated from the source rather than downloaded.
