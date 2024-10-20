Select which test to run:
1) Fast Port Scan
2) Complex Port Scan
3) Fast Host Discovery
4) Service Discovery
5) Run All
Enter your choice [1-5]: 1

### Fast port scan common ports / 1-1024
#### Running: nmap -sSV -Pn -n -T4 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.8.3.17 10.0.0.1
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-10-20 15:08 CEST
Nmap scan report for 10.255.255.251
Host is up (0.020s latency).
Not shown: 998 closed tcp ports (reset)
PORT     STATE SERVICE       VERSION
22/tcp   open  ssh           OpenSSH 9.6p1 Ubuntu 3ubuntu13.5 (Ubuntu Linux; protocol 2.0)
3389/tcp open  ms-wbt-server xrdp
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.244.170.185
Host is up (0.029s latency).
Not shown: 997 filtered tcp ports (no-response)
PORT    STATE  SERVICE VERSION
22/tcp  open   ssh     OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
25/tcp  open   smtp    Postfix smtpd
161/tcp closed snmp
Service Info: Host:  debian; OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.86.46.58
Host is up (0.0095s latency).
Not shown: 997 filtered tcp ports (no-response)
PORT    STATE SERVICE     VERSION
21/tcp  open  ftp         vsftpd 3.0.3
22/tcp  open  ssh         OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
445/tcp open  netbios-ssn Samba smbd 4.6.2
Service Info: OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.8.3.19
Host is up (0.031s latency).
Not shown: 996 filtered tcp ports (no-response)
PORT     STATE SERVICE VERSION
22/tcp   open  ssh     OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
80/tcp   open  http    Apache httpd 2.4.62 ((Debian))
443/tcp  open  http    Apache httpd 2.4.62
3306/tcp open  mysql   MariaDB (unauthorized)
Service Info: Host: 127.0.2.1; OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.8.3.17
Host is up (0.0072s latency).
Not shown: 997 filtered tcp ports (no-response)
PORT    STATE SERVICE       VERSION
135/tcp open  msrpc         Microsoft Windows RPC
139/tcp open  netbios-ssn   Microsoft Windows netbios-ssn
445/tcp open  microsoft-ds?
Service Info: OS: Windows; CPE: cpe:/o:microsoft:windows

Nmap scan report for 10.0.0.1
Host is up (0.0054s latency).
Not shown: 994 closed tcp ports (reset)
PORT     STATE SERVICE        VERSION
21/tcp   open  ftp            MikroTik router ftpd 6.49.17
22/tcp   open  ssh            MikroTik RouterOS sshd (protocol 2.0)
53/tcp   open  domain         (generic dns response: NOTIMP)
80/tcp   open  http           MikroTik router config httpd
2000/tcp open  bandwidth-test MikroTik bandwidth-test server
8291/tcp open  unknown
1 service unrecognized despite returning data. If you know the service/version, please submit the following fingerprint at https://nmap.org/cgi-bin/submit.cgi?new-service :
SF-Port53-TCP:V=7.94SVN%I=7%D=10/20%Time=67150199%P=x86_64-pc-linux-gnu%r(
SF:DNSVersionBindReqTCP,E,"\0\x0c\0\x06\x81\x84\0\0\0\0\0\0\0\0");
Service Info: Host: cgw.probe.lab; OSs: Linux, RouterOS; Device: router; CPE: cpe:/o:mikrotik:routeros

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 6 IP addresses (6 hosts up) scanned in 371.22 seconds

real    6m11.237s
user    0m0.781s
sys     0m0.727s

#### Running: masscan 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.8.3.17 10.0.0.1 -p1-1024 -i eth1 --rate 10000
Starting masscan 1.3.2 (http://bit.ly/14GZzcT) at 2024-10-20 13:14:27 GMT
Initiating SYN Stealth Scan
Scanning 6 hosts [1024 ports/host]
Discovered open port 21/tcp on 10.0.0.1                                        
Discovered open port 53/tcp on 10.0.0.1                                        
Discovered open port 22/tcp on 10.255.255.251                                  
Discovered open port 22/tcp on 10.0.0.1                                        
Discovered open port 80/tcp on 10.0.0.1                                    
                                                                             
real    0m13.056s
user    0m0.029s
sys     0m0.544s

#### Running: /opt/nscan -r 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.8.3.17 10.0.0.1
Available IPs:
10.255.255.251
10.244.170.185
10.86.46.58
10.8.3.19
10.0.0.1
10.8.3.17

Scanning IP: 10.255.255.251
|- 22 ssh

Scanning IP: 10.244.170.185
|- 22 ssh
|- 25 smtp

Scanning IP: 10.86.46.58
|- 21 ftp
|- 22 ssh
|- 445 microsoft-ds

Scanning IP: 10.8.3.19
|- 22 ssh
|- 80 http
|- 443 https

Scanning IP: 10.0.0.1

Scanning IP: 10.8.3.17

real    0m6.284s
user    0m0.108s
sys     0m0.312s
