#!/bin/bash

BASEDIR=$(dirname "$0")
cd "${BASEDIR}"/../ || exit

PROTO_LOCATION=routeguide/client/$1
PROTO_DEST=pb

mkdir -p "${PROTO_LOCATION}"
mkdir -p "${PROTO_DEST}"

protoc --go_out="${PROTO_DEST}" --go_opt=paths=source_relative \
  --go-grpc_out="${PROTO_DEST}" --go-grpc_opt=paths=source_relative \
  "${PROTO_LOCATION}/"*.proto

for f in "$PROTO_DEST"/routeguide/client/$1/*; do
  PROTO_SERVICE_DEST=${PROTO_DEST}/client/$1
  mkdir -p "${PROTO_SERVICE_DEST}"
  cp "$f" "${PROTO_SERVICE_DEST}"
done

rm -rf "$PROTO_DEST"/routeguide
