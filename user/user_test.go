package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var (
	id_1   = "1"
	name_1 = "lzb1"
	sex_1  = "男"
	id_2   = "2"
	name_2 = "lzb2"
	sex_2  = "女"
	id1    = "3"
	id2    = "4"
	id3    = "5"
	name1  = "lzb3"
	name2  = "lzb4"
	name3  = "lzb5"
	sex1   = "男"
	sex2   = "女"
	sex3   = "男"
)

type UserInfoTest struct {
	Id   string `json:"id"`   // 用户id
	Name string `json:"name"` // 用户名
	Sex  string `json:"sex"`  // 用户性别
}

var userInfoTest_1 = UserInfoTest{
	Id:   id_1,
	Name: name_1,
	Sex:  sex_1,
}

var user_1, _ = json.Marshal(userInfoTest_1)

var userInfoTest_2 = UserInfoTest{
	Id:   id_2,
	Name: name_2,
	Sex:  sex_2,
}

var user_2, _ = json.Marshal(userInfoTest_2)

var userInfoTest1 = UserInfoTest{
	Id:   id1,
	Name: name1,
	Sex:  sex1,
}

var user1, _ = json.Marshal(userInfoTest1)

var userInfoTest2 = UserInfoTest{
	Id:   id2,
	Name: name2,
	Sex:  sex2,
}

var user2, _ = json.Marshal(userInfoTest2)

var userInfoTest3 = UserInfoTest{
	Id:   id3,
	Name: name3,
	Sex:  sex3,
}

var user3, _ = json.Marshal(userInfoTest3)

func GetNewStub() *shim.MockStub {
	var scc = new(User)
	var stub = shim.NewMockStub("ex01", scc)
	stub.MockInit("init", nil)
	return stub
}

func TestUser_queryAllUser(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("queryAllUser")})
	if res.Status == shim.OK {
		var userInfoTest []UserInfoTest
		_ = json.Unmarshal(res.Payload, &userInfoTest)
		t.Log(res.Message)
		t.Log(userInfoTest)
	} else {
		t.Log(res.Message)
	}
}

func TestUser_queryOnceUser(t *testing.T) {
	stub := GetNewStub()
	userAge := UserInfoTest{
		Id: "1",
	}
	userId, _ := json.Marshal(userAge)
	res := stub.MockInvoke("1", [][]byte{[]byte("queryOnceUser"), userId})
	if res.Status == shim.OK {
		var userInfoTest UserInfoTest
		_ = json.Unmarshal(res.Payload, &userInfoTest)
		t.Log(res.Message)
		t.Log(userInfoTest)
	} else {
		t.Log(res.Message)
	}

	userAge = UserInfoTest{
		Id: "3",
	}
	userId, _ = json.Marshal(userAge)
	res = stub.MockInvoke("1", [][]byte{[]byte("queryOnceUser"), userId})
	if res.Status == shim.OK {
		var userInfoTest UserInfoTest
		_ = json.Unmarshal(res.Payload, &userInfoTest)
		t.Log(res.Message)
		t.Log(userInfoTest)
	} else {
		t.Log(res.Message)
	}
}

func TestUser_addUser(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("addUser"), user1})
	if res.Status == shim.OK {
		t.Log(res.Message)
	} else {
		t.Log(res.Message)
	}

	res1 := stub.MockInvoke("2", [][]byte{[]byte("queryOnceUser"), user1})
	if res1.Status == shim.OK {
		var userInfoTest UserInfoTest
		_ = json.Unmarshal(res1.Payload, &userInfoTest)
		t.Log(res1.Message)
		t.Log(userInfoTest)
	} else {
		t.Log(res1.Message)
	}
}

func TestUser_alterUser(t *testing.T) {
	stub := GetNewStub()
	res1 := stub.MockInvoke("1", [][]byte{[]byte("addUser"), user1})
	if res1.Status == shim.OK {
		t.Log(res1.Message)
	} else {
		t.Log(res1.Message)
	}

	res2 := stub.MockInvoke("2", [][]byte{[]byte("queryOnceUser"), user1})
	if res2.Status == shim.OK {
		var userInfoTest UserInfoTest
		_ = json.Unmarshal(res2.Payload, &userInfoTest)
		t.Log(res2.Message)
		t.Log("修改前:", userInfoTest)
	} else {
		t.Log(res2.Message)
	}

	newUserInfo := UserInfoTest{
		Id:   "3",
		Name: "test",
		Sex:  "女",
	}

	newUserInfoByte, _ := json.Marshal(newUserInfo)
	res3 := stub.MockInvoke("3", [][]byte{[]byte("alterUser"), newUserInfoByte})
	if res3.Status == shim.OK {
		t.Log(res3.Message)
	} else {
		t.Log(res3.Message)
	}

	res4 := stub.MockInvoke("4", [][]byte{[]byte("queryOnceUser"), user1})
	if res4.Status == shim.OK {
		var userInfoTest UserInfoTest
		_ = json.Unmarshal(res4.Payload, &userInfoTest)
		t.Log(res4.Message)
		t.Log("修改后:", userInfoTest)
	} else {
		t.Log(res4.Message)
	}
}

func TestUser_delUser(t *testing.T) {
	stub := GetNewStub()
	res := stub.MockInvoke("1", [][]byte{[]byte("delUser"), user_1})
	if res.Status == shim.OK {
		t.Log(res.Message)
	} else {
		t.Log(res.Message)
	}

	res2 := stub.MockInvoke("2", [][]byte{[]byte("queryOnceUser"), user_1})
	if res2.Status == shim.OK {
		var userInfoTest UserInfoTest
		_ = json.Unmarshal(res2.Payload, &userInfoTest)
		t.Log(res2.Message)
		t.Log(userInfoTest)
	} else {
		t.Log("删除后:", res2.Message)
	}

}
