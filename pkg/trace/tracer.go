package trace

import (
	"errors"
	"os"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
)

// Entry is an interface for all structs that are used to define the schema for
// traces.
type Entry interface {
	// Table defines which table the struct belongs to.
	Table() string
	// InfluxRepr returns the InfluxDB representation of the struct.
	InfluxRepr() (map[string]interface{}, error)
}

// Tracer defines the methods for a client that can write and read trace data.
type Tracer interface {
	Write(Entry)
	ReadTable(table string) (*os.File, error)
	IsCollecting(table string) bool
	Stop()
}

func NewTracer(cfg *config.Config, logger log.Logger, chainID, nodeID string) (Tracer, error) {
	return NewInfluxClient(cfg.Instrumentation, logger, chainID, nodeID)
}

func NoOpTracer() Tracer {
	return &noOpTracer{}
}

type noOpTracer struct{}

func (n *noOpTracer) Write(_ Entry) {}
func (n *noOpTracer) ReadTable(_ string) (*os.File, error) {
	return nil, errors.New("no-op tracer does not support reading")
}
func (n *noOpTracer) IsCollecting(_ string) bool { return false }
func (n *noOpTracer) Stop()                      {}
