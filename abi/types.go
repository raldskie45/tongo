package abi

// Code autogenerated. DO NOT EDIT.

import (
	"encoding/json"
	"fmt"
	"github.com/tonkeeper/tongo/tlb"
)

type DedustAsset struct {
	tlb.SumType
	Native struct{} `tlbSumType:"$0000"`
	Jetton struct {
		WorkchainId int8
		Address     tlb.Bits256
	} `tlbSumType:"$0001"`
	ExtraCurrency struct {
		CurrencyId int32
	} `tlbSumType:"$0010"`
}

func (t *DedustAsset) MarshalJSON() ([]byte, error) {
	switch t.SumType {
	case "Native":
		bytes, err := json.Marshal(t.Native)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "Native","Native":%v}`, string(bytes))), nil
	case "Jetton":
		bytes, err := json.Marshal(t.Jetton)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "Jetton","Jetton":%v}`, string(bytes))), nil
	case "ExtraCurrency":
		bytes, err := json.Marshal(t.ExtraCurrency)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "ExtraCurrency","ExtraCurrency":%v}`, string(bytes))), nil
	default:
		return nil, fmt.Errorf("unknown sum type %v", t.SumType)
	}
}

type DedustSwapParams struct {
	Deadline       uint32
	RecipientAddr  tlb.MsgAddress
	ReferralAddr   tlb.MsgAddress
	FulfillPayload *tlb.Any `tlb:"maybe^"`
	RejectPayload  *tlb.Any `tlb:"maybe^"`
}

type DedustSwapStep struct {
	PoolAddr tlb.MsgAddress
	Params   DedustSwapStepParams
}

type DedustSwapStepParams struct {
	KindOut bool
	Limit   tlb.VarUInteger16
	Next    *DedustSwapStep `tlb:"maybe^"`
}

type JettonForceAction struct {
	tlb.SumType
	SetStatus struct {
		QueryId uint64
		Status  tlb.Uint4
	} `tlbSumType:"#eed236d3"`
	Burn struct {
		QueryId             uint64
		Amount              tlb.VarUInteger16
		ResponseDestination tlb.MsgAddress
		CustomPayload       *JettonPayload `tlb:"maybe^"`
	} `tlbSumType:"#595f07bc"`
	Transfer struct {
		QueryId             uint64
		Amount              tlb.VarUInteger16
		Destination         tlb.MsgAddress
		ResponseDestination tlb.MsgAddress
		CustomPayload       *tlb.Any `tlb:"maybe^"`
		ForwardTonAmount    tlb.VarUInteger16
		ForwardPayload      tlb.EitherRef[JettonPayload]
	} `tlbSumType:"#0f8a7ea5"`
}

func (t *JettonForceAction) MarshalJSON() ([]byte, error) {
	switch t.SumType {
	case "SetStatus":
		bytes, err := json.Marshal(t.SetStatus)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "SetStatus","SetStatus":%v}`, string(bytes))), nil
	case "Burn":
		bytes, err := json.Marshal(t.Burn)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "Burn","Burn":%v}`, string(bytes))), nil
	case "Transfer":
		bytes, err := json.Marshal(t.Transfer)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "Transfer","Transfer":%v}`, string(bytes))), nil
	default:
		return nil, fmt.Errorf("unknown sum type %v", t.SumType)
	}
}

type TonstakersControllerData struct {
	ControllerId uint32
	Validator    tlb.MsgAddress
	Pool         tlb.MsgAddress
	Governor     tlb.MsgAddress
	Field4       struct {
		Approver tlb.MsgAddress
		Halter   tlb.MsgAddress
	} `tlb:"^"`
}

type MultisigOrder struct {
	Field0 tlb.Hashmap[tlb.Uint8, tlb.Ref[MultisigSendMessageAction]]
}

type MultisigProposersList struct {
	Proposers tlb.HashmapE[tlb.Uint8, tlb.MsgAddress]
}

type MultisigSendMessageAction struct {
	tlb.SumType
	SendMessage struct {
		Field0 SendMessageAction
	} `tlbSumType:"#f1381e5b"`
	UpdateMultisigParam struct {
		Threshold uint8
		Signers   tlb.Hashmap[tlb.Uint8, tlb.MsgAddress] `tlb:"^"`
		Proposers tlb.HashmapE[tlb.Uint8, tlb.MsgAddress]
	} `tlbSumType:"#1d0cfbd3"`
}

