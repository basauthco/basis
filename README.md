# BASis — Open Source IAM for Building Automation (Repo Scaffolding)

### Purpose

Commercial controls: JIT access with approvals, command allow/deny, change windows, break‑glass, key/cert lifecycle, adaptive auth hooks.

### Tree

```
basic/
├─ cmd/basicd/main.go
├─ internal/
│  ├─ approvals/
│  │  ├─ service.go
│  │  └─ connectors/
│  │     ├─ slack.go
│  │     └─ teams.go
│  ├─ controls/
│  │  ├─ allowdeny.go
│  │  └─ change_windows.go
│  ├─ breakglass/service.go
│  ├─ keys/rotator.go
│  ├─ risk/client.go   # talks to BASim
│  └─ config/config.go
├─ api/openapi.yaml
├─ policy/packs/
│  ├─ bacnet.yaml
│  └─ modbus.yaml
├─ .github/workflows/ci.yml
├─ Dockerfile
└─ README.md
```

### Key Stubs

``

```go
package approvals

type Service interface {
  Request(actor, action, target string, ctx map[string]string) (id string, err error)
  Approve(id string, approver string) error
  Status(id string) (string, error) // pending|approved|denied|expired
}
```

``

```go
package controls

type Rule struct { ActionPattern, TargetPattern string; Effect string }

type Engine struct { rules []Rule }

func (e *Engine) Evaluate(action, target string) string {
  // TODO: pattern match (glob) and return allow/deny
  return "deny"
}
```

``

```yaml
pack: bacnet-baseline
rules:
  - action: "bacnet.read_property"
    effect: allow
  - action: "bacnet.write_property"
    effect: deny
```

---
