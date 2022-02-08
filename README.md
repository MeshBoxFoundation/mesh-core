# mesh-core 

Official Golang implementation of the meshbox protocol.

<a href="https://meshbox.io/"><img src="logo/meshbox.png" height="200px"/></a>


Welcome to the official Go implementation of meshbox protocol! meshbox is building the next generation of the decentralized blockchain protocol for powering real-world information marketplace in a decentralized-yet-scalable way. Refer to meshbox [whitepaper](https://meshbox.io/research/) for details.

#### New to meshbox?

Please visit https://meshbox.io official website or [meshbox onboard pack](https://onboard.meshbox.io/) to learn more about meshbox network.

#### Run a delegate?

Please visit [meshbox Delegate Manual](https://github.com/meshboxproject/meshbox-bootstrap) for detailed setup process.

## Building the source code

### Minimum requirements

| Components | Version | Description |
|----------|-------------|-------------|
| [Golang](https://golang.org) | &ge; 1.17.3 | Go programming language |
| [Protoc](https://developers.google.com/protocol-buffers/) | &ge; 3.6.0 | Protocol buffers, required only when you rebuild protobuf messages |

### Compile

Download the code to your desired local location (doesn't have to be under `$GOPATH/src`)
```
git clone git@github.com:meshboxproject/mesh-core.git
cd mesh-core
```

If you put the project code under your `$GOPATH\src`, you will need to set up an environment variable
```
export GO111MODULE=on
set GO111MODULE=on (for windows)
```

Build the project for general purpose (server, ioctl) by

```
make
```

Build the project for broader purpose (server, ioctl, injector...) by
```
make all 
```

If the dependency needs to be updated, run

```
go get -u
go mod tidy
```
If you want learn more advanced usage about `go mod`, you can find out [here](https://github.com/golang/go/wiki/Modules).

Run unit tests only by

```
make test
```

Build the docker image by

```
make docker
```

### Run mesh-core

Start (or resume) a standalone server to operate on an blockchain by

```
make run
```

Restart the server from a clean state by

```
make reboot
```

If "make run" fails due to corrupted or missing state database while block database is in normal condition, e.g.,
failing to get factory's height from underlying DB, please try to recover state database by

```
make recover
```

Then, "make run" again.

### Use CLI

Users could interact with meshbox blockchain by

```
ioctl [command]
```

Refer to [CLI document](https://docs.meshbox.io/developer/ioctl/install.html) for more details.

## Contact

- Mailing list: [meshbox-dev](meshbox-dev@meshbox.io)
- Dev Forum: [forum](https://community.meshbox.io/c/research-development/protocol)
- Bugs: [issues](https://github.com/meshboxproject/mesh-core/issues)

## Contribution
We are glad to have contributors out of the core team; contributions, including (but not limited to) style/bug fixes,
implementation of features, proposals of schemes/algorithms, and thorough documentation, are welcomed. Please refer to
our [contribution guideline](CONTRIBUTING.md) for more
information. Development guide documentation is [here](https://github.com/meshboxproject/mesh-core/wiki/Developers%27-Guide).

For any major protocol level changes, we use [IIP](https://github.com/meshboxproject/iips) to track the proposal, decision
and etc.

## License
This project is licensed under the [Apache License 2.0](LICENSE).
