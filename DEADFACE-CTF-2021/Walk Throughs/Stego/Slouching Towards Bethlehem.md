# Slouching Towards Bethlehem

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

**Permission given to reprint if above header stays in tact.**

![[Pasted image 20211016171837.png]]

### Ghost Town Intel

Ghost Town has chatter related to "Multi-Channel Messages"

https://ghosttown.deadface.io/t/hint-for-luciafer/48/8

"The Bird is the Word" is the context sensitive password.  The poem mentions the "falcon" which cannot hear the falconer.

### Extract the File

Using "falcon" as the password in OpenStego with the image produces the following data in a file.

![[Pasted image 20211016180243.png]]

![[Pasted image 20211016180304.png]]

### Attempting to Read the File

Using strings on the file gives back no results.

Using it again with "-eb" for Unicode, gives the following:

	urning and turning in the widening gyre   
	The falcon cannot hear the falconer;
	Things fall apart; the centre cannot hold;
	Mere anarchy is loosed upon the world,
	The blood-dimmed tide is loosed, and everywhere   
	The ceremony of innocence is drowned;
	The best lack all conviction, while the worst   
	Are full of passionate intensity.
	Surely some revelation is at hand;
	Surely the Second Coming is at hand.   
	The Second Coming! Hardly are those words out   
	When a vast image out of f
	Troubles my sight: somewhere in sands of the desert   
	A shape with lion body and the head of a man,   
	A gaze blank and pitiless as the sun,   
	Is moving its slow thighs, while all about it   
	Reel shadows of the indignant desert birds.   
	The darkness drops again; but now I know   
	That twenty centuries of stony sleep
	Were vexed to nightmare by a rocking cradle,   
	And what rough beast, its hour come round at last,   
	Slouches towards Bethlehem to be born?


There are some words missing.  Bringing the file up in gedit (Linux Text Editor) gives us another view.

	Turning and turning in the widening gyre   
	The falcon cannot hear the falconer;
	Things fall apart; the centre cannot hold;
	Mere anarchy is loosed upon the world,
	The blood-dimmed tide is loosed, and everywhere   
	The ceremony of innocence is drowned;
	The best lack all conviction, while the worst   
	Are full of passionate intensity.
	Surely some revelation is at hand;
	Surely the Second Coming is at hand.   
	The Second Coming! Hardly are those words out   
	When a vast image out of ٦٬١٧ٻٰٓ٩ٲ٩ٴٵٳؠٍٵٮ٤٩ٽ
	Troubles my sight: somewhere in sands of the desert   
	A shape with lion body and the head of a man,   
	A gaze blank and pitiless as the sun,   
	Is moving its slow thighs, while all about it   
	Reel shadows of the indignant desert birds.   
	The darkness drops again; but now I know   
	That twenty centuries of stony sleep
	Were vexed to nightmare by a rocking cradle,   
	And what rough beast, its hour come round at last,   
	Slouches towards Bethlehem to be born?
	ՁդդԠԢԠԲ԰ԲԱԢԠազմեղԠԢՍյծդթԢԠմկԠգկխհլեմեԠմըեԠզլաէԮ

The garbled data can be viewed by hexdump:

	hexdump -C re-02.txt
	
The first part of the flag can be seen, but with the number "06" in the upper byte of the Unicode character.  flag{Spiritus Mundi}

![[Pasted image 20211016170528.png]]

**flag{Spiritus Mundi}** doesn't work as the flag, however.

Heeding TheZeal0t's warning to bumpyhassan in the Ghost Town chatter, we must look deeper.  Lower in the file, we see additional instructions:

![[Pasted image 20211016170620.png]]

We are told to ' Add 2021 after "Mundi" to complete the flag'

This gives us:

**flag{Spiritus Mundi 2021}** as the flag.



