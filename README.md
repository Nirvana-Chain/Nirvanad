# Nirvanad

[![Build Status](https://github.com/Nirvana-Chain/nirvanad/actions/workflows/tests.yaml/badge.svg)](https://github.com/Nirvana-Chain/nirvanad/actions/workflows/tests.yaml)
[![Code Coverage](https://codecov.io/gh/Nirvana-Chain/nirvanad/graph/badge.svg)](https://codecov.io/gh/Nirvana-Chain/nirvanad)
[![GitHub release](https://img.shields.io/github/v/release/Nirvana-Chain/nirvanad.svg)](https://github.com/Nirvana-Chain/nirvanad/releases)
[![GitHub license](https://img.shields.io/github/license/Nirvana-Chain/nirvanad.svg)](https://github.com/Nirvana-Chain/nirvanad/blob/master/LICENSE)
[![GitHub downloads](https://img.shields.io/github/downloads/Nirvana-Chain/nirvanad/total.svg)](https://github.com/Nirvana-Chain/nirvanad/releases)
[![Discord users](https://img.shields.io/discord/1169939685280337930.svg)](https://discord.gg/ZPZRvgMJDT)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/Nirvana-Chain/nirvanad/)

Nirvanad is the reference full node Nirvana implementation written in
Go (golang). It is a [blockDAG](https://en.wikipedia.org/wiki/Directed_acyclic_graph)
as a proof-of-work cryptocurrency with instant confirmations and
sub-second block times with a work-in-progress (WIP) attempt to
implement smart contracts as a side-chain. It is based on
[the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a
generalization of Nakamoto consensus.

## Overview

Nirvana is a fork of [Kaspa](https://github.com/kaspanet/kaspad)
introducing a GPU-centric fork as a solution to the dominance of ASIC
mining farms, aiming to empower small-scale miners and enhance
decentralization. We focus on bridging the gap between blockchain
technology, decentralized finance and the real world of payment systems
and traditional finance.

With Kaspa, our approach is one of friendly cohabitation. We operate
under the same protocol, enjoy similar advantages, and face
unidentified issues. Any significant improvements that need to be made
in the primary codebase will be shared back.

## Small Scale Miners

The Nirvana Network team believes in decentralization and small-scale
miners. We will ensure long-term GPU-friendly mining.

### Hashing Function

We initially started with `kHeavyHash` and `blake3` modifications
on-top. This algorithm is called `NirvanaHashv1`. However `kHeavyHash`
and `blake3` are not future proof in ASIC resistence. Therefore we've
launched already our `testnet-1` with [FishHash](https://github.com/iron-fish/fish-hash/blob/main/FishHash.pdf).
It is the worlds first implementation of FishHash with Golang in a
1bps blockchain.

`NirvanaHashv1` is currently used in [mainnet](https://github.com/Nirvana-Chain/nirvanad/releases/tag/v1.1.0)
and can be mined using the following miners maintained by the Nirvana
developers:

* Built-in CPU miner from `nirvanad`
* Nirvana [GPU miner](https://github.com/Nirvana-Chain/nirvana-miner) as reference implementation of `kHeavyHash` with `blake3`.

The following third-party miners are available and have added
`NirvanaHashv1`:

* [lolMiner](https://github.com/Lolliedieb/lolMiner-releases)
* [Team Red Miner](https://github.com/todxx/teamredminer)
* [SRBMiner](https://github.com/doktor83/SRBMiner-Multi)
* [BzMiner](https://github.com/bzminer/bzminer)
* [Rigel](https://github.com/rigelminer/rigel)
* [GMiner](https://github.com/develsoftware/GMinerRelease)

`NirvanaHashv2` is currently being investigated and tested in [testnet-1](https://github.com/Nirvana-Chain/nirvanad/releases/tag/v2.0.0-testnet-1-fishhash)
and can be mined using the following miners maintained by the Nirvana
developers:

* Built-in CPU miner from `nirvanad`
* Nirvana [GPU miner](https://github.com/wam-rd/nirvana-miner/releases/tag/v2.0.0-alpha) as bleeding edge and unoptimized reference implementation of FishHash.

There are no third-party miners available as of now.

## Smart Contracts

The Nirvana Network team is launching an R&D project to connect the
Nirvana blockchain with other blockchain ecosystems using a smart
contract layer based on the [Cosmos SDK](https://v1.cosmos.network/sdk).

This initiative aims to enhance interoperability, efficiency, and
innovation in the blockchain space. By leveraging the Cosmos SDK's
advanced features, we'll facilitate seamless transactions across
different networks, offering new opportunities for users and
developers.

### Cosmos Hub

[Cosmos](https://cosmos.network/) is a highly attractive ecosystem due
to its innovative approach to blockchain interoperability, scalability,
and usability. By enabling different blockchains to seamlessly
communicate and exchange value, Cosmos opens up vast opportunities for
businesses and developers to build and deploy decentralized
applications that can operate across multiple blockchain environments.
This interoperability fosters a more connected and efficient digital
economy, potentially driving adoption and usage across various sectors,
including finance, supply chain, and beyond.

By connecting to Cosmos, we will open the door of a web3 ecosystem
connected to a complete network of other blockchain project, making
Nirvana Network more competitive and adaptable in the rapidly evolving
landscape of decentralized technologies.

### Nirvana Sidechain

The creation of the Nirvana [sidechain](https://github.com/john-light/sidechains)
with fast transaction times, smart contract capabilities, and a dual
coin model, designed to integrate seamlessly with the Cosmos ecosystem
and utilize the Nirvana (NVA) across all interconnected platforms,
signifies a strategic advancement in Nirvana Network.

This sidechain will not only enable quick and efficient
inter-blockchain transactions but also support complex decentralized
applications through its smart contract functionality.

The ability to use Nirvana (NVA) across the entire ecosystem will
ensure a unified and streamlined user experience, promoting greater
adoption and utility within the Cosmos network. This sidechain aims
to enhance scalability, foster innovation, and provide a flexible and
user-centric blockchain solution that meets the diverse needs of
developers, users, and investors within the Cosmos ecosystem.

## Installation

### Install from Binaries

Pre-compiled binaries for Linux `x86_64` and `aarch64`, Windows `x64`
and macOS `x64` as universal binary can be downloaded at: [https://github.com/Nirvana-Chain/nirvanad/releases](https://github.com/Nirvana-Chain/nirvanad/releases)

### Build from Source

Go 1.19 or later is required. Install Go according to the installation
instructions at [http://golang.org/doc/install](http://golang.org/doc/install).
Ensure Go was installed properly and is a supported version:

```bash
go version
```

Run the following commands to obtain and install nirvanad including
all dependencies:

```bash
git clone https://github.com/Nirvana-Chain/nirvanad
cd nirvanad
go install . ./cmd/...
```

Nirvanad (and utilities) should now be installed in
`$(go env GOPATH)/bin`. If you did not already add the `bin` directory
to your system path during Go installation, you are encouraged to do
so now.


### Getting Started

Nirvanad has several configuration options available to tweak how it
runs, but all of the basic operations work with zero configuration.

- Add the following line to the end of the file:

```bash
 export PATH=$PATH:$(go env GOPATH)/bin
```

## Creating a wallet

- To create a wallet, you need to run nirvanad with utxoindex

```bash
$ nirvanad --utxoindex
```
- Open another terminal

```bash
$ nirvanawallet create
```

- You will be asked to choose a password for the wallet (a password must be at least 8 characters long, and it won't be shown on the screen you as you entering it). After that you should run this command in order to start the wallet daemon:

```bash
$ nirvanawallet start-daemon
```
- Do not close the first 2 terminals and open a new terminal and then run this in order to request an address from the wallet:

```bash
$ nirvanawallet new-address
```

- Your screen will show you something like this:

The wallet address is:
nirvana:qpff4xzlzfqertcuvztp0r8qlk59sd6y2asamplepm7nmryckumfy07nhwfur

- To see your secret seed phrase :

```bash
$ nirvanawallet dump-unencrypted-data
```

Note: Every time you ask nirvanawallet for an address you will get a different address. This is perfectly fine. Every secret key is associated with many different public addresses and there is no reason not to use a fresh one for each transaction.

At this point your can close the wallet daemon, though you should keep it running of you want to be able to check your balance and make transactions

## Telegram

- Join our Telegram server using the following link: [https://t.me/Nirvanad_Chain/](https://t.me/Nirvanad_Chain)
- 加入Telegram中文频道: [https://t.me/Nirvanad_Chain_CN/](https://t.me/Nirvanad_Chain_CN)

## Issue Tracker

The [integrated github issue tracker](https://github.com/Nirvana-Chain/nirvanad/issues)
is used for this project.

## Documentation

The [documentation](https://github.com/Nirvana-Chain/docs) is a
work-in-progress.

## License

Nirvanad is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).

## Donate:
#### Your donation is our biggest motivation for open source
- BTC：bc1qz2upj7qnw6wftpqsc5t53uexeczzdsfcxml49d
- ETH：0x07f8dea7297651f3c9720f06d856fa99ed9aa320
- NVA：nirvana:qqlk0atms6xvl266t9cs9klpcxfawlnp50xr8qesq2exa9sln9ylwv8mf30wj