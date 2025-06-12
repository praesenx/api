### About The API
Every time you visit my blog or portfolio, behind the scenes an API written in Go is hard at work delivering content,
handling requests, and powering dynamic features. The **oullin/api** repository is that engine. It bundles all the application’s
core logic, data access, and configuration into a clean, maintainable service—making it the indispensable backbone of the “Ollin” experience.

### Go-based Modular Core

* **Go Modules**
    * Managed via `go.mod` and `go.sum`, ensuring every dependency is versioned and reproducible.
* **Entry Point**
    * `main.go` initialises configuration, middleware (logging, CORS), and route registration to start the HTTP server.

### Layered, Clean Architecture
A clear separation of concerns keeps the codebase organized and easy to extend.

| Layer                   | Folder      | Responsibility                                                                        |
|-------------------------|-------------|---------------------------------------------------------------------------------------|
| **Routing & HTTP**      | `pkg/`      | Parse incoming HTTP requests, invoke services, return JSON.                           |
| **Business Logic**      | `handler/`  | Core rules and workflows (e.g. assembling blog post payloads, input validation).      |
| **Data Access**         | `models/`   | Interfaces and implementations for database operations (CRUD for posts, users, etc.). |
| **Schema & Migrations** | `database/` | SQL migration files that define tables, indexes, and seed data.                       |

### Configuration & Environment

* **Central Config**
    * A `config/` directory plus an `.env.example` file keep settings like database URLs and feature flags in one place.
* **Live Reloading**
    * An `.air.toml` file enables fast development reloads with [Air](https://github.com/cosmtrek/air), so code changes appear immediately.

## Containerization & DevOps

* **Docker Compose**

    * A `docker-compose.yml` file brings up the API alongside its dependencies (PostgreSQL, Redis), guaranteeing consistency across environments.
* **Makefile Automation**

    * Includes commands for building, testing, linting, and running migrations—streamlining both local workflows and CI pipelines.

### Why This Is the Core of The Site

Every dynamic feature on my site—from listing recent blog posts to processing contact-form submissions—calls into this API.
Because it cleanly layers routing, business logic, and data access, I can:

1. **Extend** easily with new endpoints (e.g., search, comments) by adding a handler, service, and storage method.
2. **Maintain** confidently, since each part of the code lives in its own folder with a clear purpose.
3. **Deploy** without surprises, thanks to Docker Compose and versioned Go modules.

In short, **oullin/api** isn’t just another code repository—it’s the beating heart of my blog and portfolio. It translates
every user action into data operations and returns precisely what the frontend needs. Feel free to explore the folders,
run it locally via Docker Compose, and contribute improvements with pull requests.

---
> This is where the mindful movement of “Ollin” truly comes alive, one request at a time.
