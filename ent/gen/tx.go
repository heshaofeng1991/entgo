// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Announcements is the client for interacting with the Announcements builders.
	Announcements *AnnouncementsClient
	// Attribute is the client for interacting with the Attribute builders.
	Attribute *AttributeClient
	// Channel is the client for interacting with the Channel builders.
	Channel *ChannelClient
	// ChannelCost is the client for interacting with the ChannelCost builders.
	ChannelCost *ChannelCostClient
	// ChannelCostBatche is the client for interacting with the ChannelCostBatche builders.
	ChannelCostBatche *ChannelCostBatcheClient
	// ChannelOption is the client for interacting with the ChannelOption builders.
	ChannelOption *ChannelOptionClient
	// ChannelRecommend is the client for interacting with the ChannelRecommend builders.
	ChannelRecommend *ChannelRecommendClient
	// ChannelVolumeFactor is the client for interacting with the ChannelVolumeFactor builders.
	ChannelVolumeFactor *ChannelVolumeFactorClient
	// Country is the client for interacting with the Country builders.
	Country *CountryClient
	// CountryZone is the client for interacting with the CountryZone builders.
	CountryZone *CountryZoneClient
	// CourierOrder is the client for interacting with the CourierOrder builders.
	CourierOrder *CourierOrderClient
	// CustomerConfig is the client for interacting with the CustomerConfig builders.
	CustomerConfig *CustomerConfigClient
	// Inbound is the client for interacting with the Inbound builders.
	Inbound *InboundClient
	// InboundItem is the client for interacting with the InboundItem builders.
	InboundItem *InboundItemClient
	// Inventory is the client for interacting with the Inventory builders.
	Inventory *InventoryClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// OrderHoldReason is the client for interacting with the OrderHoldReason builders.
	OrderHoldReason *OrderHoldReasonClient
	// OrderItem is the client for interacting with the OrderItem builders.
	OrderItem *OrderItemClient
	// OrderTaxation is the client for interacting with the OrderTaxation builders.
	OrderTaxation *OrderTaxationClient
	// PickupOrder is the client for interacting with the PickupOrder builders.
	PickupOrder *PickupOrderClient
	// PickupOrderItem is the client for interacting with the PickupOrderItem builders.
	PickupOrderItem *PickupOrderItemClient
	// PlatformProduct is the client for interacting with the PlatformProduct builders.
	PlatformProduct *PlatformProductClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// ProductMapping is the client for interacting with the ProductMapping builders.
	ProductMapping *ProductMappingClient
	// Sequence is the client for interacting with the Sequence builders.
	Sequence *SequenceClient
	// Store is the client for interacting with the Store builders.
	Store *StoreClient
	// TaskSchedule is the client for interacting with the TaskSchedule builders.
	TaskSchedule *TaskScheduleClient
	// Tenant is the client for interacting with the Tenant builders.
	Tenant *TenantClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// TrackMapping is the client for interacting with the TrackMapping builders.
	TrackMapping *TrackMappingClient
	// Transaction is the client for interacting with the Transaction builders.
	Transaction *TransactionClient
	// TransactionDetail is the client for interacting with the TransactionDetail builders.
	TransactionDetail *TransactionDetailClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserReferral is the client for interacting with the UserReferral builders.
	UserReferral *UserReferralClient
	// ValueAddedTax is the client for interacting with the ValueAddedTax builders.
	ValueAddedTax *ValueAddedTaxClient
	// Warehouse is the client for interacting with the Warehouse builders.
	Warehouse *WarehouseClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once
	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	txDriver.mu.Lock()
	hooks := append([]CommitHook(nil), txDriver.onCommit...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onCommit = append(txDriver.onCommit, f)
	txDriver.mu.Unlock()
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	txDriver.mu.Lock()
	hooks := append([]RollbackHook(nil), txDriver.onRollback...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onRollback = append(txDriver.onRollback, f)
	txDriver.mu.Unlock()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.Announcements = NewAnnouncementsClient(tx.config)
	tx.Attribute = NewAttributeClient(tx.config)
	tx.Channel = NewChannelClient(tx.config)
	tx.ChannelCost = NewChannelCostClient(tx.config)
	tx.ChannelCostBatche = NewChannelCostBatcheClient(tx.config)
	tx.ChannelOption = NewChannelOptionClient(tx.config)
	tx.ChannelRecommend = NewChannelRecommendClient(tx.config)
	tx.ChannelVolumeFactor = NewChannelVolumeFactorClient(tx.config)
	tx.Country = NewCountryClient(tx.config)
	tx.CountryZone = NewCountryZoneClient(tx.config)
	tx.CourierOrder = NewCourierOrderClient(tx.config)
	tx.CustomerConfig = NewCustomerConfigClient(tx.config)
	tx.Inbound = NewInboundClient(tx.config)
	tx.InboundItem = NewInboundItemClient(tx.config)
	tx.Inventory = NewInventoryClient(tx.config)
	tx.Order = NewOrderClient(tx.config)
	tx.OrderHoldReason = NewOrderHoldReasonClient(tx.config)
	tx.OrderItem = NewOrderItemClient(tx.config)
	tx.OrderTaxation = NewOrderTaxationClient(tx.config)
	tx.PickupOrder = NewPickupOrderClient(tx.config)
	tx.PickupOrderItem = NewPickupOrderItemClient(tx.config)
	tx.PlatformProduct = NewPlatformProductClient(tx.config)
	tx.Product = NewProductClient(tx.config)
	tx.ProductMapping = NewProductMappingClient(tx.config)
	tx.Sequence = NewSequenceClient(tx.config)
	tx.Store = NewStoreClient(tx.config)
	tx.TaskSchedule = NewTaskScheduleClient(tx.config)
	tx.Tenant = NewTenantClient(tx.config)
	tx.Token = NewTokenClient(tx.config)
	tx.TrackMapping = NewTrackMappingClient(tx.config)
	tx.Transaction = NewTransactionClient(tx.config)
	tx.TransactionDetail = NewTransactionDetailClient(tx.config)
	tx.User = NewUserClient(tx.config)
	tx.UserReferral = NewUserReferralClient(tx.config)
	tx.ValueAddedTax = NewValueAddedTaxClient(tx.config)
	tx.Warehouse = NewWarehouseClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Announcements.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
	// completion hooks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v any) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v any) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)