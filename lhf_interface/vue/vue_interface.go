package vue

import (
	"encoding/json"
	nr "github.com/wunder3605/noderank"
)
type Message struct {
	Code int64
	Message string
	Data interface{}
}

type OCli struct {

}

type AddAtInfo interface {
	AddAttestationInfoFunction(_data []byte)Message
	GetRankFunction(_data []byte) Message
}

func (o *OCli)AddAttestationInfoFunction(_data []byte )Message{
	mess:=Message{}

	m:=make(map[string]string)
	err := json.Unmarshal(_data, &m)
	if err!=nil{
		mess=Message{Code:0,Message:"类型转换异常"}
		return mess
	}
	info:=make([]string,3)
	info[0]=m["Attester"]
	info[1]=m["Attestee"]
	info[0]=m["Score"]
	nr.AddAttestationInfo(info)
	mess=Message{Code:1,Message:"节点添加成功"}
	return mess
}

type parameter struct {
	Period string `json:"period"`
	NumRank int64 `json:"numRank"`
}

func (o *OCli)GetRankFunction(_data []byte)Message{
	mess:=Message{}
	var para parameter
	err:=json.Unmarshal(_data,&para)
	if err!=nil{
		mess=Message{Code:0,Message:"类型转换异常"}
		return mess
	}

	tee:=nr.GetRank(para.Period,para.NumRank)  //返回值[]teectx
	if tee==nil{
		mess=Message{Code:0,Message:"查询失败"}
		return mess
	}
	mess=Message{Code:1,Message:"查询成功",Data:tee}
	return mess
}
