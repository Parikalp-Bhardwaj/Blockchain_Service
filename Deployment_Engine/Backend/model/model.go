package model

// type EthashConfig struct {
// }

type GethUTC struct {
	Address string `json:"address"`
}

type AllocEntry struct {
	Balance string            `json:"balance,omitempty"`
	Code    string            `json:"code,omitempty"`
	Storage map[string]string `json:"storage,omitempty"`
}

type YourStruct struct {
	Alloc map[string]AllocEntry `json:"alloc"`
	Utc   AllocEntry            `json:"utc"`
}

type JSONData struct {
	Config             Config                `json:"config"`
	Nonce              string                `json:"nonce"`
	Timestamp          string                `json:"timestamp"`
	ExtraData          string                `json:"extraData"`
	GasLimit           string                `json:"gasLimit"`
	Difficulty         string                `json:"difficulty"`
	MixHash            string                `json:"mixHash"`
	Coinbase           string                `json:"coinbase"`
	Alloc              map[string]AllocEntry `json:"alloc"`
	SecondsPerSlot     int                   `json:"SECONDS_PER_SLOT"`
	Eth1FollowDistance int                   `json:"ETH1_FOLLOW_DISTANCE"`
	ParentHash         string                `json:"parentHash"`
	Number             string                `json:"number"`
	GasUsed            string                `json:"gasUsed"`
}

type Config struct {
	ChainId                       int  `json:"chainId"`
	HomesteadBlock                int  `json:"homesteadBlock"`
	EIP150Block                   int  `json:"eip150Block"`
	EIP155Block                   int  `json:"eip155Block"`
	EIP158Block                   int  `json:"eip158Block"`
	ByzantiumBlock                int  `json:"byzantiumBlock"`
	ConstantinopleBlock           int  `json:"constantinopleBlock"`
	PetersburgBlock               int  `json:"petersburgBlock"`
	IstanbulBlock                 int  `json:"istanbulBlock"`
	BerlinBlock                   int  `json:"berlinBlock"`
	LondonBlock                   int  `json:"londonBlock"`
	MergeNetsplitBlock            int  `json:"mergeNetsplitBlock"`
	TerminalTotalDifficultyPassed bool `json:"terminalTotalDifficultyPassed"`
	ShanghaiTime                  int  `json:"shanghaiTime"`
	TerminalTotalDifficulty       int  `json:"terminalTotalDifficulty"`
	// ETHash                        EthashConfig `json:"ethash,omitempty"`
	// CancunTime int `json:"cancunTime"`
}

