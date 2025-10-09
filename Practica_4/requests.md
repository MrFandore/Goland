Invoke-WebRequest -Uri http://localhost:8080/api/tasks `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"title":"Выучить chi"}'

Invoke-WebRequest -Uri http://localhost:8080/api/tasks `
-Method GET

Invoke-WebRequest -Uri http://localhost:8080/api/tasks/1 `
-Method GET

Invoke-WebRequest -Uri http://localhost:8080/api/tasks/1 `
-Method PUT `
-Headers @{"Content-Type"="application/json"} `
-Body '{"title":"Выучить chi глубже","done":true}'

Invoke-WebRequest -Uri http://localhost:8080/api/tasks/1 `
-Method DELETE

Invoke-WebRequest -Uri http://localhost:8080/api/tasks `
-Method GET | Select-Object -ExpandProperty Content
