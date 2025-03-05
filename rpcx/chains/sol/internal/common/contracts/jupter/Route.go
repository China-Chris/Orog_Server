// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// route_plan Topologically sorted trade DAG
type Route struct {
	RoutePlan       *[]RoutePlanStep
	InAmount        *uint64
	QuotedOutAmount *uint64
	SlippageBps     *uint16
	PlatformFeeBps  *uint8

	// [0] = [] tokenProgram
	//
	// [1] = [SIGNER] userTransferAuthority
	//
	// [2] = [] userSourceTokenAccount
	//
	// [3] = [] userDestinationTokenAccount
	//
	// [4] = [] destinationTokenAccount
	//
	// [5] = [] destinationMint
	//
	// [6] = [WRITE] platformFeeAccount
	//
	// [7] = [] eventAuthority
	//
	// [8] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRouteInstructionBuilder creates a new `Route` instruction builder.
func NewRouteInstructionBuilder() *Route {
	nd := &Route{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetRoutePlan sets the "routePlan" parameter.
func (inst *Route) SetRoutePlan(routePlan []RoutePlanStep) *Route {
	inst.RoutePlan = &routePlan
	return inst
}

// SetInAmount sets the "inAmount" parameter.
func (inst *Route) SetInAmount(inAmount uint64) *Route {
	inst.InAmount = &inAmount
	return inst
}

// SetQuotedOutAmount sets the "quotedOutAmount" parameter.
func (inst *Route) SetQuotedOutAmount(quotedOutAmount uint64) *Route {
	inst.QuotedOutAmount = &quotedOutAmount
	return inst
}

// SetSlippageBps sets the "slippageBps" parameter.
func (inst *Route) SetSlippageBps(slippageBps uint16) *Route {
	inst.SlippageBps = &slippageBps
	return inst
}

// SetPlatformFeeBps sets the "platformFeeBps" parameter.
func (inst *Route) SetPlatformFeeBps(platformFeeBps uint8) *Route {
	inst.PlatformFeeBps = &platformFeeBps
	return inst
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Route) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Route) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserTransferAuthorityAccount sets the "userTransferAuthority" account.
func (inst *Route) SetUserTransferAuthorityAccount(userTransferAuthority ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(userTransferAuthority).SIGNER()
	return inst
}

// GetUserTransferAuthorityAccount gets the "userTransferAuthority" account.
func (inst *Route) GetUserTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserSourceTokenAccountAccount sets the "userSourceTokenAccount" account.
func (inst *Route) SetUserSourceTokenAccountAccount(userSourceTokenAccount ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userSourceTokenAccount)
	return inst
}

// GetUserSourceTokenAccountAccount gets the "userSourceTokenAccount" account.
func (inst *Route) GetUserSourceTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserDestinationTokenAccountAccount sets the "userDestinationTokenAccount" account.
func (inst *Route) SetUserDestinationTokenAccountAccount(userDestinationTokenAccount ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userDestinationTokenAccount)
	return inst
}

// GetUserDestinationTokenAccountAccount gets the "userDestinationTokenAccount" account.
func (inst *Route) GetUserDestinationTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetDestinationTokenAccountAccount sets the "destinationTokenAccount" account.
func (inst *Route) SetDestinationTokenAccountAccount(destinationTokenAccount ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(destinationTokenAccount)
	return inst
}

// GetDestinationTokenAccountAccount gets the "destinationTokenAccount" account.
func (inst *Route) GetDestinationTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetDestinationMintAccount sets the "destinationMint" account.
func (inst *Route) SetDestinationMintAccount(destinationMint ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(destinationMint)
	return inst
}

// GetDestinationMintAccount gets the "destinationMint" account.
func (inst *Route) GetDestinationMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPlatformFeeAccountAccount sets the "platformFeeAccount" account.
func (inst *Route) SetPlatformFeeAccountAccount(platformFeeAccount ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(platformFeeAccount).WRITE()
	return inst
}

