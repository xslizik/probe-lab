Select which test to run:
1) Fast Port Scan
2) Complex Port Scan
3) Fast Host Discovery
4) Service Discovery
5) Run All
Enter your choice [1-5]: 5

### Fast port scan common ports / 1-1024
#### Running: nmap -sSV -Pn -n -T4 10.0.0.1
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-10-19 23:13 CEST
Nmap scan report for 10.0.0.1
Host is up (0.0033s latency).
Not shown: 994 closed tcp ports (reset)
PORT     STATE SERVICE        VERSION
21/tcp   open  ftp            MikroTik router ftpd 6.49.17
22/tcp   open  ssh            MikroTik RouterOS sshd (protocol 2.0)
53/tcp   open  domain         (generic dns response: NOTIMP)
80/tcp   open  http           MikroTik router config httpd
2000/tcp open  bandwidth-test MikroTik bandwidth-test server
8291/tcp open  unknown
1 service unrecognized despite returning data. If you know the service/version, please submit the following fingerprint at https://nmap.org/cgi-bin/submit.cgi?new-service :
SF-Port53-TCP:V=7.94SVN%I=7%D=10/19%Time=6714211B%P=x86_64-pc-linux-gnu%r(
SF:DNSVersionBindReqTCP,E,"\0\x0c\0\x06\x81\x84\0\0\0\0\0\0\0\0");
Service Info: Host: cgw.probe.lab; OSs: Linux, RouterOS; Device: router; CPE: cpe:/o:mikrotik:routeros

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 165.34 seconds

real    2m45.381s
user    0m0.598s
sys     0m0.281s