type YAML struct {
	PresetBase                       string `json:"PRESET_BASE"`
	ConfigName                       string `json:"CONFIG_NAME"`
	MinGenesisActiveValidatorCount   int    `json:"MIN_GENESIS_ACTIVE_VALIDATOR_COUNT"`
	MinGenesisTime                   int64  `json:"MIN_GENESIS_TIME"`
	GenesisForkVersion               string `json:"GENESIS_FORK_VERSION"`
	GenesisDelay                     int    `json:"GENESIS_DELAY"`
	AltairForkVersion                string `json:"ALTAIR_FORK_VERSION"`
	AltairForkEpoch                  int    `json:"ALTAIR_FORK_EPOCH"`
	BellatrixForkVersion             string `json:"BELLATRIX_FORK_VERSION"`
	BellatrixForkEpoch               int    `json:"BELLATRIX_FORK_EPOCH"`
	TerminalTotalDifficulty          int    `json:"TERMINAL_TOTAL_DIFFICULTY"`
	TerminalBlockHash                string `json:"TERMINAL_BLOCK_HASH"`
	TerminalBlockHashActivationEpoch uint64 `json:"TERMINAL_BLOCK_HASH_ACTIVATION_EPOCH"`
	CapellaForkVersion               string `json:"CAPELLA_FORK_VERSION"`
	CapellaForkEpoch                 int    `json:"CAPELLA_FORK_EPOCH"`
	DenebForkVersion                 string `json:"DENEB_FORK_VERSION"`
	DenebForkEpoch                   int    `json:"DENEB_FORK_EPOCH"`
	SecondsPerSlot                   int    `json:"SECONDS_PER_SLOT"`
	SlotPerSecond                    int    `json:"SLOT_PER_SECOND"`
	SecondsPerEth1Block              int    `json:"SECONDS_PER_ETH1_BLOCK"`
	MinValidatorWithdrawabilityDelay int    `json:"MIN_VALIDATOR_WITHDRAWABILITY_DELAY"`
	ShardCommitteePeriod             int    `json:"SHARD_COMMITTEE_PERIOD"`
	Eth1FollowDistance               int    `json:"ETH1_FOLLOW_DISTANCE"`
	InactivityScoreBias              int    `json:"INACTIVITY_SCORE_BIAS"`
	InactivityScoreRecoveryRate      int    `json:"INACTIVITY_SCORE_RECOVERY_RATE"`
	EjectionBalance                  int    `json:"EJECTION_BALANCE"`
	MinPerEpochChurnLimit            int    `json:"MIN_PER_EPOCH_CHURN_LIMIT"`
	ChurnLimitQuotient               int    `json:"CHURN_LIMIT_QUOTIENT"`
	MaxPerEpochActivationChurnLimit  int    `json:"MAX_PER_EPOCH_ACTIVATION_CHURN_LIMIT"`
	ProposerScoreBoost               int    `json:"PROPOSER_SCORE_BOOST"`
	DepositChainID                   int    `json:"DEPOSIT_CHAIN_ID"`
	DepositNetworkID                 int    `json:"DEPOSIT_NETWORK_ID"`
	DepositContractAddress           string `json:"DEPOSIT_CONTRACT_ADDRESS"`
	GossipMaxSize                    int    `json:"GOSSIP_MAX_SIZE"`
	MaxRequestBlocks                 int    `json:"MAX_REQUEST_BLOCKS"`
	EpochsPerSubnetSubscription      int    `json:"EPOCHS_PER_SUBNET_SUBSCRIPTION"`
	MinEpochsForBlockRequests        int    `json:"MIN_EPOCHS_FOR_BLOCK_REQUESTS"`
	MaxChunkSize                     int    `json:"MAX_CHUNK_SIZE"`
	TTFBTimeout                      int    `json:"TTFB_TIMEOUT"`
	RespTimeout                      int    `json:"RESP_TIMEOUT"`
	AttestationPropagationSlotRange  int    `json:"ATTESTATION_PROPAGATION_SLOT_RANGE"`
	MaximumGossipClockDisparity      int    `json:"MAXIMUM_GOSSIP_CLOCK_DISPARITY"`
	MessageDomainInvalidSnappy       string `json:"MESSAGE_DOMAIN_INVALID_SNAPPY"`
	MessageDomainValidSnappy         string `json:"MESSAGE_DOMAIN_VALID_SNAPPY"`
	SubnetsPerNode                   int    `json:"SUBNETS_PER_NODE"`
	AttestationSubnetCount           int    `json:"ATTESTATION_SUBNET_COUNT"`
	AttestationSubnetExtraBits       int    `json:"ATTESTATION_SUBNET_EXTRA_BITS"`
	AttestationSubnetPrefixBits      int    `json:"ATTESTATION_SUBNET_PREFIX_BITS"`
	MaxRequestBlocksDeneb            int    `json:"MAX_REQUEST_BLOCKS_DENEB"`
	MaxRequestBlobSidecars           int    `json:"MAX_REQUEST_BLOB_SIDECARS"`
	MinEpochsForBlobSidecarsRequests int    `json:"MIN_EPOCHS_FOR_BLOB_SIDECARS_REQUESTS"`
	BlobSidecarSubnetCount           int    `json:"BLOB_SIDECAR_SUBNET_COUNT"`
	Eth1VotingPeriod                 int    `json:"ETH1_VOTING_PERIOD"`
	EpochPerEth1VotingPeriod         int    `json:"EPOCHS_PER_ETH1_VOTING_PERIOD"`
}

type Host struct {
	Name          string `json:"name"`
	IP            string `json:"ip"`
	Port          string `json:"port"`
	Type          string `json:"type"`
	AnsibleUser   string `json:"ansible_user"`
	AnsibleSSHKey string `json:"ansible_ssh_private_key_file"`
	Password      string `json:"password"`
}
