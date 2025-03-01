// Code generated by Kitex v0.9.1. DO NOT EDIT.

package orderservice

import (
	"context"
	order "douyin_mall/order/kitex_gen/order"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PlaceOrder": kitex.NewMethodInfo(
		placeOrderHandler,
		newPlaceOrderArgs,
		newPlaceOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListOrder": kitex.NewMethodInfo(
		listOrderHandler,
		newListOrderArgs,
		newListOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"MarkOrderPaid": kitex.NewMethodInfo(
		markOrderPaidHandler,
		newMarkOrderPaidArgs,
		newMarkOrderPaidResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetOrder": kitex.NewMethodInfo(
		getOrderHandler,
		newGetOrderArgs,
		newGetOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SmartSearchOrder": kitex.NewMethodInfo(
		smartSearchOrderHandler,
		newSmartSearchOrderArgs,
		newSmartSearchOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SmartPlaceOrder": kitex.NewMethodInfo(
		smartPlaceOrderHandler,
		newSmartPlaceOrderArgs,
		newSmartPlaceOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	orderServiceServiceInfo                = NewServiceInfo()
	orderServiceServiceInfoForClient       = NewServiceInfoForClient()
	orderServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForClient
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
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
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
		"PackageName": "order",
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

func placeOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.PlaceOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).PlaceOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PlaceOrderArgs:
		success, err := handler.(order.OrderService).PlaceOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PlaceOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPlaceOrderArgs() interface{} {
	return &PlaceOrderArgs{}
}

func newPlaceOrderResult() interface{} {
	return &PlaceOrderResult{}
}

type PlaceOrderArgs struct {
	Req *order.PlaceOrderReq
}

func (p *PlaceOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.PlaceOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PlaceOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PlaceOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PlaceOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PlaceOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PlaceOrderArgs_Req_DEFAULT *order.PlaceOrderReq

func (p *PlaceOrderArgs) GetReq() *order.PlaceOrderReq {
	if !p.IsSetReq() {
		return PlaceOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PlaceOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PlaceOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PlaceOrderResult struct {
	Success *order.PlaceOrderResp
}

var PlaceOrderResult_Success_DEFAULT *order.PlaceOrderResp

func (p *PlaceOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.PlaceOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PlaceOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PlaceOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PlaceOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PlaceOrderResult) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PlaceOrderResult) GetSuccess() *order.PlaceOrderResp {
	if !p.IsSetSuccess() {
		return PlaceOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PlaceOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.PlaceOrderResp)
}

func (p *PlaceOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PlaceOrderResult) GetResult() interface{} {
	return p.Success
}

func listOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.ListOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).ListOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListOrderArgs:
		success, err := handler.(order.OrderService).ListOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListOrderArgs() interface{} {
	return &ListOrderArgs{}
}

func newListOrderResult() interface{} {
	return &ListOrderResult{}
}

type ListOrderArgs struct {
	Req *order.ListOrderReq
}

func (p *ListOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.ListOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.ListOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListOrderArgs_Req_DEFAULT *order.ListOrderReq

func (p *ListOrderArgs) GetReq() *order.ListOrderReq {
	if !p.IsSetReq() {
		return ListOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListOrderResult struct {
	Success *order.ListOrderResp
}

var ListOrderResult_Success_DEFAULT *order.ListOrderResp

func (p *ListOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.ListOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListOrderResult) Unmarshal(in []byte) error {
	msg := new(order.ListOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListOrderResult) GetSuccess() *order.ListOrderResp {
	if !p.IsSetSuccess() {
		return ListOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.ListOrderResp)
}

func (p *ListOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListOrderResult) GetResult() interface{} {
	return p.Success
}

func markOrderPaidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.MarkOrderPaidReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).MarkOrderPaid(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MarkOrderPaidArgs:
		success, err := handler.(order.OrderService).MarkOrderPaid(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MarkOrderPaidResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMarkOrderPaidArgs() interface{} {
	return &MarkOrderPaidArgs{}
}

func newMarkOrderPaidResult() interface{} {
	return &MarkOrderPaidResult{}
}

type MarkOrderPaidArgs struct {
	Req *order.MarkOrderPaidReq
}

func (p *MarkOrderPaidArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.MarkOrderPaidReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MarkOrderPaidArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MarkOrderPaidArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MarkOrderPaidArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MarkOrderPaidArgs) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderPaidReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MarkOrderPaidArgs_Req_DEFAULT *order.MarkOrderPaidReq

func (p *MarkOrderPaidArgs) GetReq() *order.MarkOrderPaidReq {
	if !p.IsSetReq() {
		return MarkOrderPaidArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MarkOrderPaidArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MarkOrderPaidArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MarkOrderPaidResult struct {
	Success *order.MarkOrderPaidResp
}

var MarkOrderPaidResult_Success_DEFAULT *order.MarkOrderPaidResp

func (p *MarkOrderPaidResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.MarkOrderPaidResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MarkOrderPaidResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MarkOrderPaidResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MarkOrderPaidResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MarkOrderPaidResult) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderPaidResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MarkOrderPaidResult) GetSuccess() *order.MarkOrderPaidResp {
	if !p.IsSetSuccess() {
		return MarkOrderPaidResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MarkOrderPaidResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.MarkOrderPaidResp)
}

func (p *MarkOrderPaidResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MarkOrderPaidResult) GetResult() interface{} {
	return p.Success
}

func getOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.GetOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).GetOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetOrderArgs:
		success, err := handler.(order.OrderService).GetOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetOrderArgs() interface{} {
	return &GetOrderArgs{}
}

func newGetOrderResult() interface{} {
	return &GetOrderResult{}
}

type GetOrderArgs struct {
	Req *order.GetOrderReq
}

func (p *GetOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.GetOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.GetOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetOrderArgs_Req_DEFAULT *order.GetOrderReq

func (p *GetOrderArgs) GetReq() *order.GetOrderReq {
	if !p.IsSetReq() {
		return GetOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetOrderResult struct {
	Success *order.GetOrderResp
}

var GetOrderResult_Success_DEFAULT *order.GetOrderResp

func (p *GetOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.GetOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetOrderResult) Unmarshal(in []byte) error {
	msg := new(order.GetOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetOrderResult) GetSuccess() *order.GetOrderResp {
	if !p.IsSetSuccess() {
		return GetOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.GetOrderResp)
}

func (p *GetOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetOrderResult) GetResult() interface{} {
	return p.Success
}

func smartSearchOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.SmartSearchOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).SmartSearchOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SmartSearchOrderArgs:
		success, err := handler.(order.OrderService).SmartSearchOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SmartSearchOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSmartSearchOrderArgs() interface{} {
	return &SmartSearchOrderArgs{}
}

func newSmartSearchOrderResult() interface{} {
	return &SmartSearchOrderResult{}
}

type SmartSearchOrderArgs struct {
	Req *order.SmartSearchOrderReq
}

func (p *SmartSearchOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.SmartSearchOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SmartSearchOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SmartSearchOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SmartSearchOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SmartSearchOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.SmartSearchOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SmartSearchOrderArgs_Req_DEFAULT *order.SmartSearchOrderReq

func (p *SmartSearchOrderArgs) GetReq() *order.SmartSearchOrderReq {
	if !p.IsSetReq() {
		return SmartSearchOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SmartSearchOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SmartSearchOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SmartSearchOrderResult struct {
	Success *order.SmartSearchOrderResp
}

var SmartSearchOrderResult_Success_DEFAULT *order.SmartSearchOrderResp

func (p *SmartSearchOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.SmartSearchOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SmartSearchOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SmartSearchOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SmartSearchOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SmartSearchOrderResult) Unmarshal(in []byte) error {
	msg := new(order.SmartSearchOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SmartSearchOrderResult) GetSuccess() *order.SmartSearchOrderResp {
	if !p.IsSetSuccess() {
		return SmartSearchOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SmartSearchOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.SmartSearchOrderResp)
}

func (p *SmartSearchOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SmartSearchOrderResult) GetResult() interface{} {
	return p.Success
}

func smartPlaceOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.SmartPlaceOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).SmartPlaceOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SmartPlaceOrderArgs:
		success, err := handler.(order.OrderService).SmartPlaceOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SmartPlaceOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSmartPlaceOrderArgs() interface{} {
	return &SmartPlaceOrderArgs{}
}

func newSmartPlaceOrderResult() interface{} {
	return &SmartPlaceOrderResult{}
}

type SmartPlaceOrderArgs struct {
	Req *order.SmartPlaceOrderReq
}

func (p *SmartPlaceOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.SmartPlaceOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SmartPlaceOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SmartPlaceOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SmartPlaceOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SmartPlaceOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.SmartPlaceOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SmartPlaceOrderArgs_Req_DEFAULT *order.SmartPlaceOrderReq

func (p *SmartPlaceOrderArgs) GetReq() *order.SmartPlaceOrderReq {
	if !p.IsSetReq() {
		return SmartPlaceOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SmartPlaceOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SmartPlaceOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SmartPlaceOrderResult struct {
	Success *order.SmartPlaceOrderResp
}

var SmartPlaceOrderResult_Success_DEFAULT *order.SmartPlaceOrderResp

func (p *SmartPlaceOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.SmartPlaceOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SmartPlaceOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SmartPlaceOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SmartPlaceOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SmartPlaceOrderResult) Unmarshal(in []byte) error {
	msg := new(order.SmartPlaceOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SmartPlaceOrderResult) GetSuccess() *order.SmartPlaceOrderResp {
	if !p.IsSetSuccess() {
		return SmartPlaceOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SmartPlaceOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.SmartPlaceOrderResp)
}

func (p *SmartPlaceOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SmartPlaceOrderResult) GetResult() interface{} {
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

func (p *kClient) PlaceOrder(ctx context.Context, Req *order.PlaceOrderReq) (r *order.PlaceOrderResp, err error) {
	var _args PlaceOrderArgs
	_args.Req = Req
	var _result PlaceOrderResult
	if err = p.c.Call(ctx, "PlaceOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListOrder(ctx context.Context, Req *order.ListOrderReq) (r *order.ListOrderResp, err error) {
	var _args ListOrderArgs
	_args.Req = Req
	var _result ListOrderResult
	if err = p.c.Call(ctx, "ListOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MarkOrderPaid(ctx context.Context, Req *order.MarkOrderPaidReq) (r *order.MarkOrderPaidResp, err error) {
	var _args MarkOrderPaidArgs
	_args.Req = Req
	var _result MarkOrderPaidResult
	if err = p.c.Call(ctx, "MarkOrderPaid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetOrder(ctx context.Context, Req *order.GetOrderReq) (r *order.GetOrderResp, err error) {
	var _args GetOrderArgs
	_args.Req = Req
	var _result GetOrderResult
	if err = p.c.Call(ctx, "GetOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SmartSearchOrder(ctx context.Context, Req *order.SmartSearchOrderReq) (r *order.SmartSearchOrderResp, err error) {
	var _args SmartSearchOrderArgs
	_args.Req = Req
	var _result SmartSearchOrderResult
	if err = p.c.Call(ctx, "SmartSearchOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SmartPlaceOrder(ctx context.Context, Req *order.SmartPlaceOrderReq) (r *order.SmartPlaceOrderResp, err error) {
	var _args SmartPlaceOrderArgs
	_args.Req = Req
	var _result SmartPlaceOrderResult
	if err = p.c.Call(ctx, "SmartPlaceOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
