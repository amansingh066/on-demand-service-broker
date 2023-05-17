// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"log"
	"sync"

	"github.com/amansingh066/on-demand-service-broker/cf"
	"github.com/amansingh066/on-demand-service-broker/registrar"
)

type FakeRegisterBrokerCFClient struct {
	CreateServiceBrokerStub        func(string, string, string, string) error
	createServiceBrokerMutex       sync.RWMutex
	createServiceBrokerArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	createServiceBrokerReturns struct {
		result1 error
	}
	createServiceBrokerReturnsOnCall map[int]struct {
		result1 error
	}
	CreateServicePlanVisibilityStub        func(string, string, string, *log.Logger) error
	createServicePlanVisibilityMutex       sync.RWMutex
	createServicePlanVisibilityArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 *log.Logger
	}
	createServicePlanVisibilityReturns struct {
		result1 error
	}
	createServicePlanVisibilityReturnsOnCall map[int]struct {
		result1 error
	}
	DisableServiceAccessStub        func(string, string, *log.Logger) error
	disableServiceAccessMutex       sync.RWMutex
	disableServiceAccessArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}
	disableServiceAccessReturns struct {
		result1 error
	}
	disableServiceAccessReturnsOnCall map[int]struct {
		result1 error
	}
	EnableServiceAccessStub        func(string, string, *log.Logger) error
	enableServiceAccessMutex       sync.RWMutex
	enableServiceAccessArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}
	enableServiceAccessReturns struct {
		result1 error
	}
	enableServiceAccessReturnsOnCall map[int]struct {
		result1 error
	}
	ServiceBrokersStub        func() ([]cf.ServiceBroker, error)
	serviceBrokersMutex       sync.RWMutex
	serviceBrokersArgsForCall []struct {
	}
	serviceBrokersReturns struct {
		result1 []cf.ServiceBroker
		result2 error
	}
	serviceBrokersReturnsOnCall map[int]struct {
		result1 []cf.ServiceBroker
		result2 error
	}
	UpdateServiceBrokerStub        func(string, string, string, string, string) error
	updateServiceBrokerMutex       sync.RWMutex
	updateServiceBrokerArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	updateServiceBrokerReturns struct {
		result1 error
	}
	updateServiceBrokerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegisterBrokerCFClient) CreateServiceBroker(arg1 string, arg2 string, arg3 string, arg4 string) error {
	fake.createServiceBrokerMutex.Lock()
	ret, specificReturn := fake.createServiceBrokerReturnsOnCall[len(fake.createServiceBrokerArgsForCall)]
	fake.createServiceBrokerArgsForCall = append(fake.createServiceBrokerArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.CreateServiceBrokerStub
	fakeReturns := fake.createServiceBrokerReturns
	fake.recordInvocation("CreateServiceBroker", []interface{}{arg1, arg2, arg3, arg4})
	fake.createServiceBrokerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRegisterBrokerCFClient) CreateServiceBrokerCallCount() int {
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	return len(fake.createServiceBrokerArgsForCall)
}

func (fake *FakeRegisterBrokerCFClient) CreateServiceBrokerCalls(stub func(string, string, string, string) error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = stub
}

func (fake *FakeRegisterBrokerCFClient) CreateServiceBrokerArgsForCall(i int) (string, string, string, string) {
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	argsForCall := fake.createServiceBrokerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeRegisterBrokerCFClient) CreateServiceBrokerReturns(result1 error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = nil
	fake.createServiceBrokerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) CreateServiceBrokerReturnsOnCall(i int, result1 error) {
	fake.createServiceBrokerMutex.Lock()
	defer fake.createServiceBrokerMutex.Unlock()
	fake.CreateServiceBrokerStub = nil
	if fake.createServiceBrokerReturnsOnCall == nil {
		fake.createServiceBrokerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createServiceBrokerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) CreateServicePlanVisibility(arg1 string, arg2 string, arg3 string, arg4 *log.Logger) error {
	fake.createServicePlanVisibilityMutex.Lock()
	ret, specificReturn := fake.createServicePlanVisibilityReturnsOnCall[len(fake.createServicePlanVisibilityArgsForCall)]
	fake.createServicePlanVisibilityArgsForCall = append(fake.createServicePlanVisibilityArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 *log.Logger
	}{arg1, arg2, arg3, arg4})
	stub := fake.CreateServicePlanVisibilityStub
	fakeReturns := fake.createServicePlanVisibilityReturns
	fake.recordInvocation("CreateServicePlanVisibility", []interface{}{arg1, arg2, arg3, arg4})
	fake.createServicePlanVisibilityMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRegisterBrokerCFClient) CreateServicePlanVisibilityCallCount() int {
	fake.createServicePlanVisibilityMutex.RLock()
	defer fake.createServicePlanVisibilityMutex.RUnlock()
	return len(fake.createServicePlanVisibilityArgsForCall)
}

func (fake *FakeRegisterBrokerCFClient) CreateServicePlanVisibilityCalls(stub func(string, string, string, *log.Logger) error) {
	fake.createServicePlanVisibilityMutex.Lock()
	defer fake.createServicePlanVisibilityMutex.Unlock()
	fake.CreateServicePlanVisibilityStub = stub
}

func (fake *FakeRegisterBrokerCFClient) CreateServicePlanVisibilityArgsForCall(i int) (string, string, string, *log.Logger) {
	fake.createServicePlanVisibilityMutex.RLock()
	defer fake.createServicePlanVisibilityMutex.RUnlock()
	argsForCall := fake.createServicePlanVisibilityArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeRegisterBrokerCFClient) CreateServicePlanVisibilityReturns(result1 error) {
	fake.createServicePlanVisibilityMutex.Lock()
	defer fake.createServicePlanVisibilityMutex.Unlock()
	fake.CreateServicePlanVisibilityStub = nil
	fake.createServicePlanVisibilityReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) CreateServicePlanVisibilityReturnsOnCall(i int, result1 error) {
	fake.createServicePlanVisibilityMutex.Lock()
	defer fake.createServicePlanVisibilityMutex.Unlock()
	fake.CreateServicePlanVisibilityStub = nil
	if fake.createServicePlanVisibilityReturnsOnCall == nil {
		fake.createServicePlanVisibilityReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createServicePlanVisibilityReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) DisableServiceAccess(arg1 string, arg2 string, arg3 *log.Logger) error {
	fake.disableServiceAccessMutex.Lock()
	ret, specificReturn := fake.disableServiceAccessReturnsOnCall[len(fake.disableServiceAccessArgsForCall)]
	fake.disableServiceAccessArgsForCall = append(fake.disableServiceAccessArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}{arg1, arg2, arg3})
	stub := fake.DisableServiceAccessStub
	fakeReturns := fake.disableServiceAccessReturns
	fake.recordInvocation("DisableServiceAccess", []interface{}{arg1, arg2, arg3})
	fake.disableServiceAccessMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRegisterBrokerCFClient) DisableServiceAccessCallCount() int {
	fake.disableServiceAccessMutex.RLock()
	defer fake.disableServiceAccessMutex.RUnlock()
	return len(fake.disableServiceAccessArgsForCall)
}

func (fake *FakeRegisterBrokerCFClient) DisableServiceAccessCalls(stub func(string, string, *log.Logger) error) {
	fake.disableServiceAccessMutex.Lock()
	defer fake.disableServiceAccessMutex.Unlock()
	fake.DisableServiceAccessStub = stub
}

func (fake *FakeRegisterBrokerCFClient) DisableServiceAccessArgsForCall(i int) (string, string, *log.Logger) {
	fake.disableServiceAccessMutex.RLock()
	defer fake.disableServiceAccessMutex.RUnlock()
	argsForCall := fake.disableServiceAccessArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRegisterBrokerCFClient) DisableServiceAccessReturns(result1 error) {
	fake.disableServiceAccessMutex.Lock()
	defer fake.disableServiceAccessMutex.Unlock()
	fake.DisableServiceAccessStub = nil
	fake.disableServiceAccessReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) DisableServiceAccessReturnsOnCall(i int, result1 error) {
	fake.disableServiceAccessMutex.Lock()
	defer fake.disableServiceAccessMutex.Unlock()
	fake.DisableServiceAccessStub = nil
	if fake.disableServiceAccessReturnsOnCall == nil {
		fake.disableServiceAccessReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.disableServiceAccessReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) EnableServiceAccess(arg1 string, arg2 string, arg3 *log.Logger) error {
	fake.enableServiceAccessMutex.Lock()
	ret, specificReturn := fake.enableServiceAccessReturnsOnCall[len(fake.enableServiceAccessArgsForCall)]
	fake.enableServiceAccessArgsForCall = append(fake.enableServiceAccessArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}{arg1, arg2, arg3})
	stub := fake.EnableServiceAccessStub
	fakeReturns := fake.enableServiceAccessReturns
	fake.recordInvocation("EnableServiceAccess", []interface{}{arg1, arg2, arg3})
	fake.enableServiceAccessMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRegisterBrokerCFClient) EnableServiceAccessCallCount() int {
	fake.enableServiceAccessMutex.RLock()
	defer fake.enableServiceAccessMutex.RUnlock()
	return len(fake.enableServiceAccessArgsForCall)
}

func (fake *FakeRegisterBrokerCFClient) EnableServiceAccessCalls(stub func(string, string, *log.Logger) error) {
	fake.enableServiceAccessMutex.Lock()
	defer fake.enableServiceAccessMutex.Unlock()
	fake.EnableServiceAccessStub = stub
}

func (fake *FakeRegisterBrokerCFClient) EnableServiceAccessArgsForCall(i int) (string, string, *log.Logger) {
	fake.enableServiceAccessMutex.RLock()
	defer fake.enableServiceAccessMutex.RUnlock()
	argsForCall := fake.enableServiceAccessArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRegisterBrokerCFClient) EnableServiceAccessReturns(result1 error) {
	fake.enableServiceAccessMutex.Lock()
	defer fake.enableServiceAccessMutex.Unlock()
	fake.EnableServiceAccessStub = nil
	fake.enableServiceAccessReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) EnableServiceAccessReturnsOnCall(i int, result1 error) {
	fake.enableServiceAccessMutex.Lock()
	defer fake.enableServiceAccessMutex.Unlock()
	fake.EnableServiceAccessStub = nil
	if fake.enableServiceAccessReturnsOnCall == nil {
		fake.enableServiceAccessReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.enableServiceAccessReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) ServiceBrokers() ([]cf.ServiceBroker, error) {
	fake.serviceBrokersMutex.Lock()
	ret, specificReturn := fake.serviceBrokersReturnsOnCall[len(fake.serviceBrokersArgsForCall)]
	fake.serviceBrokersArgsForCall = append(fake.serviceBrokersArgsForCall, struct {
	}{})
	stub := fake.ServiceBrokersStub
	fakeReturns := fake.serviceBrokersReturns
	fake.recordInvocation("ServiceBrokers", []interface{}{})
	fake.serviceBrokersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegisterBrokerCFClient) ServiceBrokersCallCount() int {
	fake.serviceBrokersMutex.RLock()
	defer fake.serviceBrokersMutex.RUnlock()
	return len(fake.serviceBrokersArgsForCall)
}

func (fake *FakeRegisterBrokerCFClient) ServiceBrokersCalls(stub func() ([]cf.ServiceBroker, error)) {
	fake.serviceBrokersMutex.Lock()
	defer fake.serviceBrokersMutex.Unlock()
	fake.ServiceBrokersStub = stub
}

func (fake *FakeRegisterBrokerCFClient) ServiceBrokersReturns(result1 []cf.ServiceBroker, result2 error) {
	fake.serviceBrokersMutex.Lock()
	defer fake.serviceBrokersMutex.Unlock()
	fake.ServiceBrokersStub = nil
	fake.serviceBrokersReturns = struct {
		result1 []cf.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeRegisterBrokerCFClient) ServiceBrokersReturnsOnCall(i int, result1 []cf.ServiceBroker, result2 error) {
	fake.serviceBrokersMutex.Lock()
	defer fake.serviceBrokersMutex.Unlock()
	fake.ServiceBrokersStub = nil
	if fake.serviceBrokersReturnsOnCall == nil {
		fake.serviceBrokersReturnsOnCall = make(map[int]struct {
			result1 []cf.ServiceBroker
			result2 error
		})
	}
	fake.serviceBrokersReturnsOnCall[i] = struct {
		result1 []cf.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeRegisterBrokerCFClient) UpdateServiceBroker(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) error {
	fake.updateServiceBrokerMutex.Lock()
	ret, specificReturn := fake.updateServiceBrokerReturnsOnCall[len(fake.updateServiceBrokerArgsForCall)]
	fake.updateServiceBrokerArgsForCall = append(fake.updateServiceBrokerArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.UpdateServiceBrokerStub
	fakeReturns := fake.updateServiceBrokerReturns
	fake.recordInvocation("UpdateServiceBroker", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.updateServiceBrokerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRegisterBrokerCFClient) UpdateServiceBrokerCallCount() int {
	fake.updateServiceBrokerMutex.RLock()
	defer fake.updateServiceBrokerMutex.RUnlock()
	return len(fake.updateServiceBrokerArgsForCall)
}

func (fake *FakeRegisterBrokerCFClient) UpdateServiceBrokerCalls(stub func(string, string, string, string, string) error) {
	fake.updateServiceBrokerMutex.Lock()
	defer fake.updateServiceBrokerMutex.Unlock()
	fake.UpdateServiceBrokerStub = stub
}

func (fake *FakeRegisterBrokerCFClient) UpdateServiceBrokerArgsForCall(i int) (string, string, string, string, string) {
	fake.updateServiceBrokerMutex.RLock()
	defer fake.updateServiceBrokerMutex.RUnlock()
	argsForCall := fake.updateServiceBrokerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeRegisterBrokerCFClient) UpdateServiceBrokerReturns(result1 error) {
	fake.updateServiceBrokerMutex.Lock()
	defer fake.updateServiceBrokerMutex.Unlock()
	fake.UpdateServiceBrokerStub = nil
	fake.updateServiceBrokerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) UpdateServiceBrokerReturnsOnCall(i int, result1 error) {
	fake.updateServiceBrokerMutex.Lock()
	defer fake.updateServiceBrokerMutex.Unlock()
	fake.UpdateServiceBrokerStub = nil
	if fake.updateServiceBrokerReturnsOnCall == nil {
		fake.updateServiceBrokerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateServiceBrokerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegisterBrokerCFClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createServiceBrokerMutex.RLock()
	defer fake.createServiceBrokerMutex.RUnlock()
	fake.createServicePlanVisibilityMutex.RLock()
	defer fake.createServicePlanVisibilityMutex.RUnlock()
	fake.disableServiceAccessMutex.RLock()
	defer fake.disableServiceAccessMutex.RUnlock()
	fake.enableServiceAccessMutex.RLock()
	defer fake.enableServiceAccessMutex.RUnlock()
	fake.serviceBrokersMutex.RLock()
	defer fake.serviceBrokersMutex.RUnlock()
	fake.updateServiceBrokerMutex.RLock()
	defer fake.updateServiceBrokerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegisterBrokerCFClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ registrar.RegisterBrokerCFClient = new(FakeRegisterBrokerCFClient)
