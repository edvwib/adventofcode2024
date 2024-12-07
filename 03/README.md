# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.1 ms ±   0.2 ms    [User: 1.3 ms, System: 1.5 ms]
  Range (min … max):     2.8 ms …   5.2 ms    701 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.3 ms ±   0.1 ms    [User: 1.5 ms, System: 1.5 ms]
  Range (min … max):     3.1 ms …   4.2 ms    905 runs
```

> Apple M2 Pro, 32 GB
