# Cereal Killer 2

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

![[Pasted image 20211016181250.png]]

### Examining the Binary

We'll start on Linux on this one, then move to Windows.

Download the binary and run `file` on it.

![[Pasted image 20211016181512.png]]

It's a **"Mono/.Net" assembly**.  In other words, a Windows .NET executable.  .NET binaries are easily reversed using a tool like **dotPeek** or **dnSPY**.  We'll go with **dnSpy**, since you can use it to modify .NET binaries as well.  It's a free download from https://github.com/dnSpy/dnSpy/releases

Since our program was compiled with 32-bit, we'll grab that version from the download site.  Simply unzip it into a folder and run the **dnspy.exe** binary.

### Run the program Windows

The program spits out a garbled string.

![[Pasted image 20211016190239.png]]

### Examine the program in dotPeek

Load the program in **dnSpy**.  Open up the **Main** function and examine.  (Sorry, I hate all the DARK themes...)

Here, we aren't really concerned with finding any keys or that.  We are mainly interested in the function called **DecryptFromBase64ToString**.  We can also see in **Main** that the same Base64 encoded string that was spit to the screen is written with **Console.WriteLine**.  What if we wrapped it with the **DecryptFromBase64ToString** function?  Let's try it.

Right click the window, select **Edit Class (C#)**.  Make the changes.

![[Pasted image 20211016191432.png]]

Click the **Compile** button at the bottom of the **Edit Class** window.  Click **Edit | Save Module...**.  Click **OK**.

![[Pasted image 20211016191627.png]]

Now, run the program again.

![[Pasted image 20211016191721.png]]

**flag{frank3n-berry-goodness-NOM-NOM-NOM}**