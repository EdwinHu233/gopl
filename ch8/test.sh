#!/usr/bin/env bash

if [[ $1 == "8.1" ]]; then
    cd ./clock
    go build clock.go
    TZ=US/Eastern ./clock -port 8010 &
    TZ=Asian/Tokyo ./clock -port 8020 &
    TZ=Europe/London ./clock -port 8030 &

    cd ../clockwall
    go build clockwall.go
    ./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
fi
