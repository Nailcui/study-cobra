study cobra
---

### Getting Started

#### build:

```sh
go build
```

#### echo server usage:

[link to echo server code](https://github.com/Nailcui/study-cobra/blob/master/cmd/echo.go)

shou help
```
> ./study-cobra echo -h
鹦鹉学舌:
                你说什么我就回复什么

Usage:
  study-cobra echo [flags]

Flags:
  -h, --help               help for echo
  -l, --lower              转小写输出
  -s, --separator string   自定义分隔符 (default " ")
  -u, --up                 转大写输出
```

hello world
```
> ./study-cobra echo hello world
hello world
```

指定分隔符
```
> ./study-cobra echo -s - hello world
hello-world
```

转大写输出
```
> ./study-cobra echo -u hello world
HELLO WORLD
```

转小写输出
```
> ./study-cobra echo -s - Hello World
hello world
```
