#!/bin/bash

# It is important that all protos in the same namespace be run in a single command
# to avoid overlapping of package variables (primarily file descriptors)
protoc -I. pkg/crowdsound/crowdsound_service.proto pkg/crowdsound/crowdsound_admin_service.proto --go_out=plugins=grpc:.

