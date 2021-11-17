# Yeet

A drop in binary replacement for [unix builtin](https://en.wikipedia.org/wiki/Yes_(Unix)) `yes` command

## Usage

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


## Why would anyone want this? 

You don't. Please don't use this. Its stupid and pointless and you really shouldn't install it on a production server.
