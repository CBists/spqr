package local

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/pg-sharding/spqr/pkg/config"
	"github.com/pg-sharding/spqr/pkg/meta"
	"github.com/pg-sharding/spqr/pkg/models/datashards"
	"github.com/pg-sharding/spqr/pkg/models/dataspaces"
	"github.com/pg-sharding/spqr/pkg/models/kr"
	"github.com/pg-sharding/spqr/pkg/models/shrule"
	"github.com/pg-sharding/spqr/pkg/models/topology"
	"github.com/pg-sharding/spqr/pkg/spqrlog"
	"github.com/pg-sharding/spqr/qdb"
	"github.com/pg-sharding/spqr/qdb/ops"
	"github.com/pg-sharding/spqr/router/routingstate"
)

type LocalCoordinator struct {
	mu sync.Mutex

	Rules []*shrule.ShardingRule

	ColumnMapping map[string]struct{}
	LocalTables   map[string]struct{}

	// shards
	DataShardCfgs  map[string]*config.Shard
	WorldShardCfgs map[string]*config.Shard

	// not extended QDB, since the router does not need to track the installation topology
	qdb qdb.QDB
}

// TODO : unit tests
func (lc *LocalCoordinator) ListDataspace(ctx context.Context) ([]*dataspaces.Dataspace, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	resp, err := lc.qdb.ListDataspaces(ctx)
	if err != nil {
		return nil, err
	}
	var retDsp []*dataspaces.Dataspace

	for _, dsp := range resp {
		retDsp = append(retDsp, &dataspaces.Dataspace{
			Id: dsp.ID,
		})
	}
	return retDsp, nil
}

// TODO : unit tests
func (lc *LocalCoordinator) AddDataspace(ctx context.Context, ds *dataspaces.Dataspace) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	return lc.qdb.AddDataspace(ctx, &qdb.Dataspace{
		ID: ds.Id,
	})
}

// TODO : unit tests
func (lc *LocalCoordinator) AttachToDataspace(ctx context.Context, table string, ds *dataspaces.Dataspace) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	return lc.qdb.AttachToDataspace(ctx, table, ds.Id)
}

// TODO : unit tests
func (lc *LocalCoordinator) GetDataspace(ctx context.Context, table string) (*dataspaces.Dataspace, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	ret, err := lc.qdb.GetDataspace(ctx, table)
	if err != nil {
		return nil, err
	}
	return dataspaces.NewDataspace(ret.ID), nil
}

// TODO : unit tests
func (lc *LocalCoordinator) ListDataShards(ctx context.Context) []*datashards.DataShard {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	var ret []*datashards.DataShard
	for id, cfg := range lc.DataShardCfgs {
		ret = append(ret, datashards.NewDataShard(id, cfg))
	}
	return ret
}

// TODO : unit tests
func (lc *LocalCoordinator) DropDataspace(ctx context.Context, ds *dataspaces.Dataspace) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	return lc.qdb.DropDataspace(ctx, ds.Id)
}

// TODO : unit tests
func (lc *LocalCoordinator) ListShards(ctx context.Context) ([]*datashards.DataShard, error) {
	resp, err := lc.qdb.ListShards(ctx)
	if err != nil {
		return nil, err
	}
	var retShards []*datashards.DataShard

	for _, sh := range resp {
		retShards = append(retShards, &datashards.DataShard{
			ID: sh.ID,
			Cfg: &config.Shard{
				Hosts: sh.Hosts,
			},
		})
	}
	return retShards, nil
}

func (lc *LocalCoordinator) AddWorldShard(ctx context.Context, ds *datashards.DataShard) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	spqrlog.Zero.Info().
		Str("shard", ds.ID).
		Msg("adding world datashard")
	lc.WorldShardCfgs[ds.ID] = ds.Cfg

	return nil
}

// TODO : unit tests
func (lc *LocalCoordinator) DropKeyRange(ctx context.Context, id string) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	spqrlog.Zero.Info().
		Str("kr", id).
		Msg("dropping key range")
	return lc.qdb.DropKeyRange(ctx, id)
}

// TODO : unit tests
func (lc *LocalCoordinator) DropKeyRangeAll(ctx context.Context) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	spqrlog.Zero.Info().Msg("dropping all key ranges")
	return lc.qdb.DropKeyRangeAll(ctx)
}

// TODO : unit tests
func (lc *LocalCoordinator) DataShardsRoutes() []*routingstate.DataShardRoute {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	var ret []*routingstate.DataShardRoute

	for name := range lc.DataShardCfgs {
		ret = append(ret, &routingstate.DataShardRoute{
			Shkey: kr.ShardKey{
				Name: name,
				RW:   true,
			},
		})
	}

	return ret
}

