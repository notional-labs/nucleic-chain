package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewStateInfo(
	rollappId string,
	newIndex uint64,
	creator string,
	startHeight,
	numBlocks uint64,
	daPath string,
	version uint64,
	height uint64,
	blockDescs BlockDescriptors,
) *StateInfo {
	stateInfoIndex := StateInfoIndex{RollappId: rollappId, Index: newIndex}
	status := Status_PENDING
	return &StateInfo{
		StateInfoIndex: stateInfoIndex,
		Sequencer:      creator,
		StartHeight:    startHeight,
		NumBlocks:      numBlocks,
		DaPath:         daPath,
		Version:        version,
		CreationHeight: height,
		Status:         status,
		BlockDescs:     blockDescs,
	}
}

func (s *StateInfo) Finalize() {
	s.Status = Status_FINALIZED
}

func (s *StateInfo) GetIndex() StateInfoIndex {
	return s.StateInfoIndex
}

func (s *StateInfo) GetLatestHeight() uint64 {
	return s.StartHeight + s.NumBlocks - 1
}

func (s *StateInfo) GetEvents() []sdk.Attribute {
	eventAttributes := []sdk.Attribute{
		sdk.NewAttribute(AttributeKeyRollappId, s.StateInfoIndex.RollappId),
		sdk.NewAttribute(AttributeKeyStateInfoIndex, strconv.FormatUint(s.StateInfoIndex.Index, 10)),
		sdk.NewAttribute(AttributeKeyStartHeight, strconv.FormatUint(s.StartHeight, 10)),
		sdk.NewAttribute(AttributeKeyNumBlocks, strconv.FormatUint(s.NumBlocks, 10)),
		sdk.NewAttribute(AttributeKeyDAPath, s.DaPath),
		sdk.NewAttribute(AttributeKeyStatus, s.Status.String()),
	}
	return eventAttributes
}
