This package is copied from https://github.com/golang/protobuf/commit/4846b58453b3708320bdb524f25cc5a1d9cda4d4


## Customizations

1. Serialize `int64` and `uint64` to integers, not strings.
    
    - removed [lines 547-549 from `encode.go`](https://github.com/golang/protobuf/blob/4846b58453b3708320bdb524f25cc5a1d9cda4d4/jsonpb/encode.go#L547-L549)

2. Removed `decode.go` as we don't need a custom unmarshaller