#!/bin/bash

echo "=> list methods for gRPC"
grpc_cli ls localhost:50051

echo "=> grpc_cli for gRPC servcie"
grpc_cli call localhost:50051  your.service.v1.YourService.Echo 'value: "test"'

echo "=> curl gRPC gateway"
curl -X POST localhost:3001/your.service.v1.YourService/Echo -d '{"value": "echo"}'
echo

curl -X POST -k -v https://grpc-gateway-sample-oandriie-tenant.apps.ci-ln-qpxtgrt-72292.gcp-2.ci.openshift.org/your.service.v1.YourService/Echo -d '{"value": "echo"}'
