# nirvanaminer

`nirvanaminer` is a CPU-based miner for `nirvanad`.

## Requirements

Go 1.19 or later.

## Build from Source

* Install Go according to the installation instructions here:
  http://golang.org/doc/install

* Ensure Go was installed properly and is a supported version:

```bash
go version
```

* Run the following commands to obtain and install `nirvanad`
  including all dependencies:

```bash
git clone https://github.com/Nirvana-Chain/nirvanad
cd nirvanad/cmd/nirvanaminer
go install .
```

* `nirvanaminer` should now be installed in `$(go env GOPATH)/bin`.
  If you did not already add the bin directory to your system path
  during Go installation, you are encouraged to do so now.
  
## Usage

The full `nirvanaminer` configuration options can be seen with:

```bash
nirvanaminer --help
```

But the minimum configuration needed to run it is:

```bash
nirvanaminer --miningaddr=<YOUR_MINING_ADDRESS>
```
