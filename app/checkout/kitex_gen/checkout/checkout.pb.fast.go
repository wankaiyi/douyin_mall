// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package checkout

import (
	user "douyin_mall/checkout/kitex_gen/user"
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *CheckoutReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckoutReq[number], err)
}

func (x *CheckoutReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CheckoutReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v user.AddReceiveAddressReq
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Address = &v
	return offset, nil
}

func (x *CheckoutResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckoutResp[number], err)
}

func (x *CheckoutResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CheckoutResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckoutResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PaymentUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckoutProductItemsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckoutProductItemsReq[number], err)
}

func (x *CheckoutProductItemsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CheckoutProductItemsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v ProductItem
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Items = append(x.Items, &v)
	return offset, nil
}

func (x *CheckoutProductItemsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AddressId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ProductItem) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductItem[number], err)
}

func (x *ProductItem) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ProductId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ProductItem) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Quantity, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CheckoutProductItemsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckoutProductItemsResp[number], err)
}

func (x *CheckoutProductItemsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CheckoutProductItemsResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckoutProductItemsResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PaymentUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckoutReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CheckoutReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CheckoutReq) fastWriteField2(buf []byte) (offset int) {
	if x.Address == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetAddress())
	return offset
}

func (x *CheckoutResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CheckoutResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *CheckoutResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *CheckoutResp) fastWriteField3(buf []byte) (offset int) {
	if x.PaymentUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPaymentUrl())
	return offset
}

func (x *CheckoutProductItemsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CheckoutProductItemsReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CheckoutProductItemsReq) fastWriteField2(buf []byte) (offset int) {
	if x.Items == nil {
		return offset
	}
	for i := range x.GetItems() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetItems()[i])
	}
	return offset
}

func (x *CheckoutProductItemsReq) fastWriteField3(buf []byte) (offset int) {
	if x.AddressId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetAddressId())
	return offset
}

func (x *ProductItem) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ProductItem) fastWriteField1(buf []byte) (offset int) {
	if x.ProductId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetProductId())
	return offset
}

func (x *ProductItem) fastWriteField2(buf []byte) (offset int) {
	if x.Quantity == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetQuantity())
	return offset
}

func (x *CheckoutProductItemsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CheckoutProductItemsResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *CheckoutProductItemsResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *CheckoutProductItemsResp) fastWriteField3(buf []byte) (offset int) {
	if x.PaymentUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPaymentUrl())
	return offset
}

func (x *CheckoutReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CheckoutReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *CheckoutReq) sizeField2() (n int) {
	if x.Address == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetAddress())
	return n
}

func (x *CheckoutResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CheckoutResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *CheckoutResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *CheckoutResp) sizeField3() (n int) {
	if x.PaymentUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPaymentUrl())
	return n
}

func (x *CheckoutProductItemsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CheckoutProductItemsReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *CheckoutProductItemsReq) sizeField2() (n int) {
	if x.Items == nil {
		return n
	}
	for i := range x.GetItems() {
		n += fastpb.SizeMessage(2, x.GetItems()[i])
	}
	return n
}

func (x *CheckoutProductItemsReq) sizeField3() (n int) {
	if x.AddressId == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetAddressId())
	return n
}

func (x *ProductItem) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ProductItem) sizeField1() (n int) {
	if x.ProductId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetProductId())
	return n
}

func (x *ProductItem) sizeField2() (n int) {
	if x.Quantity == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetQuantity())
	return n
}

func (x *CheckoutProductItemsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CheckoutProductItemsResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *CheckoutProductItemsResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *CheckoutProductItemsResp) sizeField3() (n int) {
	if x.PaymentUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPaymentUrl())
	return n
}

var fieldIDToName_CheckoutReq = map[int32]string{
	1: "UserId",
	2: "Address",
}

var fieldIDToName_CheckoutResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "PaymentUrl",
}

var fieldIDToName_CheckoutProductItemsReq = map[int32]string{
	1: "UserId",
	2: "Items",
	3: "AddressId",
}

var fieldIDToName_ProductItem = map[int32]string{
	1: "ProductId",
	2: "Quantity",
}

var fieldIDToName_CheckoutProductItemsResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "PaymentUrl",
}

var _ = user.File_user_proto
