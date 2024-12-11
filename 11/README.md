# Part 1
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):      4.104 s ±  0.095 s    [User: 4.033 s, System: 0.048 s]
  Range (min … max):    3.952 s …  4.202 s    10 runs
```

# Part 2
```
hyperfine -N --warmup 5 ./main

Benchmark 1: ./main
  Time (mean ± σ):      17.4 ms ±   0.6 ms    [User: 14.3 ms, System: 3.3 ms]
  Range (min … max):    16.2 ms …  19.7 ms    160 runs
```

> Apple M2 Pro, 32 GB
