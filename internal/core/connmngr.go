package core

import (
	"github.com/jackc/pgproto3"
	"github.com/pg-sharding/spqr/internal/conn"
	"github.com/pg-sharding/spqr/internal/r"
	"github.com/pkg/errors"
)

type RelayState struct {
	TxActive bool

	ActiveShardConn *ShServer
	ActiveShardIndx int
}

type ConnManager interface {
	TXBeginCB(cl *ShClient, rst *RelayState) error
	TXEndCB(cl *ShClient, rst *RelayState) error

	RouteCB(cl *ShClient, rst *RelayState) error
	ValidateReRoute(rst *RelayState) bool
}

type TxConnManager struct {
}

func NewTxConnManager() *TxConnManager {
	return &TxConnManager{}
}

func (t *TxConnManager) RouteCB(cl *ShClient, rst *RelayState) error {

	shConn, err := cl.Route().GetConn("tcp6", rst.ActiveShardIndx)

	if err != nil {
		return err
	}

	cl.AssignShrdConn(shConn)

	return nil
}

func (t *TxConnManager) ValidateReRoute(rst *RelayState) bool {
	return rst.ActiveShardIndx == r.NOSHARD || !rst.TxActive
}

func (t *TxConnManager) TXBeginCB(cl *ShClient, rst *RelayState) error {
	return nil
}

func (t *TxConnManager) TXEndCB(cl *ShClient, rst *RelayState) error {

	//tracelog.InfoLogger.Println("releasing tx")

	cl.Route().Unroute(rst.ActiveShardIndx, cl)
	rst.ActiveShardIndx = r.NOSHARD

	return nil
}

var _ ConnManager = &TxConnManager{}

type SessConnManager struct {
}

func (s SessConnManager) TXBeginCB(cl *ShClient, rst *RelayState) error {
	return nil
}

func (s SessConnManager) TXEndCB(cl *ShClient, rst *RelayState) error {
	return nil
}

func (s SessConnManager) RouteCB(cl *ShClient, rst *RelayState) error {

	shConn, err := cl.Route().GetConn("tcp6", rst.ActiveShardIndx)

	if err != nil {
		return err
	}

	cl.AssignShrdConn(shConn)
	rst.ActiveShardConn = shConn

	return nil
}

func (s SessConnManager) ValidateReRoute(rst *RelayState) bool {
	return rst.ActiveShardIndx == r.NOSHARD
}

var _ ConnManager = &SessConnManager{}

func NewSessConnManager() *SessConnManager {
	return &SessConnManager{}
}

func InitClConnection(client *ShClient) (ConnManager, error) {

	var cmngr ConnManager

	//tracelog.InfoLogger.Printf("pooling mode %v", client.Rule().PoolingMode)

	switch client.rule.PoolingMode {
	case conn.PoolingModeSession:
		cmngr = NewSessConnManager()
	case conn.PoolingModeTransaction:
		cmngr = NewTxConnManager()
	default:
		for _, msg := range []pgproto3.BackendMessage{
			&pgproto3.ErrorResponse{
				Message:  "unknown pooling mode for route",
				Severity: "ERROR",
			},
		} {
			if err := client.Send(msg); err != nil {
				return nil, err
			}
		}
		return nil, errors.Errorf("unknown pooling mode %v", client.rule.PoolingMode)
	}

	return cmngr, nil
}
