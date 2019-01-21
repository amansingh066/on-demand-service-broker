package feature_flags_test

import (
	bosh "github.com/pivotal-cf/on-demand-service-broker/system_tests/bosh_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/pborman/uuid"
	cf "github.com/pivotal-cf/on-demand-service-broker/system_tests/cf_helpers"
)

var _ = Describe("FeatureFlags", func() {
	var (
		brokerInfo bosh.BrokerInfo
	)

	When("disable_ssl_cert_verification is true", func() {
		var brokerRegistered bool

		BeforeEach(func() {
			uniqueID := uuid.New()[:6]
			brokerInfo = bosh.DeployAndRegisterBroker(
				"-feature-flag-"+uniqueID,
				"update_service_catalog.yml", "disable_cf_ssl_verification.yml")
		})

		It("can run all the errands successfully", func() {
			By("running the register-broker", func() {
				bosh.RunErrand(brokerInfo.DeploymentName, "register-broker")
				brokerRegistered = true
			})

			By("running upgrade-all-service-instances", func() {
				bosh.RunErrand(brokerInfo.DeploymentName, "upgrade-all-service-instances")
			})

			By("running the orphan-deployments", func() {
				session := bosh.RunErrand(
					brokerInfo.DeploymentName,
					"orphan-deployments",
					Or(gexec.Exit(0), gexec.Exit(1)),
				)
				if session.ExitCode() == 1 {
					Expect(session.Buffer()).To(gbytes.Say("Orphan BOSH deployments detected"))
				}
			})

			By("running delete-all-service-instances", func() {
				bosh.RunErrand(brokerInfo.DeploymentName, "delete-all-service-instances")
			})

			By("running deregister-broker", func() {
				bosh.RunErrand(brokerInfo.DeploymentName, "deregister-broker")
				brokerRegistered = false
			})
		})

		AfterEach(func() {
			if brokerRegistered {
				bosh.DeregisterAndDeleteBroker(brokerInfo.DeploymentName)
			} else {
				bosh.DeleteDeployment(brokerInfo.DeploymentName)
			}
		})
	})

	When("expose_operational_errors is true", func() {
		BeforeEach(func() {
			uniqueID := uuid.New()[:6]
			brokerInfo = bosh.DeployAndRegisterBroker(
				"-feature-flag-"+uniqueID,
				"update_service_catalog.yml", "expose_operational_errors.yml")
		})

		It("correctly exposes operational errors", func() {
			bosh.RunErrand(brokerInfo.DeploymentName, "register-broker")
			serviceName := uuid.New()[:8]

			createServiceSession := cf.Cf("create-service", brokerInfo.ServiceOffering, "invalid-vm-type", serviceName)
			Eventually(createServiceSession, cf.CfTimeout).Should(gexec.Exit(0))

			cf.AwaitServiceCreationFailure(serviceName)

			s := cf.Cf("service", serviceName)
			Eventually(s.Out).Should(gbytes.Say(`Instance group 'redis-server' references an unknown vm type`))

			cf.DeleteService(serviceName)
			cf.AwaitServiceDeletion(serviceName)
		})

		AfterEach(func() {
			bosh.DeregisterAndDeleteBroker(brokerInfo.DeploymentName)
		})
	})
})
