package utils

import (
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
    err := open(url)
    if err != nil {
        panic(err)
    }
}

func open(url string) error {

	// source : https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8

    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
      cmd = "rundll32"
      args = []string{"url.dll,FileProtocolHandler"}
    case "darwin":
        cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
        cmd = "xdg-open"
    }
    args = append(args, url)
    return exec.Command(cmd, args...).Start()
}