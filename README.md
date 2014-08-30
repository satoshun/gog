# gog

this is tools for used convenience git like go. i called gog.


## Description

Go have special direcotry structure. i love it so other languages want to use it. gog achieve it.

## install

for macOS

```
$ brew tap satoshun/commands
$ brew install gog
```

## Usage

### prepare

```
# first set base directory
$ export GOG_PATH=~/git

# option variable
## if `GOG_HOOK_CMD` variable then run `GOG_HOOK_CMD` before ending.
$ export GOG_HOOK_CMD="cd {{.Directory}} && git status"
```

`GOG_HOOK_CMD` has been prepared under variables.

- Directory
- Repository
- Host
- Path
- ProjectName

### do

```
$ gog -h
NAME:
   gog - use directory like Go

USAGE:
   gog [global options] command [command options] [arguments...]

VERSION:
   0.2.0

COMMANDS:
   get, g clone repository
   update, u  update repository
   list, l  list clone repository
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --base, -b ''  define git path
   --version, -v  print the version
   --help, -h   show help

# clone
$ gog get git@github.com:satoshun/gog.git
Cloning into '/Users/satouhayabusa/git/src/github.com/satoshun/gog'...
warning: You appear to have cloned an empty repository.
Checking connectivity... done.

# update
$ gog update git@github.com:satoshun/pythonjs.git
First, rewinding head to replay your work on top of it...
Fast-forwarded master to 85d86acef567a9729e025defe895df6ee4aa35f7.
```
