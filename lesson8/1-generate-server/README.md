# Демонстрация работы генератора кода HTTP-сервера из OpenAPI-спецификации

Необходима утилита [go-swagger](https://goswagger.io/).

Для начала формируем OpenAPI-спецификацию.

Перед генерацией инициализируйте go-модуль

```sh
go mod init generate-server
```

Потом запускаем генератор

```sh
swagger generate server -A todo-list -f ./spec.yml
```

После генерации не забыть подтянуть зависимости

```sh
go mod tidy
```

После чего просто запускаем приложение

```sh
go run ./cmd/todo-list-server --port=8000
```
