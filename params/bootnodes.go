// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// SGC Go Bootnode
	"enode://c0e61b6dbb37b477ff2eb411e178154ea1f60a8bbdead903af1c853722c1a135ce3f3f6af95a77bab5c6fc2b8f9918dc0e4cae7411795cdc01e05d62c2605533@18.191.111.10:50501",
	"enode://59731ad09b749d798c74272a7cb576b42ca595603fd6be76b15a52a9fb5a9d5213dd92f0ca2fc13c2bec49e4570c8210b8501cc9fbd5009585641fdb2afa05b5@18.222.122.148:50502",
	"enode://14d0337a631ee421bc4dc81fd6d59ad1d5efe9aa8c16870341cf9f1286a860e09935c0de9d088ae1bd4f3862507704862b3e4387b325ec84a6aaacf8ce57dfad@18.218.176.226:50503",
	"enode://c3e7c0e05f8efca44a8c8b1163dc62aaabed91eb28c73dcd43eb8625ce3fc42089af72a9556e7a663e6dc7e67afa0f141374fb6abb2ec83942e9be7642ae2321@18.218.182.104:50504",
	"enode://42c5a14bb5fe178e57bc344671bb4b665a4c88b854c83025fb9d4554d873f9283994509b7efa4aa19e834bb86a809e455ad06b89d1d8eb4f093cedd59a6d50a0@3.19.215.244:50505",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
