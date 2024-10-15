@echo off
chcp 65001 >nul

REM 配置变量
set "REMOTE_USER=root"  REM 远程服务器用户名
set "REMOTE_HOST=172.16.10.138"  REM 远程服务器地址
set "REMOTE_HOME=/root/api_mock"  REM 远程工作目录
set "LOCAL_DEST_DIR=%UserProfile%\Desktop"  REM 本地存储路径

REM 解析传入的关键字参数
:parse_args
if "%~1"=="" goto done
for /f "tokens=1,2 delims==" %%A in ("%~1") do (
    if /i "%%A"=="REMOTE_USER" set "REMOTE_USER=%%B"
    if /i "%%A"=="REMOTE_HOST" set "REMOTE_HOST=%%B"
    if /i "%%A"=="REMOTE_HOME" set "REMOTE_HOME=%%B"
    if /i "%%A"=="LOCAL_DEST_DIR" set "LOCAL_DEST_DIR=%%B"
)
shift
goto parse_args

:done
echo 使用的配置：
echo 用户名: %REMOTE_USER%
echo 主机: %REMOTE_HOST%
echo 远程工作目录: %REMOTE_HOME%
echo 本地下载目录: %LOCAL_DEST_DIR%


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

echo 正在复制远程文件: %LATEST_FILE% 到本机 %LOCAL_DEST_DIR% ...
scp %REMOTE_USER%@%REMOTE_HOST%:%REMOTE_HOME%/public/collection/%LATEST_FILE% %LOCAL_DEST_DIR%

if %errorlevel% equ 0 (
    echo 文件已成功复制到 %LOCAL_DEST_DIR%
) else (
    echo 文件复制失败。
    exit /b 1
)
