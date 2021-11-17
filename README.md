# Yeet

![](https://emojis.slackmojis.com/emojis/images/1556122362/5648/yeet.png?1556122362)

A drop in binary replacement for [unix builtin](https://en.wikipedia.org/wiki/Yes_(Unix)) `yes` command

## Usage :computer:

Calling the `yeet` binary will print out `yeet` endlessly
```
$ yeet
yeet
yeet
yeet
yeet
yeet
yeet
...
```

Calling the `yeet` binary with an optional parameter will print out that parameter endlessly

```
$ yeet yes
yes
yes
yes
yes
yes
...
```

It is useful for piping into a command that prompts the user for a 'yes/no' response

```
$ yeet yes | sudo apt install foobar
````


## Why would anyone want this? :confused:

You don't. :stop_sign: Please don't use this. Its stupid and pointless and you really shouldn't install it on a production server.


## No really, why did you make this?  :see_no_evil:

- Daily challenge to learn go. 
- Wanted to publish something with github actions. 


## Installing :floppy_disk:

If you still _really_ want this, downloads are available on the releases page. Including `darwin-arm64` for Apple M1 Silicon and Raspberry Pi.

https://github.com/spuder/yeet/releases

Apple (Intel) = yeet-v0.1.0-darwin-amd64.tar.gz
Apple (M1/ARM) = yeet-v0.1.0-darwin-arm64.tar.gz
Linux (32 bit) = yeet-v0.1.0-linux-386.tar.gz
Linux (64 bit) = yeet-v0.1.0-linux-amd64.tar.gz
Linux (M1/ARM) = yeet-v0.1.0-linux-arm64.tar.gz
Windows (Intel/AMD) = yeet-v0.1.0-windows-amd64.tar.gz
