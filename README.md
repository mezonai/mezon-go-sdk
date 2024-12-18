# Mezon go sdk

🚀 A simple client-server mezon with golang

## Status

✅ Http client

✅ Socket

❌ WIP Webrtc streaming

❌ WIP Webrtc voice call

## Install

### protoc-gen-go

```shell
    go install \
        "google.golang.org/protobuf/cmd/protoc-gen-go" \
        "google.golang.org/grpc/cmd/protoc-gen-go-grpc" \
        "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway" \
        "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
```

### swagger-codegen
<https://github.com/swagger-api/swagger-codegen>

```bash
    sudo -i
    wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.8/swagger-codegen-cli-2.4.8.jar  -O /usr/local/bin/swagger-codegen-cli.jar
    echo '#!/bin/bash\njava -jar /usr/local/bin/swagger-codegen-cli.jar "$@"' > /usr/local/bin/swagger-codegen
    chmod +x /usr/local/bin/swagger-codegen
```

## Generate code

```bash
    cd mezon-protobuf
    ./protoc-gen.sh
```

```bash
    cd mezon-api
    ./oapi-gen.sh
```
