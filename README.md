# -GolangMiddleWare
Middleware Development Implementation #Golang
<br> <b> We are using Gin-Gonic </b>

![image](https://user-images.githubusercontent.com/40759486/187882367-5535fa4d-74b6-41d2-aada-b3876b698b67.png)



Gin-Gonic is a light but high-performance HTTP-Web-Framework.<br> Look another HTTP freamworks: https://deepsource.io/blog/go-web-frameworks/

<br>
- It offers Recover and Log support on the middleware side.
- Of course we can write and add our own middleware component.
- We can group routes so versioning is easy
- for the project to be easy a MongoDB service will be used. But later I'm tired of that too. A docker image will be used directly
 
 
 ```
 #Our home folder and necessary go files are created
>> cd project_dir
>> touch main.go
```

for management of gin-gonic and other modules
```
#same dir necessary
>> go mod init book-worm-api
```

Installing the required go packages for gin-gonic and mongodb
```
>> go get -u github.com/gin-gonic/gin go.mongodb.org/mongo-driver
```

for MongoDB entagration
```
>> touch quote.go
```


for the CRUD Operation
```
>> touch quotedata.go
```


Documentation of annotation declarations in service methods with Swagger 2.0
```
>> go get -u github.com/swaggo/swag/cmd/swag
```

to produce documents Swagger 2.0
```
>> swag init _ "-GolangMiddleWare/docs"
```

many times Swag has error! 
```
export PATH=$(go env GOPATH)/bin:$PAT
```

also:
```
export PATH="/usr/bin:$PATH"
```

because your machine doesn't know Swagger...

---------- <b> ---------------- </b> ---------

Then

DOCKER SIDE
Running mongodb docker container and creating database
```
>> sudo docker run --name mongodb -e MONGO_INITDB_ROOT_USERNAME=scoth -e MONGO_INITDB_ROOT_PASSWORD=tiger -e MONGO_INITDB_DATABASE=bookworms -p 27017:27017 -d mongo:latest
```



Then then then; we are looking the code...


