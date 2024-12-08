# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.3 ms ±   0.2 ms    [User: 1.4 ms, System: 1.5 ms]
  Range (min … max):     3.0 ms …   4.4 ms    865 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       3.4 ms ±   0.2 ms    [User: 1.6 ms, System: 1.5 ms]
  Range (min … max):     3.1 ms …   4.3 ms    710 runs
```

> Apple M2 Pro, 32 GB
