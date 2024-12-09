# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):      20.2 ms ±   0.5 ms    [User: 17.4 ms, System: 2.4 ms]
  Range (min … max):    19.3 ms …  22.1 ms    118 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):      61.3 ms ±   0.5 ms    [User: 58.5 ms, System: 2.4 ms]
  Range (min … max):    60.5 ms …  62.7 ms    45 runs
```

> Apple M2 Pro, 32 GB
