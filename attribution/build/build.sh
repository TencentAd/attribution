#!/bin/bash

root=$(cd "$(dirname "$0")" && cd .. && pwd)

# install attribution
cd "${root}" || { exit 1; }
go install -v pkg/modules/attribution/attribution_server.go

# install leads server
cd "${root}" || { exit 1; }
go install -v pkg/modules/leads/pull/leads_pull.go
go install -v pkg/modules/leads/receive/leads_receive_server.go

# install impression attribution server