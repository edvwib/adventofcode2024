# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.1 ms ±   0.1 ms    [User: 1.3 ms, System: 1.5 ms]
  Range (min … max):     2.9 ms …   3.9 ms    799 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.8 ms ±   0.1 ms    [User: 2.0 ms, System: 1.6 ms]
  Range (min … max):     3.6 ms …   4.6 ms    646 runs
```

> Apple M2 Pro, 32 GB
