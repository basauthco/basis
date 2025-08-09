# BASis — Open Source IAM for Building Automation (Repo Scaffolding)

---

## Repository Layout

```
BASis/
├─ README.md
├─ LICENSE
├─ CONTRIBUTING.md
├─ CODE_OF_CONDUCT.md
├─ SECURITY.md
├─ GOVERNANCE.md
├─ ROADMAP.md
├─ Makefile
├─ .gitignore
├─ .editorconfig
├─ go.mod
├─ cmd/
│  └─ basisd/
│     └─ main.go
├─ internal/
│  ├─ core/
│  │  ├─ authn/
│  │  │  └─ doc.go
│  │  ├─ authz/
│  │  │  └─ doc.go
│  │  ├─ directory/
│  │  │  └─ doc.go
│  │  └─ config/
│  │     └─ doc.go
│  ├─ gateway/
│  │  ├─ http/
│  │  │  └─ doc.go
│  │  └─ m2m/
│  │     └─ doc.go
│  └─ protocols/
│     ├─ bacnet/
│     │  └─ doc.go
│     ├─ modbus/
│     │  └─ doc.go
│     └─ lonworks/
│        └─ doc.go
├─ pkg/
│  ├─ sdk/
│  │  └─ doc.go
│  └─ telemetry/
│     └─ doc.go
├─ api/
│  ├─ openapi.yaml
│  └─ examples/
│     └─ demo.http
├─ configs/
│  ├─ basis.yaml
│  └─ example.env
├─ deployments/
│  ├─ docker/
│  │  └─ Dockerfile
│  └─ k8s/
│     ├─ namespace.yaml
│     ├─ deployment.yaml
│     ├─ service.yaml
│     └─ ingress.yaml
├─ .github/
│  ├─ ISSUE_TEMPLATE/
│  │  ├─ bug_report.md
│  │  └─ feature_request.md
│  └─ PULL_REQUEST_TEMPLATE.md
└─ docs/
   ├─ ARCHITECTURE.md
   ├─ THREAT_MODEL.md
   └─ CONTRIBUTOR_GUIDE.md
```

---

## README.md

````markdown
# BASis — Open Source IAM for Building Automation

**BASis** is the open-source identity layer for Building Automation Systems (BAS). It provides modern authentication (OIDC/OAuth2), authorization (RBAC/ABAC), and policy controls for operational technology (OT) environments — with adapters for protocols like **BACnet**, **Modbus**, and **LonWorks**.

- **Project (OSS):** BASis — community-driven core
- **Commercial (LTS):** BAS Auth — hardened builds, support, enterprise add‑ons

> _Tagline_: **BASis is the foundation of trust for building automation.**

## Goals

- **Secure access** to BAS HMIs, APIs, and gateways
- **Unify identities** across OT/IT boundaries
- **Gateway pattern** for protocol bridging and audit trails
- **Offline-friendly** with pluggable storage and policy engines

## High-Level Architecture

See [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md). In brief:

- **basisd**: the control-plane & API server
- **Gateway**: HTTP + M2M agents; policy enforcement points (PEPs)
- **Core**: authn/authz, directory, config
- **Protocols**: BACnet/Modbus/LonWorks adapters (opt‑in)

## Quick Start

### Prereqs

- Go 1.22+
- Docker (optional)

### Run in Dev Mode

```bash
make run
# or
go run ./cmd/basisd --dev
```
````

Visit `http://localhost:8080/healthz`.

### Build & Test

```bash
make build
make test
```

## Configuration

See `configs/basis.yaml` for a minimal starter. Environment variables in `configs/example.env`.

## Contributing

We welcome issues and PRs! See [`CONTRIBUTING.md`](CONTRIBUTING.md) and [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md).

## Security & Disclosure

If you find a vulnerability, please follow [`SECURITY.md`](SECURITY.md).

## License

Apache-2.0. See [`LICENSE`](LICENSE). Trademarks are property of their respective owners.

```

```
