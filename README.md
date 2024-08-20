# CITS3200 Project - Academic Contact Finder

This program is designed to make it easier to find any available email addresses for a particular academic.

## Building

**Requirements**

- [Go](https://go.dev/)
- [GCC](https://gcc.gnu.org/)
    - [w64devkit](https://github.com/skeeto/w64devkit) works well on Windows
- [Fyne](https://fyne.io/) CLI tool
    - Can be installed with `go install fyne.io/fyne/v2/cmd/fyne@latest`

Will by default compile for the current system (e.g. Windows, MacOS). For cross-compiling to other platforms, use [Fyne-cross](https://github.com/fyne-io/fyne-cross) (requires [Docker](https://www.docker.com/)).

These tools' commands must be accessible through a command line/terminal.

**Compiling for development**

Run the following command from the program's root directory:

```
go generate
go run .
```

First compilation *will* take a while. Subsequent compilations are significantly faster.

**Compiling for release**

This will package everything into a single file that can be distributed without any dependencies. Run:

```
go generate
fyne package --release
```

