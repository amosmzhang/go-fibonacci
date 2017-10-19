# go-fibonacci

Various Fibonacci sequence algorithms implemented in go.

```
Benchmark results:

BenchmarkRecursive25-8     	    3000	    483749 ns/op
BenchmarkDynamic25-8       	  500000	      3129 ns/op
BenchmarkMatrix25-8        	 1000000	      1765 ns/op
BenchmarkFastDouble25-8    	20000000	        97.2 ns/op
BenchmarkDynamic100-8      	  100000	     13494 ns/op
BenchmarkMatrix100-8       	  500000	      2563 ns/op
BenchmarkFastDouble100-8   	 3000000	       406 ns/op
BenchmarkMatrix500-8       	  500000	      3661 ns/op
BenchmarkFastDouble500-8   	 1000000	      1921 ns/op
```

Discussion on these algorithms available [here](https://www.nayuki.io/page/fast-fibonacci-algorithms).
