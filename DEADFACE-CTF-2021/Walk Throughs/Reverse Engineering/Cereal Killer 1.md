# Cereal Killer 1

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

![[Pasted image 20211016180454.png]]


### Examining the Binary

This one is pretty easy.  Pick the Windows or Linux binary, download it, and run it to see what it does.

![[Pasted image 20211016180626.png]]


Run `strings` on it.

	strings deadface_re01.bin
	
![[Pasted image 20211016180731.png]]

The flag does not appear (it is obfuscated), but a string related to High Fructose Corn Syrup laden material does: **c0unt-ch0cula**.

Run again, and provide that as the password.

![[Pasted image 20211016180907.png]]

**flag{c0unt-ch0cula-cereal-FTW}**

