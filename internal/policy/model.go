```go
package policy

type Input struct {
  Actor   string            `json:"actor"`
  Action  string            `json:"action"`
  Target  string            `json:"target"`
  Context map[string]string `json:"context"`
}

type Decision struct {
  Effect      string   `json:"effect"` // allow|deny
  Obligations []string `json:"obligations"`
}
```

``