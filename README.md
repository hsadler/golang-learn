# golang-learn
Learning Golang

---

### Requirements
- golang
- Docker

---

## Golang Example Scripts

### Run single golang example as script
```sh
go run examples/array.go
```

### Run all golang example scripts
```sh
sh examples/run_examples.sh
```

---

## Learning Golang with Tests

### Run tests for all subdirectories recursively
```sh
cd learn_with_tests/
go test ./...
```


---

## Sample JSON API Backend

### Run app
from root directory
```sh
cd web_app/
sh run_web_app.sh
```

### Run endpoint calls
once app container is running, from root directory
```sh
cd web_app/tests/
sh test_server_endpoints.sh
```

