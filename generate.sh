#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

# go run greet/greet_server/server.go