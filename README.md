# cleanarch

## Docker Dependencies

to start running mysql and rabbitmq, run

```
docker-compose up -d
```

## Initialize the application

```
go mod tidy;
go run cmd/main.go cmd/wire_gen.go
```

## REST Get/Fetch Order

execute the requests under `/api` for a few creations, pick the id, then fetch by id

## GRPC List Orders

install Evans and run

```
evans --path ./internal/infra/grpc/protofiles --proto order.proto --host localhost --port 50051 repl;
```

then

```
call ListOrders;
```

## GraphQL

open localhost:8080/playground and run the following query:

```
query {
  ListOrders{
    Data{
      id,
      Price,
      Tax,
      FinalPrice,
    }
  }
}
```
