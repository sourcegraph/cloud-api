# Cloud API

This repo contains the API specification for the proof-of-concept of Sourcegraph Cloud API.
See [RFC 765](https://docs.google.com/document/d/13SnXtM5Jpi2PMfpT6rHCkJFpXuwnCJudgvKyjGhAAZo/edit#heading=h.trqab8y0kufp) for context.

This API is only used internally at Sourcegraph - the specification is public for ease of use.

## Contents

- `cloudapi/v1/`: The Cloud API v1 proto specification.
- `cmd/`: Example programs and utilities.
- `go/cloudapi`: Generated Go bindings for the Cloud API.
