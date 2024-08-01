# CITS3200 Project - Academic Contact Finder

This program is designed to make it easier to find any available email addresses for a particular academic.

## Building

**Requirements**

- [Go](https://go.dev/)
- [GCC](https://gcc.gnu.org/)
    - [w64devkit](https://github.com/skeeto/w64devkit) works well on Windows
- [Fyne CLI tool](https://docs.fyne.io/started/packaging)

These tools' commands must be accessible through a command line/terminal.

**Compiling**

Built using the [Fyne](https://fyne.io/) GUI framework for Go, this program can be compiled to a single binary with the following commands:

```
git clone https://github.com/odddollar/CITS3200-Project.git
cd CITS3200-Project
fyne package --release
```

For cross-compiling for other platforms, use [Fyne-cross](https://github.com/fyne-io/fyne-cross) (requires [Docker](https://www.docker.com/)).
