// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PerpsAddLiquidity is the `perpsAddLiquidity` instruction.
type PerpsAddLiquidity struct {

	// [0] = [] swapProgram
	//
	// [1] = [WRITE] owner
	//
	// [2] = [WRITE] fundingOrReceivingAccount
	//
	// [3] = [WRITE] lpTokenAccount
	//
	// [4] = [] transferAuthority
	//
	// [5] = [] perpetuals
	//
	// [6] = [WRITE] pool
	//
	// [7] = [WRITE] custody
	//
	// [8] = [] custodyOracleAccount
	//
	// [9] = [WRITE] custodyTokenAccount
	//
	// [10] = [WRITE] lpTokenMint
	//
	// [11] = [] tokenProgram
	//
	// [12] = [] eventAuthority
	//
	// [13] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPerpsAddLiquidityInstructionBuilder creates a new `PerpsAddLiquidity` instruction builder.
func NewPerpsAddLiquidityInstructionBuilder() *PerpsAddLiquidity {
	nd := &PerpsAddLiquidity{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetSwapProgramAccount sets the "swapProgram" account.
func (inst *PerpsAddLiquidity) SetSwapProgramAccount(swapProgram ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(swapProgram)
	return inst
}

// GetSwapProgramAccount gets the "swapProgram" account.
func (inst *PerpsAddLiquidity) GetSwapProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *PerpsAddLiquidity) SetOwnerAccount(owner ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).WRITE()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *PerpsAddLiquidity) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundingOrReceivingAccountAccount sets the "fundingOrReceivingAccount" account.
func (inst *PerpsAddLiquidity) SetFundingOrReceivingAccountAccount(fundingOrReceivingAccount ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fundingOrReceivingAccount).WRITE()
	return inst
}

// GetFundingOrReceivingAccountAccount gets the "fundingOrReceivingAccount" account.
func (inst *PerpsAddLiquidity) GetFundingOrReceivingAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLpTokenAccountAccount sets the "lpTokenAccount" account.
func (inst *PerpsAddLiquidity) SetLpTokenAccountAccount(lpTokenAccount ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(lpTokenAccount).WRITE()
	return inst
}

// GetLpTokenAccountAccount gets the "lpTokenAccount" account.
func (inst *PerpsAddLiquidity) GetLpTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTransferAuthorityAccount sets the "transferAuthority" account.
func (inst *PerpsAddLiquidity) SetTransferAuthorityAccount(transferAuthority ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(transferAuthority)
	return inst
}

// GetTransferAuthorityAccount gets the "transferAuthority" account.
func (inst *PerpsAddLiquidity) GetTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPerpetualsAccount sets the "perpetuals" account.
func (inst *PerpsAddLiquidity) SetPerpetualsAccount(perpetuals ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(perpetuals)
	return inst
}

// GetPerpetualsAccount gets the "perpetuals" account.
func (inst *PerpsAddLiquidity) GetPerpetualsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPoolAccount sets the "pool" account.
func (inst *PerpsAddLiquidity) SetPoolAccount(pool ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *PerpsAddLiquidity) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCustodyAccount sets the "custody" account.
func (inst *PerpsAddLiquidity) SetCustodyAccount(custody ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(custody).WRITE()
	return inst
}

// GetCustodyAccount gets the "custody" account.
func (inst *PerpsAddLiquidity) GetCustodyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetCustodyOracleAccountAccount sets the "custodyOracleAccount" account.
func (inst *PerpsAddLiquidity) SetCustodyOracleAccountAccount(custodyOracleAccount ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(custodyOracleAccount)
	return inst
}

// GetCustodyOracleAccountAccount gets the "custodyOracleAccount" account.
func (inst *PerpsAddLiquidity) GetCustodyOracleAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetCustodyTokenAccountAccount sets the "custodyTokenAccount" account.
func (inst *PerpsAddLiquidity) SetCustodyTokenAccountAccount(custodyTokenAccount ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(custodyTokenAccount).WRITE()
	return inst
}

// GetCustodyTokenAccountAccount gets the "custodyTokenAccount" account.
func (inst *PerpsAddLiquidity) GetCustodyTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetLpTokenMintAccount sets the "lpTokenMint" account.
func (inst *PerpsAddLiquidity) SetLpTokenMintAccount(lpTokenMint ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(lpTokenMint).WRITE()
	return inst
}

// GetLpTokenMintAccount gets the "lpTokenMint" account.
func (inst *PerpsAddLiquidity) GetLpTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *PerpsAddLiquidity) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *PerpsAddLiquidity) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *PerpsAddLiquidity) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *PerpsAddLiquidity) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetProgramAccount sets the "program" account.
func (inst *PerpsAddLiquidity) SetProgramAccount(program ag_solanago.PublicKey) *PerpsAddLiquidity {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *PerpsAddLiquidity) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst PerpsAddLiquidity) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PerpsAddLiquidity,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PerpsAddLiquidity) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PerpsAddLiquidity) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.SwapProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FundingOrReceivingAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.LpTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TransferAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Perpetuals is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Custody is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.CustodyOracleAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.CustodyTokenAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.LpTokenMint is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *PerpsAddLiquidity) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PerpsAddLiquidity")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       swapProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("fundingOrReceiving", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           lpToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta(" transferAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        perpetuals", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("              pool", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           custody", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     custodyOracle", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      custodyToken", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       lpTokenMint", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("    eventAuthority", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj PerpsAddLiquidity) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *PerpsAddLiquidity) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewPerpsAddLiquidityInstruction declares a new PerpsAddLiquidity instruction with the provided parameters and accounts.
func NewPerpsAddLiquidityInstruction(
	// Accounts:
	swapProgram ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	fundingOrReceivingAccount ag_solanago.PublicKey,
	lpTokenAccount ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	perpetuals ag_solanago.PublicKey,
	pool ag_solanago.PublicKey,
	custody ag_solanago.PublicKey,
	custodyOracleAccount ag_solanago.PublicKey,
	custodyTokenAccount ag_solanago.PublicKey,
	lpTokenMint ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *PerpsAddLiquidity {
	return NewPerpsAddLiquidityInstructionBuilder().
		SetSwapProgramAccount(swapProgram).
		SetOwnerAccount(owner).
		SetFundingOrReceivingAccountAccount(fundingOrReceivingAccount).
		SetLpTokenAccountAccount(lpTokenAccount).
		SetTransferAuthorityAccount(transferAuthority).
		SetPerpetualsAccount(perpetuals).
		SetPoolAccount(pool).
		SetCustodyAccount(custody).
		SetCustodyOracleAccountAccount(custodyOracleAccount).
		SetCustodyTokenAccountAccount(custodyTokenAccount).
		SetLpTokenMintAccount(lpTokenMint).
		SetTokenProgramAccount(tokenProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
