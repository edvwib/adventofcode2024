# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.4 ms ±   0.1 ms    [User: 1.5 ms, System: 1.5 ms]
  Range (min … max):     3.1 ms …   4.1 ms    737 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.2 ms ±   0.1 ms    [User: 1.4 ms, System: 1.5 ms]
  Range (min … max):     3.0 ms …   3.8 ms    782 runs
```

> Apple M2 Pro, 32 GB
