Select which test to run:
1) Fast Port Scan
2) Complex Port Scan
3) Fast Host Discovery
4) Service Discovery
5) Run All
Enter your choice [1-5]: 4

### Host discovery of a specific service
#### Running: zmap -q -p 22 -i eth1 -G C4:AD:34:58:98:03 10.0.0.0/8 -r 10000
Oct 26 11:12:26.661 [INFO] zmap: output module: csv
Oct 26 11:12:26.661 [INFO] csv: no output file selected, will use stdout                                                                                                                                                                    
10.255.255.251                                                                                                                                                                                                                              
10.0.0.1
Oct 26 11:40:36.908 [INFO] zmap: completed
                                                                            
real    28m10.339s
user    70m11.211s
sys     14m12.700s

#### Running: masscan --router-mac C4-AD-34-58-98-03 -i eth1 10.0.0.0/8  -p22 --rate 10000
Starting masscan 1.3.2 (http://bit.ly/14GZzcT) at 2024-10-26 09:40:37 GMT
Initiating SYN Stealth Scan
Scanning 16777216 hosts [1 port/host]
Discovered open port 22/tcp on 10.255.255.251                                  
Discovered open port 22/tcp on 10.8.3.4                                        
                                                                             
real    27m49.236s
user    0m5.814s
sys     4m42.814s
