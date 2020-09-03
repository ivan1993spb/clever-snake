package connections

import (
	"errors"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

const firstGroupId = 1

type ConnectionGroupManager struct {
	groups      map[int]*ConnectionGroup
	groupsMutex *sync.RWMutex
	groupLimit  int
	connsLimit  int
	connsCount  int
	logger      logrus.FieldLogger
}

func NewConnectionGroupManager(logger logrus.FieldLogger, groupLimit, connsLimit int) (*ConnectionGroupManager, error) {
	if groupLimit > 0 {
		return &ConnectionGroupManager{
			groups:      map[int]*ConnectionGroup{},
			groupsMutex: &sync.RWMutex{},
			groupLimit:  groupLimit,
			connsLimit:  connsLimit,
			logger:      logger,
		}, nil
	}

	return nil, errors.New("cannot create connection group manager: invalid group limit")
}

func (m *ConnectionGroupManager) unsafeIsFull() bool {
	return len(m.groups) == m.groupLimit
}

func (m *ConnectionGroupManager) IsFull() bool {
	m.groupsMutex.RLock()
	defer m.groupsMutex.RUnlock()
	return m.unsafeIsFull()
}

type ErrAddGroup string

func (e ErrAddGroup) Error() string {
	return "cannot add group: " + string(e)
}

var (
	ErrGroupLimitReached = ErrAddGroup("limit group count reached")
	ErrCannotGetID       = ErrAddGroup("cannot get id for group")
	ErrConnsLimitReached = ErrAddGroup("cannot reserve connections for group: connections count reached")
)

func (m *ConnectionGroupManager) Add(group *ConnectionGroup) (int, error) {
	// TODO: Fix method to receive group and required conn limit.

	// TODO: Fix method to return (id int, count int, err error), where
	// id is group identifier, count is reserved connection count for the
	// group, and err is error if occurred.

	m.groupsMutex.Lock()
	defer m.groupsMutex.Unlock()

	if m.unsafeIsFull() {
		return 0, ErrGroupLimitReached
	}

	if group.GetLimit() > m.connsLimit-m.connsCount {
		if m.connsLimit-m.connsCount < 1 {
			return 0, ErrConnsLimitReached
		}
		group.SetLimit(m.connsLimit - m.connsCount)
	}

	m.connsCount += group.GetLimit()

	for id := firstGroupId; id <= len(m.groups)+firstGroupId; id++ {
		if _, occupied := m.groups[id]; !occupied {
			m.groups[id] = group
			return id, nil
		}
	}

	return 0, ErrCannotGetID
}

type ErrDeleteGroup string

func (e ErrDeleteGroup) Error() string {
	return "cannot delete group: " + string(e)
}

var (
	ErrDeleteNotEmptyGroup = ErrDeleteGroup("group is not empty")
	ErrDeleteNotFoundGroup = ErrDeleteGroup("group not found")
)

func (m *ConnectionGroupManager) Delete(group *ConnectionGroup) error {
	// TODO: Return (err error, id int).

	m.groupsMutex.Lock()
	defer m.groupsMutex.Unlock()

	// TODO: Move that checking in the core module.
	if !group.IsEmpty() {
		return ErrDeleteNotEmptyGroup
	}

	for id := range m.groups {
		if m.groups[id] == group {
			delete(m.groups, id)
			m.connsCount -= group.GetLimit()

			return nil
		}
	}

	return ErrDeleteNotFoundGroup
}

var ErrNotFoundGroup = errors.New("not found group")

func (m *ConnectionGroupManager) Get(id int) (*ConnectionGroup, error) {
	m.groupsMutex.RLock()
	defer m.groupsMutex.RUnlock()

	if group, ok := m.groups[id]; ok {
		return group, nil
	}

	return nil, ErrNotFoundGroup
}

func (m *ConnectionGroupManager) Groups() map[int]*ConnectionGroup {
	m.groupsMutex.RLock()
	defer m.groupsMutex.RUnlock()
	groups := map[int]*ConnectionGroup{}
	for id, group := range m.groups {
		groups[id] = group
	}
	return groups
}

func (m *ConnectionGroupManager) GroupLimit() int {
	return m.groupLimit
}

func (m *ConnectionGroupManager) unsafeGroupCount() int {
	return len(m.groups)
}

func (m *ConnectionGroupManager) GroupCount() int {
	m.groupsMutex.RLock()
	defer m.groupsMutex.RUnlock()
	return m.unsafeGroupCount()
}

func (m *ConnectionGroupManager) unsafeConnCount() int {
	var count = 0
	for _, group := range m.groups {
		count += group.GetCount()
	}
	return count
}

func (m *ConnectionGroupManager) unsafeCapacity() float32 {
	var count = m.unsafeConnCount()
	return float32(count) / float32(m.connsLimit)
}

func (m *ConnectionGroupManager) Capacity() float32 {
	m.groupsMutex.RLock()
	defer m.groupsMutex.RUnlock()
	return m.unsafeCapacity()
}

const (
	serverGamesPlayersGameIdLabel = "game_id"

	serverCapacityFQName     = "server_capacity"
	serverGamesFQName        = "server_games"
	serverGamesPlayersFQName = "server_games_players"

	serverCapacityHelp     = "Capacity of the server"
	serverGamesHelp        = "Games number"
	serverGamesPlayersHelp = "Players number"
)

var (
	serverCapacityDesc = prometheus.NewDesc(
		serverCapacityFQName,
		serverCapacityHelp,
		nil,
		nil,
	)
	serverGamesDesc = prometheus.NewDesc(
		serverGamesFQName,
		serverGamesHelp,
		nil,
		nil,
	)
	serverGamesPlayersDesc = prometheus.NewDesc(
		serverGamesPlayersFQName,
		serverGamesPlayersHelp,
		[]string{serverGamesPlayersGameIdLabel},
		nil,
	)
)

func (m *ConnectionGroupManager) Describe(ch chan<- *prometheus.Desc) {
	var descriptors = [...]*prometheus.Desc{
		serverCapacityDesc,
		serverGamesDesc,
		serverGamesPlayersDesc,
	}
	for _, desc := range descriptors {
		ch <- desc
	}
}

func (m *ConnectionGroupManager) Collect(ch chan<- prometheus.Metric) {
	sendMetric := func(desc *prometheus.Desc, valueType prometheus.ValueType,
		value float64, labelValues ...string) {
		metric, err := prometheus.NewConstMetric(desc, valueType, value, labelValues...)
		if err != nil {
			m.logger.Errorln("cannot create a metric:", err)
		} else {
			ch <- metric
		}
	}

	m.groupsMutex.RLock()
	defer m.groupsMutex.RUnlock()

	sendMetric(serverCapacityDesc, prometheus.GaugeValue, float64(m.unsafeCapacity()))
	sendMetric(serverGamesDesc, prometheus.GaugeValue, float64(m.unsafeGroupCount()))

	for id, group := range m.groups {
		sendMetric(serverGamesPlayersDesc, prometheus.GaugeValue, float64(group.GetCount()), strconv.Itoa(id))
	}
}
