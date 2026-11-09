# GoSlug

GoSlug is a lightweight, high-performance [Go](https://go.dev/)-based URL redirection service featuring administrative controls, flexible storage adapters, and secure session management.

<p align="center">
  <img src="./website/public/favicon.svg" width="200" alt="GoSlug Logo">
</p>

## Features

### Core Routing & Endpoints
* **Fast Redirection:** Instantly maps custom slugs (`/{slug}`) to target URLs.
* **Admin Dashboard:** Secure web console (`/admin`) for managing links.
* **REST API:** Endpoints to create, list, and delete redirects programmatically.

### Security & Validation
* **Strict Input Rules:** Enforces length constraints (max 20 chars for slugs, 100 for descriptions) and validates valid HTTP/HTTPS destinations.
* **Robust Authentication:** Protected routes via secure session tokens powered by `crypto/rand` and thread-safe concurrency controls (`sync.Mutex`).
* **Environment Config:** Configurable admin credentials and storage drivers via environment variables.

### Flexible Storage Adapters
* **Plug-and-Play Drivers:** Choose your storage engine using the `STORAGE_DRIVER` environment variable.
* **Local JSON File Driver:** Persistent, zero-setup storage (default).
* **In-Memory Driver:** Volatile, high-performance storage for testing or ephemeral use.
* *Coming soon: Remote database support (Firestore / MongoDB).*

### Developer-Friendly
* **Lightweight & Zero-Dependency:** Built entirely in Go with clean architecture and no heavy framework bloat.
* **Structured Logging:** Clear runtime logs for tracking requests and system status.


## Requirements

To run or build GoSlug, you need **either**:

* **Go 1.27+** (for local binary compilation and execution)
* **Docker** (to build container images and run without installing Go locally)

## Installation

### Opt 1. Build to source code

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/goslug.git
   cd goslug
   ```

2. Download and verify dependencies:
   ```bash
   go mod download
   ```

3. Build the binary executable:
   ```bash
   go build -o bin/goslug ./cmd/goslug
   ```

---

### Opt 2. Build Container Image (Using Docker)

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/goslug.git
   cd goslug
   ```

2. Build the Docker image:
   ```bash
   docker build -t goslug:latest .
   ```

## Execution

### 1. Environment Configuration

GoSlug supports configuration through environment variables:

| Variable | Default Value | Description |
| :--- | :--- | :--- |
| `PORT` | `8080` | Port for the HTTP server to listen on |
| `USERNAME` | `admin` | Admin username for dashboard access |
| `PASSWORD` | `password` | Admin password for dashboard access |
| `STORAGE_DRIVER` | `json` | Storage engine choice (`json` or `memory`) |

### 2.1. Running Locally

Set environment variables and launch the compiled application:

```bash
# Set environment variables
export USERNAME="myadmin"
export PASSWORD="supersecretpassword"
export STORAGE_DRIVER="json"
export PORT="8080"

# Run the application
./bin/goslug
```

Alternatively, run directly with `go run`:

```bash
USERNAME="admin" PASSWORD="password" STORAGE_DRIVER="memory" go run ./cmd/goslug
```

#### Other alternative: run the systemd service

The repo includes a goslug.service file you can use to run the complied aplication as a service. 

Just make sure you change the {user} and environment variables inside it.

Also set WorkingDirectory & ExecStart to the correct paths


### 2.2. Running with Docker

Run a container instance using your built image:

```bash
mkdir -p data
docker run -d \
  --name goslug \
  -p 8080:8080 \
  -e HOST=0.0.0.0 \
  -e USERNAME="admin" \
  -e PASSWORD="password" \
  -e STORAGE_DRIVER="json" \
  -v $(pwd)/data:/app/data \
  goslug:latest
```

You can run the dockerrun.sh script in the deployment folder too.

For docker make sure to specify the HOST=0.0.0.0 in order to be accesible from outside the container since the default goes to "localhost"

### 3. Verification

Once running, verify the server is operating by visiting:

* **Welcome Page:** `http://localhost:8080/`
* **Admin Dashboard:** `http://localhost:8080/admin`
* **Redirect Example:** `http://localhost:8080/{your-slug}`


## Recomendations

If you want to run the server on an 80 or 443 port you can use the reverse poxy service from [caddy](https://github.com/caddyserver/caddy)



## Testing Plan

- [ ] Core test suite setup
- [ ] Input validation tests
- [ ] Endpoint integration tests:
  - [ ] `GET /`
  - [ ] `POST /login`
  - [ ] `GET /admin`
  - [ ] `API /api/redirects`

## Extras Ideas:

- [ ] Option to generate random slug
- [ ] Add link access counter
- [ ] Link opening tracking?
- [ ] Single use slug?
- [ ] Time frame slug
- [ ] Tagging and group by tagging
- [ ] Multi account setup
- [ ] Add remote database driver support (Firestore / MongoDB)
