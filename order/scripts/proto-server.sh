#!/bin/bash

BASEDIR=$(dirname "$0")
cd "${BASEDIR}"/../ || exit

PROTO_LOCATION=routeguide/server
PROTO_DEST=pb

mkdir -p "${PROTO_LOCATION}"
mkdir -p "${PROTO_DEST}"

protoc --go_out="${PROTO_DEST}" --go_opt=paths=source_relative \
  --go-grpc_out="${PROTO_DEST}" --go-grpc_opt=paths=source_relative \
  "${PROTO_LOCATION}/$1.proto"

for f in "$PROTO_DEST"/routeguide/server/*; do
  PROTO_SERVICE_DEST=${PROTO_DEST}/server/$1
  mkdir -p "${PROTO_SERVICE_DEST}"
  cp "$f" "${PROTO_SERVICE_DEST}"
done

rm -rf "$PROTO_DEST"/routeguide
