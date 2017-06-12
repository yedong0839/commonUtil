package commonUtil

import (
	"testing"
	"log"
	"encoding/json"
)

func Test_MD5File(t *testing.T) {
	path := "d:/t.txt"
	md5, err := MD5File(&path)
	log.Print(md5)
	if err == nil{
		t.Log(md5)
	}else {
		t.Error(err)
	}
}

func Test_HttpPost(t *testing.T)  {


}

func Test_HttpP(t *testing.T)  {
	m := map[string]string{"name":"john","hello":"mm"}
	r,err :=json.Marshal(m)
	if err !=nil{
		t.Error(err)
	}else{
		t.Log(string(r))
	}
}

