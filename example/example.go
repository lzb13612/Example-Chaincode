// @Title		参数读取API与账本状态交互API
// @Author		lzb
// @Description	学习使用链码API
package main

// 导入 shim 链码API库,及 peer 节点响应模板库
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Example struct {
}

func (e *Example) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 创建复合主键
	userKey, err := stub.CreateCompositeKey("name", []string{"lzb"})
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create name key error:%s", err),
		}
	}
	// 序列化
	value, _ := json.Marshal("value")
	// 上传数据状态
	err = stub.PutState(userKey, value)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("put name key and info error:%s", err),
		}
	}

	// 创建复合主键
	userKey, err = stub.CreateCompositeKey("name", []string{"lzb1"})
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create name key error:%s", err),
		}
	}
	// 序列化
	value, _ = json.Marshal("value1")
	// 上传数据状态
	err = stub.PutState(userKey, value)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("put name key and info error:%s", err),
		}
	}

	// 创建复合主键
	userKey, err = stub.CreateCompositeKey("name", []string{"lzb2"})
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create name key error:%s", err),
		}
	}
	// 序列化
	value, _ = json.Marshal("value2")
	// 上传数据状态
	err = stub.PutState(userKey, value)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("put name key and info error:%s", err),
		}
	}

	return pb.Response{
		Status:  shim.OK,
		Message: "init success",
	}
}

/*=====================================================================	*
 *							参数读取系列                                 	*
 *=====================================================================	*/

//Invoke
// @title		Invoke -> 操作功能列表
// @description	提取调用链码交易中的参数，其中第一个作为被调用的函数名称，剩下的参数作为函数的执行参数。
// @auth		lzb
// @param 		stub	shim库	"包含所有链码API的库"
// @return		pb		peer库	"返回状态码和响应信息"
func (e *Example) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, _ := stub.GetFunctionAndParameters()
	switch funcName {
	case "createCompositeKey":
		return createCompositeKey(stub)
	case "putState":
		return putState(stub)
	case "delState":
		return delState(stub)
	case "getState":
		return getState(stub)
	case "getStateByPartialCompositeKey":
		return getStateByPartialCompositeKey(stub)
	case "getHistoryForKey":
		return getHistoryForKey(stub)
	case "getStateByRange":
		return getStateByRange(stub)
	default:
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("not find function:%s", funcName),
		}
	}
}

/*=====================================================================	*
 *							创建功能系列                                 	*
 *=====================================================================	*/

// @title		createCompositeKey -> 创建主键
// @description	创建一个复合键。
// @auth		lzb
// @param		objectType 	字符串	"键名"
//				attributes	字符组	"值"
//				stub	shim库	"包含所有链码API的库"
// @return		pb		peer库	"返回状态码和响应信息"
func createCompositeKey(stub shim.ChaincodeStubInterface) pb.Response {
	// 主键名称
	indexName := "sex~name"
	// 创建复合主键
	indexKey, err := stub.CreateCompositeKey(indexName, []string{"boy", "lzb3"})
	// 判断错误
	if err != nil {
		fmt.Println(err)
	}
	// 输出主键信息
	fmt.Println("indexKey:", indexKey)
	// 也可创建多个主键
	indexKey, err = stub.CreateCompositeKey(indexName, []string{"girl", "lzb4"})
	fmt.Println("indexKey:", indexKey)
	return pb.Response{
		Status:  shim.OK,
		Message: "createCompositeKey success",
	}
}

// @title		putState -> 存入数据状态
// @description	根据指定的key，将对应的value保存在分类账本中。
// @auth		lzb
// @param		key 	字符串	"键名"
//				value	字符组	"值"
//				stub	shim库	"包含所有链码API的库"
// @return		pb		peer库	"返回状态码和响应信息"
func putState(stub shim.ChaincodeStubInterface) pb.Response {
	// 创建复合主键
	indexKey, err := stub.CreateCompositeKey("name", []string{"lzb5"})
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create name key error:%s", err),
		}
	}
	// 序列化
	value, _ := json.Marshal("value")
	// 输出主键信息
	fmt.Println("indexKey:", indexKey)
	// 接收错误
	if err := stub.PutState(indexKey, value); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("put state error:%s", err),
		}
	}
	// 测试刚上传数据是否完成
	bytes, err := stub.GetState(indexKey)
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("get name key state error:%s", err),
		}
	}
	// 反序列化
	var names string
	_ = json.Unmarshal(bytes, &names)
	fmt.Println("测试获取到的name:", names)
	// 返回成功信息
	return pb.Response{
		Status:  shim.OK,
		Message: "put state success",
	}
}

/*=====================================================================	*
 *							删除功能系列                                 	*
 *=====================================================================	*/
// @title		delState -> 删除账本里某个数据状态
// @description 根据指定的key将对应的数据状态删除
// @author		lzb
// @param		key 	字符串	"键名"
//				stub	shim库	"包含所有链码API的库"
// @return		pb		peer库	"返回状态码和响应信息"
func delState(stub shim.ChaincodeStubInterface) pb.Response {
	userKey, err := stub.CreateCompositeKey("name", []string{"lzb"})
	// 接收错误
	err = stub.DelState(userKey)
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("delete state error:%s", err),
		}
	}
	return pb.Response{
		Status:  shim.OK,
		Message: "delete state success",
	}
}

/*=====================================================================	*
 *							查询功能系列                                 	*
 *=====================================================================	*/
