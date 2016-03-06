# http_server_benchmarking
This is just a benchmark of various http servers. All servers tested use a basic templating engine to display "Hello, World!" to the client, where "World" is a dynamic variable handled by the template.

The templating engines used are Jade for the Express.js app, [lwan-mustache-c](https://github.com/adam-hanna/lwan-mustache-c) for the h2o server, Blaze for meteor, and the internal templating engine for golang.

I used [wrk](https://github.com/wg/wrk) to run 30s tests.

I tested these in a VMWare virtual machine. Here is some info on the host and guest.

**Host**

* Intel Core i7-3630qm CPU @ 2.40GHz; 2401 Mhz; 4 Cores; 8 logical Processors
* 16GB RAM
* Windows 10

**Guest**

* 1 Core
* 8GB RAM
* Ubuntu 14

## Results

**Express.js**
```
$ wrk -t1 -c400 -d30s http://127.0.0.1:3000/ 
Running 30s test @ http://127.0.0.1:3000/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.88s   116.31ms   2.00s    81.40%
    Req/Sec   169.80    102.71     0.92k    88.66%
  1755 requests in 30.01s, 425.04KB read
  Socket errors: connect 0, read 0, write 0, timeout 1513
Requests/sec:     58.48
Transfer/sec:     14.16KB
```

**H2O**
```
$ wrk -t1 -c400 -d30s http://127.0.0.1:7890/chunked-test/
Running 30s test @ http://127.0.0.1:7890/chunked-test/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   141.16ms  121.63ms   1.23s    93.98%
    Req/Sec     3.20k     1.71k    6.26k    60.62%
  93380 requests in 30.07s, 16.21MB read
Requests/sec:   3105.32
Transfer/sec:    551.92KB
```

**Meteor.js**
```
$ wrk -t1 -c400 -d30s http://127.0.0.1:3000/
  Running 30s test @ http://127.0.0.1:3000/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.73s   203.83ms   2.00s    68.72%
    Req/Sec   139.77    120.98     1.22k    87.66%
  3436 requests in 30.03s, 21.63MB read
  Socket errors: connect 0, read 0, write 0, timeout 2560
Requests/sec:    114.40
Transfer/sec:    737.34KB
```

**golang**
```
$ wrk -t1 -c400 -d30s http://127.0.0.1:8080/ 
  Running 30s test @ http://127.0.0.1:8080/
  1 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.90ms    4.82ms  64.87ms   78.93%
    Req/Sec    45.77k    16.91k   81.11k    49.16%
  1364999 requests in 30.08s, 238.22MB read
  Requests/sec:  45374.55
  Transfer/sec:      7.92MB

```