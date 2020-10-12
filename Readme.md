## project begin
```
export GO111MODULE=on
export GOPROXY=https://goproxy.cn

## go mod init
go get
go mod tidy
```

### keep coding

2020-05-26 参考：https://www.bilibili.com/video/BV1CE411H7bQ?p=7

    1. register login userinfo
    2. jwt token
    3. responseFormat
    4. read config from yml

2020-03-13 参考：https://eddycjy.com/posts/go/gin/2018-02-11-api-01/


## docker-compose
```
### go build && docker build
sh build.sh

### docker-compose
docker-compose -f .docker-compose.yaml up -d
```