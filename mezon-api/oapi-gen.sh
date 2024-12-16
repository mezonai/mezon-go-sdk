#!/bin/bash

# swagger generate client -f "../../mezon-server/apigrpc/apigrpc.swagger.json" -A "Mezon"

swagger-codegen generate -i "../../mezon-server/apigrpc/apigrpc.swagger.json" -l go -o .
