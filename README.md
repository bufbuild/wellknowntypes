# Well-Known Types

This repository contains all the Well-Known Types from the [protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf) repository
for every stable version of Protobuf that contains them. This is generally all versions >=3.0.0 except for 3.4.1.

## Development

- `make printversions`: Print all available versions in SemVer order.
- `make populate`: Download the Well-Known Types and populate this repository.
- `make breaking`: Check for breaking changes between the versions.
- `make`: `make populate && make breaking`
