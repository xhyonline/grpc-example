package rpc

import (
	"context"
	"github.com/xhyonline/grpc-example/gen/basic"
	"github.com/xhyonline/grpc-example/gen/user"
	"github.com/xhyonline/grpc-example/mod"
)

// UserServer 用户服务,实现了 proto 中的所有方法
type UserServer struct {
}

// AddUser 创建一个用户
func (s *UserServer) AddUser(ctx context.Context, u *user.User) (*user.AddResponse, error) {
	conn := GetDB()
	item := &mod.User{
		Name: u.GetName(),
		Age:  int(u.GetAge()),
	}
	err := conn.Create(item).Error
	if err != nil {
		return &user.AddResponse{Success: false}, err
	}
	// 获取入参
	inPone := u.GetUserPhone()
	for _, v := range inPone {
		p := &mod.Phone{
			UserID: item.ID,
			Number: v.GetPhoneNumber(),
			Type:   int(v.GetPhoneType()),
		}
		err = conn.Create(&p).Error
		if err != nil {
			return &user.AddResponse{Success: false}, err
		}
	}
	return &user.AddResponse{ID: int64(item.ID), Success: true}, nil
}

// GetUserByID 通过用户 ID 获取一个用户
func (s *UserServer) GetUserByID(ctx context.Context, id *basic.Int) (*user.User, error) {
	conn := GetDB()
	u := new(mod.User)
	err := conn.Model(&mod.User{}).Where("id = ?", id.GetInt()).Preload("Phone").First(u).Error
	if err != nil {
		return nil, err
	}
	resp := &user.User{
		UserPhone: make([]*user.Phone, 0, len(u.Phone)),
		UserID:    int64(u.ID),
		Name:      u.Name,
		Age:       int64(u.Age),
	}
	for _, v := range u.Phone {
		resp.UserPhone = append(resp.UserPhone, &user.Phone{
			PhoneNumber: v.Number,
			PhoneType:   user.Phone_PhoneTypeEnum(v.Type),
		})
	}
	return resp, nil
}

func (s *UserServer) RemoveUserByID(ctx context.Context, id *basic.Int) (*basic.Bool, error) {
	//conn:=server.GetDB()
	//conn.Model(&mod.User{}).Delete()
	return nil, nil
}

func (s *UserServer) GetRandomUser(ctx context.Context, empty *basic.Empty) (*user.User, error) {
	return nil, nil
}

func (s *UserServer) GetAllUser(ctx context.Context, empty *basic.Empty) (*user.RepetedUser, error) {
	return nil, nil
}
