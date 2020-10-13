#!/bin/bash

function check_port_listened() {
    port=$(($1 + 0))
    bash -c "lsof -i:${port} | grep LISTEN"
    return $?
}

function wait_part_listened() {
    port=${1:-80}
    second=${2:-30}

    for ((i = 0; i < second; i++)); do
        if check_port_listened "${port}"; then
            return 0
        else
            sleep 1
        fi
    done

    return 1
}
