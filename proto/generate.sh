#!/bin/bash

echo "generating with using buf.gen.ws.yaml"
buf generate --template buf.gen.ws.yaml
echo "finish generate with using buf.gen.ws.yaml"

echo "generating with using buf.gen.api.yaml"
buf generate --template buf.gen.api.yaml
echo "finish generate with using buf.gen.api.yaml"