# Go-Journey 🚀

A personal learning repository collecting Go (Golang) experiments, demos, and small projects. Many examples were developed and tested on Linux, but most run on other platforms (macOS, Windows/WSL) as well.

---

## About

This repo catalogs learning exercises across Go fundamentals, concurrency, web development, CLIs, databases, and algorithms. Each subfolder is typically an independent mini-project or demo with its own go.mod in many cases.

---

## Selected Project Structure

Top-level folders (selected):

- API/ — REST API projects and examples (Blogging_Platform_API, Movie-Reservation-System, Ecommerce api, Caching_Proxy, image_processing_service)
- webApp/ — Web applications and related experiments (Eventify_v2, UrlShortner, crowdcity, Markdown_Notetaking App, websockets, forms, OAuth demos)
- algorithms/ — Algorithm exercises and interview-style problems
- concepts/ — Language concepts, examples and tests (goroutines, channels, interfaces, generics, protobufs)
- command_line/ — CLI utilities and small tools (Task Tracker, ExpenseTracker, file-type changer, todo)
- concurrency/ — Scheduler and concurrency-focused examples
- worker_pool/ — Worker pool demos
- database/ — SQL examples and tutorials (sqlc, others)
- Discord_Bot/ — Discord bot example(s)
- email/ — Email/SMTP demos
- Personal_Budget_Tracker/ and PersonalBudjet_App/ — Budget tracking projects
- coding interview/ — Interview-style project(s) and helpers
- test/ — Throwaway experiments and playgrounds

This is not exhaustive; browse folders to see more examples and tests.

---

## Notable runnable examples

Many subfolders contain a `main.go` with `package main`. To run an example:

```bash
cd path/to/subproject
# run the main file or module
go run .
# or
go run main.go
```

Examples include (but are not limited to):
- worker_pool (examples in /worker_pool)
- webApp/UrlShortner
- webApp/Eventify_v2
- API/Blogging_Platform_API (cmd/blogApp)
- Discord_Bot/hello_world
- command_line/ExpenseTracker
- algorithms/* (many folders with main.go)

If a project lacks a go.mod and you need modules, initialize: `go mod init example.com/name && go mod tidy`.

---

## How this repo is organized (recommended path)

1. Start in `first/` for initial starter examples
2. Read `concepts/` for foundational topics (types, interfaces, concurrency)
3. Explore `algorithms/` for problem-solving practice
4. Try `command_line/` and `worker_pool/` for runnable tools
5. Inspect `webApp/` and `API/` folders for full-stack/backend examples

---

## Development notes

- Go version: Examples were developed across Go 1.18+; Go 1.21+ is recommended where newer language features or module behavior are used.
- Many folders include their own go.mod; prefer running commands inside the subfolder.
- Some folders are experimental or incomplete — check comments and README files inside subfolders.

---

## Contributing

This is primarily a personal learning repository. Feel free to fork, adapt examples, or open issues/PRs if you have improvements or fixes. Please format Go code with `gofmt`/`go fmt` and add tests where applicable.

---

## License

This repository is open for learning and sharing. Add a LICENSE file (MIT, Apache 2.0, etc.) in the root if you want to set an explicit license.

---

## Contact

Author: [CrimsonKarma44](https://github.com/CrimsonKarma44)

Happy coding! 🎉

