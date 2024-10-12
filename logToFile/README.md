# deploy

## build

```
    cd <work dir> /path/to/api-mock
    
    win
        go build -o .\logToFile\start.exe .\logToFile\start.go
    
    linux:
        set GOOS=linux
        set GOARCH=amd64
        go build -o .\logToFile\start.exe .\logToFile\start.go

```

## migrate

打 zip 包，为 api_mock.zip
只需要 config 目录和 start.exe 就行


## start

解压到 api_mock 文件夹
```shell
    unzip -d /path/to/api_mock api_mock.zip
```


检查配置文件
```shell
    vi api_mock/logToFile/config/local.yaml
```

启动

```shell
    cd ./logToFile
    .\start.exe
```
