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

## Mocs generation
```bash
mockgen -source=main.go -destination=mocks/mocks_main.go
```