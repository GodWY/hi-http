#!/bin/bash

protoc --plugin=protoc-gen-hip=/Users/wangxinyu/go/bin/protoc-gen-hip  --plugin=protoc-gen-go=/Users/wangxinyu/go/bin/protoc-gen-go --proto_path=proto --hip_out=proto --go_out=proto greeter/greeter.proto 