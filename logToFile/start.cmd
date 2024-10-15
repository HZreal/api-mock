@echo off
REM 配置变量
set "REMOTE_USER=root"  REM 远程服务器用户名
set "REMOTE_HOST=172.16.10.138"  REM 远程服务器地址
set "REMOTE_HOME=/root/api_mock"  REM 远程工作目录
set "LOCAL_DEST_DIR=%UserProfile%\Desktop"  REM 本地存储路径

echo 正在连接远程服务器并执行程序...
ssh %REMOTE_USER%@%REMOTE_HOST% "cd %REMOTE_HOME% && ./start"

REM 查找远程目录中最新的文件
for /f "delims=" %%i in ('ssh %REMOTE_USER%@%REMOTE_HOST% "ls -t %REMOTE_HOME%/public/collection | head -n 1"') do (
    set "LATEST_FILE=%%i"
)

if "%LATEST_FILE%"=="" (
    echo 未找到任何文件，操作失败。
    exit /b 1
)

echo 正在复制文件: %LATEST_FILE% 到 %LOCAL_DEST_DIR% ...
scp %REMOTE_USER%@%REMOTE_HOST%:%REMOTE_HOME%/public/collection/%LATEST_FILE% %LOCAL_DEST_DIR%

if %errorlevel% equ 0 (
    echo 文件已成功复制到 %LOCAL_DEST_DIR%
) else (
    echo 文件复制失败。
    exit /b 1
)
