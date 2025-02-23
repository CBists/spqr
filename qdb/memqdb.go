package qdb

import (
	"context"
	"encoding/json"
	"github.com/pg-sharding/spqr/pkg/models/spqrerror"
	"os"
	"sort"
	"sync"

	"github.com/pg-sharding/spqr/pkg/spqrlog"
)

type MemQDB struct {
	ShardingSchemaKeeper
	// TODO create more mutex per map if needed
	mu           sync.RWMutex
	muDeletedKrs sync.RWMutex

	deletedKrs     map[string]bool
	Locks          map[string]*sync.RWMutex            `json:"locks"`
	Freq           map[string]bool                     `json:"freq"`
	Krs            map[string]*KeyRange                `json:"krs"`
	Shards         map[string]*Shard                   `json:"shards"`
	Shrules        map[string]*ShardingRule            `json:"shrules"`
	Dataspaces     map[string]*Dataspace               `json:"dataspaces"`
	TableDataspace map[string]string                   `json:"table_dataspace"`
	Routers        map[string]*Router                  `json:"routers"`
	Transactions   map[string]*DataTransferTransaction `json:"transactions"`
	Coordinator    string                              `json:"coordinator"`

	backupPath string
	/* caches */
}

var _ QDB = &MemQDB{}

func NewMemQDB(backupPath string) (*MemQDB, error) {
	return &MemQDB{
		Freq:           map[string]bool{},
		Krs:            map[string]*KeyRange{},
		Locks:          map[string]*sync.RWMutex{},
		Shards:         map[string]*Shard{},
		Shrules:        map[string]*ShardingRule{},
		Dataspaces:     map[string]*Dataspace{},
		TableDataspace: map[string]string{},
		Routers:        map[string]*Router{},
		Transactions:   map[string]*DataTransferTransaction{},
		deletedKrs:     map[string]bool{},

		backupPath: backupPath,
	}, nil
}

// TODO : unit tests
func RestoreQDB(backupPath string) (*MemQDB, error) {
	qdb, err := NewMemQDB(backupPath)
	if err != nil {
		return nil, err
	}
	if backupPath == "" {
		return qdb, nil
	}
	if _, err := os.Stat(backupPath); err != nil {
		spqrlog.Zero.Info().Err(err).Msg("memqdb backup file not exists. Creating new one.")
		f, err := os.Create(backupPath)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		return qdb, nil
	}
	data, err := os.ReadFile(backupPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, qdb)

	for kr, locked := range qdb.Freq {
		if locked {
			qdb.Locks[kr].Lock()
		}
	}

	if err != nil {
		return nil, err
	}
	return qdb, nil
}

// TODO : unit tests
func (q *MemQDB) DumpState() error {
	if q.backupPath == "" {
		return nil
	}
	tmpPath := q.backupPath + ".tmp"

	f, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	state, err := json.MarshalIndent(q, "", "	")

	if err != nil {
		return err
	}

	_, err = f.Write(state)
	if err != nil {
		return err
	}
	f.Close()

	err = os.Rename(tmpPath, q.backupPath)
	if err != nil {
		return err
	}

	return nil
}

// ==============================================================================
//                               SHARDING RULES
// ==============================================================================

// TODO : unit tests
func (q *MemQDB) AddShardingRule(ctx context.Context, rule *ShardingRule) error {
	spqrlog.Zero.Debug().Interface("rule", rule).Msg("memqdb: add sharding rule")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Shrules, rule.ID, rule))
}

// TODO : unit tests
func (q *MemQDB) DropShardingRule(ctx context.Context, id string) error {
	spqrlog.Zero.Debug().Str("rule", id).Msg("memqdb: drop sharding rule")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewDeleteCommand(q.Shrules, id))
}

// TODO : unit tests
func (q *MemQDB) DropShardingRuleAll(ctx context.Context) ([]*ShardingRule, error) {
	spqrlog.Zero.Debug().Msg("memqdb: drop sharding rule all")
	q.mu.Lock()
	defer q.mu.Unlock()

	var ret []*ShardingRule
	for _, v := range q.Shrules {
		ret = append(ret, v)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ID < ret[j].ID
	})

	err := ExecuteCommands(q.DumpState, NewDropCommand(q.Shrules))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// TODO : unit tests
