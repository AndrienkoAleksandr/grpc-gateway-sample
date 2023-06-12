#!/bin/bash

grpc_cli ls localhost:9090

grpc_cli call localhost:9090  your.service.v1.YourService.Echo 'value: "test"'

