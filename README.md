is `http.Server` shutting down?

```
go get github.com/bddjr/shuttingdown
```

```go
var s *http.Server
shuttingdown.IsShuttingDown(s) // true when server is in shutdown
```
