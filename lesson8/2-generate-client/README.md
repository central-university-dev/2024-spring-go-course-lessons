# Демонстрация работы генератора кода HTTP-клиента из OpenAPI-спецификации

Необходима утилита [go-swagger](https://goswagger.io/).

Для начала формируем OpenAPI-спецификацию.

Перед генерацией инициализируйте go-модуль

```sh
go mod init generate-client
```

Потом запускаем генератор

```sh
swagger generate client -A todo-list -f ./spec.yml
```

После генерации не забыть подтянуть зависимости

```sh
go mod tidy
```

После чего используем готовый код в проекте.

```sh
mkdir -p ./cmd/todo-list-client
touch ./cmd/todo-list-client/main.go
# наполняем main.go
```
