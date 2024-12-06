# Mezon go sdk

A simple client-server mezon with golang

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

### GStreamer
This example requires you have GStreamer installed, these are the supported platforms

Debian/Ubuntu
```bash
    sudo apt-get install libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev gstreamer1.0-plugins-good
```

Windows MinGW64/MSYS2
```bash
    pacman -S mingw-w64-x86_64-gstreamer mingw-w64-x86_64-gst-libav mingw-w64-x86_64-gst-plugins-good mingw-w64-x86_64-gst-plugins-bad mingw-w64-x86_64-gst-plugins-ugly
```

macOS
```bash
    brew install gst-plugins-good gst-plugins-ugly pkg-config && export PKG_CONFIG_PATH="/usr/local/opt/libffi/lib/pkgconfig"
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
