# GOLANG

## projects
#### hello -- basic go operations
#### simple-api -- api that returns hello world (understanding net/http library)  
go build   
./simple_api    
go test -cover -v  

#### Postgres connector
docker pull postgres  
docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres  
go run main.go  