// @title		getState -> 获取账本里某个数据状态
// @description	根据指定的key查询相应的数据状态
// @auth		lzb
// @param		key		字符串	"键名"
//				stub	shim库	"包含所有链码API的库"
// @return		pb		peer库	"返回状态码和响应信息"
func getState(stub shim.ChaincodeStubInterface) pb.Response {
	// 创建复合键
	userKey, err := stub.CreateCompositeKey("name", []string{"lzb"})
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("create name key error:%s", err),
		}
	}
	// 接收字符组和错误
	userBytes, err := stub.GetState(userKey)
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("get state error:%s", err),
		}
	}
	// 定义接收内容的变量
	var name string
	// 反序列化
	if err := json.Unmarshal(userBytes, &name); err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("unmarshal name error:%s", err),
		}
	}
	fmt.Println("name:", name)
	return pb.Response{
		Status:  shim.OK,
		Message: "get state success",
		Payload: userBytes,
	}
}

// @title		getStateByRange -> 起止键区间查询数据状态
// @description 查询指定范围内的键值，startKey为起始key，endKey为终止key
// @author		lzb
// @param		startKey	字符串	"开始的键名"
//				endKey		字符串	"结束的键名"
//				stub		shim库	"包含所有链码API的库"
// @return		pb			peer库	"返回状态码和响应信息"
func getStateByRange(stub shim.ChaincodeStubInterface) pb.Response {
	// 创建三个测试用的数据
	_ = stub.PutState("name1", []byte("lzb1"))
	_ = stub.PutState("name2", []byte("lzb2"))
	_ = stub.PutState("name3", []byte("lzb3"))
	// 获取name1到name3之间的值
	resultIterator, err := stub.GetStateByRange("name1", "name3")
	// 判断错误
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("get state by range error:%s", err),
		}
	}
	fmt.Println("-----start resultIterator-----")
	// 遍历迭代器
	for resultIterator.HasNext() {
		item, _ := resultIterator.Next()
		fmt.Println(string(item.Value))
	}
	// 关闭迭代器
	defer func(resultIterator shim.StateQueryIteratorInterface) {
		err := resultIterator.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resultIterator)
	fmt.Println("-----end resultIterator-----")
	return pb.Response{
		Status:  shim.OK,
		Message: "get state by range success",
	}
}

// @title		getStateByPartialCompositeKey -> 复合键查询
// @description 根据局部的复合键（前缀）返回所有匹配的键值，即与账本中的键进行前缀匹配，
//				返回结果是一个迭代器结构，可以按照字典序迭代每个键值对，最后需要调用 Close() 方法关闭
//				注意:该方法的使用需要节点配置中打开历史数据库特性
// @author		lzb
// @param		objectType	字符串	"键名"
//				keys		字符串组	"键名对应的值"
//				stub		shim库	"包含所有链码API的库"
// @return		pb			peer库	"返回状态码和响应信息"
func getStateByPartialCompositeKey(stub shim.ChaincodeStubInterface) pb.Response {
	// 通过复合键获取某键或者所有的数据状态
	resultsIterator, err := stub.GetStateByPartialCompositeKey("name", []string{})
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("get name state by partial composite key error:%s", err),
		}
	}
	// 遍历迭代器
	for resultsIterator.HasNext() {
		// 提取数据
		val, err := resultsIterator.Next()
		// 判断错误
		if err != nil {
			return pb.Response{
				Status:  shim.ERROR,
				Message: fmt.Sprintf("get name state by partial composite key error:%s", err),
			}
		}
		fmt.Println(val.Key)
		fmt.Println(string(val.Value))
	}
	// 关闭迭代器
	defer func(resultsIterator shim.StateQueryIteratorInterface) {
		err := resultsIterator.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resultsIterator)
	return pb.Response{
		Status:  shim.ERROR,
		Message: fmt.Sprintf("get state by partial composite key success"),
	}
}

// @title		GetHistoryForKey -> 获取某键的历史数据状态记录
// @description 根据指定的 key 查询所有的历史记录信息。
//				注意:该方法的使用需要节点配置中打开历史数据库特性
// @author		lzb
// @param		key	字符串	"键名"
//				stub	shim库	"包含所有链码API的库"
// @return		pb		peer库	"返回状态码和响应信息"
func getHistoryForKey(stub shim.ChaincodeStubInterface) pb.Response {
	// 获取历史数据状态
	historyIterator, err := stub.GetHistoryForKey("name")
	if err != nil {
		return pb.Response{
			Status:  shim.ERROR,
			Message: fmt.Sprintf("get history for key error:%s", err),
		}
	}

	fmt.Println("-----start historyIterator-----")
	// 遍历迭代器
	for historyIterator.HasNext() {
		item, err := historyIterator.Next()
		if err != nil {
			return pb.Response{
				Status:  shim.ERROR,
				Message: fmt.Sprintf("history iterator error:%s", err),
			}
		}
		fmt.Println(string(item.TxId))
		fmt.Println(string(item.Value))
	}
	fmt.Println("-----end historyIterator-----")
	// 关闭迭代器
	defer func(historyIterator shim.HistoryQueryIteratorInterface) {
		err := historyIterator.Close()
		if err != nil {
			fmt.Println("close iterator error")
		}
	}(historyIterator)
	return pb.Response{
		Status:  shim.OK,
		Message: "get history state by key success",
		//Payload: ,
	}
}

// title		main -> 主方法
// description	操作功能
// auth			lzb
func main() {
	err := shim.Start(new(Example))
	if err != nil {
		fmt.Println(err)
	}
}
