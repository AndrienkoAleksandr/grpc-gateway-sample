#!/bin/bash

echo "=> list methods for gRPC"
grpc_cli ls localhost:9090

echo "=> grpc_cli for gRPC servcie"
grpc_cli call localhost:9090  your.service.v1.YourService.Echo 'value: "test"'

echo "=> curl gRPC gateway"
curl -X POST localhost:3001/your.service.v1.YourService/Echo -d '{"value": "echo"}'
echo