func (t *MultisigSendMessageAction) MarshalJSON() ([]byte, error) {
	switch t.SumType {
	case "SendMessage":
		bytes, err := json.Marshal(t.SendMessage)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "SendMessage","SendMessage":%v}`, string(bytes))), nil
	case "UpdateMultisigParam":
		bytes, err := json.Marshal(t.UpdateMultisigParam)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "UpdateMultisigParam","UpdateMultisigParam":%v}`, string(bytes))), nil
	default:
		return nil, fmt.Errorf("unknown sum type %v", t.SumType)
	}
}

type MultisigSignersList struct {
	Signers tlb.Hashmap[tlb.Uint8, tlb.MsgAddress]
}

type ClosingConfig struct {
	QuarantinDuration        uint32
	MisbehaviorFine          tlb.Grams
	ConditionalCloseDuration uint32
}

type ConditionalPayment struct {
	Amount    tlb.Grams
	Condition tlb.Any
}

type PaymentConfig struct {
	ExcessFee tlb.Grams
	DestA     tlb.MsgAddress
	DestB     tlb.MsgAddress
}

type QuarantinedState struct {
	StateA           SemiChannelBody
	StateB           SemiChannelBody
	QuarantineStarts uint32
	StateCommitedByA bool
}

type SemiChannel struct {
	Magic            tlb.Magic `tlb:"#43685374"`
	ChannelId        tlb.Uint128
	Data             SemiChannelBody
	CounterpartyData *SemiChannelBody `tlb:"maybe^"`
}

type SemiChannelBody struct {
	Seqno        uint64
	Sent         tlb.Grams
	Conditionals tlb.HashmapE[tlb.Uint32, ConditionalPayment]
}

type SignedSemiChannel struct {
	Signature tlb.Bits512
	State     SemiChannel
}

type Storage struct {
	BalanceA       tlb.Grams
	BalanceB       tlb.Grams
	KeyA           tlb.Uint256
	KeyB           tlb.Uint256
	ChannelId      tlb.Uint128
	Config         ClosingConfig `tlb:"^"`
	CommitedSeqnoA uint32
	CommitedSeqnoB uint32
	Quarantin      *QuarantinedState `tlb:"maybe^"`
	Payments       PaymentConfig     `tlb:"^"`
}

type StonfiPayToParams struct {
	Amount0Out    tlb.VarUInteger16
	Token0Address tlb.MsgAddress
	Amount1Out    tlb.VarUInteger16
	Token1Address tlb.MsgAddress
}

type StonfiSwapAddrs struct {
	FromUser tlb.MsgAddress
}

type TorrentInfo struct {
	PieceSize      uint32
	FileSize       uint64
	RootHash       tlb.Uint256
	HeaderSize     uint64
	HeaderHash     tlb.Uint256
	MicrochunkHash *tlb.Uint256 `tlb:"maybe"`
	Description    tlb.Text
}

type AmmSettings struct {
	Fee                           uint32
	RolloverFee                   uint32
	FundingPeriod                 uint32
	InitMarginRatio               uint32
	MaintenanceMarginRatio        uint32
	LiquidationFeeRatio           uint32
	PartialLiquidationRatio       uint32
	SpreadLimit                   uint32
	MaxPriceImpact                uint32
	MaxPriceSpread                uint32
	MaxOpenNotional               uint32
	FeeToStakersPercent           uint32
	FundingMode                   tlb.Uint2
	MinPartialLiquidationNotional tlb.Grams
	MinLeverage                   uint32
}

type OracleData struct {
	UpdateMsg  UpdateMsg `tlb:"^"`
	Signatures tlb.Any   `tlb:"^"`
}

type Parameters struct {
	Discount uint32
	Rebate   uint32
}

type UpdateMsg struct {
	Price      tlb.Grams
	Spread     tlb.Grams
	Timestamp  uint32
	AssetIndex uint16
}

type NftRoyaltyParams struct {
	Numerator   uint16
	Denominator uint16
	Destination tlb.MsgAddress
}

type TeleitemAuctionConfig struct {
	BeneficiarAddress tlb.MsgAddress
	InitialMinBid     tlb.Grams
	MaxBid            tlb.Grams
	MinBidStep        uint8
	MinExtendTime     uint32
	Duration          uint32
}

type TelemintData struct {
	Touched           bool
	SubwalletId       uint32
	PublicKey         tlb.Bits256
	CollectionContent tlb.Any          `tlb:"^"`
	NftItemCode       tlb.Any          `tlb:"^"`
	RoyaltyParams     NftRoyaltyParams `tlb:"^"`
}

