Invoke-WebRequest -Uri http://localhost:8080/tasks `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"title":"Купить молоко"}'

curl.exe -i http://localhost:8080/tasks
curl.exe -i http://localhost:8080/tasks/1