```go
package server

import (
  "encoding/json"
  "net/http"
  "basis/internal/config"
)

type router struct { cfg config.Config }

func NewRouter(cfg config.Config) http.Handler {
  r := &router{cfg: cfg}
  mux := http.NewServeMux()

  mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
  })

  mux.HandleFunc("/v1/token/issue", r.issueToken)
  mux.HandleFunc("/v1/policy/evaluate", r.evaluatePolicy)

  return mux
}

func (r *router) issueToken(w http.ResponseWriter, req *http.Request) {
  // TODO: validate client, issue JWT (service account / user)
  json.NewEncoder(w).Encode(map[string]string{"token": "demo.jwt.token"})
}

func (r *router) evaluatePolicy(w http.ResponseWriter, req *http.Request) {
  // TODO: call policy engine with input; return allow/deny + obligations
  json.NewEncoder(w).Encode(map[string]any{"result": "allow", "obligations": []string{}})
}
```

``