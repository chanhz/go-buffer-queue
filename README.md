# go-buffer-queue
A  Go implement of buffer queue, using `channel`. <br/>

## Benchmarks

Machine: Intel i7 8750H, RAM 8G <br/>
Date: 2019/5/1 <br/>
Go Version: go1.12.1 windows/amd64 <br/>

```
BenchmarkQueue_Put   500000	      4739  ns/op
BenchmarkQueue_Pop   300000	      4956 ns/op
```
