# Luciafer's Fatal Error

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

![[Pasted image 20211016204906.png]]

To figure out the malicious binary, the player will have to do some digging around in the whole PCAP.  Dark Angel gets control of Luciafer's computer after she runs **secret_decoder.bin**, but this isn't necessarily apparent at first.  The two clues that this is the case is:

1.  Lucifer downloaded it as part of exfiltrating data from Geschickter's computer (Packet 160362)

![[Pasted image 20211016205408.png]]

2.  Dark Angel later downloads the same file, renames it, and sets it up to run at regular intervals (also seen in *Persistence Pays Off*)

  ![[Pasted image 20211016205624.png]]

Search the PCAP for **secret_decoder.bin**.  Extract it.

![[Pasted image 20211016205048.png]]

Upload it to VirusTotal.com to prove that it is malicious.

![[Pasted image 20211016205206.png]]

Take the hash

![[Pasted image 20211016205741.png]]

**flag{42e419a6391ca79dc44d7dcef1efc83b}**

