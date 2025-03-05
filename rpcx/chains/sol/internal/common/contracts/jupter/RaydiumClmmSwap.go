// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RaydiumClmmSwap is the `raydiumClmmSwap` instruction.
type RaydiumClmmSwap struct {

	// [0] = [] swapProgram
	//
	// [1] = [] payer
	//
	// [2] = [] ammConfig
	//
	// [3] = [WRITE] poolState
	//
	// [4] = [WRITE] inputTokenAccount
	//
	// [5] = [WRITE] outputTokenAccount
	//
	// [6] = [WRITE] inputVault
	//
	// [7] = [WRITE] outputVault
	//
	// [8] = [WRITE] observationState
	//
	// [9] = [] tokenProgram
	//
	// [10] = [WRITE] tickArray
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRaydiumClmmSwapInstructionBuilder creates a new `RaydiumClmmSwap` instruction builder.
func NewRaydiumClmmSwapInstructionBuilder() *RaydiumClmmSwap {
	nd := &RaydiumClmmSwap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetSwapProgramAccount sets the "swapProgram" account.
func (inst *RaydiumClmmSwap) SetSwapProgramAccount(swapProgram ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(swapProgram)
	return inst
}

// GetSwapProgramAccount gets the "swapProgram" account.
func (inst *RaydiumClmmSwap) GetSwapProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPayerAccount sets the "payer" account.
func (inst *RaydiumClmmSwap) SetPayerAccount(payer ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(payer)
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *RaydiumClmmSwap) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAmmConfigAccount sets the "ammConfig" account.
func (inst *RaydiumClmmSwap) SetAmmConfigAccount(ammConfig ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ammConfig)
	return inst
}

// GetAmmConfigAccount gets the "ammConfig" account.
func (inst *RaydiumClmmSwap) GetAmmConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolStateAccount sets the "poolState" account.
func (inst *RaydiumClmmSwap) SetPoolStateAccount(poolState ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "poolState" account.
func (inst *RaydiumClmmSwap) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetInputTokenAccountAccount sets the "inputTokenAccount" account.
func (inst *RaydiumClmmSwap) SetInputTokenAccountAccount(inputTokenAccount ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(inputTokenAccount).WRITE()
	return inst
}

// GetInputTokenAccountAccount gets the "inputTokenAccount" account.
func (inst *RaydiumClmmSwap) GetInputTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetOutputTokenAccountAccount sets the "outputTokenAccount" account.
func (inst *RaydiumClmmSwap) SetOutputTokenAccountAccount(outputTokenAccount ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(outputTokenAccount).WRITE()
	return inst
}

// GetOutputTokenAccountAccount gets the "outputTokenAccount" account.
func (inst *RaydiumClmmSwap) GetOutputTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetInputVaultAccount sets the "inputVault" account.
func (inst *RaydiumClmmSwap) SetInputVaultAccount(inputVault ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(inputVault).WRITE()
	return inst
}

// GetInputVaultAccount gets the "inputVault" account.
func (inst *RaydiumClmmSwap) GetInputVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetOutputVaultAccount sets the "outputVault" account.
func (inst *RaydiumClmmSwap) SetOutputVaultAccount(outputVault ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(outputVault).WRITE()
	return inst
}

// GetOutputVaultAccount gets the "outputVault" account.
func (inst *RaydiumClmmSwap) GetOutputVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetObservationStateAccount sets the "observationState" account.
func (inst *RaydiumClmmSwap) SetObservationStateAccount(observationState ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(observationState).WRITE()
	return inst
}

// GetObservationStateAccount gets the "observationState" account.
func (inst *RaydiumClmmSwap) GetObservationStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *RaydiumClmmSwap) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *RaydiumClmmSwap) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTickArrayAccount sets the "tickArray" account.
func (inst *RaydiumClmmSwap) SetTickArrayAccount(tickArray ag_solanago.PublicKey) *RaydiumClmmSwap {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tickArray).WRITE()
	return inst
}

// GetTickArrayAccount gets the "tickArray" account.
func (inst *RaydiumClmmSwap) GetTickArrayAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst RaydiumClmmSwap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RaydiumClmmSwap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RaydiumClmmSwap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RaydiumClmmSwap) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.SwapProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AmmConfig is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolState is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.InputTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.OutputTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.InputVault is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.OutputVault is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.ObservationState is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TickArray is not set")
		}
	}
	return nil
}

func (inst *RaydiumClmmSwap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RaydiumClmmSwap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     swapProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           payer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       ammConfig", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       poolState", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      inputToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     outputToken", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      inputVault", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("     outputVault", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("observationState", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("    tokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       tickArray", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj RaydiumClmmSwap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RaydiumClmmSwap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRaydiumClmmSwapInstruction declares a new RaydiumClmmSwap instruction with the provided parameters and accounts.
func NewRaydiumClmmSwapInstruction(
	// Accounts:
	swapProgram ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	ammConfig ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	inputTokenAccount ag_solanago.PublicKey,
	outputTokenAccount ag_solanago.PublicKey,
	inputVault ag_solanago.PublicKey,
	outputVault ag_solanago.PublicKey,
	observationState ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tickArray ag_solanago.PublicKey) *RaydiumClmmSwap {
	return NewRaydiumClmmSwapInstructionBuilder().
		SetSwapProgramAccount(swapProgram).
		SetPayerAccount(payer).
		SetAmmConfigAccount(ammConfig).
		SetPoolStateAccount(poolState).
		SetInputTokenAccountAccount(inputTokenAccount).
		SetOutputTokenAccountAccount(outputTokenAccount).
		SetInputVaultAccount(inputVault).
		SetOutputVaultAccount(outputVault).
		SetObservationStateAccount(observationState).
		SetTokenProgramAccount(tokenProgram).
		SetTickArrayAccount(tickArray)
}
