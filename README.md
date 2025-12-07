## Saga with Temporal in Go
[Temporal](https://temporal.io/)-orchestrated Saga across three Go microservices with a layout that is *close to production*: separate processes, persistence, HTTP APIs, worker, and tests. 

### Minimalistic stack
`pgx` for Postgres, `goose` for migrations, otherwise standard library + Temporal SDK. No extra frameworks to keep focus on saga flow and avoid distractions.

### Services
- `services/order`: Create order and launche the saga workflow, PostgreSQL persistence, and the Temporal workflow/worker (`internal/infrastructure/workflow/temporal_process_order.go`, worker entry `cmd/temporal_worker/main.go`).
- `services/payment`: Authorize and refund payments over HTTP.
- `services/inventory`: Reserve and unreserve stock over HTTP.

### Infra
- `docker-compose.yml` starts Temporal + UI + Postgres (`7233`, `8080`, `5432`).

### Run locally
1) `docker compose up`
2) `cd services/order && make install-bins run-migrations`
3) `go run ./services/payment/cmd/server`
4) `go run ./services/inventory/cmd/server`
5) `go run ./services/order/cmd/server`
6) `go run ./services/order/cmd/temporal_worker`

Trigger the saga with `test.http` or any HTTP tool:
- POST `/order` to create and start the workflow
- GET `/order/{order_id}/status` to observe progress/compensations

### Tests
`go test ./services/order/internal/infrastructure/workflow -run ProcessOrderWorkflow`
