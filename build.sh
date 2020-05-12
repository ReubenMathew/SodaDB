#!/bin/bash

echo "Building SodaDB..."
cd src
`go build .`
mv sodadb.exe ../build

echo "SodaDB build completed!"
echo "See 'build' directory for executable"
