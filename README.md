# CALC
This is sample calc Project.
It Access the gRPC server (written in Go) in 4 different ways
1. Using go client
2. in go test
3. python client
4. API

## Usage:
```
Start the server

CALC> go run .\server.go
Starting server
Server listening at 127.0.0.1:8080
2022/01/10 21:41:39 server.go:41: Serving gRPC on 127.0.0.1:8080      
2022/01/10 21:41:39 server.go:71: Serving gRPC-Gateway on http://0.0.0.0:8090
2022/01/10 21:42:25 server.go:99: runtime error: integer divide by zero                                                                     .0:8090   
2022/01/10 21:45:24 server.go:99: runtime error: integer divide by zeroo                                


Now you can perform CALC operations as -

(1)
CALC> go run .\client.go       
Starting client on 127.0.0.1:8080
2022/01/11 00:01:14 Add 100 20 = 120
2022/01/11 00:01:14 Sub 20 10 = 10
2022/01/11 00:01:14 Mul 10 20 = 200
2022/01/11 00:01:14 Div 20 10 = 2

(2)
CALC> go test .\calc_test.go -v
=== RUN   TestAdd
Starting client on 127.0.0.1:8080
Address 0xc00010bf10
--- PASS: TestAdd (0.01s)
=== RUN   TestSub
Address 0xc00010bf10
--- PASS: TestSub (0.00s)
=== RUN   TestMul
Address 0xc00010bf10
--- PASS: TestMul (0.00s)
=== RUN   TestDiv
Address 0xc00010bf10
--- PASS: TestDiv (0.00s)
PASS
ok      command-line-arguments  0.708s

(3)
CALC>python client.py
Start service: 2022-01-11 00:05:14.148799 on localhost:8080

Performing Operations on CalcService:
a: 10
b: 20
 Add: result: 30

a: 100
b: 20
 Sub: result: 80

a: 10
b: 20
 Mul: result: 200

a: 100
b: 20
 Div: result: 5


(4)

$ curl -X POST -k http://localhost:8090/v1/calc/add -d '{"a": 10, "b": 20}'
{"result":"30"}

$ curl -X POST -k http://localhost:8090/v1/calc/sub -d '{"a": 10, "b": 20}'
{"result":"-10"}

$ curl -X POST -k http://localhost:8090/v1/calc/mul -d '{"a": 100, "b": 20}'
{"result":"2000"}

$ curl -X POST -k http://localhost:8090/v1/calc/div -d '{"a": 10, "b": 20}'
{"result":"0"}

$ curl -X POST -k http://localhost:8090/v1/calc/div -d '{"a": 100, "b": 20}'
{"result":"5"}
```