#### Running: masscan 10.0.0.1 -p1-1024 -i eth1 --rate 10000
Starting masscan 1.3.2 (http://bit.ly/14GZzcT) at 2024-10-19 21:16:37 GMT
Initiating SYN Stealth Scan
Scanning 1 hosts [1024 ports/host]
Discovered open port 21/tcp on 10.0.0.1                                        
Discovered open port 53/tcp on 10.0.0.1           

real    0m13.594s
user    0m0.273s
sys     0m1.225s

#### Running: /opt/nscan -r 10.0.0.1
Available IPs:
10.0.0.1

Scanning IP: 10.0.0.1
|- 21 ftp
|- 22 ssh
|- 53 dns
|- 80 http

real    0m0.274s
user    0m0.083s
sys     0m0.095s

### Complex port scan
#### Running: nmap -sSV -p- -Pn -n -T4 10.0.0.1
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-10-19 23:16 CEST
Nmap scan report for 10.0.0.1
Host is up (0.0014s latency).
Not shown: 65527 closed tcp ports (reset)
PORT     STATE SERVICE          VERSION
21/tcp   open  ftp              MikroTik router ftpd 6.49.17
22/tcp   open  ssh              MikroTik RouterOS sshd (protocol 2.0)
53/tcp   open  domain           (generic dns response: NOTIMP)
80/tcp   open  http             MikroTik router config httpd
2000/tcp open  bandwidth-test   MikroTik bandwidth-test server
8291/tcp open  unknown
8728/tcp open  routeros-api     MikroTik RouterOS API
8729/tcp open  ssl/routeros-api MikroTik RouterOS API
1 service unrecognized despite returning data. If you know the service/version, please submit the following fingerprint at https://nmap.org/cgi-bin/submit.cgi?new-service :
SF-Port53-TCP:V=7.94SVN%I=7%D=10/19%Time=671421D8%P=x86_64-pc-linux-gnu%r(
SF:DNSVersionBindReqTCP,E,"\0\x0c\0\x06\x81\x84\0\0\0\0\0\0\0\0");
No exact OS matches for host (If you know what OS is #### Running on it, see https://nmap.org/submit/ ).
TCP/IP fingerprint:
OS:SCAN(V=7.94SVN%E=4%D=10/19%OT=21%CT=1%CU=42980%PV=Y%DS=2%DC=I%G=Y%TM=671
OS:4227C%P=x86_64-pc-linux-gnu)SEQ(SP=10%GCD=FA00%ISR=9B%TI=I%CI=I%II=I%SS=
OS:S%TS=U)SEQ(SP=10%GCD=FA00%ISR=9C%TI=I%CI=I%II=I%SS=S%TS=U)SEQ(SP=12%GCD=
OS:FA00%ISR=9C%TI=I%CI=I%II=I%SS=S%TS=U)SEQ(SP=E%GCD=FA00%ISR=9B%TI=I%CI=I%
OS:II=I%SS=S%TS=U)SEQ(SP=F%GCD=FA00%ISR=9B%TI=I%CI=I%II=I%SS=S%TS=U)OPS(O1=
OS:M5B4%O2=M5B4%O3=M5B4%O4=M5B4%O5=M5B4%O6=M5B4)WIN(W1=FFFF%W2=FFFF%W3=FFFF
OS:%W4=FFFF%W5=FFFF%W6=FFFF)ECN(R=Y%DF=N%T=41%W=FFFF%O=M5B4%CC=N%Q=)T1(R=Y%
OS:DF=N%T=41%S=O%A=S+%F=AS%RD=0%Q=)T2(R=Y%DF=N%T=100%W=0%S=Z%A=S%F=AR%O=%RD
OS:=0%Q=)T3(R=Y%DF=N%T=100%W=0%S=Z%A=S+%F=AR%O=%RD=0%Q=)T4(R=Y%DF=N%T=100%W
OS:=0%S=A%A=Z%F=R%O=%RD=0%Q=)T5(R=Y%DF=N%T=100%W=0%S=Z%A=S+%F=AR%O=%RD=0%Q=
OS:)T6(R=Y%DF=N%T=100%W=0%S=A%A=Z%F=R%O=%RD=0%Q=)T7(R=Y%DF=N%T=100%W=0%S=Z%
OS:A=S%F=AR%O=%RD=0%Q=)U1(R=Y%DF=N%T=33%IPL=164%UN=0%RIPL=G%RID=G%RIPCK=G%R
OS:UCK=G%RUD=G)U1(R=Y%DF=N%T=35%IPL=164%UN=0%RIPL=G%RID=G%RIPCK=G%RUCK=G%RU
OS:D=G)U1(R=Y%DF=N%T=37%IPL=164%UN=0%RIPL=G%RID=G%RIPCK=G%RUCK=G%RUD=G)U1(R
OS:=Y%DF=N%T=3A%IPL=164%UN=0%RIPL=G%RID=G%RIPCK=G%RUCK=G%RUD=G)U1(R=Y%DF=N%
OS:T=3C%IPL=164%UN=0%RIPL=G%RID=G%RIPCK=G%RUCK=G%RUD=G)IE(R=Y%DFI=S%T=25%CD
OS:=S)IE(R=Y%DFI=S%T=2B%CD=S)IE(R=Y%DFI=S%T=2C%CD=S)IE(R=Y%DFI=S%T=30%CD=S)

Network Distance: 2 hops
Service Info: Host: cgw.probe.lab; OSs: Linux, RouterOS; Device: router; CPE: cpe:/o:mikrotik:routeros

OS and Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 185.51 seconds

real    3m5.565s
user    0m1.391s
sys     0m1.834s

#### Running: masscan 10.0.0.1 -p1-65535 -i eth1 --rate 10000
Starting masscan 1.3.2 (http://bit.ly/14GZzcT) at 2024-10-19 21:19:56 GMT
Initiating SYN Stealth Scan
Scanning 1 hosts [65535 ports/host]
Discovered open port 21/tcp on 10.0.0.1                                        
Discovered open port 8291/tcp on 10.0.0.1                                      
Discovered open port 80/tcp on 10.0.0.1                                        
Discovered open port 22/tcp on 10.0.0.1                                        
Discovered open port 2000/tcp on 10.0.0.1                                      
Discovered open port 53/tcp on 10.0.0.1              

real    0m22.430s
user    0m0.157s
sys     0m2.536s

#### Running: /opt/nscan -r 10.0.0.1 -b
Available IPs:
10.0.0.1

Scanning IP: 10.0.0.1
|- 21 ftp
220 cgw.probe.lab FTP server (MikroTik 6.49.17) ready
 <nil>
|- 22 ssh
SSH-2.0-ROSSSH
 <nil>
|- 53 dns
 no banner received from 10.0.0.1:53

real    0m6.239s
user    0m0.036s
sys     0m0.089s

### Fast host discovery using ping
#### Running: masscan -i eth1 10.0.0.0/8 --ping --rate 10000
Starting masscan 1.3.2 (http://bit.ly/14GZzcT) at 2024-10-19 21:20:25 GMT
Initiating ICMP Echo Scan
Scanning 16777216 hosts

real    29m53.888s
user    0m8.664s
sys     6m40.087s

#### Running: zmap -q -p 0 --probe-module=icmp_echoscan -i eth1 -r 10000 -G C4:AD:34:58:98:08 10.0.0.0/8
Oct 19 23:50:19.446 [INFO] zmap: output module: csv
Oct 19 23:50:19.446 [INFO] csv: no output file selected, will use stdout                                                                                     
10.0.0.1                                                                               
10.255.255.251
Oct 20 00:18:36.712 [INFO] zmap: completed                                                                                  

real    28m17.344s
user    72m27.095s
sys     11m51.501s

#### Running: /opt/nscan -r 10.0.0.0/8 -d
Available IPs:
10.0.0.1
10.0.2.3
10.0.2.4
10.0.2.2
10.0.2.15

real    164m2.553s
user    280m32.979s
sys     128m13.971s

#### Running: nmap -sn -PE 10.0.0.0/8 -n -T5 --min-parallelism 100 --max-rtt-timeout 100ms --max-retries 1
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-10-20 03:02 CEST
Nmap scan report for 10.0.0.1
Host is up (0.0027s latency).
Nmap scan report for 10.0.2.2
Host is up (0.00013s latency).
MAC Address: 52:54:00:12:35:02 (QEMU virtual NIC)
Nmap scan report for 10.0.2.3
Host is up (0.00012s latency).
MAC Address: 52:54:00:12:35:03 (QEMU virtual NIC)
Nmap scan report for 10.0.2.4
Host is up (0.00014s latency).
MAC Address: 52:54:00:12:35:04 (QEMU virtual NIC)
Nmap scan report for 10.0.2.15
Host is up.
Nmap scan report for 10.255.255.251
Host is up (0.0012s latency).
Nmap done: 16777216 IP addresses (6 hosts up) scanned in 25340.32 seconds

real    422m20.367s
user    59m3.292s
sys     33m13.327s

### Host discovery of a specific service
#### Running: zmap -q -p 22 -i eth1 -G C4:AD:34:58:98:08 10.0.0.0/8 -r 10000
Oct 20 10:04:59.708 [INFO] zmap: output module: csv
Oct 20 10:04:59.709 [INFO] csv: no output file selected, will use stdout    
10.0.0.1                                                                               
10.255.255.251                                                                                                                                                           
Oct 20 10:33:07.064 [INFO] zmap: completed 

real    28m7.447s
user    72m28.121s
sys     11m51.417s

#### Running: masscan -i eth1 10.0.0.0/8 -p22 --rate 10000
Starting masscan 1.3.2 (http://bit.ly/14GZzcT) at 2024-10-20 08:33:07 GMT
Initiating SYN Stealth Scan
Scanning 16777216 hosts [1 port/host]
Discovered open port 22/tcp on 10.0.0.1
Discovered open port 22/tcp on 10.255.255.251

real    29m50.376s
user    0m8.749s
sys     6m40.375s
