# Профильное задание на Go-разработчика в VK Tech

## Запуск
```bash
docker-compose up --build app
```
* Работает скрипт по ожиданию инициализации базы данных при развертывании в докере. Просьба дождаться выполнения:
```
{"time":"2024-03-18T10:10:28.092126727Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"1","attempt":"200"}
{"time":"2024-03-18T10:10:29.09318591Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"2","attempt":"200"}
{"time":"2024-03-18T10:10:30.09421533Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"3","attempt":"200"}
{"time":"2024-03-18T10:10:31.095390433Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"4","attempt":"200"}
{"time":"2024-03-18T10:10:32.09637744Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"5","attempt":"200"}
{"time":"2024-03-18T10:10:33.097711905Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"6","attempt":"200"}
{"time":"2024-03-18T10:10:34.099060299Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"7","attempt":"200"}
{"time":"2024-03-18T10:10:35.100488265Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"8","attempt":"200"}
{"time":"2024-03-18T10:10:36.101815444Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"9","attempt":"200"}
{"time":"2024-03-18T10:10:37.102580435Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"10","attempt":"200"}
{"time":"2024-03-18T10:10:38.103867799Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"11","attempt":"200"}
{"time":"2024-03-18T10:10:39.105010637Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"12","attempt":"200"}
{"time":"2024-03-18T10:10:40.106179856Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"13","attempt":"200"}
{"time":"2024-03-18T10:10:41.107630625Z","level":"INFO","msg":"Wait for Postgres - waiting","attempt":"14","attempt":"200"}
```
```
{"time":"2024-03-18T10:10:42.159786814Z","level":"INFO","msg":"PostgreSQL is up - executing command"}
{"time":"2024-03-18T10:10:42.331774266Z","level":"INFO","msg":"Server start listening on port: ","port":"8080"}
```

## Навигация
* [Requests примеры](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/blob/main/requests.http)

* [OpenAPI спецификация](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/blob/main/specification.yaml)

* [JWT реализация](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/blob/main/auth/jwt/jwt.go)

* [Конфигурационный файл](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/blob/main/config.json)

* [Unit-тесты](https://github.com/ColdDirol/GoDeveloper-TestTask-VK/tree/main/tests)

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
