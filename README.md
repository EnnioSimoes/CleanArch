# Desafio Clean Architecture

Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order) - OK
- Service ListOrders com GRPC - OK
- Query ListOrders GraphQL - OK
  
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.  - OK

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados. - OK
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço. - OK

## Run Project
```
docker-composer up -d
```

```
cd cmd/ordersystem/
go run main.go wire_gen.go
```

## Test Rest API

Use the sample requsts in api/ folder.

## Test GraphQL

Access in browser http://localhost:8080/

Sample querys:
```
mutation createOrder {
  createOrder(input: {id: "ffff", Price: 110.1, Tax: 2.5 }){
    id
    Price
    Tax
    FinalPrice
  }
}
```
```
query listOrders {
  listOrders {
    id
    Price
    Tax
  }
}
```

## Test GRPC

```
evans --proto internal/infra/grpc/protofiles/order.proto --host localhost --port 50051

call CreateOrder
call ListOrders
```



