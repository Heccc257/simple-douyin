// Code generated by thriftgo (0.2.5). DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"simple_douyin/biz/model/common"
)

type DouyinUserRegisterRequest struct {
	UserName string `thrift:"UserName,1" json:"UserName" query:"user_name"`
	PassWord string `thrift:"PassWord,2" json:"PassWord" query:"password"`
}

func NewDouyinUserRegisterRequest() *DouyinUserRegisterRequest {
	return &DouyinUserRegisterRequest{}
}

func (p *DouyinUserRegisterRequest) GetUserName() (v string) {
	return p.UserName
}

func (p *DouyinUserRegisterRequest) GetPassWord() (v string) {
	return p.PassWord
}

var fieldIDToName_DouyinUserRegisterRequest = map[int16]string{
	1: "UserName",
	2: "PassWord",
}

func (p *DouyinUserRegisterRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DouyinUserRegisterRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DouyinUserRegisterRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.UserName = v
	}
	return nil
}

func (p *DouyinUserRegisterRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PassWord = v
	}
	return nil
}

func (p *DouyinUserRegisterRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("DouyinUserRegisterRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DouyinUserRegisterRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("UserName", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UserName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DouyinUserRegisterRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("PassWord", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PassWord); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DouyinUserRegisterRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DouyinUserRegisterRequest(%+v)", *p)
}

type DouyinUserRegisterResponse struct {
	StatusCode int32  `thrift:"StatusCode,1" form:"status_code" json:"status_code"`
	StatusMsg  string `thrift:"StatusMsg,2" form:"status_msg" json:"status_msg"`
	UserID     int64  `thrift:"UserID,3" form:"user_id" json:"user_id"`
	Token      string `thrift:"Token,4" form:"token" json:"token"`
}

func NewDouyinUserRegisterResponse() *DouyinUserRegisterResponse {
	return &DouyinUserRegisterResponse{}
}

func (p *DouyinUserRegisterResponse) GetStatusCode() (v int32) {
	return p.StatusCode
}

func (p *DouyinUserRegisterResponse) GetStatusMsg() (v string) {
	return p.StatusMsg
}

func (p *DouyinUserRegisterResponse) GetUserID() (v int64) {
	return p.UserID
}

func (p *DouyinUserRegisterResponse) GetToken() (v string) {
	return p.Token
}

var fieldIDToName_DouyinUserRegisterResponse = map[int16]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserID",
	4: "Token",
}

func (p *DouyinUserRegisterResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DouyinUserRegisterResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DouyinUserRegisterResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *DouyinUserRegisterResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.StatusMsg = v
	}
	return nil
}

func (p *DouyinUserRegisterResponse) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.UserID = v
	}
	return nil
}

func (p *DouyinUserRegisterResponse) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Token = v
	}
	return nil
}

func (p *DouyinUserRegisterResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("DouyinUserRegisterResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DouyinUserRegisterResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusCode", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.StatusCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DouyinUserRegisterResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusMsg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.StatusMsg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DouyinUserRegisterResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("UserID", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UserID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *DouyinUserRegisterResponse) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Token", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Token); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *DouyinUserRegisterResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DouyinUserRegisterResponse(%+v)", *p)
}

type DouyinUserLoginRequest struct {
	UserName string `thrift:"UserName,1" json:"UserName" query:"user_name"`
	PassWord string `thrift:"PassWord,2" json:"PassWord" query:"password"`
}

func NewDouyinUserLoginRequest() *DouyinUserLoginRequest {
	return &DouyinUserLoginRequest{}
}

func (p *DouyinUserLoginRequest) GetUserName() (v string) {
	return p.UserName
}

func (p *DouyinUserLoginRequest) GetPassWord() (v string) {
	return p.PassWord
}

var fieldIDToName_DouyinUserLoginRequest = map[int16]string{
	1: "UserName",
	2: "PassWord",
}

func (p *DouyinUserLoginRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DouyinUserLoginRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DouyinUserLoginRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.UserName = v
	}
	return nil
}

func (p *DouyinUserLoginRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PassWord = v
	}
	return nil
}

func (p *DouyinUserLoginRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("DouyinUserLoginRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DouyinUserLoginRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("UserName", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UserName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DouyinUserLoginRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("PassWord", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PassWord); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DouyinUserLoginRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DouyinUserLoginRequest(%+v)", *p)
}

