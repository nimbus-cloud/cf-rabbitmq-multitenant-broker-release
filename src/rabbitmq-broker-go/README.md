Create a docker network:
`docker network create test`

Compile the go code for a linux distribution:
`env GOOS=linux GOARCH=386 go build`

Run the binary in a container connected to the docker network
```
docker run -v ${PWD}:/tmp/rabbitmq-broker --rm --name go-broker --net test cflondonservices/london-services-ci-rabbitmq /tmp/rabbitmq-broker/rabbitmq-broker-go
```
