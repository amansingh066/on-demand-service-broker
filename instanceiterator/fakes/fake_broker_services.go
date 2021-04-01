// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/brokerapi/v8/domain"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-service-broker/broker/services"
	"github.com/pivotal-cf/on-demand-service-broker/instanceiterator"
	"github.com/pivotal-cf/on-demand-service-broker/service"
)

type FakeBrokerServices struct {
	InstancesStub        func(map[string]string) ([]service.Instance, error)
	instancesMutex       sync.RWMutex
	instancesArgsForCall []struct {
		arg1 map[string]string
	}
	instancesReturns struct {
		result1 []service.Instance
		result2 error
	}
	instancesReturnsOnCall map[int]struct {
		result1 []service.Instance
		result2 error
	}
	LastOperationStub        func(string, broker.OperationData) (domain.LastOperation, error)
	lastOperationMutex       sync.RWMutex
	lastOperationArgsForCall []struct {
		arg1 string
		arg2 broker.OperationData
	}
	lastOperationReturns struct {
		result1 domain.LastOperation
		result2 error
	}
	lastOperationReturnsOnCall map[int]struct {
		result1 domain.LastOperation
		result2 error
	}
	LatestInstanceInfoStub        func(service.Instance) (service.Instance, error)
	latestInstanceInfoMutex       sync.RWMutex
	latestInstanceInfoArgsForCall []struct {
		arg1 service.Instance
	}
	latestInstanceInfoReturns struct {
		result1 service.Instance
		result2 error
	}
	latestInstanceInfoReturnsOnCall map[int]struct {
		result1 service.Instance
		result2 error
	}
	ProcessInstanceStub        func(service.Instance, string) (services.BOSHOperation, error)
	processInstanceMutex       sync.RWMutex
	processInstanceArgsForCall []struct {
		arg1 service.Instance
		arg2 string
	}
	processInstanceReturns struct {
		result1 services.BOSHOperation
		result2 error
	}
	processInstanceReturnsOnCall map[int]struct {
		result1 services.BOSHOperation
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBrokerServices) Instances(arg1 map[string]string) ([]service.Instance, error) {
	fake.instancesMutex.Lock()
	ret, specificReturn := fake.instancesReturnsOnCall[len(fake.instancesArgsForCall)]
	fake.instancesArgsForCall = append(fake.instancesArgsForCall, struct {
		arg1 map[string]string
	}{arg1})
	fake.recordInvocation("Instances", []interface{}{arg1})
	fake.instancesMutex.Unlock()
	if fake.InstancesStub != nil {
		return fake.InstancesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.instancesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerServices) InstancesCallCount() int {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return len(fake.instancesArgsForCall)
}

func (fake *FakeBrokerServices) InstancesCalls(stub func(map[string]string) ([]service.Instance, error)) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = stub
}

func (fake *FakeBrokerServices) InstancesArgsForCall(i int) map[string]string {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	argsForCall := fake.instancesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBrokerServices) InstancesReturns(result1 []service.Instance, result2 error) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = nil
	fake.instancesReturns = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) InstancesReturnsOnCall(i int, result1 []service.Instance, result2 error) {
	fake.instancesMutex.Lock()
	defer fake.instancesMutex.Unlock()
	fake.InstancesStub = nil
	if fake.instancesReturnsOnCall == nil {
		fake.instancesReturnsOnCall = make(map[int]struct {
			result1 []service.Instance
			result2 error
		})
	}
	fake.instancesReturnsOnCall[i] = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) LastOperation(arg1 string, arg2 broker.OperationData) (domain.LastOperation, error) {
	fake.lastOperationMutex.Lock()
	ret, specificReturn := fake.lastOperationReturnsOnCall[len(fake.lastOperationArgsForCall)]
	fake.lastOperationArgsForCall = append(fake.lastOperationArgsForCall, struct {
		arg1 string
		arg2 broker.OperationData
	}{arg1, arg2})
	fake.recordInvocation("LastOperation", []interface{}{arg1, arg2})
	fake.lastOperationMutex.Unlock()
	if fake.LastOperationStub != nil {
		return fake.LastOperationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.lastOperationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerServices) LastOperationCallCount() int {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return len(fake.lastOperationArgsForCall)
}

func (fake *FakeBrokerServices) LastOperationCalls(stub func(string, broker.OperationData) (domain.LastOperation, error)) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = stub
}

