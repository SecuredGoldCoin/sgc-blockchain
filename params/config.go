// Copyright 2016 The go-ethereum Authors
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

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0x35ac7a156d894756c089aa8deaf1574cc523d1573f37c55da41663f16cde55a3")
	TestnetGenesisHash = common.HexToHash("")
	RinkebyGenesisHash = common.HexToHash("")
)

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(786),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.Hash{},
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),

		//Premine Checks for First 5 Years
		Premine2019Start: big.NewInt(50),
		Premine2020Start: big.NewInt(613788),
		Premine2021Start: big.NewInt(2788684),
		Premine2022Start: big.NewInt(4963581),
		Premine2023Start: big.NewInt(7138478),

		// Premine Checks for Year 2024-2042
		Premine2024Start: big.NewInt(9313374),
		Premine2025Start: big.NewInt(11488271),
		Premine2026Start: big.NewInt(13663167),
		Premine2027Start: big.NewInt(15838064),
		Premine2028Start: big.NewInt(18012960),
		Premine2029Start: big.NewInt(20187857),
		Premine2030Start: big.NewInt(22362753),
		Premine2031Start: big.NewInt(24537650),
		Premine2032Start: big.NewInt(26712547),
		Premine2033Start: big.NewInt(28887443),
		Premine2034Start: big.NewInt(31062340),
		Premine2035Start: big.NewInt(33237236),
		Premine2036Start: big.NewInt(35412133),
		Premine2037Start: big.NewInt(37587029),
		Premine2038Start: big.NewInt(39761926),
		Premine2039Start: big.NewInt(41936822),
		Premine2040Start: big.NewInt(44111719),
		Premine2041Start: big.NewInt(46286616),
		Premine2042Start: big.NewInt(48461512),

		//Premine Checks for Year 2043-2062
		Premine2043Start: big.NewInt(50636409),
		Premine2044Start: big.NewInt(52811305),
		Premine2045Start: big.NewInt(54986202),
		Premine2046Start: big.NewInt(57161098),
		Premine2047Start: big.NewInt(59335995),
		Premine2048Start: big.NewInt(61510891),
		Premine2049Start: big.NewInt(63685788),
		Premine2050Start: big.NewInt(65860684),
		Premine2051Start: big.NewInt(68035581),
		Premine2052Start: big.NewInt(70210478),
		Premine2053Start: big.NewInt(72385374),
		Premine2054Start: big.NewInt(74560271),
		Premine2055Start: big.NewInt(76735167),
		Premine2056Start: big.NewInt(78910064),
		Premine2057Start: big.NewInt(81084960),
		Premine2058Start: big.NewInt(83259857),
		Premine2059Start: big.NewInt(85434753),
		Premine2060Start: big.NewInt(87609650),
		Premine2061Start: big.NewInt(89784547),
		Premine2062Start: big.NewInt(91959443),
		PremineEnd:       big.NewInt(91970000),
		Ethash:           new(EthashConfig),
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "mainnet",
		SectionIndex: 208,
		SectionHead:  common.HexToHash(""),
		CHTRoot:      common.HexToHash(""),
		BloomRoot:    common.HexToHash(""),
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil}
	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, &CliqueConfig{Period: 0, Epoch: 30000}}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	Name         string      `json:"-"`
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	Premine2019Start    *big.Int `json:"premine2019Start,omitempty"`    // Premine2019Start switch block (nil = no fork, 0 = already activated)
	Premine2020Start    *big.Int `json:"premine2020Start,omitempty"`    // Premine2020Start switch block (nil = no fork, 0 = already activated)
	Premine2021Start    *big.Int `json:"premine2021Start,omitempty"`    // Premine2021Start switch block (nil = no fork, 0 = already activated)
	Premine2022Start    *big.Int `json:"premine2022Start,omitempty"`    // Premine2022Start switch block (nil = no fork, 0 = already activated)
	Premine2023Start    *big.Int `json:"premine2023Start,omitempty"`    // Premine2023Start switch block (nil = no fork, 0 = already activated)

	//Premine Conditions for Year 2024 - 2042
	Premine2024Start *big.Int `json:"premine2024Start,omitempty"` // Premine2024Start switch block (nil = no fork, 0 = already activated)
	Premine2025Start *big.Int `json:"premine2025Start,omitempty"` // Premine2025Start switch block (nil = no fork, 0 = already activated)
	Premine2026Start *big.Int `json:"premine2026Start,omitempty"` // Premine2026Start switch block (nil = no fork, 0 = already activated)
	Premine2027Start *big.Int `json:"premine2027Start,omitempty"` // Premine2027Start switch block (nil = no fork, 0 = already activated)
	Premine2028Start *big.Int `json:"premine2028Start,omitempty"` // Premine2028Start switch block (nil = no fork, 0 = already activated)
	Premine2029Start *big.Int `json:"premine2029Start,omitempty"` // Premine2029Start switch block (nil = no fork, 0 = already activated)
	Premine2030Start *big.Int `json:"premine2030Start,omitempty"` // Premine2030Start switch block (nil = no fork, 0 = already activated)
	Premine2031Start *big.Int `json:"premine2031Start,omitempty"` // Premine2031Start switch block (nil = no fork, 0 = already activated)
	Premine2032Start *big.Int `json:"premine2032Start,omitempty"` // Premine2032Start switch block (nil = no fork, 0 = already activated)
	Premine2033Start *big.Int `json:"premine2033Start,omitempty"` // Premine2033Start switch block (nil = no fork, 0 = already activated)
	Premine2034Start *big.Int `json:"premine2034Start,omitempty"` // Premine2034Start switch block (nil = no fork, 0 = already activated)
	Premine2035Start *big.Int `json:"premine2035Start,omitempty"` // Premine2035Start switch block (nil = no fork, 0 = already activated)
	Premine2036Start *big.Int `json:"premine2036Start,omitempty"` // Premine2036Start switch block (nil = no fork, 0 = already activated)
	Premine2037Start *big.Int `json:"premine2037Start,omitempty"` // Premine2037Start switch block (nil = no fork, 0 = already activated)
	Premine2038Start *big.Int `json:"premine2038Start,omitempty"` // Premine2038Start switch block (nil = no fork, 0 = already activated)
	Premine2039Start *big.Int `json:"premine2039Start,omitempty"` // Premine2039Start switch block (nil = no fork, 0 = already activated)
	Premine2040Start *big.Int `json:"premine2040Start,omitempty"` // Premine2040Start switch block (nil = no fork, 0 = already activated)
	Premine2041Start *big.Int `json:"premine2041Start,omitempty"` // Premine2041Start switch block (nil = no fork, 0 = already activated)
	Premine2042Start *big.Int `json:"premine2042Start,omitempty"` // Premine2042Start switch block (nil = no fork, 0 = already activated)

	//Premine Conditions for Year 2043-2062
	Premine2043Start *big.Int `json:"premine2043Start,omitempty"` // Premine2043Start switch block (nil = no fork, 0 = already activated)
	Premine2044Start *big.Int `json:"premine2044Start,omitempty"` // Premine2044Start switch block (nil = no fork, 0 = already activated)
	Premine2045Start *big.Int `json:"premine2045Start,omitempty"` // Premine2045Start switch block (nil = no fork, 0 = already activated)
	Premine2046Start *big.Int `json:"premine2046Start,omitempty"` // Premine2046Start switch block (nil = no fork, 0 = already activated)
	Premine2047Start *big.Int `json:"premine2047Start,omitempty"` // Premine2047Start switch block (nil = no fork, 0 = already activated)
	Premine2048Start *big.Int `json:"premine2048Start,omitempty"` // Premine2048Start switch block (nil = no fork, 0 = already activated)
	Premine2049Start *big.Int `json:"premine2049Start,omitempty"` // Premine2049Start switch block (nil = no fork, 0 = already activated)
	Premine2050Start *big.Int `json:"premine2050Start,omitempty"` // Premine2050Start switch block (nil = no fork, 0 = already activated)
	Premine2051Start *big.Int `json:"premine2051Start,omitempty"` // Premine2051Start switch block (nil = no fork, 0 = already activated)
	Premine2052Start *big.Int `json:"premine2052Start,omitempty"` // Premine2052Start switch block (nil = no fork, 0 = already activated)
	Premine2053Start *big.Int `json:"premine2053Start,omitempty"` // Premine2053Start switch block (nil = no fork, 0 = already activated)
	Premine2054Start *big.Int `json:"premine2054Start,omitempty"` // Premine2054Start switch block (nil = no fork, 0 = already activated)
	Premine2055Start *big.Int `json:"premine2055Start,omitempty"` // Premine2055Start switch block (nil = no fork, 0 = already activated)
	Premine2056Start *big.Int `json:"premine2056Start,omitempty"` // Premine2056Start switch block (nil = no fork, 0 = already activated)
	Premine2057Start *big.Int `json:"premine2057Start,omitempty"` // Premine2057Start switch block (nil = no fork, 0 = already activated)
	Premine2058Start *big.Int `json:"premine2058Start,omitempty"` // Premine2058Start switch block (nil = no fork, 0 = already activated)
	Premine2059Start *big.Int `json:"premine2059Start,omitempty"` // Premine2059Start switch block (nil = no fork, 0 = already activated)
	Premine2060Start *big.Int `json:"premine2060Start,omitempty"` // Premine2060Start switch block (nil = no fork, 0 = already activated)
	Premine2061Start *big.Int `json:"premine2061Start,omitempty"` // Premine2061Start switch block (nil = no fork, 0 = already activated)
	Premine2062Start *big.Int `json:"premine2062Start,omitempty"` // Premine2062Start switch block (nil = no fork, 0 = already activated)
	PremineEnd       *big.Int `json:"premineEnd,omitempty"`       // PremineEnd switch block (nil = no fork, 0 = already activated)
	EWASMBlock       *big.Int `json:"ewasmBlock,omitempty"`       // EWASM switch block (nil = no fork, 0 = already activated)

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	default:
		engine = "unknown"
	}
	return fmt.Sprintf(
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.Premine2019Start,
		c.Premine2020Start,
		c.Premine2021Start,
		c.Premine2022Start,
		c.Premine2023Start,
		c.Premine2024Start,
		c.Premine2025Start,
		c.Premine2026Start,
		c.Premine2027Start,
		c.Premine2028Start,
		c.Premine2029Start,
		c.Premine2030Start,
		c.Premine2031Start,
		c.Premine2032Start,
		c.Premine2033Start,
		c.Premine2034Start,
		c.Premine2035Start,
		c.Premine2036Start,
		c.Premine2037Start,
		c.Premine2038Start,
		c.Premine2039Start,
		c.Premine2040Start,
		c.Premine2041Start,
		c.Premine2042Start,
		c.Premine2043Start,
		c.Premine2044Start,
		c.Premine2045Start,
		c.Premine2046Start,
		c.Premine2047Start,
		c.Premine2048Start,
		c.Premine2049Start,
		c.Premine2050Start,
		c.Premine2051Start,
		c.Premine2052Start,
		c.Premine2053Start,
		c.Premine2054Start,
		c.Premine2055Start,
		c.Premine2056Start,
		c.Premine2057Start,
		c.Premine2058Start,
		c.Premine2059Start,
		c.Premine2060Start,
		c.Premine2061Start,
		c.Premine2062Start,
		c.PremineEnd,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsPremine2019Start
func (c *ChainConfig) IsPremine2019Start(num *big.Int) bool {
	return isForked(c.Premine2019Start, num)
}

// IsPremine2020Start
func (c *ChainConfig) IsPremine2020Start(num *big.Int) bool {
	return isForked(c.Premine2020Start, num)
}

// IsPremine2021Start
func (c *ChainConfig) IsPremine2021Start(num *big.Int) bool {
	return isForked(c.Premine2021Start, num)
}

// IsPremine2022Start
func (c *ChainConfig) IsPremine2022Start(num *big.Int) bool {
	return isForked(c.Premine2022Start, num)
}

// IsPremine2023Start
func (c *ChainConfig) IsPremine2023Start(num *big.Int) bool {
	return isForked(c.Premine2023Start, num)
}

// IsPremine2024Start
func (c *ChainConfig) IsPremine2024Start(num *big.Int) bool {
	return isForked(c.Premine2024Start, num)
}

// IsPremine2025Start
func (c *ChainConfig) IsPremine2025Start(num *big.Int) bool {
	return isForked(c.Premine2025Start, num)
}

// IsPremine2026Start
func (c *ChainConfig) IsPremine2026Start(num *big.Int) bool {
	return isForked(c.Premine2026Start, num)
}

// IsPremine2027Start
func (c *ChainConfig) IsPremine2027Start(num *big.Int) bool {
	return isForked(c.Premine2027Start, num)
}

// IsPremine2028Start
func (c *ChainConfig) IsPremine2028Start(num *big.Int) bool {
	return isForked(c.Premine2028Start, num)
}

// IsPremine2029Start
func (c *ChainConfig) IsPremine2029Start(num *big.Int) bool {
	return isForked(c.Premine2029Start, num)
}

// IsPremine2030Start
func (c *ChainConfig) IsPremine2030Start(num *big.Int) bool {
	return isForked(c.Premine2030Start, num)
}

// IsPremine2031Start
func (c *ChainConfig) IsPremine2031Start(num *big.Int) bool {
	return isForked(c.Premine2031Start, num)
}

// IsPremine2032Start
func (c *ChainConfig) IsPremine2032Start(num *big.Int) bool {
	return isForked(c.Premine2032Start, num)
}

// IsPremine2033Start
func (c *ChainConfig) IsPremine2033Start(num *big.Int) bool {
	return isForked(c.Premine2033Start, num)
}

// IsPremine2034Start
func (c *ChainConfig) IsPremine2034Start(num *big.Int) bool {
	return isForked(c.Premine2034Start, num)
}

// IsPremine2035Start
func (c *ChainConfig) IsPremine2035Start(num *big.Int) bool {
	return isForked(c.Premine2035Start, num)
}

// IsPremine2036Start
func (c *ChainConfig) IsPremine2036Start(num *big.Int) bool {
	return isForked(c.Premine2036Start, num)
}

// IsPremine2037Start
func (c *ChainConfig) IsPremine2037Start(num *big.Int) bool {
	return isForked(c.Premine2037Start, num)
}

// IsPremine2038Start
func (c *ChainConfig) IsPremine2038Start(num *big.Int) bool {
	return isForked(c.Premine2038Start, num)
}

// IsPremine2039Start
func (c *ChainConfig) IsPremine2039Start(num *big.Int) bool {
	return isForked(c.Premine2039Start, num)
}

// IsPremine2040Start
func (c *ChainConfig) IsPremine2040Start(num *big.Int) bool {
	return isForked(c.Premine2040Start, num)
}

// IsPremine2041Start
func (c *ChainConfig) IsPremine2041Start(num *big.Int) bool {
	return isForked(c.Premine2041Start, num)
}

// IsPremine2042Start
func (c *ChainConfig) IsPremine2042Start(num *big.Int) bool {
	return isForked(c.Premine2042Start, num)
}

// IsPremine2043Start
func (c *ChainConfig) IsPremine2043Start(num *big.Int) bool {
	return isForked(c.Premine2043Start, num)
}

// IsPremine2044Start
func (c *ChainConfig) IsPremine2044Start(num *big.Int) bool {
	return isForked(c.Premine2044Start, num)
}

// IsPremine2045Start
func (c *ChainConfig) IsPremine2045Start(num *big.Int) bool {
	return isForked(c.Premine2045Start, num)
}

// IsPremine2046Start
func (c *ChainConfig) IsPremine2046Start(num *big.Int) bool {
	return isForked(c.Premine2046Start, num)
}

// IsPremine2047Start
func (c *ChainConfig) IsPremine2047Start(num *big.Int) bool {
	return isForked(c.Premine2047Start, num)
}

// IsPremine2048Start
func (c *ChainConfig) IsPremine2048Start(num *big.Int) bool {
	return isForked(c.Premine2048Start, num)
}

// IsPremine2049Start
func (c *ChainConfig) IsPremine2049Start(num *big.Int) bool {
	return isForked(c.Premine2049Start, num)
}

// IsPremine2050Start
func (c *ChainConfig) IsPremine2050Start(num *big.Int) bool {
	return isForked(c.Premine2050Start, num)
}

// IsPremine2051Start
func (c *ChainConfig) IsPremine2051Start(num *big.Int) bool {
	return isForked(c.Premine2051Start, num)
}

// IsPremine2052Start
func (c *ChainConfig) IsPremine2052Start(num *big.Int) bool {
	return isForked(c.Premine2052Start, num)
}

// IsPremine2053Start
func (c *ChainConfig) IsPremine2053Start(num *big.Int) bool {
	return isForked(c.Premine2053Start, num)
}

// IsPremine2054Start
func (c *ChainConfig) IsPremine2054Start(num *big.Int) bool {
	return isForked(c.Premine2054Start, num)
}

// IsPremine2055Start
func (c *ChainConfig) IsPremine2055Start(num *big.Int) bool {
	return isForked(c.Premine2055Start, num)
}

// IsPremine2056Start
func (c *ChainConfig) IsPremine2056Start(num *big.Int) bool {
	return isForked(c.Premine2056Start, num)
}

// IsPremine2057Start
func (c *ChainConfig) IsPremine2057Start(num *big.Int) bool {
	return isForked(c.Premine2057Start, num)
}

// IsPremine2058Start
func (c *ChainConfig) IsPremine2058Start(num *big.Int) bool {
	return isForked(c.Premine2058Start, num)
}

// IsPremine2059Start
func (c *ChainConfig) IsPremine2059Start(num *big.Int) bool {
	return isForked(c.Premine2059Start, num)
}

// IsPremine2060Start
func (c *ChainConfig) IsPremine2060Start(num *big.Int) bool {
	return isForked(c.Premine2060Start, num)
}

// IsPremine2061Start
func (c *ChainConfig) IsPremine2061Start(num *big.Int) bool {
	return isForked(c.Premine2061Start, num)
}

// IsPremine2062Start
func (c *ChainConfig) IsPremine2062Start(num *big.Int) bool {
	return isForked(c.Premine2062Start, num)
}

// IsPremineEnd
func (c *ChainConfig) IsPremineEnd(num *big.Int) bool {
	return isForked(c.PremineEnd, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsConstantinople(num):
		return GasTableConstantinople
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                   *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158 bool
	IsByzantium, IsConstantinople             bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
	}
}
