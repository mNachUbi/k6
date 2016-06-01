package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/loadimpact/speedboat/sampler"
	stdlog "log"
	"time"
)

func printMetrics(l *stdlog.Logger) {
	for name, m := range sampler.DefaultSampler.Metrics {
		l.Printf("%s\n", name)
		switch m.Type {
		case sampler.GaugeType:
			l.Printf("  value=%s\n", applyIntent(m, m.Last()))
		case sampler.CounterType:
			l.Printf("  num=%s\n", applyIntent(m, m.Sum()))
		case sampler.StatsType:
			l.Printf("  min=%s\n", applyIntent(m, m.Min()))
			l.Printf("  max=%s\n", applyIntent(m, m.Max()))
			l.Printf("  avg=%s\n", applyIntent(m, m.Avg()))
			l.Printf("  med=%s\n", applyIntent(m, m.Med()))
		}
	}
}

func commitMetrics() {
	if err := sampler.DefaultSampler.Commit(); err != nil {
		log.WithError(err).Error("Couldn't write samples!")
	}
}

func applyIntent(m *sampler.Metric, v int64) interface{} {
	if m.Intent == sampler.TimeIntent {
		return time.Duration(v)
	}
	return fmt.Sprint(v)
}