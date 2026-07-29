# Embedded iPXE binaries

This directory is embedded into the booter binary and holds the iPXE boot binaries under
`amd64/` and `arm64/` subdirectories.

The docker build populates it automatically from the pinned iPXE container image (see the
generated Dockerfile). For a native `go build`, populate it first:

```bash
make fetch-source-assets
```

The populated subdirectories are gitignored, and the embedded files are listed explicitly in
the source, so building without them fails at compile time. A truncated or zero-length file
still compiles, and fails at startup when the binaries are patched.
