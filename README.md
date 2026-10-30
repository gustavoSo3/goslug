# GoSlug

## Overview
GoSlug is a Go-based URL redirection service that features administrative controls, flexible storage adapters, and session management.

## Project Roadmap / TODOs

### 1. Routing & Endpoints
- [x] `/` - Serves the welcome page
- [x] `/admin` - Accesses the slug administration dashboard
- [x] `/{slug}` - Redirects to stored destination URL
- [x] `POST /api/redirects` - Creates a new redirect (prevents duplicate slugs)
- [x] `DELETE /api/redirects` - Deletes a redirect referenced by slug

### 2. Request & Input Validation
- [x] Validate all fields (ensures required fields are non-empty) For both front and back end
  - [x] Front
  - [x] Back
- [x] Ensure slugs contain no spaces
- [x] Enforce field length limits:
  - Description: Maximum 100 characters
  - Slug: Max 20 characters
  - Destination: Valid http:// or https:// url
- [ ] Implement Frontend handeling of bad request
  - [ ] login
  - [ ] api/redirects

### 3. Storage Drivers
- [x] Implement persistent storage support
- [x] Design generic Database Interface
- [x] In-memory database driver (volatile, lost on restart)
- [x] Local JSON file storage driver
- [x] Implement storage engine selection via environment variable STORAGE_DRIVER
  - Current options:
    - memory
    - json
  - If ENV not set will default to json
  - If invalid STORAGE_DRIVER option will exit program
  - If selected DB fails will exit program
- [ ] Add remote database driver support (Firestore / MongoDB)

### 4. Authentication & Authorization
- [x] Environment variable configuration for `USERNAME` and `PASSWORD`
  - Default fallback credentials: `admin` / `password`
- [x] Dedicated `/login` endpoint for authentication actions
- [x] Authentication middleware to protect `/admin` and `/api/redirects` routes
- [x] Single admin account flow (initial setup requirement)

### 5. Logging & UI
- [x] Implement structured runtime logging
- [x] Build simple Welcome landing page

---

## Refactoring & Security Enhancements

### Session & Concurrency Refactor
- [ ] Fix single session bottleneck: Replace `AUTH_KEY` single string with a `map[string]time.Time` session store.
- [ ] Prevent data races: Add `sync.Mutex` (`sessionsMu`) around reads/writes to the session map.
  - [x] Also added on local db reads and writes (inMemory and JSON)
- [x] Fix accidental logouts: Remove `credentials.AUTH_KEY = "UNAUTORIZED"` on failed login attempts.
- [ ] Use secure tokens: Replace hashed timestamp with `crypto/rand` token generation (`CreateSession()`).
- [ ] Fix session deletion: Add `RevokeSession(key)` to delete tokens from the map upon logout (`DELETE` request).

### Redirect Array to Map
- [x] Refactor Redirects to use an map, insted of an array, and maybe just get a list of it when asking GetRedirects()

---

## Testing Plan

- [ ] Core test suite setup
- [ ] Input validation tests
- [ ] Endpoint integration tests:
  - [ ] `GET /`
  - [ ] `POST /login`
  - [ ] `GET /admin`
  - [ ] `API /api/redirects`


---

### Extras Ideas:

- [ ] Option to generate random slug
- [ ] Add link access counter
- [ ] Link opening tracking?
- [ ] Single use slug?
- [ ] Time frame slug