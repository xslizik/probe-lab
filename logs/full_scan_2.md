Select which test to run:
1) Fast Port Scan
2) Complex Port Scan
3) Fast Host Discovery
4) Service Discovery
5) Run All
Enter your choice [1-5]: 5

### Fast port scan common ports / 1-1024
#### Running: nmap -sSV -Pn -n -T4 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.0.0.1
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-10-25 15:58 CEST
Nmap scan report for 10.255.255.251
Host is up (0.000068s latency).
Not shown: 998 closed tcp ports (reset)
PORT     STATE SERVICE       VERSION
22/tcp   open  ssh           OpenSSH 9.6p1 Ubuntu 3ubuntu13.5 (Ubuntu Linux; protocol 2.0)
3389/tcp open  ms-wbt-server xrdp
MAC Address: D4:5D:DF:13:8D:66 (Pegatron)
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.244.170.185
Host is up (0.00043s latency).
Not shown: 997 filtered tcp ports (no-response)
PORT    STATE  SERVICE VERSION
22/tcp  open   ssh     OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
25/tcp  open   smtp    Postfix smtpd
161/tcp closed snmp
MAC Address: 08:00:27:64:FD:0B (Oracle VirtualBox virtual NIC)
Service Info: Host:  debian; OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.86.46.58
Host is up (0.00034s latency).
Not shown: 997 filtered tcp ports (no-response)
PORT    STATE SERVICE     VERSION
21/tcp  open  ftp         vsftpd 3.0.3
22/tcp  open  ssh         OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
445/tcp open  netbios-ssn Samba smbd 4.6.2
MAC Address: 08:00:27:AB:8F:DD (Oracle VirtualBox virtual NIC)
Service Info: OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.8.3.19
Host is up (0.00035s latency).
Not shown: 996 filtered tcp ports (no-response)
PORT     STATE SERVICE VERSION
22/tcp   open  ssh     OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
80/tcp   open  http    Apache httpd 2.4.62 ((Debian))
443/tcp  open  http    Apache httpd 2.4.62
3306/tcp open  mysql   MariaDB (unauthorized)
MAC Address: 08:00:27:09:E0:F0 (Oracle VirtualBox virtual NIC)
Service Info: Host: 127.0.2.1; OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.0.0.1
Host is up (0.00088s latency).
Not shown: 994 closed tcp ports (reset)
PORT     STATE SERVICE        VERSION
21/tcp   open  ftp            MikroTik router ftpd 6.49.17
22/tcp   open  ssh            MikroTik RouterOS sshd (protocol 2.0)
53/tcp   open  domain         (generic dns response: NOTIMP)
80/tcp   open  http           MikroTik router config httpd
2000/tcp open  bandwidth-test MikroTik bandwidth-test server
8291/tcp open  unknown
1 service unrecognized despite returning data. If you know the service/version, please submit the following fingerprint at https://nmap.org/cgi-bin/submit.cgi?new-service :
SF-Port53-TCP:V=7.94SVN%I=7%D=10/25%Time=671BA405%P=x86_64-pc-linux-gnu%r(
SF:DNSVersionBindReqTCP,E,"\0\x0c\0\x06\x81\x84\0\0\0\0\0\0\0\0");
MAC Address: C4:AD:34:58:98:03 (Routerboard.com)
Service Info: Host: cgw.probe.lab; OSs: Linux, RouterOS; Device: router; CPE: cpe:/o:mikrotik:routeros

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 5 IP addresses (5 hosts up) scanned in 171.24 seconds

real    2m51.296s
user    0m0.669s
sys     0m0.663s

#### Running: masscan 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.0.0.1 -p1-1024 -i eth1 --rate 10000
[+] resolving router 0.0.0.0 with ARP (may take some time)...
[-] FAIL: ARP timed-out resolving MAC address for router eth1: "0.0.0.0"
    [hint] try "--router ip 192.0.2.1" to specify different router
    [hint] try "--router-mac 66-55-44-33-22-11" instead to bypass ARP
    [hint] try "--interface eth0" to change interface

real    0m19.654s
user    0m0.000s
sys     0m0.009s

#### Running: /opt/nscan -r 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.0.0.1
Available IPs:
10.255.255.251
10.0.0.1
10.8.3.19
10.86.46.58
10.244.170.185

Scanning IP: 10.255.255.251
|- 22 ssh

Scanning IP: 10.0.0.1
|- 21 ftp
|- 22 ssh
|- 53 dns

Scanning IP: 10.8.3.19
|- 22 ssh
|- 80 http
|- 443 https

Scanning IP: 10.86.46.58
|- 21 ftp
|- 22 ssh
|- 445 microsoft-ds

Scanning IP: 10.244.170.185
|- 22 ssh
|- 25 smtp

real    0m3.875s
user    0m0.110s
sys     0m0.352s

