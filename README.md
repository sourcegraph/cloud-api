# Cloud API

[![pipeline](https://github.com/sourcegraph/cloud-api/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/sourcegraph/cloud-api/actions/workflows/pipeline.yaml)

This repo contains the API specification for the proof-of-concept of Sourcegraph Cloud API.
See [RFC 765](https://docs.google.com/document/d/13SnXtM5Jpi2PMfpT6rHCkJFpXuwnCJudgvKyjGhAAZo/edit#heading=h.trqab8y0kufp) for context.
The Cloud API server implementation is being developed in [`sourcegraph/controller/cmd/apiserver`](https://github.com/sourcegraph/controller/tree/main/cmd/apiserver).

This API is only used internally at Sourcegraph - the specification is public for ease of use.

## Contents

- `cloudapi/v1/`: The Cloud API v1 proto specification.
- `cmd/`: Example programs and utilities.
- `go/cloudapi`: Generated Go bindings for the Cloud API.

## Development

To update or add endpoints, edit [`cloudapi/v1/cloudapi.proto`](./cloudapi/v1/cloudapi.proto).

Code generation is configured in [`buf.yaml`](buf.yaml) and [`buf.gen.yaml`](buf.gen.yaml). To update or regenerate the language bindings:

```sh
make generate
```

Note that this will also install tools into `./bin` - to remove them for a reinstall, run:

```sh
make clean
```

See `cmd/example-client` for a simple example of a client connecting to an implementation of the Cloud API.