func (q *MemQDB) GetShardingRule(ctx context.Context, id string) (*ShardingRule, error) {
	spqrlog.Zero.Debug().Str("rule", id).Msg("memqdb: get sharding rule")
	q.mu.RLock()
	defer q.mu.RUnlock()
	rule, ok := q.Shrules[id]
	if ok {
		return rule, nil
	}
	return nil, spqrerror.Newf(spqrerror.SPQR_SHARDING_RULE_ERROR, "rule with id %s not found", id)
}

// TODO : unit tests
func (q *MemQDB) ListShardingRules(ctx context.Context, dataspace string) ([]*ShardingRule, error) {
	spqrlog.Zero.Debug().
		Str("dataspace", dataspace).
		Msg("memqdb: list sharding rules")
	q.mu.RLock()
	defer q.mu.RUnlock()
	var ret []*ShardingRule
	for _, v := range q.Shrules {
		if dataspace == v.DataspaceId {
			ret = append(ret, v)
		}
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ID < ret[j].ID
	})

	return ret, nil
}

// TODO : unit tests
func (q *MemQDB) ListAllShardingRules(ctx context.Context) ([]*ShardingRule, error) {
	spqrlog.Zero.Debug().Msg("memqdb: list all sharding rules")
	q.mu.RLock()
	defer q.mu.RUnlock()
	var ret []*ShardingRule
	for _, v := range q.Shrules {
		ret = append(ret, v)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ID < ret[j].ID
	})

	return ret, nil
}

// TODO : unit tests
func (q *MemQDB) MatchShardingRules(ctx context.Context, m func(shrules map[string]*ShardingRule) error) error {
	spqrlog.Zero.Debug().Msg("memqdb: match sharding rules")
	q.mu.RLock()
	defer q.mu.RUnlock()
	return m(q.Shrules)
}

// ==============================================================================
//                                 KEY RANGES
// ==============================================================================

// TODO : unit tests
func (q *MemQDB) AddKeyRange(ctx context.Context, keyRange *KeyRange) error {
	spqrlog.Zero.Debug().Interface("key-range", keyRange).Msg("memqdb: add key range")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Krs, keyRange.KeyRangeID, keyRange),
		NewUpdateCommand(q.Locks, keyRange.KeyRangeID, &sync.RWMutex{}),
		NewUpdateCommand(q.Freq, keyRange.KeyRangeID, false))
}

// TODO : unit tests
func (q *MemQDB) GetKeyRange(ctx context.Context, id string) (*KeyRange, error) {
	spqrlog.Zero.Debug().Str("key-range", id).Msg("memqdb: get key range")
	q.mu.RLock()
	defer q.mu.RUnlock()

	krs, ok := q.Krs[id]
	if !ok {
		return nil, spqrerror.Newf(spqrerror.SPQR_KEYRANGE_ERROR, "there is no key range %s", id)
	}

	return krs, nil
}

// TODO : unit tests
func (q *MemQDB) UpdateKeyRange(ctx context.Context, keyRange *KeyRange) error {
	spqrlog.Zero.Debug().Interface("key-range", keyRange).Msg("memqdb: update key range")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Krs, keyRange.KeyRangeID, keyRange))
}

// TODO : unit tests
func (q *MemQDB) DropKeyRange(ctx context.Context, id string) error {
	spqrlog.Zero.Debug().Str("key-range", id).Msg("memqdb: drop key range")

	// Do not allow new locks on key range we want to delete
	q.muDeletedKrs.Lock()
	if q.deletedKrs[id] {
		q.muDeletedKrs.Unlock()
		return spqrerror.Newf(spqrerror.SPQR_KEYRANGE_ERROR, "key range '%s' already deleted", id)
	}
	q.deletedKrs[id] = true
	q.muDeletedKrs.Unlock()

	defer func() {
		q.muDeletedKrs.Lock()
		defer q.muDeletedKrs.Unlock()
		delete(q.deletedKrs, id)
	}()

	q.mu.RLock()

	// Wait until key range will be unlocked
	if lock, ok := q.Locks[id]; ok {
		lock.Lock()
		defer lock.Unlock()
	}
	q.mu.RUnlock()

	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewDeleteCommand(q.Krs, id),
		NewDeleteCommand(q.Freq, id), NewDeleteCommand(q.Locks, id))
}

