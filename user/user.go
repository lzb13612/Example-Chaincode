package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// User 案例
type User struct {
}

// UserInfo 用户对象结构体 -> 定义了用户的基础信息
type UserInfo struct {
	Id   string `json:"id"`   // 用户id
	Name string `json:"name"` // 用户名
	Sex  string `json:"sex"`  // 用户性别
}

// Init
// @title		Init -> 初始化
// @description	对用户和管理员进行初始化一个账户。
// @auth		lzb
// @param 		stub	shim库	"包含所有链码API的库"
// @return		pb		peer库	"返回状态码和响应信息"
func (e *User) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 创建一个用户
	user := UserInfo{
		Id:   "1",
		Name: "lzb1",
		Sex:  "男",
	}
	// 创建复合主键
	userKey, err := stub.CreateCompositeKey("user", []string{user.Id})
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create user key error:%s", err),
		}
	}
	// 序列化
	userBytes, err := json.Marshal(user)
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("marshal user error:%s", err),
		}
	}
	// 上传数据状态
	if err := stub.PutState(userKey, userBytes); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("put user key and info error:%s", err),
		}
	}
	// 创建一个管理员
	user = UserInfo{
		Id:   "2",
		Name: "lzb2",
		Sex:  "女",
	}
	// 创建复合主键
	userKey, err = stub.CreateCompositeKey("user", []string{user.Id})
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create user key error:%s", err),
		}
	}
	// 序列化
	userBytes, err = json.Marshal(user)
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("marshal user error:%s", err),
		}
	}
	// 上传数据状态
	if err := stub.PutState(userKey, userBytes); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("put userKey error:%s", err),
		}
	}
	return pb.Response{
		Status:  shim.OK,
		Message: "Init success",
	}
}

// Invoke
// @title		Invoke -> 调用方法
// @description	对通过fabric网络传会的参数进行判断,以此调用对应方法。
// @auth		lzb
// @param 		stub	shim库	"包含所有链码API的库"
// @return		pb		peer库	"返回状态码和响应信息"
func (e *User) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// 获取方法名与参数
	funcName, args := stub.GetFunctionAndParameters()
	// 选择方法
	switch funcName {
	case "addUser":
		return addUser(stub, args)
	case "queryOnceUser":
		return queryOnceUser(stub, args)
	case "queryAllUser":
		return queryAllUser(stub)
	case "alterUser":
		return alterUser(stub, args)
	case "delUser":
		return delUser(stub, args)
	default:
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("not find function %s", funcName),
		}
	}
}

func addUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return pb.Response{
			Status:  shim.ERRORTHRESHOLD,
			Message: "no enough args",
		}
	}
	var userInfo UserInfo
	if err := json.Unmarshal([]byte(args[0]), &userInfo); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("unmarshal error:%s", err),
		}
	}
	key, err := stub.CreateCompositeKey("user", []string{userInfo.Id})
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create user key error:%s", err),
		}
	}
	verifyByte, err := stub.GetState(key)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("get user state error:%s", err),
		}
	}
	if len(verifyByte) != 0 {
		return pb.Response{
			Status:  shim.ERRORTHRESHOLD,
			Message: "user exist",
		}
	}
	userByte, err := json.Marshal(userInfo)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("marshal user info error:%s", err),
		}
	}
	err = stub.PutState(key, userByte)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("put user %s state error:%s", userInfo.Id, err),
		}
	}
	return pb.Response{
		Status:  shim.OK,
		Message: "add user success",
	}
}

func queryOnceUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return pb.Response{
			Status:  shim.ERRORTHRESHOLD,
			Message: "no enough args",
		}
	}
	var userInfo UserInfo
	if err := json.Unmarshal([]byte(args[0]), &userInfo); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("unmarshal user error:%s", err),
		}
	}
	key, err := stub.CreateCompositeKey("user", []string{userInfo.Id})
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create user key error:%s", err),
		}
	}
	userByte, err := stub.GetState(key)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("get user %s state error:%s", userInfo.Id, err),
		}
	}
	if len(userByte) == 0 {
		return pb.Response{
			Status:  shim.ERRORTHRESHOLD,
			Message: fmt.Sprintf("user %s does not exist", userInfo.Id),
		}
	}
	return pb.Response{
		Status:  shim.OK,
		Message: "get once user success",
		Payload: userByte,
	}
}

func queryAllUser(stub shim.ChaincodeStubInterface) pb.Response {
	userInfos := make([]*UserInfo, 0)
	resultIterator, err := stub.GetStateByPartialCompositeKey("user", []string{})
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("get user info by partial composite key error:%s", err),
		}
	}
	defer resultIterator.Close()
	for resultIterator.HasNext() {
		val, _ := resultIterator.Next()
		userInfo := new(UserInfo)
		if err := json.Unmarshal(val.GetValue(), &userInfo); err != nil {
			return pb.Response{
				Status:  shim.ERROR,
				Message: fmt.Sprintf("unmarshal user info error:%s", err),
			}
		}
		userInfos = append(userInfos, userInfo)
	}
	userByte, err := json.Marshal(userInfos)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: "marshal user info error",
		}
	}
	return pb.Response{
		Status:  shim.OK,
		Message: "get all user info success",
		Payload: userByte,
	}
}

func alterUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return pb.Response{
			Status:  shim.ERRORTHRESHOLD,
			Message: "no enough args",
		}
	}
	var newUserInfo UserInfo
	if err := json.Unmarshal([]byte(args[0]), &newUserInfo); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("unmarshal user error:%s", err),
		}
	}
	oldUserInfoKey, err := stub.CreateCompositeKey("user", []string{newUserInfo.Id})
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: "create key error",
		}
	}
	oldUserInfoByte, err := stub.GetState(oldUserInfoKey)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: "get user error",
		}
	}
	var oldUserInfo UserInfo
	if err := json.Unmarshal(oldUserInfoByte, &oldUserInfo); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("unmarshal user error:%s", err),
		}
	}
	if oldUserInfo.Name != newUserInfo.Name {
		oldUserInfo.Name = newUserInfo.Name
	}
	if oldUserInfo.Sex != newUserInfo.Sex {
		oldUserInfo.Sex = newUserInfo.Sex
	}
	userByte, err := json.Marshal(oldUserInfo)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: "marshal user error",
		}
	}
	err = stub.PutState(oldUserInfoKey, userByte)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("put user error:%S", err),
		}
	}
	return pb.Response{
		Status:  shim.OK,
		Message: "alt user success",
	}
}

func delUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return pb.Response{
			Status:  shim.ERRORTHRESHOLD,
			Message: "no enough args",
		}
	}
	var userInfo UserInfo
	if err := json.Unmarshal([]byte(args[0]), &userInfo); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("unmarshal user error:%s", err),
		}
	}
	userKey, err := stub.CreateCompositeKey("user", []string{userInfo.Id})
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: "create key error",
		}
	}
	err = stub.DelState(userKey)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("del user error:%s", err),
		}
	}
	return pb.Response{
		Status:  shim.OK,
		Message: "del user state success",
	}
}

// title		main -> 启动
// description	启动合约
// auth			lzb
func main() {
	err := shim.Start(new(User))
	if err != nil {
		fmt.Printf("User start error:%s", err)
	}
}