func (fake *FakeBrokerServices) LastOperationArgsForCall(i int) (string, broker.OperationData) {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	argsForCall := fake.lastOperationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBrokerServices) LastOperationReturns(result1 domain.LastOperation, result2 error) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = nil
	fake.lastOperationReturns = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) LastOperationReturnsOnCall(i int, result1 domain.LastOperation, result2 error) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = nil
	if fake.lastOperationReturnsOnCall == nil {
		fake.lastOperationReturnsOnCall = make(map[int]struct {
			result1 domain.LastOperation
			result2 error
		})
	}
	fake.lastOperationReturnsOnCall[i] = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) LatestInstanceInfo(arg1 service.Instance) (service.Instance, error) {
	fake.latestInstanceInfoMutex.Lock()
	ret, specificReturn := fake.latestInstanceInfoReturnsOnCall[len(fake.latestInstanceInfoArgsForCall)]
	fake.latestInstanceInfoArgsForCall = append(fake.latestInstanceInfoArgsForCall, struct {
		arg1 service.Instance
	}{arg1})
	fake.recordInvocation("LatestInstanceInfo", []interface{}{arg1})
	fake.latestInstanceInfoMutex.Unlock()
	if fake.LatestInstanceInfoStub != nil {
		return fake.LatestInstanceInfoStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.latestInstanceInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerServices) LatestInstanceInfoCallCount() int {
	fake.latestInstanceInfoMutex.RLock()
	defer fake.latestInstanceInfoMutex.RUnlock()
	return len(fake.latestInstanceInfoArgsForCall)
}

func (fake *FakeBrokerServices) LatestInstanceInfoCalls(stub func(service.Instance) (service.Instance, error)) {
	fake.latestInstanceInfoMutex.Lock()
	defer fake.latestInstanceInfoMutex.Unlock()
	fake.LatestInstanceInfoStub = stub
}

func (fake *FakeBrokerServices) LatestInstanceInfoArgsForCall(i int) service.Instance {
	fake.latestInstanceInfoMutex.RLock()
	defer fake.latestInstanceInfoMutex.RUnlock()
	argsForCall := fake.latestInstanceInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBrokerServices) LatestInstanceInfoReturns(result1 service.Instance, result2 error) {
	fake.latestInstanceInfoMutex.Lock()
	defer fake.latestInstanceInfoMutex.Unlock()
	fake.LatestInstanceInfoStub = nil
	fake.latestInstanceInfoReturns = struct {
		result1 service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) LatestInstanceInfoReturnsOnCall(i int, result1 service.Instance, result2 error) {
	fake.latestInstanceInfoMutex.Lock()
	defer fake.latestInstanceInfoMutex.Unlock()
	fake.LatestInstanceInfoStub = nil
	if fake.latestInstanceInfoReturnsOnCall == nil {
		fake.latestInstanceInfoReturnsOnCall = make(map[int]struct {
			result1 service.Instance
			result2 error
		})
	}
	fake.latestInstanceInfoReturnsOnCall[i] = struct {
		result1 service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) ProcessInstance(arg1 service.Instance, arg2 string) (services.BOSHOperation, error) {
	fake.processInstanceMutex.Lock()
	ret, specificReturn := fake.processInstanceReturnsOnCall[len(fake.processInstanceArgsForCall)]
	fake.processInstanceArgsForCall = append(fake.processInstanceArgsForCall, struct {
		arg1 service.Instance
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ProcessInstance", []interface{}{arg1, arg2})
	fake.processInstanceMutex.Unlock()
	if fake.ProcessInstanceStub != nil {
		return fake.ProcessInstanceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.processInstanceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerServices) ProcessInstanceCallCount() int {
	fake.processInstanceMutex.RLock()
	defer fake.processInstanceMutex.RUnlock()
	return len(fake.processInstanceArgsForCall)
}

func (fake *FakeBrokerServices) ProcessInstanceCalls(stub func(service.Instance, string) (services.BOSHOperation, error)) {
	fake.processInstanceMutex.Lock()
	defer fake.processInstanceMutex.Unlock()
	fake.ProcessInstanceStub = stub
}

func (fake *FakeBrokerServices) ProcessInstanceArgsForCall(i int) (service.Instance, string) {
	fake.processInstanceMutex.RLock()
	defer fake.processInstanceMutex.RUnlock()
	argsForCall := fake.processInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBrokerServices) ProcessInstanceReturns(result1 services.BOSHOperation, result2 error) {
	fake.processInstanceMutex.Lock()
	defer fake.processInstanceMutex.Unlock()
	fake.ProcessInstanceStub = nil
	fake.processInstanceReturns = struct {
		result1 services.BOSHOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) ProcessInstanceReturnsOnCall(i int, result1 services.BOSHOperation, result2 error) {
	fake.processInstanceMutex.Lock()
	defer fake.processInstanceMutex.Unlock()
	fake.ProcessInstanceStub = nil
	if fake.processInstanceReturnsOnCall == nil {
		fake.processInstanceReturnsOnCall = make(map[int]struct {
			result1 services.BOSHOperation
			result2 error
		})
	}
	fake.processInstanceReturnsOnCall[i] = struct {
		result1 services.BOSHOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	fake.latestInstanceInfoMutex.RLock()
	defer fake.latestInstanceInfoMutex.RUnlock()
	fake.processInstanceMutex.RLock()
	defer fake.processInstanceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBrokerServices) recordInvocation(key string, args []interface{}) {
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

var _ instanceiterator.BrokerServices = new(FakeBrokerServices)
