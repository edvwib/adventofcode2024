# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.2 ms ±   0.1 ms    [User: 1.4 ms, System: 1.5 ms]
  Range (min … max):     3.0 ms …   4.1 ms    734 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):      4.404 s ±  0.016 s    [User: 4.413 s, System: 0.205 s]
  Range (min … max):    4.389 s …  4.447 s    10 runs
```

> Apple M2 Pro, 32 GB
