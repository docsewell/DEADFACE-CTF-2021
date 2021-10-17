# The Sum of All Fears

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

![[Pasted image 20211016203150.png]]

Players will search for binaries with identical names in the packet capture file, but different extensions: **lytton-crypt.exe** and **lytton-crypt.bin**.  They need to extract these and take the md5 has of them.

1.  Filter on `ftp || ftp-data`.
2.  Select the packet in the `ftp-data` after its filename appears in `ftp` packets.
3.  Right click and select **Follow | TCP Stream**.  In the window that appears, there will be the file raw data.  Select **Raw** for the format, then **Save As**.

![[Pasted image 20211016203357.png]]

**flag{9cb9b11484369b95ce35904c691a5b28|4da8e81ee5b08777871e347a6b296953}**