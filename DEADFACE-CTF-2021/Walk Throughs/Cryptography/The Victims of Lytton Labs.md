# The Victims of Lytton Labs

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

![[Pasted image 20211016192505.png]]

This challenge is dependent on the Network Traffic Analysis (PCAP) challenge(s).  

There are several encrypted ".lcr" files, most of which have encrypted data, which, when unencrypted, tell the stories of victims of Lytton Labs.  A few are actually just random data as red herrings.  One of these files, **daniel-s.txt.lcr** contains the flag at the end of it.

Players will search for encrypted files in the packet capture (PCAP) file.  Luciafer transfered them to her computer over FTP, so they can all be extracted from the PCAP using either **Wireshark**.   In **Wireshark**, you have to isolate the packets for the file you want to export.  The best way is to do:

1.  Filter on `ftp || ftp-data`.
2.  Select the packet in the `ftp-data` after its filename appears in `ftp` packets.
3.  Right click and select **Follow | TCP Stream**.  In the window that appears, there will be the file raw data.  Select **Raw** for the format, then **Save As**.

Do this for each of the ***.lcr*** files you want to explore.  The flag is in **daniel-s.txt.lcr**

![[Pasted image 20211016195336.png]]

Get the password file.  This can be exported as ASCII.

![[Pasted image 20211016195841.png]]

Get one of the decryptor files.  We'll get the .bin version for Linux.

![[Pasted image 20211016200050.png]]

The Players also need to find a file that contains a hashed key, along with the encryption programs that can decrypt the files.

The password is found in **/SECRET/encryption-password-cgeschickter.txt**.  It is a **SHA-256** hash of the word "**demagorgon**".  It can be cracked at:
https://crackstation.net or by using **john**, **hashcat**, etc.

![[Pasted image 20211016200543.png]]

Password is "**demagorgon**" (for all of the files... others have "stories" relating to the victims of Lytton Labs).

Flag is found in **/MKSEARCH/daniel-s.txt.lcr**

![[Pasted image 20211016200959.png]]

**flag{D0nt-ME$$-with-The-ZEAL0t!!!}**
