# Introduce

<h5> Welcome API Nabitu 1.0, build with golang version 1.22.2 and use extension GORM, migration, fastergoding(hot reload) and mux</h5>

# Migrate Run

```bash
migrate -database "ConnectionDB" -path database/migrations up
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
