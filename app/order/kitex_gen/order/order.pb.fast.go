// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package order

import (
	cart "douyin_mall/order/kitex_gen/cart"
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Address) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Address[number], err)
}

func (x *Address) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PhoneNumber, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Province, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.City, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Region, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.DetailAddress, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Product[number], err)
}

func (x *Product) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Product) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Picture, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Price, offset, err = fastpb.ReadDouble(buf, _type)
	return offset, err
}

func (x *Product) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Quantity, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PlaceOrderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PlaceOrderReq[number], err)
}

func (x *PlaceOrderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PlaceOrderReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Address
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Address = &v
	return offset, nil
}

func (x *PlaceOrderReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v OrderItem
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.OrderItems = append(x.OrderItems, &v)
	return offset, nil
}

func (x *PlaceOrderReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.TotalCost, offset, err = fastpb.ReadDouble(buf, _type)
	return offset, err
}

func (x *OrderItem) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderItem[number], err)
}

func (x *OrderItem) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v cart.CartItem
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Item = &v
	return offset, nil
}

func (x *OrderItem) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Cost, offset, err = fastpb.ReadDouble(buf, _type)
	return offset, err
}

func (x *OrderResult) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderResult[number], err)
}

func (x *OrderResult) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PlaceOrderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PlaceOrderResp[number], err)
}

func (x *PlaceOrderResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v OrderResult
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Order = &v
	return offset, nil
}

func (x *PlaceOrderResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PlaceOrderResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ListOrderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListOrderReq[number], err)
}

func (x *ListOrderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Order) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Order[number], err)
}

func (x *Order) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Order) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Address
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Address = &v
	return offset, nil
}

func (x *Order) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Product
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Products = append(x.Products, &v)
	return offset, nil
}

func (x *Order) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Cost, offset, err = fastpb.ReadDouble(buf, _type)
	return offset, err
}

func (x *Order) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CreatedAt, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Order) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Status, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListOrderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListOrderResp[number], err)
}

func (x *ListOrderResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Order
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Orders = append(x.Orders, &v)
	return offset, nil
}

func (x *ListOrderResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListOrderResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MarkOrderPaidReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MarkOrderPaidReq[number], err)
}

func (x *MarkOrderPaidReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *MarkOrderPaidReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MarkOrderPaidResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MarkOrderPaidResp[number], err)
}

func (x *MarkOrderPaidResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *MarkOrderPaidResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *Address) fastWriteField1(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetName())
	return offset
}

func (x *Address) fastWriteField2(buf []byte) (offset int) {
	if x.PhoneNumber == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPhoneNumber())
	return offset
}

func (x *Address) fastWriteField3(buf []byte) (offset int) {
	if x.Province == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetProvince())
	return offset
}

func (x *Address) fastWriteField4(buf []byte) (offset int) {
	if x.City == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCity())
	return offset
}

func (x *Address) fastWriteField5(buf []byte) (offset int) {
	if x.Region == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetRegion())
	return offset
}

func (x *Address) fastWriteField6(buf []byte) (offset int) {
	if x.DetailAddress == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetDetailAddress())
	return offset
}

func (x *Product) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *Product) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Product) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *Product) fastWriteField3(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *Product) fastWriteField4(buf []byte) (offset int) {
	if x.Picture == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPicture())
	return offset
}

func (x *Product) fastWriteField5(buf []byte) (offset int) {
	if x.Price == 0 {
		return offset
	}
	offset += fastpb.WriteDouble(buf[offset:], 5, x.GetPrice())
	return offset
}

func (x *Product) fastWriteField6(buf []byte) (offset int) {
	if x.Quantity == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.GetQuantity())
	return offset
}

func (x *PlaceOrderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *PlaceOrderReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *PlaceOrderReq) fastWriteField2(buf []byte) (offset int) {
	if x.Address == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetAddress())
	return offset
}

func (x *PlaceOrderReq) fastWriteField3(buf []byte) (offset int) {
	if x.OrderItems == nil {
		return offset
	}
	for i := range x.GetOrderItems() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetOrderItems()[i])
	}
	return offset
}

