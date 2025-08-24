package consumer

import "github.com/prometheus/client_golang/prometheus"

// # HELP save_event_error_count 事件入库失败个数统计
// # TYPE save_event_error_count gauge
// save_event_error_count{service="devcloud"} 0
func NewEventCollector() *EventCollector {
	return &EventCollector{
		errCountDesc: prometheus.NewDesc(
			"save_event_error_count",
			"事件入库失败个数统计",
			[]string{},
			prometheus.Labels{"service": "devcloud"},
		),
	}
}

// 收集事件指标的采集器
type EventCollector struct {
	errCountDesc *prometheus.Desc
	// 需要自己根据实践情况来维护这个变量
	errCount int
}

func (c *EventCollector) Inc() {
	c.errCount++
}

// 指标元数据注册
func (c *EventCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.errCountDesc
}

// 指标的值的采集
func (c *EventCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.errCountDesc, prometheus.GaugeValue, float64(c.errCount))
}
