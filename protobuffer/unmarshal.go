package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"

	"github.com/GoFunc/Tailang-Li/protobuffer/exampleprotobuffer"
)

const (
	testOutputFileName = "output.pb"
)

func main() {

	buf, err := ioutil.ReadFile(testOutputFileName)
	if err != nil {
		glog.Fatal(err)
	}

	req := exampleprotobuffer.GetSomethingReq{}
	err = proto.Unmarshal(buf, &req)
	if err != nil {
		glog.Fatal(err)
	}

	fmt.Printf("unmarshaled proto: %+v\n", req)
}
