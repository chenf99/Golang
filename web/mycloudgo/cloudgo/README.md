# CloudGo
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

# 任务要求
1. 编程 web 服务程序 类似 cloudgo 应用。
    - 要求有详细的注释
    - 是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
2. 使用 curl 测试，将测试结果写入 README.md
3. 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

# 框架说明
本实验使用了Go内置的net/http框架和negroni库。
negroni是一个GoLang的http中间件库，它定义了中间件的框架和风格，我们可以基于它开发出我们自己的中间件，并且可以集成到Negroni中。

negroni兼容原生的http.Handler,我们可以把自己的http.Handler加入到negroni的中间件链中，negroni会自动调用他们来处理HTTP Request。

negroni的特点是非常小，不复杂，又十分优雅地设计了中间件调用链

# curl测试

## 测试命令：
`$ curl -v http://localhost:8080/`

## 返回结果:
```bash
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.58.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 14 Nov 2018 09:53:44 GMT
< Content-Length: 49
< Content-Type: text/plain; charset=utf-8
< 
Welcome to the home page!
This is a test server.
* Connection #0 to host localhost left intact

```

服务器输出:
```bash
[negroni] listening on :8080
[negroni] 2018-11-14T17:53:44+08:00 | 200 |      127.726µs | localhost:8080 | GET /
```
# ab测试

## 测试命令：
`$ ab -n 1000 -c 100 http://localhost:8080/`
- -n请求数量
- -c并发数量

## 返回结果
```bash
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        49 bytes

Concurrency Level:      100
Time taken for tests:   0.207 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      166000 bytes
HTML transferred:       49000 bytes
Requests per second:    4835.94 [#/sec] (mean)
Time per request:       20.678 [ms] (mean)
Time per request:       0.207 [ms] (mean, across all concurrent requests)
Transfer rate:          783.95 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   2.2      1      12
Processing:     0   18  12.4     20      55
Waiting:        0   18  12.3     20      55
Total:          0   20  12.7     20      56

Percentage of the requests served within a certain time (ms)
  50%     20
  66%     25
  75%     28
  80%     29
  90%     37
  95%     45
  98%     54
  99%     55
 100%     56 (longest request)

```

## 参数解释
