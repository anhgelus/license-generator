# License Generator

License generator is a cli application helping to generate the `LICENSE` file for any open source project.

## Installation

To install the application, you need Go 1.22.

And run this command to install it:
```bash
$ go install github.com/anhgelus/license-generator@latest
```

If after installing the application your shell is saying "Command not found", it means that your `$PATH` is not correctly set.
To fix this issue, just enter this command (it will fix your path by adding the binaries installed by `go install`):
```bash
$ export PATH=${PATH}:`go env GOPATH`/bin
```

## Build

To build this application, you need Go 1.19

First of all, install every dependency

```bash
$ go get
```

Now you can build the project

```bash
$ go build .
```

Use the script `setup.sh` to install the binaries!

```bash
$ sh setup.sh
```

And use it!

## How to use

After installing the application, you can use it with a terminal.

To see the help, just use `license-generator -h`.

Now, let's see how to create a new license.

### Creating a new license

To create a new license, just run `license-generator` and answer every question.

The first question is the name of your application. Here it's *license-generator*: this is basically the name of the project.

Then, it will ask you each license do you want to use. Answer with the identifier of the license. For example, it's `gpl` *for GPLv3*.

To check every available license, use `license-generator -l`.

Next, just answer with the authors of the program and separate each with a coma (,) if there are more than one author.

Finally, answer with the year. It can be `2023` or `2020 - 2023`. 

After answering these questions, the LICENSE should be generated!

### Creating a new license with arguments

You can also use arguments to create a new license.

- `--name` is for the name
- `--license` is for the license identifier
- `--year` is for the year
- `--authors` if for authors, actually, it's not working with double quote (")!

The uninformed arguments will be asked with the same questions as in the previous part. 

### Custom license

If you want to add custom license to this program, you can! But you must follow this wiki.

#### Create the configuration

Create a new folder.
Create a new file entitled `config.toml`. This file is not required but recommended.
Paste this content inside:
```toml
customLicenses = []
```
In the field `customLicenses`, you will put every enabled licenses. When you want to enable a license, just add the file name in the array. When you want to disable a license, just remove the file name from the array.

Create a new file entitled `your_license.toml`.
Paste this content inside:
```toml
path = "./cc0.license"
name = "Creative Commons 0"
identifier = "cc0"
```
In `path`, put a path (relative or absolute) to the license file. In `name`, put the name of the license. In `identifier`, put the identifier of the license (example: `cc0` for Creative Commons 0 or `gpl` for GPLv3), the identifier is used when we ask you wich license do you want to use.

Create a new file according to the path put in the `path` variable.
Put your license inside and replace the year(s) by `{{ .Year }}`, the author(s) by `{{ .Authors }}` and the project name by `{{ .AppName }}`.

#### Use the configuration

Now, to use this configuration, you just need to add  `--config-path PATH_TO_YOUR_CONFIG` when using the command. Just, do not forget to replace the `PATH_TO_YOUR_CONFIG` by the path to your custom config (relative or absolute).
When the program will ask you wich license do you want to use, just use the identifier you put inside the file.

## Technologies

- Go 1.19
