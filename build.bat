@echo off

echo Building Windows64
set GOOS=windows
set GOARCH=amd64
go build -o "bin/tiktokdl-win64.exe"

echo Building Windows32
set GOARCH=386
go build -o "bin/tiktokdl-win32.exe"

echo Building Linux64
set GOOS=linux
set GOARCH=amd64
go build -o "bin/tiktokdl-linux64"

echo Building Linux32
set GOARCH=386
go build -o "bin/tiktokdl-linux32"

echo Building LinuxARM
set GOARCH=arm
go build -o "bin/tiktokdl-linuxARM"

echo Building LinuxARM64
set GOARCH=arm64
go build -o "bin/tiktokdl-linuxARM64"

echo Building Mac64
set GOOS=darwin
set GOARCH=amd64
go build -o "bin/tiktokdl-mac64"

echo Building Mac32
set GOARCH=amd64
go build -o "bin/tiktokdl-mac32"

echo Building MacARM64
set GOARCH=arm64
go build -o "bin/tiktokdl-macARM64"