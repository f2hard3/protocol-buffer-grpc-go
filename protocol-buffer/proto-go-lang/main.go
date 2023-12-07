package main

import (
	"fmt"
	"reflect"

	pb "github.com/f2hard3/proto-go-lang/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   42,
			Name: "name1",
		},
		MultipleDummies: []*pb.Dummy{
			{Id: 43, Name: "name2"},
			{Id: 44, Name: "name3"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: 2, //pb.EyeColor_EYE_COLOR_BLUE
	}
}

func doOneof(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Printf("message has unexpeced type: %v\n", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"
	writeTofile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)

	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)

	return message
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doComplex())
	fmt.Println(doEnum())
	doOneof(&pb.Result_Id{Id: 42})
	doOneof(&pb.Result_Message{Message: "a message"})
	doOneof("sad")
	fmt.Println(doMap())
	doFile(doSimple())

	jsonString := doToJSON(doSimple())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	jsonString = doToJSON(doComplex())
	message = doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	fmt.Println(doFromJSON(`{"id": 42, "unknown": "test"}`, reflect.TypeOf(pb.Simple{})))
}
