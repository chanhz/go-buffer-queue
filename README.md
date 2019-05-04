# go-buffer-queue
A  Go implement of buffer queue, using `sync.Pool`. <br/>

## Benchmarks

Machine: Intel i7 8750H, RAM 8G <br/>
Date: 2019/5/4 <br/>
Go Version: go1.12.1 windows/amd64 <br/>

```
BenchmarkQueue_Put   20000000	   92.6 ns/op
BenchmarkQueue_Pop   500000	      4952 ns/op
```
