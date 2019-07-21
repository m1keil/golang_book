Using the manderblot from ex8-5 with 4096x4096 PNG resolution.

Results on i5-6200U CPU @ 2.30GHz:

GOMAXPROCS=1

    $ export GOMAXPROCS=1
    $ perf stat -r 10  ./main >/dev/null

    3.81614 +- 0.00502 seconds time elapsed  ( +-  0.13% )

GOMAXPROCS=2

    $ export GOMAXPROCS=2
    $ perf stat -r 10  ./main >/dev/null

    2.41020 +- 0.00472 seconds time elapsed  ( +-  0.20% )

GOMAXPROCS=4

    $ export GOMAXPROCS=4
    $ perf stat -r 10  ./main >/dev/null
 
    2.27412 +- 0.00347 seconds time elapsed  ( +-  0.15% )

GOMAXPROCS=8

    $ export GOMAXPROCS=8
    $ perf stat -r 10  ./main >/dev/null

    2.27529 +- 0.00357 seconds time elapsed  ( +-  0.16% )

GOMAXPROCS=16

    $ export GOMAXPROCS=16
    $ perf stat -r 10  ./main >/dev/null

    2.27094 +- 0.00458 seconds time elapsed  ( +-  0.20% )

GOMAXPROCS=32

    $ export GOMAXPROCS=32
    $ perf stat -r 10  ./main >/dev/null

    2.27034 +- 0.00322 seconds time elapsed  ( +-  0.14% )


Ideal GOMAXPROCS is CPU count. Anything above it have very small or even negative impact.