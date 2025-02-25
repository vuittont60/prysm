package apimiddleware

type ListKeystoresResponseJson struct {
	Keystores []*KeystoreJson `json:"data"`
}

type KeystoreJson struct {
	ValidatingPubkey string `json:"validating_pubkey" hex:"true"`
	DerivationPath   string `json:"derivation_path"`
}

type ImportKeystoresRequestJson struct {
	Keystores          []string `json:"keystores"`
	Passwords          []string `json:"passwords"`
	SlashingProtection string   `json:"slashing_protection"`
}

type ImportKeystoresResponseJson struct {
	Statuses []*StatusJson `json:"data"`
}

type DeleteKeystoresRequestJson struct {
	PublicKeys []string `json:"pubkeys" hex:"true"`
}

type StatusJson struct {
	Status  string `json:"status" enum:"true"`
	Message string `json:"message"`
}

type DeleteKeystoresResponseJson struct {
	Statuses           []*StatusJson `json:"data"`
	SlashingProtection string        `json:"slashing_protection"`
}

type FeeRecipientJson struct {
	Pubkey     string `json:"pubkey" hex:"true"`
	Ethaddress string `json:"ethaddress" address:"true"`
}

type GasLimitJson struct {
	Pubkey   string `json:"pubkey" hex:"true"`
	GasLimit string `json:"gas_limit"`
}

type GetFeeRecipientByPubkeyResponseJson struct {
	Data *FeeRecipientJson `json:"data"`
}

type SetFeeRecipientByPubkeyRequestJson struct {
	Ethaddress string `json:"ethaddress" hex:"true"`
}

type DeleteFeeRecipientByPubkeyRequestJson struct {
	Pubkey string `json:"pubkey" hex:"true"`
}

type GetGasLimitResponseJson struct {
	Data *GasLimitJson `json:"data"`
}

type SetGasLimitRequestJson struct {
	GasLimit string `json:"gas_limit"`
}

type DeleteGasLimitRequestJson struct {
	Pubkey string `json:"pubkey" hex:"true"`
}
