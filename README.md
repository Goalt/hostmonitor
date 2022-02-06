# Hostmonitor

## Build
It builds docker image with application.
```bash
make build
```

## Run
It runs services.
```bash
make run
```

## All
Clean, build, run.
```bash
make all
```

## Generate (generate swagger.json, stubs)
```bash
make generate
```

1. Код для Grpc сервера
2. Код для http proxy сервера
3. interactor
4. ssh клиент
5. ci/cd lint, test
6. deploy
7. тесты