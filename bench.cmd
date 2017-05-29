@echo off
if "%1" EQU "" (
    ECHO package name required!
) else if "%2" EQU "" (
    go test github.com/Ganelon13/go-algorithms/"%1" -run=- -bench=. -benchmem > "%1"/benchmarks.txt
) else (
    go test github.com/Ganelon13/go-algorithms/"%1" -run=- -bench=. -benchmem
)
