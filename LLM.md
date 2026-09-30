# Hanzo S3 Cli

## Overview
Go module: github.com/hanzoai/s3-cli

## Tech Stack
- **Language**: Go

## Build & Run
```bash
go build ./...
go test ./...
```

## Structure
```
s3-cli/
  CONFLICT.md
  CONTRIBUTING.md
  CREDITS
  Dockerfile
  Dockerfile.dev
  Dockerfile.hotfix
  Dockerfile.release
  Dockerfile.release.old_cpu
  LICENSE
  Makefile
  NOTICE
  README.md
  README_zh_CN.md
  _config.yml
  buildscripts/
```

## Key Files
- `README.md` -- Project documentation
- `go.mod` -- Go module definition
- `Makefile` -- Build automation
- `Dockerfile` -- Container build

## Scope
Object and bucket commands only. Server administration (admin, idp, quota, replicate,
batch, license, support) is not part of this client; IAM owns identity. The vendored
CLI framework and helpers live in `pkg/` (cli, console, env, policy, quick, ...).
`TestNoForeignProductNames` (`cmd/names_test.go`) fails if another product's name is
tracked outside LICENSE, NOTICE and CREDITS.