### Complex port scan
#### Running: nmap -sSV -p- -Pn -n -T4 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.0.0.1
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-10-25 16:01 CEST
Nmap scan report for 10.255.255.251
Host is up (0.00011s latency).
Not shown: 65533 closed tcp ports (reset)
PORT     STATE SERVICE       VERSION
22/tcp   open  ssh           OpenSSH 9.6p1 Ubuntu 3ubuntu13.5 (Ubuntu Linux; protocol 2.0)
3389/tcp open  ms-wbt-server xrdp
MAC Address: D4:5D:DF:13:8D:66 (Pegatron)
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.244.170.185
Host is up (0.00039s latency).
Not shown: 65532 filtered tcp ports (no-response)
PORT    STATE  SERVICE VERSION
22/tcp  open   ssh     OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
25/tcp  open   smtp    Postfix smtpd
161/tcp closed snmp
MAC Address: 08:00:27:64:FD:0B (Oracle VirtualBox virtual NIC)
Service Info: Host:  debian; OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.86.46.58
Host is up (0.00038s latency).
Not shown: 65531 filtered tcp ports (no-response)
PORT    STATE  SERVICE     VERSION
21/tcp  open   ftp         vsftpd 3.0.3
22/tcp  open   ssh         OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
137/tcp closed netbios-ns
445/tcp open   netbios-ssn Samba smbd 4.6.2
MAC Address: 08:00:27:AB:8F:DD (Oracle VirtualBox virtual NIC)
Service Info: OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.8.3.19
Host is up (0.00042s latency).
Not shown: 65531 filtered tcp ports (no-response)
PORT     STATE SERVICE VERSION
22/tcp   open  ssh     OpenSSH 9.2p1 Debian 2+deb12u2 (protocol 2.0)
80/tcp   open  http    Apache httpd 2.4.62 ((Debian))
443/tcp  open  http    Apache httpd 2.4.62
3306/tcp open  mysql   MariaDB (unauthorized)
MAC Address: 08:00:27:09:E0:F0 (Oracle VirtualBox virtual NIC)
Service Info: Host: 127.0.2.1; OS: Linux; CPE: cpe:/o:linux:linux_kernel

Nmap scan report for 10.0.0.1
Host is up (0.00047s latency).
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
SF-Port53-TCP:V=7.94SVN%I=7%D=10/25%Time=671BA54C%P=x86_64-pc-linux-gnu%r(
SF:DNSVersionBindReqTCP,E,"\0\x0c\0\x06\x81\x84\0\0\0\0\0\0\0\0");
MAC Address: C4:AD:34:58:98:03 (Routerboard.com)
Service Info: Host: cgw.probe.lab; OSs: Linux, RouterOS; Device: router; CPE: cpe:/o:mikrotik:routeros

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 5 IP addresses (5 hosts up) scanned in 303.85 seconds

real    5m3.906s
user    0m9.247s
sys     0m23.178s

#### Running: masscan 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.0.0.1 --banners -p1-65535 -i eth1 --rate 10000
[+] resolving router 0.0.0.0 with ARP (may take some time)...
[-] FAIL: ARP timed-out resolving MAC address for router eth1: "0.0.0.0"
    [hint] try "--router ip 192.0.2.1" to specify different router
    [hint] try "--router-mac 66-55-44-33-22-11" instead to bypass ARP
    [hint] try "--interface eth0" to change interface

real    0m16.623s
user    0m0.004s
sys     0m0.003s

#### Running: /opt/nscan -r 10.255.255.251 10.244.170.185 10.86.46.58 10.8.3.19 10.0.0.1 -p 0-65535 -b
Available IPs:
10.8.3.19
10.244.170.185
10.255.255.251
10.0.0.1
10.86.46.58

Scanning IP: 10.8.3.19
|- 22 ssh
SSH-2.0-OpenSSH_9.2p1 Debian-2+deb12u2
 <nil>
|- 80 http
 no banner received from 10.8.3.19:80
|- 443 https
 no banner received from 10.8.3.19:443
|- 3306 mysql
C�jHost '10.8.3.7' is not allowed to connect to this MariaDB server <nil>

Scanning IP: 10.244.170.185
|- 22 ssh
SSH-2.0-OpenSSH_9.2p1 Debian-2+deb12u2
 <nil>
|- 25 smtp
220 debian ESMTP Postfix (Debian/GNU)
 <nil>

Scanning IP: 10.255.255.251
|- 22 ssh
SSH-2.0-OpenSSH_9.6p1 Ubuntu-3ubuntu13.5
 <nil>
|- 3389 rdp
 no banner received from 10.255.255.251:3389

Scanning IP: 10.0.0.1
|- 21 ftp
220 cgw.probe.lab FTP server (MikroTik 6.49.17) ready
 <nil>
|- 22 ssh
SSH-2.0-ROSSSH
 <nil>
|- 53 dns
 no banner received from 10.0.0.1:53
|- 80 http
 no banner received from 10.0.0.1:80
