
# Start
```shell
go mod download
```

User Service
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

Video Service
```shell
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```

Gateway
```shell
cd cmd/api
go run .
```
启动顺序：user/video Service > Gateway

使用上面的命令只能在 Linux 环境下运行，windows 环境下应顺序执行里面的 main.go 文件
