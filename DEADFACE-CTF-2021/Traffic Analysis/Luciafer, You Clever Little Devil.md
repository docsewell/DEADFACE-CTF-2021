# Luciafer, You Clever Little Devil

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

![[Pasted image 20211016204252.png]]

Inspecting the PCAP, the players can see that Luciafer is trying to crack the FTP protocol.  A bit of Googling will come back with FTP Code 230 for successful log in.  Using `ftp.response.code==230` will show all successful logins with FTP.

![[Pasted image 20211016203539.png]]

Pick the second one, since the first one belongs to the packet capture.  The packet number is on the left (159765).

![[Pasted image 20211016204204.png]]

**flag{159765}**