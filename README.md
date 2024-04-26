# GoTrash 🗑️

GoTrash is a simple command-line tool written in Go that moves files to the trash instead of deleting them permanently. It provides a safer alternative to the `rm` command, allowing users to recover files if needed.

## Features ✨

- Move files to the trash instead of permanently deleting them
- Simple and easy-to-use command-line tool, with a similar interface to `rm`
- Compatible with most shell environments
- Option to display help information

## Installation 🚀

To use GoTrash, you can either compile the source code yourself or download the precompiled binaries from the [Releases](https://github.com/h4ckm1n-dev/gotrash/releases) page.

### From Source

1. Clone this repository:

```bash
git clone https://github.com/h4ckm1n-dev/gotrash.git
```
Navigate to the repository directory:
```bash
cd gotrash
```
Compile the source code:
```bash
go build
```
Move the generated binary to a directory in your PATH, for example:
```bash
sudo mv gotrash /usr/local/bin
```
Optionaly you can alias gotrash to rm
```bash
alias rm="gotrash"
```
Usage 🛠️
```bash
gotrash [OPTION]... [FILE]...
[OPTION]: Optional arguments, including -h or --help to display help information.
[FILE]...: List of files to move to the trash.
```
Examples
Move a single file to the trash:
```bash
gotrash example.txt
```
Move multiple files to the trash:

```bash
gotrash file1.txt file2.txt file3.txt
```
For more options, run:

```bash
gotrash --help
```
License 📝
This project is licensed under the MIT License - see the LICENSE file for details.

Contributing 🤝
Contributions are welcome! Feel free to open issues or submit pull requests.

Acknowledgments 🙏
This project was inspired by rmtrash by Sindre Sorhus.
