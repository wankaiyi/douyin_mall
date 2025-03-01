// Code generated by Kitex v0.9.1. DO NOT EDIT.

package checkoutservice

import (
	"context"
	checkout "douyin_mall/checkout/kitex_gen/checkout"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Checkout": kitex.NewMethodInfo(
		checkoutHandler,
		newCheckoutArgs,
		newCheckoutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CheckoutProductItems": kitex.NewMethodInfo(
		checkoutProductItemsHandler,
		newCheckoutProductItemsArgs,
		newCheckoutProductItemsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	checkoutServiceServiceInfo                = NewServiceInfo()
	checkoutServiceServiceInfoForClient       = NewServiceInfoForClient()
	checkoutServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return checkoutServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return checkoutServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return checkoutServiceServiceInfoForClient
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
	serviceName := "CheckoutService"
	handlerType := (*checkout.CheckoutService)(nil)
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
		"PackageName": "checkout",
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

func checkoutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(checkout.CheckoutReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(checkout.CheckoutService).Checkout(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CheckoutArgs:
		success, err := handler.(checkout.CheckoutService).Checkout(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckoutResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCheckoutArgs() interface{} {
	return &CheckoutArgs{}
}

func newCheckoutResult() interface{} {
	return &CheckoutResult{}
}

type CheckoutArgs struct {
	Req *checkout.CheckoutReq
}

func (p *CheckoutArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(checkout.CheckoutReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckoutArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckoutArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckoutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CheckoutArgs) Unmarshal(in []byte) error {
	msg := new(checkout.CheckoutReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckoutArgs_Req_DEFAULT *checkout.CheckoutReq

func (p *CheckoutArgs) GetReq() *checkout.CheckoutReq {
	if !p.IsSetReq() {
		return CheckoutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckoutArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CheckoutArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CheckoutResult struct {
	Success *checkout.CheckoutResp
}

var CheckoutResult_Success_DEFAULT *checkout.CheckoutResp

func (p *CheckoutResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(checkout.CheckoutResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckoutResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckoutResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckoutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CheckoutResult) Unmarshal(in []byte) error {
	msg := new(checkout.CheckoutResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckoutResult) GetSuccess() *checkout.CheckoutResp {
	if !p.IsSetSuccess() {
		return CheckoutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckoutResult) SetSuccess(x interface{}) {
	p.Success = x.(*checkout.CheckoutResp)
}

func (p *CheckoutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckoutResult) GetResult() interface{} {
	return p.Success
}

func checkoutProductItemsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(checkout.CheckoutProductItemsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(checkout.CheckoutService).CheckoutProductItems(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CheckoutProductItemsArgs:
		success, err := handler.(checkout.CheckoutService).CheckoutProductItems(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckoutProductItemsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCheckoutProductItemsArgs() interface{} {
	return &CheckoutProductItemsArgs{}
}

func newCheckoutProductItemsResult() interface{} {
	return &CheckoutProductItemsResult{}
}

type CheckoutProductItemsArgs struct {
	Req *checkout.CheckoutProductItemsReq
}

func (p *CheckoutProductItemsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(checkout.CheckoutProductItemsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckoutProductItemsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckoutProductItemsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckoutProductItemsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CheckoutProductItemsArgs) Unmarshal(in []byte) error {
	msg := new(checkout.CheckoutProductItemsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckoutProductItemsArgs_Req_DEFAULT *checkout.CheckoutProductItemsReq

func (p *CheckoutProductItemsArgs) GetReq() *checkout.CheckoutProductItemsReq {
	if !p.IsSetReq() {
		return CheckoutProductItemsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckoutProductItemsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CheckoutProductItemsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CheckoutProductItemsResult struct {
	Success *checkout.CheckoutProductItemsResp
}

var CheckoutProductItemsResult_Success_DEFAULT *checkout.CheckoutProductItemsResp

func (p *CheckoutProductItemsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(checkout.CheckoutProductItemsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckoutProductItemsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckoutProductItemsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckoutProductItemsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CheckoutProductItemsResult) Unmarshal(in []byte) error {
	msg := new(checkout.CheckoutProductItemsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckoutProductItemsResult) GetSuccess() *checkout.CheckoutProductItemsResp {
	if !p.IsSetSuccess() {
		return CheckoutProductItemsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckoutProductItemsResult) SetSuccess(x interface{}) {
	p.Success = x.(*checkout.CheckoutProductItemsResp)
}

func (p *CheckoutProductItemsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckoutProductItemsResult) GetResult() interface{} {
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

func (p *kClient) Checkout(ctx context.Context, Req *checkout.CheckoutReq) (r *checkout.CheckoutResp, err error) {
	var _args CheckoutArgs
	_args.Req = Req
	var _result CheckoutResult
	if err = p.c.Call(ctx, "Checkout", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckoutProductItems(ctx context.Context, Req *checkout.CheckoutProductItemsReq) (r *checkout.CheckoutProductItemsResp, err error) {
	var _args CheckoutProductItemsArgs
	_args.Req = Req
	var _result CheckoutProductItemsResult
	if err = p.c.Call(ctx, "CheckoutProductItems", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
