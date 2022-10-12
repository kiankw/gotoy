#!/bin/bash  

for value in `seq 1 3 100`
    do  
        let begin=$value
        let end=$value+2
        go run os/args.go $begin $end
    done