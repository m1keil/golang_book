results were tested with "perf":


original mandelbrot: 336ms 
goroutines:
2: 210ms
4: 242ms
8: 125ms
16: 115ms
32: 107ms
64: 104ms
128: 103ms
256: 103ms
512: 103ms


anything beyond 16-32 goroutines seem to have deminishing returns on my 2 core HT processor
