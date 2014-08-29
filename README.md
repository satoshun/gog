# go-git

this is tools for used convenience git like go. i called go-git.


## Description

Go have special direcotry structure. i love it so other languages want to use it. go-git achieve it.

## install

for macOS

```
$ brew tap satoshun/commands
$ brew install go-git
```

## Usage

### prepare

```
# first set base directory
$ export GO_GIT_PATH=~/git

# option variable
## if `GO_GIT_HOOK_CMD` variable then run `GO_GIT_HOOK_CMD` before ending.
$ export GO_GIT_HOOK_CMD="cd {{.Directory}} && git status"
```

`GO_GIT_HOOK_CMD` has been prepared under variables. 

- Directory
- Repository
- Host
- Path
- ProjectName

### do

```
$ go-git -h
NAME:
   go-git - use directory like Go

USAGE:
   go-git [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --repository, -r ''  repository url
   --version, -v  print the version

# clone
$ go-git -r git@github.com:satoshun/go-git.git
Cloning into '/Users/satouhayabusa/git/src/github.com/satoshun/go-git'...
warning: You appear to have cloned an empty repository.
Checking connectivity... done.

# update
$ go run main.go -r git@github.com:satoshun/pythonjs.git -u
First, rewinding head to replay your work on top of it...
Fast-forwarded master to 85d86acef567a9729e025defe895df6ee4aa35f7.
```
