package main

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var (
	id1      = "1"
	id2      = "2"
	newName1 = "lzb_new1"
	newSex1  = ""
	newAge1  = "50"
	newName2 = "lzb_new2"
	newSex2  = "女"
	newAge2  = "5"
)

type User1 struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  string `json:"age"`
}

func GetNewStub() *shim.MockStub {
	var scc = new(Example)
	var stub = shim.NewMockStub("ex01", scc)
	stub.MockInit("init", nil)
	return stub
}

func TestExample_createCompositeKey(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("createCompositeKey")})
	if res.Status == shim.OK {
		t.Log(res.Message)
	}

}

func TestExample_putState(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("putState")})
	if res.Status == shim.OK {
		t.Log(res.Message)
	}
}

func TestExample_delState(t *testing.T) {
	stub := GetNewStub()
	res1 := stub.MockInvoke("1", [][]byte{[]byte("getState")})
	if res1.Status != shim.OK {
		t.Log(res1.Message)
	}
	res2 := stub.MockInvoke("2", [][]byte{[]byte("delState")})
	if res2.Status != shim.OK {
		t.Log(res2.Message)
	}
	t.Log("删除成功")
	res3 := stub.MockInvoke("3", [][]byte{[]byte("getState")})
	if res3.Status != shim.OK {
		t.Log("查询失败,不存在")
	}
}

func TestExample_getState(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("getState")})
	if res.Status != shim.OK {
		t.Log(res.Message)
	}
}

func TestExample_getStateByPartialCompositeKey(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("getStateByPartialCompositeKey")})
	if res.Status != shim.OK {
		t.Log(res.Message)
	}
}

func TestExample_getHistoryForKey(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("getHistoryForKey")})
	if res.Status != shim.OK {
		t.Log(res.Message)
	}
}

func TestExample_getStateByRange(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("getStateByRange")})
	if res.Status != shim.OK {
		t.Log(res.Message)
	}
}
