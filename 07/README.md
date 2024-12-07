# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):       6.7 ms ±   0.2 ms    [User: 4.9 ms, System: 1.6 ms]
  Range (min … max):     6.4 ms …   8.0 ms    376 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):     684.3 ms ±   3.1 ms    [User: 673.5 ms, System: 15.8 ms]
  Range (min … max):   681.1 ms … 690.9 ms    10 runs
```

> Apple M2 Pro, 32 GB
