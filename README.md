# gin-hello-world
Con Go/Golang iniciamos un 'hello world' con [Gin Web framework](https://github.com/gin-gonic/gin#installation).

## Modo de ejecutar. 
Ingresa a la carpeta base del proyecto y ejecuta
```bash 
go mod download
```

Y ejecuta el proyecto como prueba
```bash 
go run src/main.go
```

El proyecto se iniciará en el puerto 8080, puedes revisarlo en la siguiente URL [http://localhost:8080/hello-world](http://localhost:8080/hello-world) y se desplegará un mensaje en JSON.
```bash 
{
    "message": "Hello world"
}
```