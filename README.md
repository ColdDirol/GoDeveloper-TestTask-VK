# Профильное задание на Go-разработчика в VK Tech

## Запуск
```bash
docker-compose up --build app
```

## Навигация
* [Requests example](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/blob/main/requests.http)

* [OpenAPI specification](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/blob/main/specification.yaml)

* [JWT logic](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/blob/main/auth/jwt/jwt.go)

* [Configuration file](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/blob/main/config.json)

* [Init-test](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/tree/main/tests)

## Правила конфигурации
### Конфигурационный файл
```json
{
  "env": "dev",
  "http_server": {
    "host": "0.0.0.0",
    "port": "8080"
  },
  "database": {
    "host": "db",
    "port": "5432",
    "username": "postgres",
    "password": "postgres",
    "db_name": "postgres",
    "ssl_mode": "disable"
  },
  "auth": {
    "secret_key": "secret_key",
    "salt": "salt"
  }
}
```

#### ENV:
В переменную env можно установить значения: `local`, `dev`, `prod`.
```json
"env": "dev"
```
От значения зависит уровень и способ логирования: `local` - уровень логирования - Debug, `dev` - уровень логирования - Debug, `prod` - уровень огирования - Info.

#### HTTP_SERVER:
```json
"http_server": {
  "host": "0.0.0.0",
  "port": "8080"
}
```
`host` - хост адреса, на котором поднимается сервер (`localhost` для локального запуска, `0.0.0.0` для запуска в докере)

`port` - порт адреса, на котором поднимается сервер

#### DATABASE:
```json
"database": {
  "host": "db",
  "port": "5432",
  "username": "postgres",
  "password": "postgres",
  "db_name": "postgres",
  "ssl_mode": "disable"
}
```
`host` - хост, на котором расположена база данных (`localhost` для локального запуска, `db` для запуска в докере)

`post` - порт базы данных

`username`, `password`, `db_name` - креды от базы данных

`ssl_mode` - протокол шифрования: `disable` - не используется

#### AUTH:
```json
"auth": {
    "secret_key": "secret_key",
    "salt": "salt"
}
```
`secret_key` - secret key для составления jwt токена

`salt` - солб при шифровании пароля в *SHA256*
