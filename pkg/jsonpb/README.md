This package is copied from github.com/gogo/protobuf at commit deb6fe8ca7c6d06584bfbd40ca407bf69d9fd2aa.


## Customizations

1. Serialize `int64` and `uint64` to integers, not strings.
    
    - removed [lines 739-742 and 744-746 from jsonpb.go](https://github.com/gogo/protobuf/blob/deb6fe8ca7c6d06584bfbd40ca407bf69d9fd2aa/jsonpb/jsonpb.go#L739-L742)


## Notes

1. If we transition from gogo/protobuf to golang/protobuf, we'll need to replace
this package with the `jsonpb` from github.com/golang/protobuf, and carry over
the changes listed above.
