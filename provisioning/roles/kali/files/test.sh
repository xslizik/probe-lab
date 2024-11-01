#!/bin/bash

coma_ips="10.255.255.251,10.244.170.185,10.86.46.58,10.8.3.19,10.8.3.17,10.0.0.1"
#coma_ips="10.0.0.1"
ips=(10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.8.3.17 10.0.0.1)
#ips=(10.0.0.1)

run_fast_port_scan() {
    echo -e "\n### Fast port scan common ports / 1-1024"
    echo "#### Running: nmap -sSV -Pn -n -T4 ${ips[@]}"
    time nmap -sSV -Pn -n -T4 "${ips[@]}"

    echo -e "\n#### Running: masscan ${ips[@]} -p1-1024 -i eth1 --rate 10000"
    time masscan "${ips[@]}" -p1-1024 -i eth1 --rate 10000

    echo -e "\n#### Running: /opt/nscan -r ${ips[@]}"
    time /opt/nscan -r "$coma_ips"
}

run_complex_port_scan() {
    echo -e "\n### Complex port scan"
    echo "#### Running: nmap -sSV -p- -Pn -n -T4 ${ips[@]}"
    time nmap -sSV -p- -Pn -n -T4 "${ips[@]}"

    echo -e "\n#### Running: masscan ${ips[@]} --banners -p1-65535 -i eth1 --rate 10000"
    time masscan "${ips[@]}" --banners -p1-65535 -i eth1 --rate 10000

    echo -e "\n#### Running: /opt/nscan -r ${ips[@]} -p 0-65535 -b"
    time /opt/nscan -r "$coma_ips" -p 0-65535 -b
}

run_fast_host_discovery() {
    echo -e "\n### Fast host discovery"
    echo "#### Running: masscan -i eth1 10.0.0.0/8 --ping --rate 10000"
    time masscan -i eth1 10.0.0.0/8 --ping --rate 10000

    echo -e "\n#### Running: zmap -q -p 0 --probe-module=icmp_echoscan -i eth1 -r 10000 -G C4:AD:34:58:98:08 10.0.0.0/8"
    time zmap -q -p 0 --probe-module=icmp_echoscan -i eth1 -r 10000 -G C4:AD:34:58:98:08 10.0.0.0/8

    echo -e "\n#### Running: /opt/nscan -r 10.0.0.0/8 -d"
    time /opt/nscan -r 10.0.0.0/8 -d

    echo -e "\n#### Running: nmap -sn -PE 10.0.0.0/8 -n -T5 --min-parallelism 100 --max-rtt-timeout 100ms --max-retries 1"
    time nmap -sn -PE 10.0.0.0/8 -n -T5 --min-parallelism 100 --max-rtt-timeout 100ms --max-retries 1
}

run_service_discovery() {
    echo -e "\n### Host discovery of a specific service"
    echo "#### Running: zmap -q -p 22 -i eth1 -G C4:AD:34:58:98:08 10.0.0.0/8 -r 10000"
    time zmap -q -p 22 -i eth1 -G C4:AD:34:58:98:08 10.0.0.0/8 -r 10000

    echo -e "\n#### Running: masscan -i eth1 10.0.0.0/8 -p22 --rate 10000"
    time masscan -i eth1 10.0.0.0/8 -p22 --rate 10000
}

echo "Select which test to run:"
echo "1) Fast Port Scan"
echo "2) Complex Port Scan"
echo "3) Fast Host Discovery"
echo "4) Service Discovery"
echo "5) Run All"
read -rp "Enter your choice [1-5]: " choice

case $choice in
    1) run_fast_port_scan ;;
    2) run_complex_port_scan ;;
    3) run_fast_host_discovery ;;
    4) run_service_discovery ;;
    5) 
        run_fast_port_scan
        run_complex_port_scan
        run_fast_host_discovery
        run_service_discovery
        ;;
    *) echo "Invalid choice!" ;;
esac
