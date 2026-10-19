# Roadmap Belajar Go (Backend/Web API)

Level: fresh Go, background bahasa lain (JS/PHP/Java).
Toolchain: go1.26.1. Target: backend REST API production-ready.
Estimasi: 6-10 minggu @ 5-10 jam/minggu. Kerjakan berurutan; tiap fase ada latihan wajib.

Cara pakai: tiap fase punya folder latihan. Isi kode di situ. Centang `- [x]` setelah selesai.

---

## Fase 1 - Refresher (`01-refresher/`)

Ulang cepat, sudah bisa bahasa lain. Fokus kebede Go dari bahasa lain.

- [x] Variabel, zero value, `:=` vs `var`, konstanta
- [x] Control flow: `for` (satu-satunya loop), `if` init statement, `switch`
- [x] Composite: array vs slice, map, struct, `append`, `make`, `len`/`cap`
- [x] Pointer: kapan perlu, kapan tidak, receiver value vs pointer
- [x] Error handling: `error`, `errors.Is`/`errors.As`, custom error type
- [x] Fungsi: multiple return, named return, defer, closure

Latihan wajib:
- [x] CLI word counter: baca file/stdin, hitung jumlah kata & baris (mirip `wc`)

## Fase 2 - Idiomatic Go (`02-idiomatic/`)

Gaya nulis Go yang benar. Pembeda coder Go macet vs jalan.

- [ ] Methods & receiver (value vs pointer receiver)
- [ ] Interface: implicit, small interface, `io.Reader`/`io.Writer` pattern
- [ ] Composition over inheritance, embedding
- [ ] Generics dasar: type parameter, constraint
- [ ] Effective Go + Go idioms: error wrapping, sentinel error, fmt.Errorf `%w`

Latihan wajib:
- [ ] Buat `Stack[T]` generik + pake di program kecil

## Fase 3 - Stdlib inti (`03-stdlib/`)

Bekal sebelum nyentuh framework. Go API production kebanyakan pakai ini.

- [ ] `net/http`: jadi server, routing, `mux`, handler signature
- [ ] `encoding/json`: marshal/unmarshal, struct tag, omitempty, custom marshal
- [ ] `context`: timeout, cancellasi, passed down ke handler/IO
- [ ] `time`: duration, timer, ticker, formatting/layout
- [ ] `io` & `bufio`: reader/writer interface, streaming
- [ ] Testing: table-driven test, `t.Run`, benchmark, `-race`

Latihan wajib:
- [ ] Strukt JSON API: `/health`, `/echo` (POST, balikin JSON), pakai stdlib only + test

## Fase 4 - REST API + DB (`04-rest-api/`)

Inti tujuan belajar. API nyata dengan database.

- [ ] Router: chi atau stdlib mux, middleware chain
- [ ] CRUD resource: handler -> service -> repo, DTO
- [ ] `database/sql` + `pgx` driver, connection pool, query param (hindari SQLi)
- [ ] Migrasi: `golang-migrate` atau embed SQL
- [ ] Validasi input: manual atau validator lib
- [ ] Error mapping: DB error -> HTTP status (404, 409, 500)

Latihan wajib (project utama):
- [ ] REST API "todo/book/library": CRUD + PostgreSQL + migrasi + test integration

## Fase 5 - Production basics (`05-production/`)

Bikin API layak jalan di server beneran.

- [ ] Auth: JWT (middleware, refresh token), atau session
- [ ] Password hashing: bcrypt/argon2
- [ ] Middleware: logging (`slog`), recover panic, request ID, CORS
- [ ] Config: env-based (`os.Getenv` atau env lib), graceful shutdown
- [ ] Observability: structured log, health/readiness endpoint, basic metrics
- [ ] Deploy: Dockerfile multi-stage, `docker compose` (app + Postgres)

Latihan wajib:
- [ ] Tambah auth + Docker ke project Fase 4

## Fase 6 - Concurrency (`06-concurrency/`)

Nilai jual utama Go. Terapkan ke API.

- [ ] Goroutine: launch, sync.WaitGroup, lifecycle
- [ ] Channel: buffered/unbuffered, select, close, nil channel gotcha
- [ ] Pattern: worker pool, fan-in/fan-out, pipeline, generator
- [ ] `sync`: Mutex vs RWMutex, atomic, Once, race detector
- [ ] Context-based cancellation di pipeline
- [ ] Rate limiter, timeout pattern di HTTP client

Latihan wajib:
- [ ] Worker pool: scrape/download N URL paralel dgn batas concurrency + graceful stop

---

## Sumber belajar

- [Effective Go](https://go.dev/doc/effective_go) - wajib baca Fase 1-2
- [Go by Example](https://gobyexample.com) - cepat, per topik
- [A Tour of Go](https://go.dev/tour) - refresh interaktif
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) - alternatif TDD, cocok Fase 3-4
- [Go Web Examples](https://gowebexamples.com) - Fase 3-5
- [100 Go Mistakes and How to Avoid Them](https://100go.co) - Fase 5-6
- [Let's Go Further](https://lets-go-further.alexedwards.net) - Fase 4-5 (paid, opsional)

## Kriteria selesai

- Semua fase checkbox centang, project Fase 4/5 jalan di Docker end-to-end
- Bisa jelasin kenapa pakai interface, pointer receiver, channel pattern di kode sendiri
- `-race` bersih di semua test