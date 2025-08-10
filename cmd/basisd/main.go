``

```go
package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "basis/internal/server"
  "basis/internal/config"
)

func main() {
  cfg := config.Load()
  mux := server.NewRouter(cfg)

  srv := &http.Server{Addr: cfg.HTTPAddr, Handler: mux}

  go func() {
    log.Printf("BASis listening on %s", cfg.HTTPAddr)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("server error: %v", err)
    }
  }()

  // graceful shutdown
  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  _ = srv.Shutdown(ctx)
}
```

``