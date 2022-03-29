### Development environment

```
docker build -f dev.Dockerfile -t lidani/go-votes:develop .

docker run --rm -it -v $pwd\:/app -p 8080:9090 lidani/go-votes:develop
```

### Build environment

```
docker build -t lidani/go-votes:latest .

docker run --rm -d --name go-votes -p 8080:9090 lidani/go-votes:latest

docker container logs go-votes -f
```
