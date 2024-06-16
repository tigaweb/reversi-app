#!/bin/bash

docker compose exec -it backend go run migrate/migrate.go
