```go
package config

import "os"

type Config struct {
  HTTPAddr    string
  IssuerURL   string // OIDC issuer you proxy to or expose
  SigningKey  string // demo only; use KMS/HSM in prod
}

func Load() Config {
  return Config{
    HTTPAddr:   env("BASIS_HTTP_ADDR", ":8080"),
    IssuerURL:  env("BASIS_ISSUER_URL", "http://localhost:8080"),
    SigningKey: env("BASIS_SIGNING_KEY", "dev-insecure-key"),
  }
}

func env(k, def string) string {
  if v := os.Getenv(k); v != "" { return v }
  return def
}
```

``