package monitoring

import (
	"errors"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

// Visibility instances will be called especially in Metrics
type Visibility interface {
	RecordMetric(name string, value float64)
	RecordEvent(name string, event map[string]interface{})
	StartTransaction(name string) Transaction
}

// Visibility stores an instance of a newrelic application
type newrelicMonitoring struct {
	app     *newrelic.Application
	enabled bool
}

// NewVisibility returns an instance of a new newrelic application
func NewVisibility() Visibility {
	a := new(newrelicMonitoring)
	a.enabled = false

	monitoringAppName := ""
	newrelicEnabled := os.Getenv("ENABLE_NEWRELIC")
	if newrelicEnabled != "" {
		monitoringAppName = os.Getenv("NEWRELIC_APP_NAME")
		if monitoringAppName == "" {
			panic(errors.New("environment variable NEWRELIC_APP_NAME must be specified"))
		}

		key := os.Getenv("NEWRELIC_LICENSE_KEY")
		if key == "" {
			logrus.Fatalf("env NEWRELIC_LICENSE_KEY undefined")
		}
		app, err := newrelic.NewApplication(
			newrelic.ConfigAppName(monitoringAppName),
			newrelic.ConfigLicense(key),
			func(cfg *newrelic.Config) {
				cfg.ErrorCollector.RecordPanics = true
			},
			nrlogrus.ConfigStandardLogger(),
		)
		if err != nil {
			logrus.Fatalf("could not create newrelic app; err: %v", err)
		}
		a.enabled = true
		a.app = app
	} else {
		logrus.Warn("ENABLE_NEWRELIC is not defined, disabling monitoring")
	}
	return a
}

func (nRelic *newrelicMonitoring) RecordMetric(name string, value float64) {
	if !nRelic.enabled {
		return
	}
	nRelic.app.RecordCustomMetric(name, value)
}

func (nRelic *newrelicMonitoring) RecordEvent(name string, event map[string]interface{}) {
	if !nRelic.enabled {
		return
	}
	nRelic.app.RecordCustomEvent(name, event)
}

// Transaction will store an instance of a newrelic transaction
type Transaction interface {
	RegisterNewSegment(name string, fun func())
	RegisterError(err error)
	End()
}

type newRelicTransaction struct {
	Txn     *newrelic.Transaction
	enabled bool
}

// StartTransaction starts a new visibility transaction
func (nRelic *newrelicMonitoring) StartTransaction(name string) Transaction {
	t := new(newRelicTransaction)
	t.enabled = true
	if !nRelic.enabled {
		t.enabled = false
		return t
	}
	txn := nRelic.app.StartTransaction(name)
	t.Txn = txn
	return t
}

// End stops the transaction
func (txn *newRelicTransaction) End() {
	if !txn.enabled {
		return
	}
	txn.Txn.End()
}

// RegisterNewSegment is the decorator method to register the latency of ApplicationFunction
func (txn *newRelicTransaction) RegisterNewSegment(name string, fun func()) {
	if !txn.enabled {
		fun()
		return
	}
	logrus.Debugf("Starting a new transaction for goroutine for segment: %s", name)
	newTxn := txn.Txn.NewGoroutine()
	defer newTxn.StartSegment(name).End()
	fun()
	// TODO: add functionality of returning values returned by fun
	// TODO: have means for this method to accept parameters
}

// RegisterError registers an error given a transaction
func (txn *newRelicTransaction) RegisterError(err error) {
	if !txn.enabled {
		return
	}
	txn.Txn.NoticeError(err)
}
