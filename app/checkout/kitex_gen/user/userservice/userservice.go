// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	user "douyin_mall/checkout/kitex_gen/user"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newRegisterArgs,
		newRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newLoginArgs,
		newLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUser": kitex.NewMethodInfo(
		getUserHandler,
		newGetUserArgs,
		newGetUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateUser": kitex.NewMethodInfo(
		updateUserHandler,
		newUpdateUserArgs,
		newUpdateUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteUser": kitex.NewMethodInfo(
		deleteUserHandler,
		newDeleteUserArgs,
		newDeleteUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUserRoleById": kitex.NewMethodInfo(
		getUserRoleByIdHandler,
		newGetUserRoleByIdArgs,
		newGetUserRoleByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AddReceiveAddress": kitex.NewMethodInfo(
		addReceiveAddressHandler,
		newAddReceiveAddressArgs,
		newAddReceiveAddressResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetReceiveAddress": kitex.NewMethodInfo(
		getReceiveAddressHandler,
		newGetReceiveAddressArgs,
		newGetReceiveAddressResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.RegisterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RegisterArgs:
		success, err := handler.(user.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *user.RegisterReq
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.RegisterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.RegisterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *user.RegisterReq

func (p *RegisterArgs) GetReq() *user.RegisterReq {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *user.RegisterResp
}

var RegisterResult_Success_DEFAULT *user.RegisterResp

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.RegisterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(user.RegisterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *user.RegisterResp {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.RegisterResp)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.LoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *LoginArgs:
		success, err := handler.(user.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *user.LoginReq
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.LoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(user.LoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *user.LoginReq

func (p *LoginArgs) GetReq() *user.LoginReq {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *user.LoginResp
}

var LoginResult_Success_DEFAULT *user.LoginResp

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.LoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(user.LoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *user.LoginResp {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.LoginResp)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserArgs:
		success, err := handler.(user.UserService).GetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserArgs() interface{} {
	return &GetUserArgs{}
}

func newGetUserResult() interface{} {
	return &GetUserResult{}
}

type GetUserArgs struct {
	Req *user.GetUserReq
}

func (p *GetUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserArgs_Req_DEFAULT *user.GetUserReq

func (p *GetUserArgs) GetReq() *user.GetUserReq {
	if !p.IsSetReq() {
		return GetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserResult struct {
	Success *user.GetUserResp
}

var GetUserResult_Success_DEFAULT *user.GetUserResp

func (p *GetUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserResult) GetSuccess() *user.GetUserResp {
	if !p.IsSetSuccess() {
		return GetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserResp)
}

func (p *GetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserResult) GetResult() interface{} {
	return p.Success
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateUserArgs:
		success, err := handler.(user.UserService).UpdateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateUserArgs() interface{} {
	return &UpdateUserArgs{}
}

func newUpdateUserResult() interface{} {
	return &UpdateUserResult{}
}

type UpdateUserArgs struct {
	Req *user.UpdateUserReq
}

func (p *UpdateUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserArgs_Req_DEFAULT *user.UpdateUserReq

func (p *UpdateUserArgs) GetReq() *user.UpdateUserReq {
	if !p.IsSetReq() {
		return UpdateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserResult struct {
	Success *user.UpdateUserResp
}

var UpdateUserResult_Success_DEFAULT *user.UpdateUserResp

func (p *UpdateUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserResult) GetSuccess() *user.UpdateUserResp {
	if !p.IsSetSuccess() {
		return UpdateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateUserResp)
}

func (p *UpdateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserResult) GetResult() interface{} {
	return p.Success
}

func deleteUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DeleteUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).DeleteUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteUserArgs:
		success, err := handler.(user.UserService).DeleteUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteUserArgs() interface{} {
	return &DeleteUserArgs{}
}

func newDeleteUserResult() interface{} {
	return &DeleteUserResult{}
}

type DeleteUserArgs struct {
	Req *user.DeleteUserReq
}

func (p *DeleteUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DeleteUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteUserArgs) Unmarshal(in []byte) error {
	msg := new(user.DeleteUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteUserArgs_Req_DEFAULT *user.DeleteUserReq

func (p *DeleteUserArgs) GetReq() *user.DeleteUserReq {
	if !p.IsSetReq() {
		return DeleteUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteUserResult struct {
	Success *user.DeleteUserResp
}

var DeleteUserResult_Success_DEFAULT *user.DeleteUserResp

func (p *DeleteUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DeleteUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteUserResult) Unmarshal(in []byte) error {
	msg := new(user.DeleteUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUserResult) GetSuccess() *user.DeleteUserResp {
	if !p.IsSetSuccess() {
		return DeleteUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DeleteUserResp)
}

func (p *DeleteUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteUserResult) GetResult() interface{} {
	return p.Success
}

func getUserRoleByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserRoleByIdReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserRoleById(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserRoleByIdArgs:
		success, err := handler.(user.UserService).GetUserRoleById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserRoleByIdResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserRoleByIdArgs() interface{} {
	return &GetUserRoleByIdArgs{}
}

func newGetUserRoleByIdResult() interface{} {
	return &GetUserRoleByIdResult{}
}

type GetUserRoleByIdArgs struct {
	Req *user.GetUserRoleByIdReq
}

func (p *GetUserRoleByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserRoleByIdReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserRoleByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserRoleByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserRoleByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserRoleByIdArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserRoleByIdReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserRoleByIdArgs_Req_DEFAULT *user.GetUserRoleByIdReq

func (p *GetUserRoleByIdArgs) GetReq() *user.GetUserRoleByIdReq {
	if !p.IsSetReq() {
		return GetUserRoleByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserRoleByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserRoleByIdArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserRoleByIdResult struct {
	Success *user.GetUserRoleByIdResp
}

var GetUserRoleByIdResult_Success_DEFAULT *user.GetUserRoleByIdResp

func (p *GetUserRoleByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserRoleByIdResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserRoleByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserRoleByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserRoleByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserRoleByIdResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserRoleByIdResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserRoleByIdResult) GetSuccess() *user.GetUserRoleByIdResp {
	if !p.IsSetSuccess() {
		return GetUserRoleByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserRoleByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserRoleByIdResp)
}

func (p *GetUserRoleByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserRoleByIdResult) GetResult() interface{} {
	return p.Success
}

func addReceiveAddressHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.AddReceiveAddressReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).AddReceiveAddress(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddReceiveAddressArgs:
		success, err := handler.(user.UserService).AddReceiveAddress(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddReceiveAddressResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddReceiveAddressArgs() interface{} {
	return &AddReceiveAddressArgs{}
}

func newAddReceiveAddressResult() interface{} {
	return &AddReceiveAddressResult{}
}

type AddReceiveAddressArgs struct {
	Req *user.AddReceiveAddressReq
}

func (p *AddReceiveAddressArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.AddReceiveAddressReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddReceiveAddressArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddReceiveAddressArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddReceiveAddressArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddReceiveAddressArgs) Unmarshal(in []byte) error {
	msg := new(user.AddReceiveAddressReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddReceiveAddressArgs_Req_DEFAULT *user.AddReceiveAddressReq

func (p *AddReceiveAddressArgs) GetReq() *user.AddReceiveAddressReq {
	if !p.IsSetReq() {
		return AddReceiveAddressArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddReceiveAddressArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddReceiveAddressArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddReceiveAddressResult struct {
	Success *user.AddReceiveAddressResp
}

var AddReceiveAddressResult_Success_DEFAULT *user.AddReceiveAddressResp

func (p *AddReceiveAddressResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.AddReceiveAddressResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddReceiveAddressResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddReceiveAddressResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddReceiveAddressResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddReceiveAddressResult) Unmarshal(in []byte) error {
	msg := new(user.AddReceiveAddressResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddReceiveAddressResult) GetSuccess() *user.AddReceiveAddressResp {
	if !p.IsSetSuccess() {
		return AddReceiveAddressResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddReceiveAddressResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.AddReceiveAddressResp)
}

func (p *AddReceiveAddressResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddReceiveAddressResult) GetResult() interface{} {
	return p.Success
}

func getReceiveAddressHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetReceiveAddressReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetReceiveAddress(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetReceiveAddressArgs:
		success, err := handler.(user.UserService).GetReceiveAddress(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetReceiveAddressResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetReceiveAddressArgs() interface{} {
	return &GetReceiveAddressArgs{}
}

func newGetReceiveAddressResult() interface{} {
	return &GetReceiveAddressResult{}
}

type GetReceiveAddressArgs struct {
	Req *user.GetReceiveAddressReq
}

func (p *GetReceiveAddressArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetReceiveAddressReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetReceiveAddressArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetReceiveAddressArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetReceiveAddressArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetReceiveAddressArgs) Unmarshal(in []byte) error {
	msg := new(user.GetReceiveAddressReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetReceiveAddressArgs_Req_DEFAULT *user.GetReceiveAddressReq

func (p *GetReceiveAddressArgs) GetReq() *user.GetReceiveAddressReq {
	if !p.IsSetReq() {
		return GetReceiveAddressArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetReceiveAddressArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetReceiveAddressArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetReceiveAddressResult struct {
	Success *user.GetReceiveAddressResp
}

var GetReceiveAddressResult_Success_DEFAULT *user.GetReceiveAddressResp

func (p *GetReceiveAddressResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetReceiveAddressResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetReceiveAddressResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetReceiveAddressResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetReceiveAddressResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetReceiveAddressResult) Unmarshal(in []byte) error {
	msg := new(user.GetReceiveAddressResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetReceiveAddressResult) GetSuccess() *user.GetReceiveAddressResp {
	if !p.IsSetSuccess() {
		return GetReceiveAddressResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetReceiveAddressResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetReceiveAddressResp)
}

func (p *GetReceiveAddressResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetReceiveAddressResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, Req *user.RegisterReq) (r *user.RegisterResp, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *user.LoginReq) (r *user.LoginResp, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, Req *user.GetUserReq) (r *user.GetUserResp, err error) {
	var _args GetUserArgs
	_args.Req = Req
	var _result GetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, Req *user.UpdateUserReq) (r *user.UpdateUserResp, err error) {
	var _args UpdateUserArgs
	_args.Req = Req
	var _result UpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteUser(ctx context.Context, Req *user.DeleteUserReq) (r *user.DeleteUserResp, err error) {
	var _args DeleteUserArgs
	_args.Req = Req
	var _result DeleteUserResult
	if err = p.c.Call(ctx, "DeleteUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserRoleById(ctx context.Context, Req *user.GetUserRoleByIdReq) (r *user.GetUserRoleByIdResp, err error) {
	var _args GetUserRoleByIdArgs
	_args.Req = Req
	var _result GetUserRoleByIdResult
	if err = p.c.Call(ctx, "GetUserRoleById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddReceiveAddress(ctx context.Context, Req *user.AddReceiveAddressReq) (r *user.AddReceiveAddressResp, err error) {
	var _args AddReceiveAddressArgs
	_args.Req = Req
	var _result AddReceiveAddressResult
	if err = p.c.Call(ctx, "AddReceiveAddress", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetReceiveAddress(ctx context.Context, Req *user.GetReceiveAddressReq) (r *user.GetReceiveAddressResp, err error) {
	var _args GetReceiveAddressArgs
	_args.Req = Req
	var _result GetReceiveAddressResult
	if err = p.c.Call(ctx, "GetReceiveAddress", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
