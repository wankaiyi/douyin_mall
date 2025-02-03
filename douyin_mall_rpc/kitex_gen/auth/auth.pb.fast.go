// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package auth

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *DeliverTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeliverTokenReq[number], err)
}

func (x *DeliverTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *VerifyTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VerifyTokenReq[number], err)
}

func (x *VerifyTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AccessToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyTokenReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.RefreshToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeliveryResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeliveryResp[number], err)
}

func (x *DeliveryResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DeliveryResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeliveryResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AccessToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeliveryResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.RefreshToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VerifyResp[number], err)
}

func (x *VerifyResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *VerifyResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RefreshTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RefreshTokenReq[number], err)
}

func (x *RefreshTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.RefreshToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshTokenResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RefreshTokenResp[number], err)
}

func (x *RefreshTokenResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RefreshTokenResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshTokenResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AccessToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshTokenResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.RefreshToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RevokeTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RevokeTokenReq[number], err)
}

func (x *RevokeTokenReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RevokeResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RevokeResp[number], err)
}

func (x *RevokeResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RevokeResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeliverTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeliverTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *VerifyTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *VerifyTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.AccessToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAccessToken())
	return offset
}

func (x *VerifyTokenReq) fastWriteField2(buf []byte) (offset int) {
	if x.RefreshToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetRefreshToken())
	return offset
}

func (x *DeliveryResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *DeliveryResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *DeliveryResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *DeliveryResp) fastWriteField3(buf []byte) (offset int) {
	if x.AccessToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAccessToken())
	return offset
}

func (x *DeliveryResp) fastWriteField4(buf []byte) (offset int) {
	if x.RefreshToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetRefreshToken())
	return offset
}

func (x *VerifyResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *VerifyResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *VerifyResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *VerifyResp) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *RefreshTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RefreshTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.RefreshToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetRefreshToken())
	return offset
}

func (x *RefreshTokenResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *RefreshTokenResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *RefreshTokenResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *RefreshTokenResp) fastWriteField3(buf []byte) (offset int) {
	if x.AccessToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAccessToken())
	return offset
}

func (x *RefreshTokenResp) fastWriteField4(buf []byte) (offset int) {
	if x.RefreshToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetRefreshToken())
	return offset
}

func (x *RevokeTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RevokeTokenReq) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *RevokeResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RevokeResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *RevokeResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *DeliverTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeliverTokenReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *VerifyTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *VerifyTokenReq) sizeField1() (n int) {
	if x.AccessToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAccessToken())
	return n
}

func (x *VerifyTokenReq) sizeField2() (n int) {
	if x.RefreshToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetRefreshToken())
	return n
}

func (x *DeliveryResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *DeliveryResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *DeliveryResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *DeliveryResp) sizeField3() (n int) {
	if x.AccessToken == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAccessToken())
	return n
}

func (x *DeliveryResp) sizeField4() (n int) {
	if x.RefreshToken == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetRefreshToken())
	return n
}

func (x *VerifyResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *VerifyResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *VerifyResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *VerifyResp) sizeField3() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetUserId())
	return n
}

func (x *RefreshTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RefreshTokenReq) sizeField1() (n int) {
	if x.RefreshToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetRefreshToken())
	return n
}

func (x *RefreshTokenResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *RefreshTokenResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *RefreshTokenResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *RefreshTokenResp) sizeField3() (n int) {
	if x.AccessToken == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAccessToken())
	return n
}

func (x *RefreshTokenResp) sizeField4() (n int) {
	if x.RefreshToken == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetRefreshToken())
	return n
}

func (x *RevokeTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField2()
	return n
}

func (x *RevokeTokenReq) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *RevokeResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RevokeResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *RevokeResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

var fieldIDToName_DeliverTokenReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_VerifyTokenReq = map[int32]string{
	1: "AccessToken",
	2: "RefreshToken",
}

var fieldIDToName_DeliveryResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "AccessToken",
	4: "RefreshToken",
}

var fieldIDToName_VerifyResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserId",
}

var fieldIDToName_RefreshTokenReq = map[int32]string{
	1: "RefreshToken",
}

var fieldIDToName_RefreshTokenResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "AccessToken",
	4: "RefreshToken",
}

var fieldIDToName_RevokeTokenReq = map[int32]string{
	2: "UserId",
}

var fieldIDToName_RevokeResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}
