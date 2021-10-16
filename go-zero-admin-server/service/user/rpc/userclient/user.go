// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"go-zero-admin-server/service/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	IdReq          = user.IdReq
	RoleReq        = user.RoleReq
	UsernameReq    = user.UsernameReq
	UserInfoReply  = user.UserInfoReply
	IdsReq         = user.IdsReq
	UserUpdateReq  = user.UserUpdateReq
	RolesReq       = user.RolesReq
	UsersReq       = user.UsersReq
	IsExistReply   = user.IsExistReply
	IsSuccessReply = user.IsSuccessReply
	EmptyReq       = user.EmptyReq
	UserAddReq     = user.UserAddReq
	RoleReply      = user.RoleReply
	RolesReply     = user.RolesReply
	UsersInfoReply = user.UsersInfoReply

	User interface {
		GetUserById(ctx context.Context, in *IdReq) (*UserInfoReply, error)
		GetUserByUsername(ctx context.Context, in *UsernameReq) (*UserInfoReply, error)
		GetUserByUsernameLike(ctx context.Context, in *UsernameReq) (*UsersInfoReply, error)
		IsExistUserById(ctx context.Context, in *IdReq) (*IsExistReply, error)
		IsExistUserByUsername(ctx context.Context, in *UsernameReq) (*IsExistReply, error)
		GetUsers(ctx context.Context, in *UsersReq) (*UsersInfoReply, error)
		AddUser(ctx context.Context, in *UserAddReq) (*IsSuccessReply, error)
		UpdateUser(ctx context.Context, in *UserUpdateReq) (*IsSuccessReply, error)
		DelUser(ctx context.Context, in *IdReq) (*IsSuccessReply, error)
		DelSomeUser(ctx context.Context, in *IdsReq) (*IsSuccessReply, error)
		GetAllRole(ctx context.Context, in *EmptyReq) (*RolesReply, error)
		EditUserRoles(ctx context.Context, in *RolesReq) (*IsSuccessReply, error)
		IsExistRoleById(ctx context.Context, in *IdReq) (*IsExistReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUserById(ctx context.Context, in *IdReq) (*UserInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserById(ctx, in)
}

func (m *defaultUser) GetUserByUsername(ctx context.Context, in *UsernameReq) (*UserInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserByUsername(ctx, in)
}

func (m *defaultUser) GetUserByUsernameLike(ctx context.Context, in *UsernameReq) (*UsersInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserByUsernameLike(ctx, in)
}

func (m *defaultUser) IsExistUserById(ctx context.Context, in *IdReq) (*IsExistReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.IsExistUserById(ctx, in)
}

func (m *defaultUser) IsExistUserByUsername(ctx context.Context, in *UsernameReq) (*IsExistReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.IsExistUserByUsername(ctx, in)
}

func (m *defaultUser) GetUsers(ctx context.Context, in *UsersReq) (*UsersInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUsers(ctx, in)
}

func (m *defaultUser) AddUser(ctx context.Context, in *UserAddReq) (*IsSuccessReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddUser(ctx, in)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UserUpdateReq) (*IsSuccessReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in)
}

func (m *defaultUser) DelUser(ctx context.Context, in *IdReq) (*IsSuccessReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DelUser(ctx, in)
}

func (m *defaultUser) DelSomeUser(ctx context.Context, in *IdsReq) (*IsSuccessReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DelSomeUser(ctx, in)
}

func (m *defaultUser) GetAllRole(ctx context.Context, in *EmptyReq) (*RolesReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetAllRole(ctx, in)
}

func (m *defaultUser) EditUserRoles(ctx context.Context, in *RolesReq) (*IsSuccessReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EditUserRoles(ctx, in)
}

func (m *defaultUser) IsExistRoleById(ctx context.Context, in *IdReq) (*IsExistReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.IsExistRoleById(ctx, in)
}
