# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       2.7 ms ±   0.3 ms    [User: 1.1 ms, System: 1.4 ms]
  Range (min … max):     2.2 ms …   3.8 ms    364 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.1 ms ±   0.2 ms    [User: 1.2 ms, System: 1.5 ms]
  Range (min … max):     2.8 ms …   4.3 ms    735 runs
```

> Apple M2 Pro, 32 GB
