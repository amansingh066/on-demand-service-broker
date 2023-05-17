// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"log"
	"sync"

	"github.com/amansingh066/on-demand-service-broker/registrar"
)

type FakeCloudFoundryClient struct {
	DeregisterBrokerStub        func(string, *log.Logger) error
	deregisterBrokerMutex       sync.RWMutex
	deregisterBrokerArgsForCall []struct {
		arg1 string
		arg2 *log.Logger
	}
	deregisterBrokerReturns struct {
		result1 error
	}
	deregisterBrokerReturnsOnCall map[int]struct {
		result1 error
	}
	GetServiceOfferingGUIDStub        func(string, *log.Logger) (string, error)
	getServiceOfferingGUIDMutex       sync.RWMutex
	getServiceOfferingGUIDArgsForCall []struct {
		arg1 string
		arg2 *log.Logger
	}
	getServiceOfferingGUIDReturns struct {
		result1 string
		result2 error
	}
	getServiceOfferingGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCloudFoundryClient) DeregisterBroker(arg1 string, arg2 *log.Logger) error {
	fake.deregisterBrokerMutex.Lock()
	ret, specificReturn := fake.deregisterBrokerReturnsOnCall[len(fake.deregisterBrokerArgsForCall)]
	fake.deregisterBrokerArgsForCall = append(fake.deregisterBrokerArgsForCall, struct {
		arg1 string
		arg2 *log.Logger
	}{arg1, arg2})
	stub := fake.DeregisterBrokerStub
	fakeReturns := fake.deregisterBrokerReturns
	fake.recordInvocation("DeregisterBroker", []interface{}{arg1, arg2})
	fake.deregisterBrokerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCloudFoundryClient) DeregisterBrokerCallCount() int {
	fake.deregisterBrokerMutex.RLock()
	defer fake.deregisterBrokerMutex.RUnlock()
	return len(fake.deregisterBrokerArgsForCall)
}

func (fake *FakeCloudFoundryClient) DeregisterBrokerCalls(stub func(string, *log.Logger) error) {
	fake.deregisterBrokerMutex.Lock()
	defer fake.deregisterBrokerMutex.Unlock()
	fake.DeregisterBrokerStub = stub
}

func (fake *FakeCloudFoundryClient) DeregisterBrokerArgsForCall(i int) (string, *log.Logger) {
	fake.deregisterBrokerMutex.RLock()
	defer fake.deregisterBrokerMutex.RUnlock()
	argsForCall := fake.deregisterBrokerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCloudFoundryClient) DeregisterBrokerReturns(result1 error) {
	fake.deregisterBrokerMutex.Lock()
	defer fake.deregisterBrokerMutex.Unlock()
	fake.DeregisterBrokerStub = nil
	fake.deregisterBrokerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCloudFoundryClient) DeregisterBrokerReturnsOnCall(i int, result1 error) {
	fake.deregisterBrokerMutex.Lock()
	defer fake.deregisterBrokerMutex.Unlock()
	fake.DeregisterBrokerStub = nil
	if fake.deregisterBrokerReturnsOnCall == nil {
		fake.deregisterBrokerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deregisterBrokerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCloudFoundryClient) GetServiceOfferingGUID(arg1 string, arg2 *log.Logger) (string, error) {
	fake.getServiceOfferingGUIDMutex.Lock()
	ret, specificReturn := fake.getServiceOfferingGUIDReturnsOnCall[len(fake.getServiceOfferingGUIDArgsForCall)]
	fake.getServiceOfferingGUIDArgsForCall = append(fake.getServiceOfferingGUIDArgsForCall, struct {
		arg1 string
		arg2 *log.Logger
	}{arg1, arg2})
	stub := fake.GetServiceOfferingGUIDStub
	fakeReturns := fake.getServiceOfferingGUIDReturns
	fake.recordInvocation("GetServiceOfferingGUID", []interface{}{arg1, arg2})
	fake.getServiceOfferingGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCloudFoundryClient) GetServiceOfferingGUIDCallCount() int {
	fake.getServiceOfferingGUIDMutex.RLock()
	defer fake.getServiceOfferingGUIDMutex.RUnlock()
	return len(fake.getServiceOfferingGUIDArgsForCall)
}

func (fake *FakeCloudFoundryClient) GetServiceOfferingGUIDCalls(stub func(string, *log.Logger) (string, error)) {
	fake.getServiceOfferingGUIDMutex.Lock()
	defer fake.getServiceOfferingGUIDMutex.Unlock()
	fake.GetServiceOfferingGUIDStub = stub
}

func (fake *FakeCloudFoundryClient) GetServiceOfferingGUIDArgsForCall(i int) (string, *log.Logger) {
	fake.getServiceOfferingGUIDMutex.RLock()
	defer fake.getServiceOfferingGUIDMutex.RUnlock()
	argsForCall := fake.getServiceOfferingGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCloudFoundryClient) GetServiceOfferingGUIDReturns(result1 string, result2 error) {
	fake.getServiceOfferingGUIDMutex.Lock()
	defer fake.getServiceOfferingGUIDMutex.Unlock()
	fake.GetServiceOfferingGUIDStub = nil
	fake.getServiceOfferingGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) GetServiceOfferingGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getServiceOfferingGUIDMutex.Lock()
	defer fake.getServiceOfferingGUIDMutex.Unlock()
	fake.GetServiceOfferingGUIDStub = nil
	if fake.getServiceOfferingGUIDReturnsOnCall == nil {
		fake.getServiceOfferingGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getServiceOfferingGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCloudFoundryClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deregisterBrokerMutex.RLock()
	defer fake.deregisterBrokerMutex.RUnlock()
	fake.getServiceOfferingGUIDMutex.RLock()
	defer fake.getServiceOfferingGUIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCloudFoundryClient) recordInvocation(key string, args []interface{}) {
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

var _ registrar.CloudFoundryClient = new(FakeCloudFoundryClient)
