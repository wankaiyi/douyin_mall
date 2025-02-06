// Code generated by Kitex v0.9.1. DO NOT EDIT.

package paymentservice

import (
	"context"
	payment "douyin_mall/rpc/kitex_gen/payment"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Charge": kitex.NewMethodInfo(
		chargeHandler,
		newChargeArgs,
		newChargeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CancelCharge": kitex.NewMethodInfo(
		cancelChargeHandler,
		newCancelChargeArgs,
		newCancelChargeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"PaymentOrderRecord": kitex.NewMethodInfo(
		paymentOrderRecordHandler,
		newPaymentOrderRecordArgs,
		newPaymentOrderRecordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"PaymentTransactionRecord": kitex.NewMethodInfo(
		paymentTransactionRecordHandler,
		newPaymentTransactionRecordArgs,
		newPaymentTransactionRecordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"idempotentControl": kitex.NewMethodInfo(
		idempotentControlHandler,
		newIdempotentControlArgs,
		newIdempotentControlResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	paymentServiceServiceInfo                = NewServiceInfo()
	paymentServiceServiceInfoForClient       = NewServiceInfoForClient()
	paymentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return paymentServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return paymentServiceServiceInfoForClient
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
	serviceName := "PaymentService"
	handlerType := (*payment.PaymentService)(nil)
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
		"PackageName": "payment",
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

func chargeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.ChargeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).Charge(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ChargeArgs:
		success, err := handler.(payment.PaymentService).Charge(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChargeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newChargeArgs() interface{} {
	return &ChargeArgs{}
}

func newChargeResult() interface{} {
	return &ChargeResult{}
}

type ChargeArgs struct {
	Req *payment.ChargeReq
}

func (p *ChargeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.ChargeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChargeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChargeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChargeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ChargeArgs) Unmarshal(in []byte) error {
	msg := new(payment.ChargeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChargeArgs_Req_DEFAULT *payment.ChargeReq

func (p *ChargeArgs) GetReq() *payment.ChargeReq {
	if !p.IsSetReq() {
		return ChargeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChargeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChargeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChargeResult struct {
	Success *payment.ChargeResp
}

var ChargeResult_Success_DEFAULT *payment.ChargeResp

func (p *ChargeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.ChargeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChargeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChargeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChargeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ChargeResult) Unmarshal(in []byte) error {
	msg := new(payment.ChargeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChargeResult) GetSuccess() *payment.ChargeResp {
	if !p.IsSetSuccess() {
		return ChargeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChargeResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.ChargeResp)
}

func (p *ChargeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChargeResult) GetResult() interface{} {
	return p.Success
}

func cancelChargeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.CancelChargeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).CancelCharge(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CancelChargeArgs:
		success, err := handler.(payment.PaymentService).CancelCharge(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CancelChargeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCancelChargeArgs() interface{} {
	return &CancelChargeArgs{}
}

func newCancelChargeResult() interface{} {
	return &CancelChargeResult{}
}

type CancelChargeArgs struct {
	Req *payment.CancelChargeReq
}

func (p *CancelChargeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.CancelChargeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CancelChargeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CancelChargeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CancelChargeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CancelChargeArgs) Unmarshal(in []byte) error {
	msg := new(payment.CancelChargeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CancelChargeArgs_Req_DEFAULT *payment.CancelChargeReq

func (p *CancelChargeArgs) GetReq() *payment.CancelChargeReq {
	if !p.IsSetReq() {
		return CancelChargeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CancelChargeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CancelChargeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CancelChargeResult struct {
	Success *payment.CancelChargeResp
}

var CancelChargeResult_Success_DEFAULT *payment.CancelChargeResp

func (p *CancelChargeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.CancelChargeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CancelChargeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CancelChargeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CancelChargeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CancelChargeResult) Unmarshal(in []byte) error {
	msg := new(payment.CancelChargeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CancelChargeResult) GetSuccess() *payment.CancelChargeResp {
	if !p.IsSetSuccess() {
		return CancelChargeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CancelChargeResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.CancelChargeResp)
}

func (p *CancelChargeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CancelChargeResult) GetResult() interface{} {
	return p.Success
}

func paymentOrderRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.PaymentOrderRecordReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).PaymentOrderRecord(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PaymentOrderRecordArgs:
		success, err := handler.(payment.PaymentService).PaymentOrderRecord(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PaymentOrderRecordResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPaymentOrderRecordArgs() interface{} {
	return &PaymentOrderRecordArgs{}
}

func newPaymentOrderRecordResult() interface{} {
	return &PaymentOrderRecordResult{}
}

type PaymentOrderRecordArgs struct {
	Req *payment.PaymentOrderRecordReq
}

func (p *PaymentOrderRecordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.PaymentOrderRecordReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PaymentOrderRecordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PaymentOrderRecordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PaymentOrderRecordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PaymentOrderRecordArgs) Unmarshal(in []byte) error {
	msg := new(payment.PaymentOrderRecordReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PaymentOrderRecordArgs_Req_DEFAULT *payment.PaymentOrderRecordReq

func (p *PaymentOrderRecordArgs) GetReq() *payment.PaymentOrderRecordReq {
	if !p.IsSetReq() {
		return PaymentOrderRecordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PaymentOrderRecordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PaymentOrderRecordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PaymentOrderRecordResult struct {
	Success *payment.PaymentOrderRecordResp
}

var PaymentOrderRecordResult_Success_DEFAULT *payment.PaymentOrderRecordResp

func (p *PaymentOrderRecordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.PaymentOrderRecordResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PaymentOrderRecordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PaymentOrderRecordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PaymentOrderRecordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PaymentOrderRecordResult) Unmarshal(in []byte) error {
	msg := new(payment.PaymentOrderRecordResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PaymentOrderRecordResult) GetSuccess() *payment.PaymentOrderRecordResp {
	if !p.IsSetSuccess() {
		return PaymentOrderRecordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PaymentOrderRecordResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.PaymentOrderRecordResp)
}

func (p *PaymentOrderRecordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentOrderRecordResult) GetResult() interface{} {
	return p.Success
}

func paymentTransactionRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.PaymentTransactionRecordReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).PaymentTransactionRecord(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PaymentTransactionRecordArgs:
		success, err := handler.(payment.PaymentService).PaymentTransactionRecord(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PaymentTransactionRecordResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPaymentTransactionRecordArgs() interface{} {
	return &PaymentTransactionRecordArgs{}
}

func newPaymentTransactionRecordResult() interface{} {
	return &PaymentTransactionRecordResult{}
}

type PaymentTransactionRecordArgs struct {
	Req *payment.PaymentTransactionRecordReq
}

func (p *PaymentTransactionRecordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.PaymentTransactionRecordReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PaymentTransactionRecordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PaymentTransactionRecordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PaymentTransactionRecordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PaymentTransactionRecordArgs) Unmarshal(in []byte) error {
	msg := new(payment.PaymentTransactionRecordReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PaymentTransactionRecordArgs_Req_DEFAULT *payment.PaymentTransactionRecordReq

func (p *PaymentTransactionRecordArgs) GetReq() *payment.PaymentTransactionRecordReq {
	if !p.IsSetReq() {
		return PaymentTransactionRecordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PaymentTransactionRecordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PaymentTransactionRecordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PaymentTransactionRecordResult struct {
	Success *payment.PaymentTransactionRecordResp
}

var PaymentTransactionRecordResult_Success_DEFAULT *payment.PaymentTransactionRecordResp

func (p *PaymentTransactionRecordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.PaymentTransactionRecordResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PaymentTransactionRecordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PaymentTransactionRecordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PaymentTransactionRecordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PaymentTransactionRecordResult) Unmarshal(in []byte) error {
	msg := new(payment.PaymentTransactionRecordResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PaymentTransactionRecordResult) GetSuccess() *payment.PaymentTransactionRecordResp {
	if !p.IsSetSuccess() {
		return PaymentTransactionRecordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PaymentTransactionRecordResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.PaymentTransactionRecordResp)
}

func (p *PaymentTransactionRecordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentTransactionRecordResult) GetResult() interface{} {
	return p.Success
}

func idempotentControlHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.IdempotentControlReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).IdempotentControl(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *IdempotentControlArgs:
		success, err := handler.(payment.PaymentService).IdempotentControl(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IdempotentControlResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newIdempotentControlArgs() interface{} {
	return &IdempotentControlArgs{}
}

func newIdempotentControlResult() interface{} {
	return &IdempotentControlResult{}
}

type IdempotentControlArgs struct {
	Req *payment.IdempotentControlReq
}

func (p *IdempotentControlArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.IdempotentControlReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IdempotentControlArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IdempotentControlArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IdempotentControlArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *IdempotentControlArgs) Unmarshal(in []byte) error {
	msg := new(payment.IdempotentControlReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IdempotentControlArgs_Req_DEFAULT *payment.IdempotentControlReq

func (p *IdempotentControlArgs) GetReq() *payment.IdempotentControlReq {
	if !p.IsSetReq() {
		return IdempotentControlArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IdempotentControlArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *IdempotentControlArgs) GetFirstArgument() interface{} {
	return p.Req
}

type IdempotentControlResult struct {
	Success *payment.IdempotentControlResp
}

var IdempotentControlResult_Success_DEFAULT *payment.IdempotentControlResp

func (p *IdempotentControlResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.IdempotentControlResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IdempotentControlResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IdempotentControlResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IdempotentControlResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *IdempotentControlResult) Unmarshal(in []byte) error {
	msg := new(payment.IdempotentControlResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IdempotentControlResult) GetSuccess() *payment.IdempotentControlResp {
	if !p.IsSetSuccess() {
		return IdempotentControlResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IdempotentControlResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.IdempotentControlResp)
}

func (p *IdempotentControlResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IdempotentControlResult) GetResult() interface{} {
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

func (p *kClient) Charge(ctx context.Context, Req *payment.ChargeReq) (r *payment.ChargeResp, err error) {
	var _args ChargeArgs
	_args.Req = Req
	var _result ChargeResult
	if err = p.c.Call(ctx, "Charge", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CancelCharge(ctx context.Context, Req *payment.CancelChargeReq) (r *payment.CancelChargeResp, err error) {
	var _args CancelChargeArgs
	_args.Req = Req
	var _result CancelChargeResult
	if err = p.c.Call(ctx, "CancelCharge", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PaymentOrderRecord(ctx context.Context, Req *payment.PaymentOrderRecordReq) (r *payment.PaymentOrderRecordResp, err error) {
	var _args PaymentOrderRecordArgs
	_args.Req = Req
	var _result PaymentOrderRecordResult
	if err = p.c.Call(ctx, "PaymentOrderRecord", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PaymentTransactionRecord(ctx context.Context, Req *payment.PaymentTransactionRecordReq) (r *payment.PaymentTransactionRecordResp, err error) {
	var _args PaymentTransactionRecordArgs
	_args.Req = Req
	var _result PaymentTransactionRecordResult
	if err = p.c.Call(ctx, "PaymentTransactionRecord", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IdempotentControl(ctx context.Context, Req *payment.IdempotentControlReq) (r *payment.IdempotentControlResp, err error) {
	var _args IdempotentControlArgs
	_args.Req = Req
	var _result IdempotentControlResult
	if err = p.c.Call(ctx, "idempotentControl", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
