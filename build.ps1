# 关闭命令回显，提升美观度
$ErrorActionPreference = "Stop"  # 如果出现错误则立即停止

# 配置变量
$remoteUser = "root"                   # 远程服务器用户名
$remoteHost = "172.16.10.138"          # 远程服务器地址
$remotePath = "/root/api_mock/start"   # 远程服务器存储路径
$localStartPath = ".\logToFile\start"  # 本地拷贝出的可执行程序路径

# 输出信息函数
function Log($message) {
    Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss') - $message"
}

# 1. 构建 Docker 镜像
Log "正在构建 Docker 镜像 api_mock:1.0..."
docker build -t api_mock:1.0 . | Out-Host

if ($LASTEXITCODE -ne 0) {
    Log "Docker 镜像构建失败！"
    exit 1
}

# 2. 创建 Docker 容器（未启动）
Log "正在创建 Docker 容器..."
$containerId = docker create api_mock:1.0

if ($LASTEXITCODE -ne 0 -or -not $containerId) {
    Log "Docker 容器创建失败！"
    exit 1
}

Log "Docker 容器创建成功，容器ID: ${containerId}"

# 3. 将文件从 Docker 容器中复制到本地目录
$sourcePath = "/app/start"

Log "正在从容器 ${containerId} 复制文件 ${sourcePath} 到 ${localStartPath}..."
docker cp "${containerId}:${sourcePath}" $localStartPath

if ($LASTEXITCODE -ne 0) {
    Log "文件复制失败！"
    exit 1
}

Log "文件复制成功！"

# 4. 清理创建的 Docker 容器
Log "正在删除临时容器 ${containerId}..."
docker rm $containerId | Out-Host

if ($LASTEXITCODE -ne 0) {
    Log "删除容器 ${containerId} 失败！"
} else {
    Log "临时容器 ${containerId} 已成功删除。"
}

# 5. 将文件上传到远程服务器
Log "正在将文件 ${localStartPath} 上传到远程服务器 ${remoteHost}:${remotePath}..."
scp $localStartPath "${remoteUser}@${remoteHost}:${remotePath}"

if ($LASTEXITCODE -ne 0) {
    Log "文件上传失败！"
    exit 1
}

Log @"
文件已成功上传至远程服务器 ${remoteHost}:${remotePath}。
脚本执行完成。
"@