type TelemintRestrictions struct {
	ForceSenderAddress   *tlb.MsgAddress `tlb:"maybe"`
	RewriteSenderAddress *tlb.MsgAddress `tlb:"maybe"`
}

type TelemintTokenInfo struct {
	Name   tlb.FixedLengthText
	Domain tlb.FixedLengthText
}

type TelemintUnsignedDeploy struct {
	SubwalletId   uint32
	ValidSince    uint32
	ValidTill     uint32
	Username      tlb.FixedLengthText
	Content       tlb.Any               `tlb:"^"`
	AuctionConfig TeleitemAuctionConfig `tlb:"^"`
	RoyaltyParams *NftRoyaltyParams     `tlb:"maybe^"`
}

type TelemintUnsignedDeployV2 struct {
	SubwalletId   uint32
	ValidSince    uint32
	ValidTill     uint32
	TokenName     tlb.FixedLengthText
	Content       tlb.Any               `tlb:"^"`
	AuctionConfig TeleitemAuctionConfig `tlb:"^"`
	RoyaltyParams *NftRoyaltyParams     `tlb:"maybe^"`
	Restrictions  *TelemintRestrictions `tlb:"maybe^"`
}

type AccountLists struct {
	List tlb.Hashmap[tlb.Bits256, tlb.Any]
}

type CommonMsgInfoRelaxed struct {
	Magic       tlb.Magic `tlb:"$0"`
	IhrDisabled bool
	Bounce      bool
	Bounced     bool
	Src         tlb.MsgAddress
	Dest        tlb.MsgAddress
	Value       tlb.CurrencyCollection
	IhrFee      tlb.Grams
	FwdFee      tlb.Grams
	CreatedLt   uint64
	CreatedAt   uint32
}

type HighloadV3MsgInner struct {
	SubwalletId   uint32
	MessageToSend MessageRelaxed `tlb:"^"`
	SendMode      uint8
	QueryId       HighloadV3QueryId
	CreatedAt     uint64
	Timeout       tlb.Uint22
}

type HighloadV3QueryId struct {
	Shift     tlb.Uint13
	BitNumber tlb.Uint10
}

type HighloadWalletV3MessageRelaxed struct {
	Info CommonMsgInfoRelaxed
	Init *tlb.EitherRef[tlb.StateInit] `tlb:"maybe"`
	Body tlb.EitherRef[MessageRelaxed]
}

type MessageRelaxed struct {
	tlb.SumType
	MessageInternal struct {
		IhrDisabled bool
		Bounce      bool
		Bounced     bool
		Src         tlb.MsgAddress
		Dest        tlb.MsgAddress
		Value       tlb.CurrencyCollection
		IhrFee      tlb.Grams
		FwdFee      tlb.Grams
		CreatedLt   uint64
		CreatedAt   uint32
		Init        *tlb.EitherRef[tlb.StateInit] `tlb:"maybe"`
		Body        tlb.EitherRef[InMsgBody]
	} `tlbSumType:"$0"`
	MessageExtOut struct {
		Src       tlb.MsgAddress
		Dest      tlb.MsgAddress
		CreatedLt uint64
		CreatedAt uint32
		Init      *tlb.EitherRef[tlb.StateInit] `tlb:"maybe"`
		Body      tlb.EitherRef[ExtOutMsgBody]
	} `tlbSumType:"$11"`
}

func (t *MessageRelaxed) MarshalJSON() ([]byte, error) {
	switch t.SumType {
	case "MessageInternal":
		bytes, err := json.Marshal(t.MessageInternal)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "MessageInternal","MessageInternal":%v}`, string(bytes))), nil
	case "MessageExtOut":
		bytes, err := json.Marshal(t.MessageExtOut)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`{"SumType": "MessageExtOut","MessageExtOut":%v}`, string(bytes))), nil
	default:
		return nil, fmt.Errorf("unknown sum type %v", t.SumType)
	}
}

type SendMessageAction struct {
	Mode    uint8
	Message MessageRelaxed `tlb:"^"`
}

type WhalesNominatorsMember struct {
	ProfitPerCoin      tlb.Int128
	Balance            tlb.Grams
	PendingWithdraw    tlb.Grams
	PendingWithdrawAll bool
	PendingDeposit     tlb.Grams
	MemberWithdraw     tlb.Grams
}

type WhalesNominatorsMembersList struct {
	List tlb.Hashmap[tlb.Bits256, WhalesNominatorsMember]
}
