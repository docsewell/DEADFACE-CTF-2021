# Release the Crackin'

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

![[Pasted image 20211016203729.png]]

Inspecting the PCAP, the players can see that Luciafer is trying to crack the FTP protocol.  A bit of Googling will come back with FTP Code 230 for successful log in.  Using `ftp.response.code==230` will show all successful logins with FTP.

![[Pasted image 20211016203539.png]]

Pick one (in this case, it doesn't matter which), and remove the filter.  In this case, the packet immediately above shows the password that gained successful login (**darkangel**).

![[Pasted image 20211016203608.png]]

**flag{darkangel}**

