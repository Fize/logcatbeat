#!/bin/sh

################################################
# start logcatbeat process.
# eg. ./logcatbeat -c logcatbeat.yml
#            -E seccomp.enabled=false
################################################

set -e

export CONFIG_FILE="/system/etc/logcatbeat.yml"

log() {
    echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $@"
}

file_exist() {
    if [ ! -f "${CONFIG_FILE}" ]; then
        log "can't start logcatbeat process, I will sleep."
        while :
        do
            sleep 3600
        done
    fi
}

start() {
    file_exist
    log "running process..."
    $1 -c ${CONFIG_FILE} -E seccomp.enabled=false -e
}

main() {
    while
        [ $# -gt 0 ]
    do
        key="$1"
        case $key in
            --path)
                export BIN=$2
                shift
                ;;
        esac
    shift
    done
    start ${BIN}
}

main "$@"
