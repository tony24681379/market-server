# market-server

## How to run

```
$ ./market-server --help
Usage of market-server:
      --alsologtostderr                  log to standard error as well as files (default true)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --port string                      serve port (default "3000")
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs (default 2)
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### Run

```
$ ./market-server
```

### Run with release mode

```
$ export GIN_MODE=release
$ ./market-server
```

### Update vendor directory using [govendor](https://github.com/kardianos/govendor)

```
$ govendor add +external
```

### Get Rt-Mart category

```
GET 127.0.0.1:3000/rt-mart/category
```

### Get Rt-Mart product

```
GET 127.0.0.1:3000/rt-mart/product?category=氣泡礦泉水
```

### Get shopping category

```
GET 127.0.0.1:3000/shopping/category
```

### Get shopping product

```
GET 127.0.0.1:3000/rt-mart/product?category=蘇打餅乾
```