// GetPlatformFeeAccountAccount gets the "platformFeeAccount" account.
func (inst *Route) GetPlatformFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *Route) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *Route) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetProgramAccount sets the "program" account.
func (inst *Route) SetProgramAccount(program ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Route) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst Route) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Route,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Route) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Route) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RoutePlan == nil {
			return errors.New("RoutePlan parameter is not set")
		}
		if inst.InAmount == nil {
			return errors.New("InAmount parameter is not set")
		}
		if inst.QuotedOutAmount == nil {
			return errors.New("QuotedOutAmount parameter is not set")
		}
		if inst.SlippageBps == nil {
			return errors.New("SlippageBps parameter is not set")
		}
		if inst.PlatformFeeBps == nil {
			return errors.New("PlatformFeeBps parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UserTransferAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UserSourceTokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserDestinationTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.DestinationTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.DestinationMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PlatformFeeAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *Route) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Route")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("      RoutePlan", *inst.RoutePlan))
						paramsBranch.Child(ag_format.Param("       InAmount", *inst.InAmount))
						paramsBranch.Child(ag_format.Param("QuotedOutAmount", *inst.QuotedOutAmount))
						paramsBranch.Child(ag_format.Param("    SlippageBps", *inst.SlippageBps))
						paramsBranch.Child(ag_format.Param(" PlatformFeeBps", *inst.PlatformFeeBps))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("userTransferAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      userSourceToken", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" userDestinationToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     destinationToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      destinationMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          platformFee", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       eventAuthority", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              program", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj Route) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RoutePlan` param:
	err = encoder.Encode(obj.RoutePlan)
	if err != nil {
		return err
	}
	// Serialize `InAmount` param:
	err = encoder.Encode(obj.InAmount)
	if err != nil {
		return err
	}
	// Serialize `QuotedOutAmount` param:
	err = encoder.Encode(obj.QuotedOutAmount)
	if err != nil {
		return err
	}
	// Serialize `SlippageBps` param:
	err = encoder.Encode(obj.SlippageBps)
	if err != nil {
		return err
	}
	// Serialize `PlatformFeeBps` param:
	err = encoder.Encode(obj.PlatformFeeBps)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Route) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RoutePlan`:
	err = decoder.Decode(&obj.RoutePlan)
	if err != nil {
		return err
	}
	// Deserialize `InAmount`:
	err = decoder.Decode(&obj.InAmount)
	if err != nil {
		return err
	}
	// Deserialize `QuotedOutAmount`:
	err = decoder.Decode(&obj.QuotedOutAmount)
	if err != nil {
		return err
	}
	// Deserialize `SlippageBps`:
	err = decoder.Decode(&obj.SlippageBps)
	if err != nil {
		return err
	}
	// Deserialize `PlatformFeeBps`:
	err = decoder.Decode(&obj.PlatformFeeBps)
	if err != nil {
		return err
	}
	return nil
}

// NewRouteInstruction declares a new Route instruction with the provided parameters and accounts.
func NewRouteInstruction(
	// Parameters:
	routePlan []RoutePlanStep,
	inAmount uint64,
	quotedOutAmount uint64,
	slippageBps uint16,
	platformFeeBps uint8,
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	userTransferAuthority ag_solanago.PublicKey,
	userSourceTokenAccount ag_solanago.PublicKey,
	userDestinationTokenAccount ag_solanago.PublicKey,
	destinationTokenAccount ag_solanago.PublicKey,
	destinationMint ag_solanago.PublicKey,
	platformFeeAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *Route {
	return NewRouteInstructionBuilder().
		SetRoutePlan(routePlan).
		SetInAmount(inAmount).
		SetQuotedOutAmount(quotedOutAmount).
		SetSlippageBps(slippageBps).
		SetPlatformFeeBps(platformFeeBps).
		SetTokenProgramAccount(tokenProgram).
		SetUserTransferAuthorityAccount(userTransferAuthority).
		SetUserSourceTokenAccountAccount(userSourceTokenAccount).
		SetUserDestinationTokenAccountAccount(userDestinationTokenAccount).
		SetDestinationTokenAccountAccount(destinationTokenAccount).
		SetDestinationMintAccount(destinationMint).
		SetPlatformFeeAccountAccount(platformFeeAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
