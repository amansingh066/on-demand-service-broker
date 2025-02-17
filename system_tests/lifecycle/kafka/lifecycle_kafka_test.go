package lifecycle_test

import (
	. "github.com/amansingh066/on-demand-service-broker/system_tests/lifecycle/all_lifecycle_tests"
	"github.com/amansingh066/on-demand-service-broker/system_tests/test_helpers/service_helpers"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Kafka Lifecycle Tests", func() {
	It("can perform the basic service lifecycle", func() {
		BasicLifecycleTest(
			service_helpers.Kafka,
			brokerInfo,
			"kafka-small",
			"kafka-medium",
			`{"auto_create_topics": true}`)
	})
})
