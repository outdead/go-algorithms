Algorithms in golang
====================

Implementation of basic algorithms and data structures with Go programming language. Tests and benchmarks are included

### Data structures

* [Binary Heap](/structures/binaryheap) ([benchmarks](/structures/binaryheap/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Binary_heap))

### Sort algorithms

* [Bubble sort](/sort/bubblesort) ([benchmarks](/sort/bubblesort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Bubble_sort))
* [Insertion sort](/sort/insertionsort) ([benchmarks](/sort/insertionsort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Insertion_sort))
* [Shellsort](/sort/shellsort) ([benchmarks](/sort/shellsort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Shellsort))
* [Selection sort](/sort/selectionsort) ([benchmarks](/sort/selectionsort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Selection_sort))
* [Merge sort](/sort/mergesort) ([benchmarks](/sort/mergesort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Merge_sort))
* [Quicksort](/sort/quicksort) ([benchmarks](/sort/quicksort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Quicksort))
* [Counting sort](/sort/countingsort) ([benchmarks](/sort/countingsort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Counting_sort))
* [Radix sort](/sort/radixsort) ([benchmarks](/sort/radixsort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Radix_sort))
* [Heapsort](/sort/heapsort) ([benchmarks](/sort/heapsort/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Heapsort))

### Search algorithms

* [Linear search](/search/linearsearch) ([benchmarks](/search/linearsearch/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Linear_search))
* [Interpolation search](/search/interpolationsearch) ([benchmarks](/search/interpolationsearch/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Interpolation_search))
* [Binary search](/search/binarysearch) ([benchmarks](/search/binarysearch/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Binary_search_algorithm))

### Math

* [Fibonacci number](/math/fibonacci) ([benchmarks](/math/fibonacci/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Fibonacci_number))
* [Prime number](/math/primenumber) ([benchmarks](/math/primenumber/benchmarks.txt)) ([wiki](https://en.wikipedia.org/wiki/Prime_number))
    * Trial division ([wiki](https://en.wikipedia.org/wiki/Trial_division))
    * Sieve of Eratosthenes ([wiki](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes))
    * Sieve of Atkin ([wiki](https://en.wikipedia.org/wiki/Sieve_of_Atkin))

## Run benchmarks

To start benchmarks is needed to use the standard utility `go test`. It allows to execute the following commands:

    go test -run=- -bench=. -benchmem
    go test -run=- -bench=. -benchmem > benchmarks.txt

Script `bench.sh` allows to run tests without visiting the directory with the package:

    ./bench.sh $1 $2

where `$1` is the name of the package, for example `sort/bubblesort`, And `$2` an optional value - if it's not
specified, the result is written to file `benchmarks.txt`. If parameter is set, result of execution is output
to the console. For example:

    ./bench.sh sort/heapsort false
    ./bench.sh sort/heapsort

is equivalent to the following commands:

    go test github.com/ganelon13/go-algorithms/sort/heapsort -run=- -bench=. -benchmem
    go test github.com/ganelon13/go-algorithms/sort/heapsort -run=- -bench=. -benchmem > sort/heapsort/benchmarks.txt

Under the control of the Windows operating system script `bench.sh` can be executed provided that the system has a
unix terminal emulator installed, for example [cygwin](https://www.cygwin.com/). Then the following command format is
available from cmd:

    echo ./bench.sh sort/bubblesort false | bash

## Materials used

    https://youtu.be/eqWzZGNO_XM?list=PLrCZzMib1e9pDxHYzmEzMmnMMUK-dz0_7
    https://rosettacode.org/wiki/Sorting_algorithms