|- 2000 
 <nil>
|- 8291 
 no banner received from 10.0.0.1:8291
|- 8728 
 no banner received from 10.0.0.1:8728
|- 8729 
 no banner received from 10.0.0.1:8729

Scanning IP: 10.86.46.58
|- 21 ftp
220 (vsFTPd 3.0.3)
 <nil>
|- 22 ssh
SSH-2.0-OpenSSH_9.2p1 Debian-2+deb12u2
 <nil>
|- 445 microsoft-ds
 no banner received from 10.86.46.58:445

real    4m0.729s
user    0m7.936s
sys     0m23.416s

### Fast host discovery using
#### Running: masscan -i eth1 10.0.0.0/8 --ping --rate 10000
[+] resolving router 0.0.0.0 with ARP (may take some time)...
[-] FAIL: ARP timed-out resolving MAC address for router eth1: "0.0.0.0"
    [hint] try "--router ip 192.0.2.1" to specify different router
    [hint] try "--router-mac 66-55-44-33-22-11" instead to bypass ARP
    [hint] try "--interface eth0" to change interface

real    0m17.943s
user    0m0.000s
sys     0m0.007s

#### Running: zmap -q -p 0 --probe-module=icmp_echoscan -i eth1 -r 10000 -G C4:AD:34:58:98:08 10.0.0.0/8
Oct 25 16:11:05.183 [INFO] zmap: output module: csv
Oct 25 16:11:05.183 [INFO] csv: no output file selected, will use stdout    
Oct 25 16:39:11.648 [INFO] zmap: completed                                  
                                                                            
real    28m6.544s
user    72m49.546s
sys     11m24.331s

#### Running: /opt/nscan -r 10.0.0.0/8 -d
Available IPs:
10.0.0.1
10.0.2.2
10.0.2.3
10.0.2.4
10.0.2.15
10.8.3.7
10.8.3.19

real    172m2.942s
user    72m30.827s
sys     40m26.248s

#### Running: nmap -sn -PE 10.0.0.0/8 -n -T5 --min-parallelism 100 --max-rtt-timeout 100ms --max-retries 1
Starting Nmap 7.94SVN ( https://nmap.org ) at 2024-10-25 19:31 CEST
Nmap scan report for 10.0.0.1
Host is up (0.00049s latency).
MAC Address: C4:AD:34:58:98:03 (Routerboard.com)
Nmap scan report for 10.0.2.2
Host is up (0.00023s latency).
MAC Address: 52:54:00:12:35:02 (QEMU virtual NIC)
Nmap scan report for 10.0.2.3
Host is up (0.00022s latency).
MAC Address: 52:54:00:12:35:03 (QEMU virtual NIC)
Nmap scan report for 10.0.2.4
Host is up (0.00020s latency).
MAC Address: 52:54:00:12:35:04 (QEMU virtual NIC)
Nmap scan report for 10.0.2.15
Host is up.
Nmap scan report for 10.8.3.19
Host is up (0.00032s latency).
MAC Address: 08:00:27:09:E0:F0 (Oracle VirtualBox virtual NIC)
Nmap scan report for 10.8.3.7
Host is up.
Nmap scan report for 10.86.46.58
Host is up (0.00034s latency).
MAC Address: 08:00:27:AB:8F:DD (Oracle VirtualBox virtual NIC)
Nmap scan report for 10.244.170.185
Host is up (0.00066s latency).
MAC Address: 08:00:27:64:FD:0B (Oracle VirtualBox virtual NIC)
Nmap scan report for 10.255.255.251
Host is up (0.00018s latency).
MAC Address: D4:5D:DF:13:8D:66 (Pegatron)
Nmap done: 16777216 IP addresses (10 hosts up) scanned in 35538.76 seconds

real    592m18.791s
user    28m7.843s
sys     24m16.176s

### Host discovery of a specific service
#### Running: zmap -q -p 22 -i eth1 -G C4:AD:34:58:98:08 10.0.0.0/8 -r 10000
Oct 26 05:23:33.443 [INFO] zmap: output module: csv
Oct 26 05:23:33.443 [INFO] csv: no output file selected, will use stdout                                                                                                                                                                    
Oct 26 05:51:39.843 [INFO] zmap: completed                                                                                                                                                                                                  
                                                                                                                                                                                                                                            
real    28m6.478s
user    72m57.247s
sys     11m20.630s

#### Running: masscan -i eth1 10.0.0.0/8 -p22 --rate 10000
[+] resolving router 0.0.0.0 with ARP (may take some time)...
[-] FAIL: ARP timed-out resolving MAC address for router eth1: "0.0.0.0"
    [hint] try "--router ip 192.0.2.1" to specify different router
    [hint] try "--router-mac 66-55-44-33-22-11" instead to bypass ARP
    [hint] try "--interface eth0" to change interface

real    0m20.176s
user    0m0.007s
sys     0m0.000s
