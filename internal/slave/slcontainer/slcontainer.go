// Package slcontainer provides the container for the slave node
package slcontainer

import (
	"fmt"
	"sync"
)

// SlaveContainer represents the container for the slave node
type SlaveContainer struct {
	Auth                          *Auth
	Store                         *Store
	Target                        *Target
	Loader                        *Loader
	CommandMap                    *sync.Map
	ReceiveChanelRequestContainer *ReceiveChanelRequestContainer
}

// NewSlaveContainer creates a new container for the slave node
func NewSlaveContainer() *SlaveContainer {
	return &SlaveContainer{
		Auth:                          NewAuth(),                          // DON'T CHANGE POINTER TO VALUE
		Store:                         NewStore(),                         // DON'T CHANGE POINTER TO VALUE
		Target:                        NewTarget(),                        // DON'T CHANGE POINTER TO VALUE
		Loader:                        NewLoader(),                        // DON'T CHANGE POINTER TO VALUE
		CommandMap:                    &sync.Map{},                        // DON'T CHANGE POINTER TO VALUE
		ReceiveChanelRequestContainer: NewReceiveChanelRequestContainer(), // DON'T CHANGE POINTER TO VALUE
	}
}

// CommandMapData represents the data for the command map
type CommandMapData struct {
	LoaderID         string
	OutputRoot       string
	StrMap           *sync.Map
	ThreadOnlyStrMap *sync.Map
	SlaveValues      map[string]any
}

// AddCommandMap adds a command map to the slave container
func (s *SlaveContainer) AddCommandMap(cmdID string, data CommandMapData) {
	s.CommandMap.Store(cmdID, data)
}

// GetCommandMap returns the command map from the slave container
func (s *SlaveContainer) GetCommandMap(cmdID string) (CommandMapData, bool) {
	if loaderID, ok := s.CommandMap.Load(cmdID); ok {
		if v, ok := loaderID.(CommandMapData); ok {
			return v, true
		}
	}
	return CommandMapData{}, false
}

// DeleteCommandMap deletes the command map from the slave container
func (s *SlaveContainer) DeleteCommandMap(cmdID string) {
	s.CommandMap.Delete(cmdID)
}

// SetStrMap sets the str map for the command map
func (s *SlaveContainer) SetStrMap(cmdID string, strMap map[string]any) error {
	syncMap := &sync.Map{}
	for k, v := range strMap {
		syncMap.Store(k, v)
	}
	if data, ok := s.GetCommandMap(cmdID); ok {
		data.StrMap = syncMap
		s.AddCommandMap(cmdID, data)
	} else {
		return fmt.Errorf("CommandMap not found")
	}
	return nil
}

// SetThreadOnlyStrMap sets the thread only str map for the command map
func (s *SlaveContainer) SetThreadOnlyStrMap(cmdID string, strMap map[string]any) error {
	syncMap := &sync.Map{}
	for k, v := range strMap {
		syncMap.Store(k, v)
	}
	if data, ok := s.GetCommandMap(cmdID); ok {
		data.ThreadOnlyStrMap = syncMap
		s.AddCommandMap(cmdID, data)
	} else {
		return fmt.Errorf("CommandMap not found")
	}
	return nil
}

// SetSlaveValues sets the slave values for the command map
func (s *SlaveContainer) SetSlaveValues(cmdID string, slaveValues map[string]any) error {
	if data, ok := s.GetCommandMap(cmdID); ok {
		data.SlaveValues = slaveValues
		s.AddCommandMap(cmdID, data)
	} else {
		return fmt.Errorf("CommandMap not found")
	}
	return nil
}
