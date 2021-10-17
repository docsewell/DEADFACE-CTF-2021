# Decrypting Lytton Labs Cryptoware 2 (AES)

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

**Permission given to reprint if above header stays in tact.**

### Looking for network IOCs

There is a separate challenge for the Network IOC  That challenge indicates that the cryptoware talks on the network. One of the IOCs is the file that is downloaded from the network.

The cryptoware downloads a file from: http://insidious.deadface.io/zealotcrypt-aes-key.txt  It contains the string `scoobiedoobiedoo`

The cryptoware takes the MD5 hash of the text, and uses that as the encryption key (as raw bytes) for the data.  This can be seen by doing static analysis with a tool like IDA (Free).

Following the program flow, there is a method called **main_fetchKey**.

![[Pasted image 20211016071147.png]]

Clicking into that method, and following it down, it can be seen that an MD5 sum is performed on the text string.

![[Pasted image 20211016071338.png]]

Using the ASCII hex version of the MD5 sum with the decryptor will decrypt the file.


### Pulling the key from memory

Alternatively, players could simply put a breakpoint on the encryption algorithm and read the key from memory.

