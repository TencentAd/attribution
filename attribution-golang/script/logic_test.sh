#!/bin/bash

root=$(cd "$(dirname "$0")" && cd .. && pwd)



cd "${root}" || { exit 1; }
mkdir -p .build
go build -o batch_example "${root}/example/batch/batch_example.go"

cd "${root}/.build" || { exit 1; }
.build/batch_example -click_data_path="${root}/data/click.dat" -conv_data_path="${root}/data/conversion.dat"

