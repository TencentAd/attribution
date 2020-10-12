#!/bin/bash

root=$(cd "$(dirname "$0")" && cd .. && pwd)

cd "${root}" || { exit 1; }
mkdir -p .build
cd .build || { exit 1; }
go build -o batch_example "${root}/example/batch/batch_example.go" || { echo "failed to build batch_example"; exit 1; }

./batch_example -click_data_path="${root}/data/click.dat" -conversion_data_path="${root}/data/conversion.dat"
