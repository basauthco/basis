```go
package policy

// Engine is a placeholder; later swap with OPA/Rego or Cedar.

type Engine struct{}

func New() *Engine { return &Engine{} }

func (e *Engine) Evaluate(in Input) Decision {
  // default‑deny example: allow read, deny write unless window or role etc.
  if in.Action == "read" { return Decision{Effect: "allow"} }
  return Decision{Effect: "deny"}
}
```

``