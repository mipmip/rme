#!/usr/bin/env bash
#(C)2019-2022 Pim Snel - https://github.com/mipmip/RUNME.sh
CMDS=();DESC=();NARGS=$#;ARG1=$1;make_command(){ CMDS+=($1);DESC+=("$2");};usage(){ printf "\nUsage: %s [command]\n\nCommands:\n" $0;line="              ";for((i=0;i<=$(( ${#CMDS[*]} -1));i++));do printf "  %s %s ${DESC[$i]}\n" ${CMDS[$i]} "${line:${#CMDS[$i]}}";done;echo;};runme(){ if test $NARGS -eq 1;then eval "$ARG1"||usage;else usage;fi;}

##### PLACE YOUR COMMANDS BELOW #####

make_command "build" "Build rme binary with version injection"
build(){
  go build -ldflags "-s -w -X main.version=0.1.0" -o rme .
}

make_command "run_tests" "Run Go test suite"
run_tests(){
  go test ./...
}

make_command "fmt" "Format Go source files"
fmt(){
  gofmt -w .
}

make_command "vet" "Run static analysis"
vet(){
  go vet ./...
}

##### PLACE YOUR COMMANDS ABOVE #####

runme
