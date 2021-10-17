# Cereal Killer 3

## DEADFACE CTF 2021
## Cyber Hacktics
## by Daniel "Doc" Sewell, a.k.a. TheZeal0t (Author)

![[Pasted image 20211016183411.png]]

## Examining the Binary

Run `file` on it.

![[Pasted image 20211016183504.png]]

## Run the Program

![[Pasted image 20211016183751.png]]

So it gathers input from the user, compares it to the real password, then if correct, displays the flag to the user.

## Reversing it with IDA (or tool of your choice)

**IDA (FREE) 7.6** (https://hex-rays.com/ida-free/) works really well for this, but Ghidra or other debugging tools could work as well.  You could use dynamic debugging to find the password, but I chose an easier path with minimal knowledge of RE required.

Use IDA to load the file.  Look for C functions where input is gathered from the user (in this case, `isoc99_scanf`), followed by a `strcmp`, followed by a branch instruction `jnz`,

![[Pasted image 20211016183921.png]]

We would like the program to give us the flag with a bad password, so we can simply patch the instruction from `jnz` to `jz` to reverse the logic.

## Patching the binary

Select the `jnz` instruction after the `_strcmp` function and immediately after the `cmp`.  Select from the main menu: **Edit | Patch Program | Assemble...**.

![[Pasted image 20211016185213.png]]

Change the `jnz` instruction to `jz`.

![[Pasted image 20211016185239.png]]

Click **OK** then **Cancel** when the **Assemble instruction** window appears again.

Now select from the main menu **Edit | Patch Program | Apply patches to input file...**.  Make sure you have already changed permissions on the binary with **chmod** before trying this, or it will likely fail.  Click **OK**.

![[Pasted image 20211016185537.png]]

## Getting the flag

Run the program again.  Enter anything EXCEPT the correct password, and the flag will be displayed!

![[Pasted image 20211016185631.png]]

**flag{B00-B00-B00-Bury-IZ-DA-BOMB}**