// TODO : unit tests
func (q *MemQDB) DropKeyRangeAll(ctx context.Context) error {
	spqrlog.Zero.Debug().Msg("memqdb: drop all key ranges")
	q.mu.RLock()

	// Do not allow new locks on key range we want to delete
	q.muDeletedKrs.Lock()
	ids := make([]string, 0)
	for id := range q.Locks {
		if q.deletedKrs[id] {
			q.muDeletedKrs.Unlock()
			q.mu.RUnlock()
			return spqrerror.Newf(spqrerror.SPQR_KEYRANGE_ERROR, "key range '%s' already deleted", id)
		}
		ids = append(ids, id)
		q.deletedKrs[id] = true
	}
	q.muDeletedKrs.Unlock()
	defer func() {
		q.muDeletedKrs.Lock()
		defer q.muDeletedKrs.Unlock()
		spqrlog.Zero.Debug().Msg("delete previous marks")

		for _, id := range ids {
			delete(q.deletedKrs, id)
		}
	}()

	// Wait until key range will be unlocked
	var locks []*sync.RWMutex
	for _, l := range q.Locks {
		l.Lock()
		locks = append(locks, l)
	}
	defer func() {
		for _, l := range locks {
			l.Unlock()
		}
	}()
	spqrlog.Zero.Debug().Msg("memqdb: acquired all locks")

	q.mu.RUnlock()

	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewDropCommand(q.Krs), NewDropCommand(q.Locks))
}

// TODO : unit tests
func (q *MemQDB) ListKeyRanges(_ context.Context, dataspace string) ([]*KeyRange, error) {
	spqrlog.Zero.Debug().
		Str("dataspace", dataspace).
		Msg("memqdb: list key ranges")
	q.mu.RLock()
	defer q.mu.RUnlock()

	var ret []*KeyRange

	for _, el := range q.Krs {
		if el.DataspaceId == dataspace {
			ret = append(ret, el)
		}
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].KeyRangeID < ret[j].KeyRangeID
	})

	return ret, nil
}

// TODO : unit tests
func (q *MemQDB) ListAllKeyRanges(_ context.Context) ([]*KeyRange, error) {
	spqrlog.Zero.Debug().Msg("memqdb: list all key ranges")
	q.mu.RLock()
	defer q.mu.RUnlock()

	var ret []*KeyRange

	for _, el := range q.Krs {
		ret = append(ret, el)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].KeyRangeID < ret[j].KeyRangeID
	})

	return ret, nil
}

// TODO : unit tests
func (q *MemQDB) TryLockKeyRange(lock *sync.RWMutex, id string, read bool) error {
	q.muDeletedKrs.RLock()

	if _, ok := q.deletedKrs[id]; ok {
		q.muDeletedKrs.RUnlock()
		return spqrerror.Newf(spqrerror.SPQR_KEYRANGE_ERROR, "key range '%s' deleted", id)
	}
	q.muDeletedKrs.RUnlock()

	if read {
		lock.RLock()
	} else {
		lock.Lock()
	}

	if _, ok := q.Krs[id]; !ok {
		return spqrerror.Newf(spqrerror.SPQR_KEYRANGE_ERROR, "key range '%s' deleted after lock acuired", id)
	}
	return nil
}

// TODO : unit tests
func (q *MemQDB) LockKeyRange(_ context.Context, id string) (*KeyRange, error) {
	spqrlog.Zero.Debug().Str("key-range", id).Msg("memqdb: lock key range")
	q.mu.RLock()
	defer q.mu.RUnlock()
	defer spqrlog.Zero.Debug().Str("key-range", id).Msg("memqdb: exit: lock key range")

	krs, ok := q.Krs[id]
	if !ok {
		return nil, spqrerror.Newf(spqrerror.SPQR_KEYRANGE_ERROR, "key range '%s' does not exist", id)
	}

	err := ExecuteCommands(q.DumpState, NewUpdateCommand(q.Freq, id, true),
		NewCustomCommand(func() error {
			if lock, ok := q.Locks[id]; ok {
				return q.TryLockKeyRange(lock, id, false)
			}
			return nil
		}, func() error {
			if lock, ok := q.Locks[id]; ok {
				lock.Unlock()
			}
			return nil
		}))
	if err != nil {
		return nil, err
	}

	return krs, nil
}

