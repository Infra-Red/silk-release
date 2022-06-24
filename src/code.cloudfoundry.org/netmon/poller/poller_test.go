package poller_test

import (
	"os"
	"time"

	"code.cloudfoundry.org/lager/lagertest"
	"code.cloudfoundry.org/netmon/fakes"
	"code.cloudfoundry.org/netmon/poller"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Poller Run", func() {
	var (
		networkStatsFetcher *fakes.NetworkStatsFetcher
		logger              *lagertest.TestLogger

		metrics      *poller.SystemMetrics
		pollInterval time.Duration
	)

	BeforeEach(func() {
		networkStatsFetcher = &fakes.NetworkStatsFetcher{}
		logger = lagertest.NewTestLogger("test")
		pollInterval = 1 * time.Second

		networkStatsFetcher.CountIPTablesRulesReturnsOnCall(0, 4, nil)

		metrics = &poller.SystemMetrics{
			Logger:              logger,
			PollInterval:        pollInterval,
			InterfaceName:       "meow",
			NetworkStatsFetcher: networkStatsFetcher,
		}
	})

	It("should report measurements once within single interval", func() {
		runTest(metrics, pollInterval)
		Expect(logger.LogMessages()).To(Equal([]string{
			"test.measure.measure-start",
			"test.measure.metric-sent",
			"test.measure.metric-sent",
			"test.measure.read-tx-bytes",
			"test.measure.measure-complete",
		}))
	})

	It("should use the network stats fetcher  when checking the rules", func() {
		runTest(metrics, pollInterval)

		iptablesLog := logger.Logs()[2]
		Expect(iptablesLog.Data["IPTablesRuleCount"]).To(Equal(float64(4)))
	})

	Context("when a telemetry logger is configured", func() {
		var (
			telemetryLogger *lagertest.TestLogger
		)

		BeforeEach(func() {
			telemetryLogger = lagertest.NewTestLogger("telemetry")
			metrics = &poller.SystemMetrics{
				Logger:              logger,
				TelemetryLogger:     telemetryLogger,
				PollInterval:        pollInterval,
				InterfaceName:       "meow",
				NetworkStatsFetcher: networkStatsFetcher,
			}
		})

		It("logs the number of iptables rules in the telemetry log", func() {
			runTest(metrics, pollInterval)
			Expect(telemetryLogger.LogMessages()).To(Equal([]string{"telemetry.count-iptables-rules"}))
		})
	})
})

func runTest(metrics *poller.SystemMetrics, pollInterval time.Duration) {
	doneCh := make(chan os.Signal)
	readyCh := make(chan struct{})

	go metrics.Run(doneCh, readyCh)

	<-readyCh
	<-time.After(pollInterval + 10*time.Millisecond)
	doneCh <- os.Interrupt
}
