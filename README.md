<p align="center">
  <img src="./assets/x.png" alt="Logo" height=170>
</p>

<p align="center">

<div align="center"><i>interact with x api and perform tasks like posting tweets or whatever</i></div>

## Table of Contents

- [Usage](#usage)
- [Installation](#installation)
- [Authentication](#authentication)

## Usage

### Main Usage
- post a tweet from terminal like this

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
  - for example assume the file name is `x-linux-amd64.tar.gz`
  - remove any previous downloaded version of x 

  	```bash
  	rm -rf /usr/local/x
  	```
  - run the following command to extract the file 
  	```bash
	tar -C /usr/local -xzf x-linux-amd64.tar.gz
	```
  - Add `/usr/local/x` to the PATH environment variable
  - Do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
	```bash 
	export PATH=$PATH:/usr/local/x
	```
  - open a new terminal or run the following command to update the current session 

  	```bash
	source $HOME/.profile 
  	```
  - verify the installation by running the following command
  	```bash
	x --version
	```
  - if the installation was successful, you should see the version of x that was installed

### Windows
  - download the verion that matches your archeticture from [latest release](https://github.com/devhindo/x/releases/latest)

### Mac


## Authentication

- run `x auth` 
- after authorizing the app
- run `x auth --verify` or `x auth -v`

## Posting tweets

`x -t "tweet text"` or `x --tweet "tweet text"`
