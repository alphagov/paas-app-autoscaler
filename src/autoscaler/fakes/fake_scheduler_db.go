// This file was generated by counterfeiter
package fakes

import (
	"autoscaler/db"
	"autoscaler/models"
	"database/sql"
	"sync"
)

type FakeSchedulerDB struct {
	GetDBStatusStub        func() sql.DBStats
	getDBStatusMutex       sync.RWMutex
	getDBStatusArgsForCall []struct{}
	getDBStatusReturns     struct {
		result1 sql.DBStats
	}
	GetActiveSchedulesStub        func() (map[string]*models.ActiveSchedule, error)
	getActiveSchedulesMutex       sync.RWMutex
	getActiveSchedulesArgsForCall []struct{}
	getActiveSchedulesReturns     struct {
		result1 map[string]*models.ActiveSchedule
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSchedulerDB) GetDBStatus() sql.DBStats {
	fake.getDBStatusMutex.Lock()
	fake.getDBStatusArgsForCall = append(fake.getDBStatusArgsForCall, struct{}{})
	fake.recordInvocation("GetDBStatus", []interface{}{})
	fake.getDBStatusMutex.Unlock()
	if fake.GetDBStatusStub != nil {
		return fake.GetDBStatusStub()
	}
	return fake.getDBStatusReturns.result1
}

func (fake *FakeSchedulerDB) GetDBStatusCallCount() int {
	fake.getDBStatusMutex.RLock()
	defer fake.getDBStatusMutex.RUnlock()
	return len(fake.getDBStatusArgsForCall)
}

func (fake *FakeSchedulerDB) GetDBStatusReturns(result1 sql.DBStats) {
	fake.GetDBStatusStub = nil
	fake.getDBStatusReturns = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeSchedulerDB) GetActiveSchedules() (map[string]*models.ActiveSchedule, error) {
	fake.getActiveSchedulesMutex.Lock()
	fake.getActiveSchedulesArgsForCall = append(fake.getActiveSchedulesArgsForCall, struct{}{})
	fake.recordInvocation("GetActiveSchedules", []interface{}{})
	fake.getActiveSchedulesMutex.Unlock()
	if fake.GetActiveSchedulesStub != nil {
		return fake.GetActiveSchedulesStub()
	}
	return fake.getActiveSchedulesReturns.result1, fake.getActiveSchedulesReturns.result2
}

func (fake *FakeSchedulerDB) GetActiveSchedulesCallCount() int {
	fake.getActiveSchedulesMutex.RLock()
	defer fake.getActiveSchedulesMutex.RUnlock()
	return len(fake.getActiveSchedulesArgsForCall)
}

func (fake *FakeSchedulerDB) GetActiveSchedulesReturns(result1 map[string]*models.ActiveSchedule, result2 error) {
	fake.GetActiveSchedulesStub = nil
	fake.getActiveSchedulesReturns = struct {
		result1 map[string]*models.ActiveSchedule
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	return fake.closeReturns.result1
}

func (fake *FakeSchedulerDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSchedulerDB) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSchedulerDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDBStatusMutex.RLock()
	defer fake.getDBStatusMutex.RUnlock()
	fake.getActiveSchedulesMutex.RLock()
	defer fake.getActiveSchedulesMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSchedulerDB) recordInvocation(key string, args []interface{}) {
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

var _ db.SchedulerDB = new(FakeSchedulerDB)
