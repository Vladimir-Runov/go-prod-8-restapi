@echo off

@echo 'hello - page not exist, should be 404'
curl http://localhost:8088/hello
timeout /t 1 > nul

@echo 'about - page '
curl http://localhost:8088/about
timeout /t 1 > nul

@echo 'list of tasks'
curl -X GET http://localhost:8088/tasks
timeout /t 1 > nul

@echo '+ new task'
curl -X POST -H "Content-Type: application/json" -d "{\"title\": \"new task name\"}" http://localhost:8088/tasks

@echo 'list of tasks after addition'
curl -X GET http://localhost:8088/tasks
timeout /t 1 > nul

@echo 'get task [2]'
curl -X GET http://localhost:8088/tasks/2
timeout /t 1 > nul

@echo 'get task [21] - should be error msg!'
curl -X GET http://localhost:8088/tasks/21
timeout /t 1 > nul

@echo 'change  task [2] - done = true'
curl -X PUT -H "Content-Type: application/json" -d "{\"title\": \"task 2 completed!!\", \"done\": true}" http://localhost:8088/tasks/2

@echo 'get task [2]'
curl -X GET http://localhost:8088/tasks/2
timeout /t 1 > nul

curl http://localhost:8088/exit

