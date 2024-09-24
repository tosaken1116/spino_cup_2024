#!/bin/bash

echo "generating with using buf.gen.api.yaml"
buf generate --template buf.gen.api.yaml
echo "finish generate with using buf.gen.api.yaml"