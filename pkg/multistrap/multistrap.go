package multistrap

import (
	"os"
	"os/exec"
	"strings"
)

type Options struct {
	Arch       string
	Directory  string
	Suite      string
	Components []string
	Packages   []string
}

func Run(options Options) error {
	cmd := exec.Command("multistrap", "-a", options.Arch, "-d", options.Directory, "-f", "/dev/stdin")
	cmd.Stdout = os.Stdout // cmd.Stdout -> stdout
   	cmd.Stderr = os.Stderr // cmd.Stderr -> stderr
	cmd.Stdin = strings.NewReader(`
[General]
noauth=true
unpack=true
allowrecommends=true
debootstrap=Debian
aptsources=Debian

[Debian]
source=http://cdn-fastly.deb.debian.org/debian/
keyring=debian-archive-keyring
components=` + strings.Join(options.Components, " ") + `
suite=` + options.Suite + `
packages=` + strings.Join(options.Packages, " "))

	return cmd.Run()
}
