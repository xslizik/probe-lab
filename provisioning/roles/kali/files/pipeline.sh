#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: [-r <p_rate>] [-a set for UDP] [-o <output>] [-p <ports>] [IP range in \"\"]"
    exit 1
fi

RATE=1000   
FLAG_A=false
NAME="results"
PORTS=""

while getopts ":r:a:o:p:" opt; do
  case ${opt} in
    r )
      RATE=$OPTARG
      ;;
    a ) 
      FLAG_A=true
      ;;
    o ) 
      NAME=$OPTARG   
      ;;
    p ) 
      PORTS=$OPTARG  
      ;;
    \? ) 
      echo "Invalid option: -$OPTARG" 1>&2
      exit 1
      ;;
  esac  
done    

shift $((OPTIND -1))
timestamp=$(date +"%Y-%m-%d-%H:%M:%S")

IP_RANGE="$1"     
ZMAP_OUTPUT="/tmp/webmap/zmap_${NAME}${timestamp}.txt"    
NMAP_OUTPUT_XML="/tmp/webmap/nmap_${NAME}${timestamp}.xml"

echo "Running zmap -q --probe-module=icmp_echoscan --target-port=0 -r $RATE -o $ZMAP_OUTPUT $IP_RANGE"
zmap -q --probe-module=icmp_echoscan --target-port=0 -r "$RATE" -o "$ZMAP_OUTPUT" "$IP_RANGE"

if [ "$FLAG_A" = true ]; then
    echo "Running nmap -T4 -sC -sSV -sUV -iL $ZMAP_OUTPUT $PORTS -Pn -n -oX $NMAP_OUTPUT_XML"
    nmap -T4 -sC -sSV -sUV -iL "$ZMAP_OUTPUT" "$PORTS" -Pn -n -O -oX "$NMAP_OUTPUT_XML"
else 
    echo "Running nmap -T4 -sC -sSV -iL $ZMAP_OUTPUT $PORTS -Pn -n -oX $NMAP_OUTPUT_XML"
    nmap -T4 -sC -sSV -iL "$ZMAP_OUTPUT" "$PORTS" -Pn -n -O -oX "$NMAP_OUTPUT_XML"
fi