func (x *PlaceOrderReq) fastWriteField4(buf []byte) (offset int) {
	if x.TotalCost == 0 {
		return offset
	}
	offset += fastpb.WriteDouble(buf[offset:], 4, x.GetTotalCost())
	return offset
}

func (x *OrderItem) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *OrderItem) fastWriteField1(buf []byte) (offset int) {
	if x.Item == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetItem())
	return offset
}

func (x *OrderItem) fastWriteField2(buf []byte) (offset int) {
	if x.Cost == 0 {
		return offset
	}
	offset += fastpb.WriteDouble(buf[offset:], 2, x.GetCost())
	return offset
}

func (x *OrderResult) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *OrderResult) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *PlaceOrderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *PlaceOrderResp) fastWriteField1(buf []byte) (offset int) {
	if x.Order == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetOrder())
	return offset
}

func (x *PlaceOrderResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetStatusCode())
	return offset
}

func (x *PlaceOrderResp) fastWriteField3(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetStatusMsg())
	return offset
}

func (x *ListOrderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ListOrderReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *Order) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *Order) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *Order) fastWriteField2(buf []byte) (offset int) {
	if x.Address == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetAddress())
	return offset
}

func (x *Order) fastWriteField3(buf []byte) (offset int) {
	if x.Products == nil {
		return offset
	}
	for i := range x.GetProducts() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetProducts()[i])
	}
	return offset
}

func (x *Order) fastWriteField4(buf []byte) (offset int) {
	if x.Cost == 0 {
		return offset
	}
	offset += fastpb.WriteDouble(buf[offset:], 4, x.GetCost())
	return offset
}

func (x *Order) fastWriteField5(buf []byte) (offset int) {
	if x.CreatedAt == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetCreatedAt())
	return offset
}

func (x *Order) fastWriteField6(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.GetStatus())
	return offset
}

func (x *ListOrderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ListOrderResp) fastWriteField1(buf []byte) (offset int) {
	if x.Orders == nil {
		return offset
	}
	for i := range x.GetOrders() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetOrders()[i])
	}
	return offset
}

func (x *ListOrderResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetStatusCode())
	return offset
}

func (x *ListOrderResp) fastWriteField3(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetStatusMsg())
	return offset
}

func (x *MarkOrderPaidReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *MarkOrderPaidReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *MarkOrderPaidReq) fastWriteField2(buf []byte) (offset int) {
	if x.OrderId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetOrderId())
	return offset
}

func (x *MarkOrderPaidResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *MarkOrderPaidResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *MarkOrderPaidResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *Address) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *Address) sizeField1() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetName())
	return n
}

func (x *Address) sizeField2() (n int) {
	if x.PhoneNumber == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPhoneNumber())
	return n
}

func (x *Address) sizeField3() (n int) {
	if x.Province == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetProvince())
	return n
}

func (x *Address) sizeField4() (n int) {
	if x.City == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetCity())
	return n
}

func (x *Address) sizeField5() (n int) {
	if x.Region == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetRegion())
	return n
}

func (x *Address) sizeField6() (n int) {
	if x.DetailAddress == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetDetailAddress())
	return n
}

func (x *Product) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *Product) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetId())
	return n
}

func (x *Product) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *Product) sizeField3() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *Product) sizeField4() (n int) {
	if x.Picture == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPicture())
	return n
}

func (x *Product) sizeField5() (n int) {
	if x.Price == 0 {
		return n
	}
	n += fastpb.SizeDouble(5, x.GetPrice())
	return n
}

func (x *Product) sizeField6() (n int) {
	if x.Quantity == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.GetQuantity())
	return n
}

func (x *PlaceOrderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *PlaceOrderReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *PlaceOrderReq) sizeField2() (n int) {
	if x.Address == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetAddress())
	return n
}

func (x *PlaceOrderReq) sizeField3() (n int) {
	if x.OrderItems == nil {
		return n
	}
	for i := range x.GetOrderItems() {
		n += fastpb.SizeMessage(3, x.GetOrderItems()[i])
	}
	return n
}

