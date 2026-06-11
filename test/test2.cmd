@echo off
for /l %%i in (1,1,5) do (
    curl http://localhost:8088/hello
    timeout /t 1 > nul
)