# deploy

## build

```
    cd <work dir> /path/to/api-mock
    
    win
        go build -o .\logToFile\start.exe .\logToFile\start.go
    
    linux:
        cmd
            set GOOS=linux
            set GOARCH=amd64
            go build -o .\logToFile\start .\logToFile\start.go
        powershell
            $env:GOOS="linux"
            $env:GOARCH="amd64"

```

## migrate

打 zip 包，为 api_mock.zip
只需要 config 目录和 start 可执行程序就行
后续只需要更新可执行程序


## start

进入 /root/api_mock，解压到当前文件夹
```shell
    unzip api_mock.zip
```


检查配置文件: 数据库配置 或 其他可调整设置
```shell
    vi config/local.yaml
    
    数据库配置：dipcc dipcc@2024
```

增加可执行程序 start 权限
```shell
    chmod 777 start
```

启动
```shell
    .\start
```

## 拷贝文件
```shell
    scp root@172.16.10.138:/root/api_mock/public/collection/postman_collection_xxxxx.json <主机目录>
```
