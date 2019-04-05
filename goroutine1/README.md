#go best practices

- Don't create go routine in libraries, let the consumer control concurrency
- When you create a go routine, know how it will end. Avoid subtle memory leaks
- Check for race condition at compile time <br />
  $ go run -race main.go