// TODO : unit tests
func (lc *LocalCoordinator) WorldShardsRoutes() []*routingstate.DataShardRoute {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	var ret []*routingstate.DataShardRoute

	for name := range lc.WorldShardCfgs {
		ret = append(ret, &routingstate.DataShardRoute{
			Shkey: kr.ShardKey{
				Name: name,
				RW:   true,
			},
		})
	}

	// a sort of round robin

	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

func (lc *LocalCoordinator) WorldShards() []string {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	var ret []string

	for name := range lc.WorldShardCfgs {
		ret = append(ret, name)
	}

	return ret
}

// TODO : unit tests
func (qr *LocalCoordinator) Move(ctx context.Context, req *kr.MoveKeyRange) error {
	var krmv *qdb.KeyRange
	var err error
	if krmv, err = qr.qdb.CheckLockedKeyRange(ctx, req.Krid); err != nil {
		return err
	}

	var reqKr = kr.KeyRangeFromDB(krmv)
	reqKr.ShardID = req.ShardId
	return ops.ModifyKeyRangeWithChecks(ctx, qr.qdb, reqKr)
}

// TODO : unit tests
func (qr *LocalCoordinator) Unite(ctx context.Context, req *kr.UniteKeyRange) error {
	var krleft *qdb.KeyRange
	var krright *qdb.KeyRange
	var err error

	if krleft, err = qr.qdb.LockKeyRange(ctx, req.KeyRangeIDLeft); err != nil { //nolint:all TODO
		return err
	}
	defer func(qdb qdb.QDB, ctx context.Context, keyRangeID string) {
		err := qdb.UnlockKeyRange(ctx, keyRangeID)
		if err != nil {
			spqrlog.Zero.Error().Err(err).Msg("")
			return
		}
	}(qr.qdb, ctx, req.KeyRangeIDLeft)

	// TODO: krRight seems to be empty.
	if krright, err = qr.qdb.GetKeyRange(ctx, req.KeyRangeIDRight); err != nil {
		return err
	}

	if err = qr.qdb.DropKeyRange(ctx, krright.KeyRangeID); err != nil {
		return err
	}

	united := &kr.KeyRange{
		LowerBound: krleft.LowerBound,
		UpperBound: krright.UpperBound,
		ShardID:    krleft.ShardID,
		Dataspace:  krleft.DataspaceId,
		ID:         krleft.KeyRangeID,
	}

	return ops.ModifyKeyRangeWithChecks(ctx, qr.qdb, united)
}

// TODO : unit tests
func (qr *LocalCoordinator) Split(ctx context.Context, req *kr.SplitKeyRange) error {
	var krOld *qdb.KeyRange
	var err error

	spqrlog.Zero.Debug().
		Str("krid", req.Krid).
		Interface("bound", req.Bound).
		Str("source-id", req.SourceID).
		Msg("split request is")

	if krOld, err = qr.qdb.LockKeyRange(ctx, req.SourceID); err != nil {
		return err
	}

	defer func() {
		if err := qr.qdb.UnlockKeyRange(ctx, req.SourceID); err != nil {
			spqrlog.Zero.Error().Err(err).Msg("")
		}
	}()

	krNew := kr.KeyRangeFromDB(
		&qdb.KeyRange{
			LowerBound:  req.Bound,
			UpperBound:  krOld.UpperBound,
			KeyRangeID:  req.Krid,
			ShardID:     krOld.ShardID,
			DataspaceId: krOld.DataspaceId,
		},
	)

	spqrlog.Zero.Debug().
		Bytes("lower-bound", krNew.LowerBound).
		Bytes("upper-bound", krNew.UpperBound).
		Str("shard-id", krNew.ShardID).
		Str("id", krNew.ID).
		Msg("new key range")

	krOld.UpperBound = req.Bound
	if err := ops.ModifyKeyRangeWithChecks(ctx, qr.qdb, kr.KeyRangeFromDB(krOld)); err != nil {
		return err
	}

	if err := ops.AddKeyRangeWithChecks(ctx, qr.qdb, krNew); err != nil {
		return fmt.Errorf("failed to add a new key range: %w", err)
	}

	return nil
}

// TODO : unit tests
func (qr *LocalCoordinator) LockKeyRange(ctx context.Context, krid string) (*kr.KeyRange, error) {
	keyRangeDB, err := qr.qdb.LockKeyRange(ctx, krid)
	if err != nil {
		return nil, err
	}

	return kr.KeyRangeFromDB(keyRangeDB), nil
}

// TODO : unit tests
func (qr *LocalCoordinator) UnlockKeyRange(ctx context.Context, krid string) error {
	return qr.qdb.UnlockKeyRange(ctx, krid)
}

// TODO : unit tests
func (lc *LocalCoordinator) AddDataShard(ctx context.Context, ds *datashards.DataShard) error {
	spqrlog.Zero.Info().
		Str("node", ds.ID).
		Msg("adding node")

	lc.DataShardCfgs[ds.ID] = ds.Cfg

	return lc.qdb.AddShard(ctx, &qdb.Shard{
		ID:    ds.ID,
		Hosts: ds.Cfg.Hosts,
	})
}

// TODO : unit tests
func (qr *LocalCoordinator) Shards() []string {
	var ret []string

	for name := range qr.DataShardCfgs {
		ret = append(ret, name)
	}

	return ret
}

// TODO : unit tests
func (qr *LocalCoordinator) ListKeyRanges(ctx context.Context, dataspace string) ([]*kr.KeyRange, error) {
	var ret []*kr.KeyRange
	if krs, err := qr.qdb.ListKeyRanges(ctx, dataspace); err != nil {
		return nil, err
	} else {
		for _, keyRange := range krs {
			ret = append(ret, kr.KeyRangeFromDB(keyRange))

		}
	}

	return ret, nil
}

// TODO : unit tests
func (qr *LocalCoordinator) ListAllKeyRanges(ctx context.Context) ([]*kr.KeyRange, error) {
	var ret []*kr.KeyRange
	if krs, err := qr.qdb.ListAllKeyRanges(ctx); err != nil {
		return nil, err
	} else {
		for _, keyRange := range krs {
			ret = append(ret, kr.KeyRangeFromDB(keyRange))

		}
	}

	return ret, nil
}

// TODO : unit tests
func (qr *LocalCoordinator) ListRouters(ctx context.Context) ([]*topology.Router, error) {
	return []*topology.Router{{
		ID: "local",
	}}, nil
}

func (qr *LocalCoordinator) AddShardingRule(ctx context.Context, rule *shrule.ShardingRule) error {
	return ops.AddShardingRuleWithChecks(ctx, qr.qdb, rule)
}

// TODO : unit tests
func (qr *LocalCoordinator) ListShardingRules(ctx context.Context, dataspace string) ([]*shrule.ShardingRule, error) {
	rules, err := qr.qdb.ListShardingRules(ctx, dataspace)
	if err != nil {
		return nil, err
	}
	var resp []*shrule.ShardingRule
	for _, v := range rules {
		resp = append(resp, shrule.ShardingRuleFromDB(v))

	}

	return resp, nil
}

// TODO : unit tests
func (qr *LocalCoordinator) ListAllShardingRules(ctx context.Context) ([]*shrule.ShardingRule, error) {
	rules, err := qr.qdb.ListAllShardingRules(ctx)
	if err != nil {
		return nil, err
	}
	var resp []*shrule.ShardingRule
	for _, v := range rules {
		resp = append(resp, shrule.ShardingRuleFromDB(v))

	}

	return resp, nil
}

func (qr *LocalCoordinator) DropShardingRule(ctx context.Context, id string) error {
	return qr.qdb.DropShardingRule(ctx, id)
}

func (qr *LocalCoordinator) AddKeyRange(ctx context.Context, kr *kr.KeyRange) error {
	return ops.AddKeyRangeWithChecks(ctx, qr.qdb, kr)
}

func (qr *LocalCoordinator) MoveKeyRange(ctx context.Context, kr *kr.KeyRange) error {
	return ops.ModifyKeyRangeWithChecks(ctx, qr.qdb, kr)
}

var ErrNotCoordinator = fmt.Errorf("request is unprocessable in route")

// TODO : unit tests
func (qr *LocalCoordinator) DropShardingRuleAll(ctx context.Context) ([]*shrule.ShardingRule, error) {
	rules, err := qr.qdb.DropShardingRuleAll(ctx)
	if err != nil {
		return nil, err
	}
	var retRules []*shrule.ShardingRule

	for _, r := range rules {
		retRules = append(retRules, shrule.ShardingRuleFromDB(r))
	}

	return retRules, nil
}

func (qr *LocalCoordinator) RegisterRouter(ctx context.Context, r *topology.Router) error {
	return ErrNotCoordinator
}

func (qr *LocalCoordinator) UnregisterRouter(ctx context.Context, id string) error {
	return ErrNotCoordinator
}

func (qr *LocalCoordinator) SyncRouterMetadata(ctx context.Context, router *topology.Router) error {
	return ErrNotCoordinator
}

func (qr *LocalCoordinator) UpdateCoordinator(ctx context.Context, addr string) error {
	return qr.qdb.UpdateCoordinator(ctx, addr)
}

func (qr *LocalCoordinator) GetCoordinator(ctx context.Context) (string, error) {
	addr, err := qr.qdb.GetCoordinator(ctx)
	spqrlog.Zero.Debug().Str("address", addr).Msg("resp local coordiantor: get coordinator")
	return addr, err
}

func (qr *LocalCoordinator) GetShardInfo(ctx context.Context, shardID string) (*datashards.DataShard, error) {
	return nil, ErrNotCoordinator
}

func (lc *LocalCoordinator) ShareKeyRange(id string) error {
	return lc.qdb.ShareKeyRange(id)
}

func (lc *LocalCoordinator) QDB() qdb.QDB {
	return lc.qdb
}

func NewLocalCoordinator(db qdb.QDB) meta.EntityMgr {
	return &LocalCoordinator{
		DataShardCfgs:  map[string]*config.Shard{},
		WorldShardCfgs: map[string]*config.Shard{},
		qdb:            db,
	}
}
