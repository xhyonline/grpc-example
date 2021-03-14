package main

import (
	"context"
	"fmt"
	"github.com/xhyonline/grpc-example/gen/basic"
	"github.com/xhyonline/grpc-example/gen/user"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := user.NewUserServerClient(conn)
	//u := &user.User{
	//	Name:      "兰陵美酒郁金香",
	//	Age:       24,
	//	UserPhone: make([]*user.Phone, 0, 2),
	//}
	//u.UserPhone = append(u.UserPhone, &user.Phone{PhoneType: 1, PhoneNumber: "13819720378"})
	//u.UserPhone = append(u.UserPhone, &user.Phone{PhoneNumber: "86506859"})
	//
	//result, err := client.AddUser(context.Background(), u)
	//if err != nil {
	//	log.Fatalf("插入失败 %s", err)
	//}
	//fmt.Println("新增的数据主键ID是", result.GetID())

	result, err := client.GetUserByID(context.Background(), &basic.Int{
		Int: 1,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("%+v", result)

}
