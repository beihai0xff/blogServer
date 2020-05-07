@echo off
:loop
@echo off&amp;color 0A
cls
echo,
echo 请选择要编译的系统环境：
echo,
echo 1. Windows_amd64
echo 2. Linux_amd64
echo 3. Darwin_amd64
echo 4. All
echo 0. Quit
echo,
::清空release目录...
rmdir /s build /Q

set/p action=请选择目标平台:
if %action% == 1 goto build_Windows_amd64
if %action% == 2 goto build_Linux_amd64
if %action% == 4 goto build_Darwin_amd64
if %action% == 4 goto all
if %action% == 0 goto end
cls &amp; goto :loop

:build_Windows_amd64
echo 编译 Windows 64 位版本
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -v -o build/blogServer.exe -tags=jsoniter .
goto end

:build_Linux_amd64
echo 编译 Linux 64 位版本
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -v -o build/blogServer -tags=jsoniter .
goto end

:build_Darwin_amd64
echo 编译 Darwin 64 位版本
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -v -o build/blogServer -tags=jsoniter .
goto end

:all
SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=windows
go build -v -o build/blogServer.exe -tags=jsoniter .
SET GOOS=linux
go build -v -o build/blogServer -tags=jsoniter .
SET GOOS=darwin
go build -v -o build/blogServer -tags=jsoniter .
goto end

:end