package kr

import (
	proto "github.com/pg-sharding/spqr/pkg/protos"
	"github.com/pg-sharding/spqr/qdb"
	spqrparser "github.com/pg-sharding/spqr/yacc/console"
)

type KeyRangeBound []byte

type ShardKey struct {
	Name string
	RW   bool
}

type KeyRange struct {
	LowerBound []byte
	UpperBound []byte
	ShardID    string
	ID         string
	Dataspace  string
}

// TODO : unit tests
func CmpRangesLess(kr []byte, other []byte) bool {
	if len(kr) == len(other) {
		return string(kr) < string(other)
	}

	return len(kr) < len(other)
}

// TODO : unit tests
func CmpRangesLessEqual(kr []byte, other []byte) bool {
	if len(kr) == len(other) {
		return string(kr) <= string(other)
	}

	return len(kr) < len(other)
}

// TODO : unit tests
func CmpRangesEqual(kr []byte, other []byte) bool {
	if len(kr) == len(other) {
		return string(kr) == string(other)
	}

	return false
}

// TODO : unit tests
func KeyRangeFromDB(kr *qdb.KeyRange) *KeyRange {
	return &KeyRange{
		LowerBound: kr.LowerBound,
		UpperBound: kr.UpperBound,
		ShardID:    kr.ShardID,
		ID:         kr.KeyRangeID,
		Dataspace:  kr.DataspaceId,
	}
}

// TODO : unit tests
func KeyRangeFromSQL(kr *spqrparser.KeyRangeDefinition) *KeyRange {
	if kr == nil {
		return nil
	}
	return &KeyRange{
		LowerBound: kr.LowerBound,
		UpperBound: kr.UpperBound,
		ShardID:    kr.ShardID,
		ID:         kr.KeyRangeID,
		Dataspace:  kr.Dataspace,
	}
}

// TODO : unit tests
func KeyRangeFromProto(kr *proto.KeyRangeInfo) *KeyRange {
	if kr == nil {
		return nil
	}
	return &KeyRange{
		LowerBound: []byte(kr.KeyRange.LowerBound),
		UpperBound: []byte(kr.KeyRange.UpperBound),
		ShardID:    kr.ShardId,
		ID:         kr.Krid,
		Dataspace:  kr.DataspaceId,
	}
}

// TODO : unit tests
func (kr *KeyRange) ToDB() *qdb.KeyRange {
	return &qdb.KeyRange{
		LowerBound:  kr.LowerBound,
		UpperBound:  kr.UpperBound,
		ShardID:     kr.ShardID,
		KeyRangeID:  kr.ID,
		DataspaceId: kr.Dataspace,
	}
}

// TODO : unit tests
func (kr *KeyRange) ToProto() *proto.KeyRangeInfo {
	return &proto.KeyRangeInfo{
		KeyRange: &proto.KeyRange{
			LowerBound: string(kr.LowerBound),
			UpperBound: string(kr.UpperBound),
		},
		ShardId:     kr.ShardID,
		Krid:        kr.ID,
		DataspaceId: kr.Dataspace,
	}
}
