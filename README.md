# License Generator

License generator is a cli application helping to generate the `LICENSE` file for any open source project.

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

Put the binary inside your `/usr/bin` and give you the right to execute it

```bash
$ sudo cp ./license-generator /usr/bin/license-generator
$ sudo chmod +x /usr/bin/license-generator
```

And use it!

## Technologies

- Go 1.19