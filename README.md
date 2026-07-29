# booter

A tool to easily boot Talos machines using PXE.

## Usage

Run `booter` in a container on the host network:

```bash
docker run --rm --network host \
  ghcr.io/siderolabs/booter:v0.1.0
```

Then, power on machines in the **same subnet** as `booter` with **PXE boot** enabled (both UEFI and legacy BIOS are supported).  
Recommended boot order: **disk first, then network**.

To connect the machines to **Omni**, go to the **Omni Overview** page, click **“Copy Kernel Parameters”**, and run `booter` with the copied arguments:

```bash
docker run --rm --network host \
  ghcr.io/siderolabs/booter:v0.1.0 \
  <KERNEL_ARGS>
```

To see more options:

```bash
docker run --rm --network host ghcr.io/siderolabs/booter:v0.1.0 --help
```

## Development

`booter` also builds and runs natively without docker.
On a fresh clone, fetch the iPXE binaries which get embedded into the binary once, then regular Go commands work:

```bash
make fetch-source-assets
go build ./...
```
