@echo off
echo : Generate src/tool/const.go
python .\generate_const.py
cd ..\src

echo : Build
go build
if not ERRORLEVEL 0 exit

echo : Delete old executable
del ..\testing\ximfect.exe

echo : Move new executable
move .\ximfect.exe ..\testing > NUL
cd ..\testing
