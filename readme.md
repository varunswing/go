# How I Learned Go

## Resources Used

### ✅ Completed
- [Go Tour - Basics](https://go.dev/tour/welcome/1)
- [GeeksforGeeks Go Tutorials](https://www.geeksforgeeks.org/go/)
- [Youtube - Coder's Gyan](https://www.youtube.com/playlist?list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6)

### 🚧 In Progress / To-Do
- [ ] Rob Pike's YouTube Talks on Go
- [ ] [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [ ] [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

## Learning Journey
I started learning Go by following the official Go Tour, which helped me grasp the basics like variables, functions, control flow, and concurrency. GeeksforGeeks provided additional explanations and coding examples to reinforce my understanding.

### Next Steps
- Explore Generics in Go to understand type parameterization.
- Watch Rob Pike's talks to get deeper insights into Go's design philosophy.
- Read and follow best practices from the Go Code Review Comments and Uber Go Style Guide to write clean, maintainable code.

## Notes & Takeaways

### 1. Fundamentals of Go
- Go is statically typed and compiled.
- Go has garbage collection, but efficient memory management is key.
- Simplicity and readability are core principles of the language.

### 2. Key Language Features
- **Concurrency with Goroutines**: Lightweight threads managed by the Go runtime.
- **Channels**: Used for safe communication between Goroutines.
- **Interfaces**: Structural typing rather than nominal typing.
- **Slices vs Arrays**: Slices are more powerful and flexible than arrays.
- **Defer, Panic, and Recover**: Exception handling mechanism.

### 3. Best Practices
- Use `go fmt` for formatting.
- Write tests using Go’s built-in `testing` package.
- Handle errors explicitly and gracefully.
- Keep interfaces small and focused.
- Use struct embedding for composition instead of inheritance.
- Leverage Go’s standard library before relying on third-party packages.

### 4. Performance Optimizations
- Use Goroutines efficiently instead of OS threads.
- Minimize garbage collection impact by managing memory wisely.
- Optimize slice and map allocations to reduce reallocation overhead.
- Use sync.Pool for high-performance object reuse.

### 5. Tooling & Ecosystem
- `go run`, `go build`, `go test`, and `go mod` for package management.
- Profiling tools: `pprof` for CPU & memory profiling.
- `golangci-lint` for linting and static analysis.

## Next Steps
- Dive deeper into **microservices** using Go.
- Learn more about **Go in production**: monitoring, logging, and tracing.
- Contribute to **open-source Go projects**.
- Explore **Go frameworks** like Gin (for web apps) and gRPC (for APIs).

---
This README will evolve as I progress in my Go learning journey. 🚀
