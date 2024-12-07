# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.0 ms ±   0.1 ms    [User: 1.2 ms, System: 1.5 ms]
  Range (min … max):     2.7 ms …   4.1 ms    737 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.3 ms ±   0.1 ms    [User: 1.5 ms, System: 1.5 ms]
  Range (min … max):     3.1 ms …   4.0 ms    742 runs
```

> Apple M2 Pro, 32 GB