// TODO : unit tests
func (q *MemQDB) UnlockKeyRange(_ context.Context, id string) error {
	spqrlog.Zero.Debug().Str("key-range", id).Msg("memqdb: unlock key range")
	q.mu.RLock()
	defer q.mu.RUnlock()
	defer spqrlog.Zero.Debug().Str("key-range", id).Msg("memqdb: exit: unlock key range")

	if !q.Freq[id] {
		return spqrerror.Newf(spqrerror.SPQR_KEYRANGE_ERROR, "key range %v not locked", id)
	}

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Freq, id, false),
		NewCustomCommand(func() error {
			if lock, ok := q.Locks[id]; ok {
				lock.Unlock()
			}
			return nil
		}, func() error {
			if lock, ok := q.Locks[id]; ok {
				return q.TryLockKeyRange(lock, id, false)
			}
			return nil
		}))
}

// TODO : unit tests
func (q *MemQDB) CheckLockedKeyRange(ctx context.Context, id string) (*KeyRange, error) {
	spqrlog.Zero.Debug().Str("key-range", id).Msg("memqdb: check locked key range")
	q.mu.RLock()
	defer q.mu.RUnlock()

	krs, err := q.GetKeyRange(ctx, id)
	if err != nil {
		return nil, err
	}

	if !q.Freq[id] {
		return nil, spqrerror.Newf(spqrerror.SPQR_KEYRANGE_ERROR, "key range %v not locked", id)
	}

	return krs, nil
}

// TODO : unit tests
func (q *MemQDB) ShareKeyRange(id string) error {
	spqrlog.Zero.Debug().Str("key-range", id).Msg("memqdb: sharing key with key")

	q.mu.RLock()
	defer q.mu.RUnlock()

	lock, ok := q.Locks[id]
	if !ok {
		return spqrerror.New(spqrerror.SPQR_KEYRANGE_ERROR, "no such key")
	}

	err := q.TryLockKeyRange(lock, id, true)
	if err != nil {
		return err
	}
	defer lock.RUnlock()

	return nil
}

// ==============================================================================
//                           Transfer transactions
// ==============================================================================

// TODO : unit tests
func (q *MemQDB) RecordTransferTx(ctx context.Context, key string, info *DataTransferTransaction) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Transactions, key, info))
}

// TODO : unit tests
func (q *MemQDB) GetTransferTx(ctx context.Context, key string) (*DataTransferTransaction, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	ans, ok := q.Transactions[key]
	if !ok {
		return nil, spqrerror.Newf(spqrerror.SPQR_TRANSFER_ERROR, "no tx with key %s", key)
	}
	return ans, nil
}

// TODO : unit tests
func (q *MemQDB) RemoveTransferTx(ctx context.Context, key string) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	return ExecuteCommands(q.DumpState, NewDeleteCommand(q.Transactions, key))
}

// ==============================================================================
//	                           COORDINATOR LOCK
// ==============================================================================

func (q *MemQDB) TryCoordinatorLock(ctx context.Context) error {
	return nil
}

// TODO : unit tests
func (q *MemQDB) UpdateCoordinator(ctx context.Context, address string) error {
	spqrlog.Zero.Debug().Str("address", address).Msg("memqdb: update coordinator address")

	q.mu.Lock()
	defer q.mu.Unlock()

	q.Coordinator = address
	return nil
}

func (q *MemQDB) GetCoordinator(ctx context.Context) (string, error) {
	return q.Coordinator, nil
}

// ==============================================================================
//                                  ROUTERS
// ==============================================================================

// TODO : unit tests
func (q *MemQDB) AddRouter(ctx context.Context, r *Router) error {
	spqrlog.Zero.Debug().Interface("router", r).Msg("memqdb: add router")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Routers, r.ID, r))
}

// TODO : unit tests
func (q *MemQDB) DeleteRouter(ctx context.Context, id string) error {
	spqrlog.Zero.Debug().Str("router", id).Msg("memqdb: delete router")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewDeleteCommand(q.Routers, id))
}

// TODO : unit tests
func (q *MemQDB) OpenRouter(ctx context.Context, id string) error {
	spqrlog.Zero.Debug().
		Str("router", id).
		Msg("memqdb: open router")
	q.mu.Lock()
	defer q.mu.Unlock()

	q.Routers[id].State = OPENED

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Routers, id, q.Routers[id]))
}

