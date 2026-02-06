<div align="center"><i>CLI tool to post tweets on x (twitter) from terminal</i></div>

<br>

![preview](./assets/preview.gif)

## Table of Contents

- [Usage](#usage)
- [Installation](#installation)
  - [Linux](#linux)
  - [Windows](#windows)
  - [Mac](#mac)
- [Authentication](#authentication)
- [Docs](#docs)

## Usage


### Main Usage

- post tweets from terminal like this

```bash
x "first tweet from terminal!"
```

### Full Docs

```bash

interact with x (twitter) from terminal.

USAGE
  x <command>

Main usage (posting a tweet)
  x <tweet-text>
    Examples:
      one word tweet (no spaces)  x hi
      multiple words tweet        x "hi from terminal"
      with optional arg:          x -t "hi there"
      extended optional arg:      x --tweet "hi x"

Commands
  -h             show this help
  auth           start authorizing your X account
  auth --url     get auth url if it didn't open browser after running 'x auth'
  auth -v        verify authorization after running 'x auth'
  -t "text"      post a tweet
  -v             show version
  -c             clear authorized account
```

## Installation

### Linux

- download the verion that matches computer's operating system from [latest release](https://github.com/devhindo/x/releases/latest)
- open a terminal in the directory where the file was downloaded
- for example assume the file name is `x_Linux_x86_64.tar.gz`
- remove any previous downloaded version of x

```bash
sudo rm -rf /usr/local/x
```

- run the following command to extract the file

```bash
sudo tar -C /usr/local -xzf x_Linux_x86_64.tar.gz
```

- Add `/usr/local/x` to the PATH environment variable
- Do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

```bash
# opening the file
nano $HOME/.profile

# and add this line at the end of the file
export PATH=$PATH:/usr/local/x
```

- run the following command to update the current session

```bash
source $HOME/.profile 
```

- verify the installation by running the following command

```bash
x version
```

- if the installation was successful, you should see the version of x that was installed

### Windows

- download the verion that matches your archeticture from [latest release](https://github.com/devhindo/x/releases/latest)
- extract the file and get the `x.exe` file and place it in a directory that is in your PATH environment variable. a common one is `C:\Windows\system32`
- open a terminal and run `x --version` to verify the installation
- if the installation was successful, you should see the version of x that was installed

### Mac

- idk never used it before. but I have darwin binaries in the [latest release](https://github.com/devhindo/x/releases/latest) so you can try it out

## Authentication

- run `x auth`
- an auth url will be opened in your browser, if it doesn't run `x auth --url` to get the url
- verify `X CLI` in the opened twitter page
- after authorizing the app return to the terminal and run `x auth -v`

Note: running `x auth -v` windows might flag the tool as a threat this is because it reads a config file that was previously generated [here](https://github.com/devhindo/x/blob/b00b1a911e1e7ac364dd2d10f941e95d76bfb3ac/src/cli/lock/key.go#L13) to identify the user by a random string that's unique for your device. you can check the function [here](https://github.com/devhindo/x/blob/b00b1a911e1e7ac364dd2d10f941e95d76bfb3ac/src/cli/lock/key.go#L42C1-L42C1). so if the threat shows up, allow it and continue.

- if the authorization was successful, you should see a `verified` message
- now you can post tweets using `x -t "tweet text"`
- if anything went wrong run `x -c` to clear any settings and start over
- if the problem persists, kindly open an issue [here](https://github.com/devhindo/x/issues/new) and describe the problem and I'll be happy to help!


## Posting tweets

`x -t "tweet text"` or `x --tweet "tweet text"` or `x "hi there!"` or even one word tweet like `x spotifyURL`

## Docs
[Docs](https://deepwiki.com/devhindo/x)

## Detailed Documentation

The CLI has been updated to use a more robust command structure.

### Global Options
- `-h, --help`: Help for any command

### Commands

#### `x [message]`
Post a tweet directly.
```bash
x "Hello World"
```

#### `tweet` (alias: `t`)
Post a tweet.
```bash
x tweet "Hello World"
x t "Hello World"
```

#### `auth`
Manage authentication.
- `x auth`: Start authentication flow
- `x auth --verify` (or `-v`): Verify authentication status
- `x auth --clear` (or `-c`): Clear stored credentials
- `x auth --url`: Display the authorization URL

#### `future` (alias: `f`)
Schedule a tweet for later.
```bash
x future "Tweet later" 2h30m
x f "Tweet later" 5h
```
Arguments:
1. Message: The tweet content
2. Duration: Time to wait (e.g., "1h", "30m", "1h30m")

#### `version` (alias: `v`)
Print the CLI version.
```bash
x version
```

### Legacy Support
The following legacy flag styles are still supported for backward compatibility:
- `x -t "msg"` -> `x t "msg"`
- `x -f "msg" "time"` -> `x f "msg" "time"`
- `x -v` -> `x version`
