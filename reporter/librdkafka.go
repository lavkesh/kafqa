package reporter

type LibrdKafkaStatsHandler struct {
	counters map[string][]string
	gauges   map[string][]string
}

func GetStats(stats []string, statsJSON, level, statType, topic string) {
	switch level {
	case "top-level":
		TopLevelStats(stats, statsJSON, statType, topic)
	case "brokers":
		BrokersStats(stats, statsJSON, statType, topic)
	}
}

func (stats LibrdKafkaStatsHandler) HandleStats(statJSON, topic string) {
	for k, counterStat := range stats.counters {
		GetStats(counterStat, statJSON, k, "counter", topic)
	}

	for k, gaugeStat := range stats.gauges {
		GetStats(gaugeStat, statJSON, k, "gauge", topic)
	}
}

func NewlibrdKafkaStat() LibrdKafkaStatsHandler {
	return LibrdKafkaStatsHandler{defaultCounters(), defaultGauges()}
}

func defaultCounters() map[string][]string {
	return map[string][]string{
		"top-level": {"tx", "rx", "txmsgs", "rxmsgs"},
		"brokers":   {"tx", "rx"},
	}
}

func defaultGauges() map[string][]string {
	return map[string][]string{
		"top-level": {"msg_cnt", "msg_size"},
		"brokers": {"outbuf_msg_cnt", "int_latency.p99", "int_latency.avg",
			"outbuf_latency.p99", "outbuf_latency.avg", "throttle.avg", "throttle.p99",
			"rtt.avg", "rtt.p99"},
	}
}
