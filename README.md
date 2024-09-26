# cleanarch

## REST Fetch Order

execute the requests under `/api` for creations, pick the id, then fetch by id

## GRPC List Orders

install Evans and run

```
evans --path ./internal/infra/grpc/protofiles --proto order.proto --host localhost --port 50051 repl;
```

then

```
call ListOrders;
```
