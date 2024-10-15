@echo off
chcp 65001 >nul
REM 启用延迟变量扩展
setlocal enabledelayedexpansion

REM 配置变量
set "REMOTE_USER=root"
set "REMOTE_HOST=172.16.10.138"
set "REMOTE_PATH=/root/api_mock/start"
set "LOCAL_START_PATH=logToFile\start"

echo [INFO] 正在构建 Docker 镜像 api_mock:1.0...
docker build -t api_mock:1.0 .
if %ERRORLEVEL% neq 0 (
    echo [ERROR] Docker 镜像构建失败！
    pause
    exit /b 1
)

echo [INFO] 正在创建 Docker 容器...
for /f "delims=" %%i in ('docker create api_mock:1.0') do set CONTAINER_ID=%%i
if not defined CONTAINER_ID (
    echo [ERROR] Docker 容器创建失败！
    pause
    exit /b 1
)

echo [INFO] Docker 容器创建成功，容器ID: %CONTAINER_ID%

REM 确保目标目录存在
if not exist logToFile (
    mkdir logToFile
)

echo [INFO] 正在从容器 %CONTAINER_ID% 复制文件 /app/start 到 %LOCAL_START_PATH%...
docker cp %CONTAINER_ID%:/app/start %LOCAL_START_PATH%
if %ERRORLEVEL% neq 0 (
    echo [ERROR] 文件复制失败！
    pause
    exit /b 1
)

echo [INFO] 文件复制成功！

echo [INFO] 正在删除临时容器 %CONTAINER_ID%...
docker rm %CONTAINER_ID%
if %ERRORLEVEL% neq 0 (
    echo [ERROR] 删除容器失败！
) else (
    echo [INFO] 临时容器 %CONTAINER_ID% 已成功删除。
)

echo [INFO] 正在将文件 %LOCAL_START_PATH% 上传到远程服务器 %REMOTE_HOST%:%REMOTE_PATH%...
scp %LOCAL_START_PATH% %REMOTE_USER%@%REMOTE_HOST%:%REMOTE_PATH%
if %ERRORLEVEL% neq 0 (
    echo [ERROR] 文件上传失败！
    pause
    exit /b 1
)

echo [INFO] 文件已成功上传至远程服务器 %REMOTE_HOST%:%REMOTE_PATH%。

echo [INFO] 脚本执行完成。按任意键退出...
pause
