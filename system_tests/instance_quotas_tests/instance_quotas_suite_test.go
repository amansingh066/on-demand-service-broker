// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package instance_quotas_tests

import (
	"testing"

	bosh "github.com/amansingh066/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"
	"github.com/amansingh066/on-demand-service-broker/system_tests/test_helpers/service_helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"
)

func TestQuotasTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Instance Quotas Suite")
}

var brokerInfo bosh.BrokerInfo

var _ = BeforeSuite(func() {
	uniqueID := uuid.New()[:6]
	brokerInfo = bosh.DeployAndRegisterBroker(
		"-instance-quotas-"+uniqueID,
		bosh.BrokerDeploymentOptions{BrokerTLS: true},
		service_helpers.Redis,
		[]string{"update_service_catalog.yml"},
	)
})

var _ = AfterSuite(func() {
	bosh.DeregisterAndDeleteBroker(brokerInfo.DeploymentName)
})
