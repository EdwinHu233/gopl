#!/usr/bin/env bash

declare -a child_pid

trap clean_up INT

function clean_up() {
    for pid in "${child_pid[@]}"; do
        kill $pid
    done
}

if [[ $1 == "8.1" ]]; then
    cd ./clock
    go build -o ../build/clock
    cd ../clockwall
    go build -o ../build/clockwall
    cd ..

    TZ=US/Eastern ./build/clock -port 8010 &
    child_pid=("${child_pid[@]}" "$!")
    TZ=Asian/Tokyo ./build/clock -port 8020 &
    child_pid=("${child_pid[@]}" "$!")
    TZ=Europe/London ./build/clock -port 8030 &
    child_pid=("${child_pid[@]}" "$!")
    ./build/clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
    child_pid=("${child_pid[@]}" "$!")
elif [[ $1 == "8.2" ]]; then
    cd ./netcat
    go build -o ../build/netcat
    cd ../reverb
    go build -o ../build/reverb
    cd ..

    ./build/reverb &
    child_pid=("${child_pid[@]}" "$!")
    ./build/netcat
    child_pid=("${child_pid[@]}" "$!")
fi
