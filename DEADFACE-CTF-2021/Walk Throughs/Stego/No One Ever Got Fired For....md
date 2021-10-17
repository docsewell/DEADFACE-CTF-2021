# No One Ever Got Fired For...

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

**Permission given to reprint if above header stays in tact.**

![[Pasted image 20211016173340.png]]

### Some Googling

The first words that pop up when you Google: "No One Ever Got Fired For"

![[Pasted image 20211016173428.png]]

The cover image shows people working around an old mainframe.  The challenge description describes frustration at not being able to read the data.

A quick Google search on "IBM data format" comes back with:

![[Pasted image 20211016173753.png]]

A reference to **EBCDIC**.  Searching for EBCDIC "epsuh-dick" (sorry, that's the pronunciation) gives this:

https://en.wikipedia.org/wiki/EBCDIC

![[Pasted image 20211016173947.png]]

### Solving the Challenge (A couple different ways)

Now, you could write a program to convert every ASCII character to its EBCDIC equivalent, but there is a much easier way.

Use **hexeditor** on Linux.

Install it with (on Debian/Ubuntu versions):

	sudo apt install ncurses-hexedit
	
Edit the file

	hexeditor no-one-ever-got-fired.dat
	
Switch from ASCII representation to EBCDIC:

	CTRL-R
	
Search for "FLAG"

	CTRL-W
	
	FLAG
	
![[Pasted image 20211016174404.png]]

**FLAG{BUILD-PCS-WITH-LINUX-INSTEAD}**

(Set to case insensitive, just in case someone set flag to lower case)

Here is an alternate way:

https://www.unix.com/unix-for-dummies-questions-and-answers/93738-viewing-ebcidic-files.html

![[Pasted image 20211016174810.png]]