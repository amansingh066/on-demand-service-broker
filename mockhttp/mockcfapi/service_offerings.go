// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package mockcfapi

import (
	"fmt"

	"github.com/amansingh066/on-demand-service-broker/mockhttp"
)

type serviceOfferingsMock struct {
	*mockhttp.Handler
}

func ListServiceOfferings() *serviceOfferingsMock {
	return &serviceOfferingsMock{
		mockhttp.NewMockedHttpRequest("GET", "/v2/services?results-per-page=100"),
	}
}

func ListServiceOfferingsForPage(page int) *serviceOfferingsMock {
	return &serviceOfferingsMock{
		mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("/v2/services?order-direction=asc&page=%d&results-per-page=100", page)),
	}
}

func (m *serviceOfferingsMock) RespondsWithNoServiceOfferings() *mockhttp.Handler {
	return m.RespondsOKWith(`{"next_url": null, "resources": []}`)
}

func (m *serviceOfferingsMock) RespondsWithServiceOffering(serviceOfferingID, ccServiceOfferingGUID string) *mockhttp.Handler {
	return m.RespondsOKWith(listServiceOfferingResponse(serviceOfferingID, ccServiceOfferingGUID))
}

func listServiceOfferingResponse(serviceOfferingID, ccServiceOfferingGUID string) string {
	return `{
    "total_results": 1,
	  "total_pages": 1,
	  "prev_url": null,
	  "next_url": null,
	  "resources": [
	    {
	      "metadata": {
	   	    "guid": "` + ccServiceOfferingGUID + `"
	   	  },
        "entity": {
          "unique_id": "` + serviceOfferingID + `",
          "service_plans_url": "/v2/services/` + ccServiceOfferingGUID + `/service_plans"
        }
	    }
	  ]
	}`
}