type DouyinUserLoginResponse struct {
	StatusCode int32  `thrift:"StatusCode,1" form:"status_code" json:"status_code"`
	StatusMsg  string `thrift:"StatusMsg,2" form:"status_msg" json:"status_msg"`
	UserID     int64  `thrift:"UserID,3" form:"user_id" json:"user_id"`
	Token      string `thrift:"Token,4" form:"token" json:"token"`
}

func NewDouyinUserLoginResponse() *DouyinUserLoginResponse {
	return &DouyinUserLoginResponse{}
}

func (p *DouyinUserLoginResponse) GetStatusCode() (v int32) {
	return p.StatusCode
}

func (p *DouyinUserLoginResponse) GetStatusMsg() (v string) {
	return p.StatusMsg
}

func (p *DouyinUserLoginResponse) GetUserID() (v int64) {
	return p.UserID
}

func (p *DouyinUserLoginResponse) GetToken() (v string) {
	return p.Token
}

var fieldIDToName_DouyinUserLoginResponse = map[int16]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserID",
	4: "Token",
}

func (p *DouyinUserLoginResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DouyinUserLoginResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DouyinUserLoginResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *DouyinUserLoginResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.StatusMsg = v
	}
	return nil
}

func (p *DouyinUserLoginResponse) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.UserID = v
	}
	return nil
}

func (p *DouyinUserLoginResponse) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Token = v
	}
	return nil
}

func (p *DouyinUserLoginResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("DouyinUserLoginResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DouyinUserLoginResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusCode", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.StatusCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DouyinUserLoginResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusMsg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.StatusMsg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DouyinUserLoginResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("UserID", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UserID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *DouyinUserLoginResponse) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Token", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Token); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *DouyinUserLoginResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DouyinUserLoginResponse(%+v)", *p)
}

type DouyinUserRequest struct {
	UserName string `thrift:"UserName,1" json:"UserName" query:"user_name"`
	PassWord string `thrift:"PassWord,2" json:"PassWord" query:"password"`
}

func NewDouyinUserRequest() *DouyinUserRequest {
	return &DouyinUserRequest{}
}

func (p *DouyinUserRequest) GetUserName() (v string) {
	return p.UserName
}

func (p *DouyinUserRequest) GetPassWord() (v string) {
	return p.PassWord
}

var fieldIDToName_DouyinUserRequest = map[int16]string{
	1: "UserName",
	2: "PassWord",
}

func (p *DouyinUserRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DouyinUserRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DouyinUserRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.UserName = v
	}
	return nil
}

func (p *DouyinUserRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PassWord = v
	}
	return nil
}

func (p *DouyinUserRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("DouyinUserRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DouyinUserRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("UserName", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UserName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DouyinUserRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("PassWord", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PassWord); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DouyinUserRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DouyinUserRequest(%+v)", *p)
}

type DouyinUserResponse struct {
	StatusCode int32        `thrift:"StatusCode,1" form:"status_code" json:"status_code"`
	StatusMsg  string       `thrift:"StatusMsg,2" form:"status_msg" json:"status_msg"`
	User       *common.User `thrift:"User,3" form:"user" json:"user"`
}

func NewDouyinUserResponse() *DouyinUserResponse {
	return &DouyinUserResponse{}
}

func (p *DouyinUserResponse) GetStatusCode() (v int32) {
	return p.StatusCode
}

func (p *DouyinUserResponse) GetStatusMsg() (v string) {
	return p.StatusMsg
}

var DouyinUserResponse_User_DEFAULT *common.User

func (p *DouyinUserResponse) GetUser() (v *common.User) {
	if !p.IsSetUser() {
		return DouyinUserResponse_User_DEFAULT
	}
	return p.User
}

var fieldIDToName_DouyinUserResponse = map[int16]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "User",
}

func (p *DouyinUserResponse) IsSetUser() bool {
	return p.User != nil
}

func (p *DouyinUserResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DouyinUserResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DouyinUserResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *DouyinUserResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.StatusMsg = v
	}
	return nil
}

func (p *DouyinUserResponse) ReadField3(iprot thrift.TProtocol) error {
	p.User = common.NewUser()
	if err := p.User.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *DouyinUserResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("DouyinUserResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DouyinUserResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusCode", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.StatusCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DouyinUserResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusMsg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.StatusMsg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DouyinUserResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("User", thrift.STRUCT, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.User.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *DouyinUserResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DouyinUserResponse(%+v)", *p)
}