func (x *PlaceOrderReq) sizeField4() (n int) {
	if x.TotalCost == 0 {
		return n
	}
	n += fastpb.SizeDouble(4, x.GetTotalCost())
	return n
}

func (x *OrderItem) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *OrderItem) sizeField1() (n int) {
	if x.Item == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetItem())
	return n
}

func (x *OrderItem) sizeField2() (n int) {
	if x.Cost == 0 {
		return n
	}
	n += fastpb.SizeDouble(2, x.GetCost())
	return n
}

func (x *OrderResult) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *OrderResult) sizeField1() (n int) {
	if x.OrderId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetOrderId())
	return n
}

func (x *PlaceOrderResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *PlaceOrderResp) sizeField1() (n int) {
	if x.Order == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetOrder())
	return n
}

func (x *PlaceOrderResp) sizeField2() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetStatusCode())
	return n
}

func (x *PlaceOrderResp) sizeField3() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetStatusMsg())
	return n
}

func (x *ListOrderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ListOrderReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *Order) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *Order) sizeField1() (n int) {
	if x.OrderId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetOrderId())
	return n
}

func (x *Order) sizeField2() (n int) {
	if x.Address == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetAddress())
	return n
}

func (x *Order) sizeField3() (n int) {
	if x.Products == nil {
		return n
	}
	for i := range x.GetProducts() {
		n += fastpb.SizeMessage(3, x.GetProducts()[i])
	}
	return n
}

func (x *Order) sizeField4() (n int) {
	if x.Cost == 0 {
		return n
	}
	n += fastpb.SizeDouble(4, x.GetCost())
	return n
}

func (x *Order) sizeField5() (n int) {
	if x.CreatedAt == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetCreatedAt())
	return n
}

func (x *Order) sizeField6() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.GetStatus())
	return n
}

func (x *ListOrderResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ListOrderResp) sizeField1() (n int) {
	if x.Orders == nil {
		return n
	}
	for i := range x.GetOrders() {
		n += fastpb.SizeMessage(1, x.GetOrders()[i])
	}
	return n
}

func (x *ListOrderResp) sizeField2() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetStatusCode())
	return n
}

func (x *ListOrderResp) sizeField3() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetStatusMsg())
	return n
}

func (x *MarkOrderPaidReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *MarkOrderPaidReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *MarkOrderPaidReq) sizeField2() (n int) {
	if x.OrderId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetOrderId())
	return n
}

func (x *MarkOrderPaidResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *MarkOrderPaidResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *MarkOrderPaidResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

var fieldIDToName_Address = map[int32]string{
	1: "Name",
	2: "PhoneNumber",
	3: "Province",
	4: "City",
	5: "Region",
	6: "DetailAddress",
}

var fieldIDToName_Product = map[int32]string{
	1: "Id",
	2: "Name",
	3: "Description",
	4: "Picture",
	5: "Price",
	6: "Quantity",
}

var fieldIDToName_PlaceOrderReq = map[int32]string{
	1: "UserId",
	2: "Address",
	3: "OrderItems",
	4: "TotalCost",
}

var fieldIDToName_OrderItem = map[int32]string{
	1: "Item",
	2: "Cost",
}

var fieldIDToName_OrderResult = map[int32]string{
	1: "OrderId",
}

var fieldIDToName_PlaceOrderResp = map[int32]string{
	1: "Order",
	2: "StatusCode",
	3: "StatusMsg",
}

var fieldIDToName_ListOrderReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_Order = map[int32]string{
	1: "OrderId",
	2: "Address",
	3: "Products",
	4: "Cost",
	5: "CreatedAt",
	6: "Status",
}

var fieldIDToName_ListOrderResp = map[int32]string{
	1: "Orders",
	2: "StatusCode",
	3: "StatusMsg",
}

var fieldIDToName_MarkOrderPaidReq = map[int32]string{
	1: "UserId",
	2: "OrderId",
}

var fieldIDToName_MarkOrderPaidResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var _ = cart.File_cart_proto
