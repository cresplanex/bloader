package runner

import (
	"context"

	"github.com/cresplanex/bloader/internal/utils"
)

// Event represents the flow step flow depends on event
type Event string

const (
	// RunnerEventStart represents the event start
	RunnerEventStart Event = "sys:start"
	// RunnerEventStoreImporting represents the event store importing
	RunnerEventStoreImporting Event = "sys:store:importing"
	// RunnerEventStoreImported represents the event store imported
	RunnerEventStoreImported Event = "sys:store:imported"
	// RunnerEventValidating represents the event validating
	RunnerEventValidating Event = "sys:validating"
	// RunnerEventValidated represents the event validated
	RunnerEventValidated Event = "sys:validated"
	// RunnerEventTerminated represents the event terminated
	RunnerEventTerminated Event = "sys:terminated"
)

// EventCaster is an interface for casting event
type EventCaster interface {
	// CastEvent casts the event
	CastEvent(ctx context.Context, event Event) error
	// CastEventWithWait casts the event with wait
	CastEventWithWait(ctx context.Context, event Event) error
	// Subscribe subscribes to the event
	Subscribe(ctx context.Context) error
	// Unsubscribe unsubscribes to the event
	Unsubscribe(ctx context.Context, ch chan Event) error
	// Close closes the event caster
	Close(ctx context.Context) error
}

// DefaultEventCaster is a struct that holds the event caster information
type DefaultEventCaster struct {
	// Caster is an event caster
	Caster *utils.Broadcaster[Event]
}

// NewDefaultEventCaster creates a new DefaultEventCaster
func NewDefaultEventCaster() *DefaultEventCaster {
	return &DefaultEventCaster{
		Caster: utils.NewBroadcaster[Event](),
	}
}

// NewDefaultEventCasterWithBroadcaster creates a new DefaultEventCaster with broadcaster
func NewDefaultEventCasterWithBroadcaster(broadcaster *utils.Broadcaster[Event]) *DefaultEventCaster {
	return &DefaultEventCaster{
		Caster: broadcaster,
	}
}

// CastEvent casts the event
func (ec *DefaultEventCaster) CastEvent(_ context.Context, event Event) error {
	ec.Caster.Broadcast(event)
	return nil
}

// CastEventWithWait casts the event with wait
func (ec *DefaultEventCaster) CastEventWithWait(ctx context.Context, event Event) error {
	waitChan := ec.Caster.Broadcast(event)
	select {
	case <-waitChan:
	case <-ctx.Done():
	}
	return nil
}

// Subscribe subscribes to the event
func (ec *DefaultEventCaster) Subscribe(_ context.Context) error {
	ec.Caster.Subscribe()
	return nil
}

// Unsubscribe unsubscribes to the event
func (ec *DefaultEventCaster) Unsubscribe(_ context.Context, ch chan Event) error {
	ec.Caster.Unsubscribe(ch)
	return nil
}

// Close closes the event caster
func (ec *DefaultEventCaster) Close(_ context.Context) error {
	ec.Caster.Close()
	return nil
}
