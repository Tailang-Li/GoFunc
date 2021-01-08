package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"

	"github.com/GoFunc/Tailang-Li/protobuffer/exampleprotobuffer"
)

const (
	testOutputFileName          = "output.pb"
	testOutputTextProtoFileName = "output.pb.txt"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	req := exampleprotobuffer.GetSomethingReq{
		SomeInt64: r.Int63(),
		SomeStr:   fmt.Sprintf("some_str_%d", r.Int()),
	}

	// change pb into bytes
	buf, err := proto.Marshal(&req)
	if err != nil {
		glog.Fatal(err)
	}

	// change pb bytes into file
	err = ioutil.WriteFile(testOutputFileName, buf, 0644)
	if err != nil {
		glog.Fatal(err)
	}

	// change pb bytes into text format
	pbStr := proto.MarshalTextString(&req)
	fmt.Printf("pb text string:\n%s", pbStr)

	// change pb bytes into text format file
	f, err := os.Create(testOutputTextProtoFileName)
	if err != nil {
		glog.Fatal(err)
	}
	defer f.Close()

	err = proto.MarshalText(f, &req)
	if err != nil {
		glog.Fatal(err)
	}
}
