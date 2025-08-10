# BASis — Open Source IAM for Building Automation

### Purpose

Open-source identity core: OIDC provider façade, token service accounts, basic RBAC/ABAC, protocol adapter plugin system, minimal audit events.

### Tree

```
basis/
├─ cmd/
│  └─ basisd/
│     └─ main.go
├─ internal/
│  ├─ server/
│  │  ├─ http.go
│  │  └─ middleware.go
│  ├─ auth/
│  │  ├─ jwt.go
│  │  └─ oidc_proxy.go
│  ├─ policy/
│  │  ├─ engine.go
│  │  └─ model.go
│  ├─ adapters/
│  │  ├─ manager.go
│  │  └─ registry.go
│  ├─ audit/
│  │  └─ sink.go
│  └─ config/
│     └─ config.go
├─ adapters/
│  ├─ reference-http-proxy/
│  │  ├─ main.go
│  │  └─ Dockerfile
│  └─ sim-bacnet/
│     └─ main.go
├─ api/
│  ├─ openapi.yaml
│  └─ policy-examples/
│     └─ sample-policy.yaml
├─ proto/
│  └─ adapter.proto
├─ pkg/
│  └─ sdk/
│     └─ client.go
├─ .github/
│  └─ workflows/
│     └─ ci.yml
├─ .gitignore
├─ CODE_OF_CONDUCT.md
├─ CODEOWNERS
├─ CONTRIBUTING.md
├─ LICENSE
├─ SECURITY.md
└─ README.md
```
