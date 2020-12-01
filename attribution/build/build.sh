#!/bin/bash

root=$(cd "$(dirname "$0")" && cd .. && pwd)

cd "${root}" || { exit 1; }

# install attribution
go install -v pkg/modules/attribution/attribution_server.go

# install leads server
go install -v pkg/modules/leads/pull/leads_pull.go
go install -v pkg/modules/leads/receive/leads_receive_server.go

# install impression server
go install -v pkg/modules/impression/impression_server.go

# imp-attribution
go install -v pkg/modules/crypto/crypto_server.go
go install -v pkg/modules/ia/imp_attribution_server.go