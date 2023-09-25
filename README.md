![preview](./assets/preview.gif)

<div align="center"><i>CLI tool to interact with x (twitter) from terminal</i></div>

## Table of Contents

- [Usage](#usage)
- [Installation](#installation)
  - [Linux](#linux)
  - [Windows](#windows)
  - [Mac](#mac)
- [Authentication](#authentication)

## Usage


### Main Usage

- post tweets from terminal like this

```bash
x -t "first tweet from terminal!"
```

### Full Docs

```text

interact with x (twitter) from terminal.

USAGE
  x <command>

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
    rm -rf /usr/local/x
```

- run the following command to extract the file

```bash
tar -C /usr/local -xzf x_Linux_x86_64.tar.gz
```

- Add `/usr/local/x` to the PATH environment variable
- Do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

```bash
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

`x -t "tweet text"` or `x --tweet "tweet text"`