// TODO : unit tests
func (q *MemQDB) CloseRouter(ctx context.Context, id string) error {
	spqrlog.Zero.Debug().
		Str("router", id).
		Msg("memqdb: open router")
	q.mu.Lock()
	defer q.mu.Unlock()

	q.Routers[id].State = CLOSED

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Routers, id, q.Routers[id]))
}

// TODO : unit tests
func (q *MemQDB) ListRouters(ctx context.Context) ([]*Router, error) {
	spqrlog.Zero.Debug().Msg("memqdb: list routers")
	q.mu.RLock()
	defer q.mu.RUnlock()

	var ret []*Router
	for _, v := range q.Routers {
		// TODO replace with new
		ret = append(ret, v)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ID < ret[j].ID
	})

	return ret, nil
}

// ==============================================================================
//                                  SHARDS
// ==============================================================================

// TODO : unit tests
func (q *MemQDB) AddShard(ctx context.Context, shard *Shard) error {
	spqrlog.Zero.Debug().Interface("shard", shard).Msg("memqdb: add shard")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Shards, shard.ID, shard))
}

// TODO : unit tests
func (q *MemQDB) ListShards(ctx context.Context) ([]*Shard, error) {
	spqrlog.Zero.Debug().Msg("memqdb: list shards")
	q.mu.RLock()
	defer q.mu.RUnlock()

	var ret []*Shard
	for _, v := range q.Shards {
		// TODO replace with new
		ret = append(ret, v)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ID < ret[j].ID
	})

	return ret, nil
}

// TODO : unit tests
func (q *MemQDB) GetShard(ctx context.Context, id string) (*Shard, error) {
	spqrlog.Zero.Debug().Str("shard", id).Msg("memqdb: get shard")
	q.mu.RLock()
	defer q.mu.RUnlock()

	if _, ok := q.Shards[id]; ok {
		return &Shard{ID: id}, nil
	}

	return nil, spqrerror.Newf(spqrerror.SPQR_NO_DATASHARD, "unknown shard %s", id)
}

// ==============================================================================
//                                 DATASPACES
// ==============================================================================

// TODO : unit tests
func (q *MemQDB) AddDataspace(ctx context.Context, dataspace *Dataspace) error {
	spqrlog.Zero.Debug().Interface("dataspace", dataspace).Msg("memqdb: add dataspace")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.Dataspaces, dataspace.ID, dataspace))
}

// TODO : unit tests
func (q *MemQDB) ListDataspaces(ctx context.Context) ([]*Dataspace, error) {
	spqrlog.Zero.Debug().Msg("memqdb: list dataspaces")
	q.mu.RLock()
	defer q.mu.RUnlock()
	var ret []*Dataspace
	ret = append(ret, &Dataspace{ID: "default"})
	for _, v := range q.Dataspaces {
		ret = append(ret, v)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ID < ret[j].ID
	})

	return ret, nil
}

// TODO : unit tests
func (q *MemQDB) DropDataspace(ctx context.Context, id string) error {
	spqrlog.Zero.Debug().Str("dataspace", id).Msg("memqdb: delete dataspace")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewDeleteCommand(q.Dataspaces, id))
}

// TODO : unit tests
func (q *MemQDB) AttachToDataspace(ctx context.Context, table string, id string) error {
	spqrlog.Zero.Debug().Str("dataspace", id).Msg("memqdb: attach table to dataspace")
	q.mu.Lock()
	defer q.mu.Unlock()

	return ExecuteCommands(q.DumpState, NewUpdateCommand(q.TableDataspace, table, id))
}

// TODO : unit tests
func (q *MemQDB) GetDataspace(ctx context.Context, table string) (*Dataspace, error) {
	spqrlog.Zero.Debug().Msg("memqdb: get dataspace for table")
	q.mu.RLock()
	defer q.mu.RUnlock()

	if _, ok := q.TableDataspace[table]; !ok {
		return &Dataspace{ID: "default"}, nil
	}

	if dataspace, ok := q.Dataspaces[q.TableDataspace[table]]; ok {
		return dataspace, nil
	} else {
		return nil, spqrerror.Newf(spqrerror.SPQR_NO_DATASPACE, "dataspace with id \"%s\" not found", q.TableDataspace[table])
	}
}
