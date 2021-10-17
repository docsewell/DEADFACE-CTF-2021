# Decrypting Lytton Labs Cryptoware 1 (RC4)

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

**Permission given to reprint if above header stays in tact.**

![[Pasted image 20211016175036.png]]

### The Hard Way

A lot of folks reverse engineered this binary using some advanced static and dynamic analysis.  Kudos to them!  That must have been a lot of work.

You can figure out that the cryptoware downloads a file from:

https://insidious.deadface.io/pretty-lady.gif

It takes the SHA512 hash of the file, then selects out 16 bytes from 14 to 30 (index 13-28) to form the 16-byte (128 bit) key.  This key is used in the RC4 algorithm to do the encryption.

Confusingly (mostly because I forgot to change it), the name of the function that hashes the file is called **getMd5HashByte**.

This key works with CyberChef to decrypt the data.

### The Easy Way

In Ghost Town, there is the following thread:

https://ghosttown.deadface.io/t/custom-ransomware/40/3

In the thread, TheZeal0t scolds Luciafer for selecting RC4 as the algorithm for the ransomware.  He tells her that it is flawed, and that she should have used something else.

The flaw occurs because the same algorithm is used for encryption and decryption.  That's why there is no "decryptor" for Luciafer's cryptoware!

In the challenge called *Luciafer's Cryptoware IOC 2*, the player must decrypt the filenames as encrypted by the cryptoware.  It uses Caesar Cipher +3 to encrypt and -3 to decrypt the file names.  The filenames end in **oodev**, which when decrypted, are **llabs**.

To decrypt the files, there is a backdoor (a "logic error") in Luciafer's malware.  Simply change the file extension from ## Luciafer's Cryptoware IOC 2**oodev** to **llabs** and run the cryptoware again.  The filename will further be scrambled, but the contents of the file will now be readable.

![[Pasted image 20211016064642.png]]




