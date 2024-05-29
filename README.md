# Introduce

<h5> Welcome API Nabitu 1.0, build with golang version 1.22.2 and use extension GORM, migration, fastergoding(hot reload) and mux</h5>

# Run Migrations

```bash
# migrate (push new structure table database)
migrate -database "mysql://nabituapi:nabituapi@tcp(localhost)/nabitu" -path database/migrations up
#migrate rollback (back to structure table database)
migrate -database "mysql://nabituapi:nabituapi@tcp(localhost)/nabitu" -path database/migrations down
```

# Setup ENV

```bash
1. copy .env.example .env
2. define port, credential database and smtp
```

# Run Project

```bash
cd /app
go run main.go

#base url (port api run base on config port of .env)
http://localhost:port
```
