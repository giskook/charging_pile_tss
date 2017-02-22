// Code generated by protoc-gen-go.
// source: charge_pile_status.proto
// DO NOT EDIT!

/*
Package Report is a generated protocol buffer package.

It is generated from these files:
	charge_pile_status.proto
	charge_station_status.proto

It has these top-level messages:
	ChargingPileStatus
	ChargingStationStatus
*/
package Report

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChargingPileStatus struct {
	// 通信相关
	DasUuid   string `protobuf:"bytes,1,opt,name=das_uuid" json:"das_uuid,omitempty"`
	Cpid      uint64 `protobuf:"varint,2,opt,name=cpid" json:"cpid,omitempty"`
	Timestamp uint64 `protobuf:"varint,3,opt,name=Timestamp" json:"Timestamp,omitempty"`
	// --------------  状态相关 ------------
	// 充电桩状态 0 可用 !0不可用
	Status uint32 `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	// 实时电流 安
	RealTimeCurrent float32 `protobuf:"fixed32,5,opt,name=RealTimeCurrent" json:"RealTimeCurrent,omitempty"`
	// 实时电压 1福特
	RealTimeVoltage float32 `protobuf:"fixed32,6,opt,name=RealTimeVoltage" json:"RealTimeVoltage,omitempty"`
	// --------------  订单相关 ------------
	// 当前订单的订单号 30位（14位日期+16为cpid）
	CurrentOrderNumber string `protobuf:"bytes,7,opt,name=currentOrderNumber" json:"currentOrderNumber,omitempty"`
	// 充电时长 单位 秒
	ChargingDuration uint32 `protobuf:"varint,8,opt,name=ChargingDuration" json:"ChargingDuration,omitempty"`
	// 充电电量 单位 度
	ChargingCapacity float32 `protobuf:"fixed32,9,opt,name=ChargingCapacity" json:"ChargingCapacity,omitempty"`
	// 充电价格 单位 元
	ChargingCost float32 `protobuf:"fixed32,10,opt,name=ChargingCost" json:"ChargingCost,omitempty"`
	// 充电起始时间 1970流失的秒数
	StartTime uint64 `protobuf:"varint,11,opt,name=StartTime" json:"StartTime,omitempty"`
	// 充电结束时间 1970流失的秒数
	EndTime uint64 `protobuf:"varint,12,opt,name=EndTime" json:"EndTime,omitempty"`
	// 充电起始电表读数 单位度
	StartMeterReading float32 `protobuf:"fixed32,13,opt,name=StartMeterReading" json:"StartMeterReading,omitempty"`
	// 充电终止电表读数 单位度
	EndMeterReading float32 `protobuf:"fixed32,14,opt,name=EndMeterReading" json:"EndMeterReading,omitempty"`
	// --------------  属性相关 ------------
	// id
	Id uint32 `protobuf:"varint,15,opt,name=id" json:"id,omitempty"`
	// 所属充电站id
	StationId uint32 `protobuf:"varint,16,opt,name=stationId" json:"stationId,omitempty"`
	// 终端类型id
	TerminalTypeId uint32 `protobuf:"varint,17,opt,name=terminalTypeId" json:"terminalTypeId,omitempty"`
	// 额定功率 单位：KW
	RatedPower float64 `protobuf:"fixed64,18,opt,name=ratedPower" json:"ratedPower,omitempty"`
	// 电流类型 0：交流，1：直流
	ElectricCurrentType uint32 `protobuf:"varint,19,opt,name=electricCurrentType" json:"electricCurrentType,omitempty"`
	// 输入电压 单位：伏特
	VoltageInput uint32 `protobuf:"varint,20,opt,name=voltageInput" json:"voltageInput,omitempty"`
	// 输出电压 单位：伏特
	VoltageOutput uint32 `protobuf:"varint,21,opt,name=voltageOutput" json:"voltageOutput,omitempty"`
	// 输出电流 单位：安培
	ElectricCurrentOutput uint32 `protobuf:"varint,22,opt,name=electricCurrentOutput" json:"electricCurrentOutput,omitempty"`
	// 枪个数
	GunNumber uint32 `protobuf:"varint,23,opt,name=gunNumber" json:"gunNumber,omitempty"`
	// 充电桩编码 车位号
	Code uint32 `protobuf:"varint,24,opt,name=code" json:"code,omitempty"`
	// 接口 0:RS232,1:RS485,2:CAN,3:USB,4:RJ45,5:RS232(DEBUG)
	InterfaceType uint32 `protobuf:"varint,25,opt,name=interfaceType" json:"interfaceType,omitempty"`
	// 波特率 0:9600,1:14400,2:19200,3:38400,4:576005,5:115200
	BaudRate uint32 `protobuf:"varint,26,opt,name=baudRate" json:"baudRate,omitempty"`
}

func (m *ChargingPileStatus) Reset()                    { *m = ChargingPileStatus{} }
func (m *ChargingPileStatus) String() string            { return proto.CompactTextString(m) }
func (*ChargingPileStatus) ProtoMessage()               {}
func (*ChargingPileStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChargingPileStatus) GetDasUuid() string {
	if m != nil {
		return m.DasUuid
	}
	return ""
}

func (m *ChargingPileStatus) GetCpid() uint64 {
	if m != nil {
		return m.Cpid
	}
	return 0
}

func (m *ChargingPileStatus) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ChargingPileStatus) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ChargingPileStatus) GetRealTimeCurrent() float32 {
	if m != nil {
		return m.RealTimeCurrent
	}
	return 0
}

func (m *ChargingPileStatus) GetRealTimeVoltage() float32 {
	if m != nil {
		return m.RealTimeVoltage
	}
	return 0
}

func (m *ChargingPileStatus) GetCurrentOrderNumber() string {
	if m != nil {
		return m.CurrentOrderNumber
	}
	return ""
}

func (m *ChargingPileStatus) GetChargingDuration() uint32 {
	if m != nil {
		return m.ChargingDuration
	}
	return 0
}

func (m *ChargingPileStatus) GetChargingCapacity() float32 {
	if m != nil {
		return m.ChargingCapacity
	}
	return 0
}

func (m *ChargingPileStatus) GetChargingCost() float32 {
	if m != nil {
		return m.ChargingCost
	}
	return 0
}

func (m *ChargingPileStatus) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ChargingPileStatus) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *ChargingPileStatus) GetStartMeterReading() float32 {
	if m != nil {
		return m.StartMeterReading
	}
	return 0
}

func (m *ChargingPileStatus) GetEndMeterReading() float32 {
	if m != nil {
		return m.EndMeterReading
	}
	return 0
}

func (m *ChargingPileStatus) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChargingPileStatus) GetStationId() uint32 {
	if m != nil {
		return m.StationId
	}
	return 0
}

func (m *ChargingPileStatus) GetTerminalTypeId() uint32 {
	if m != nil {
		return m.TerminalTypeId
	}
	return 0
}

func (m *ChargingPileStatus) GetRatedPower() float64 {
	if m != nil {
		return m.RatedPower
	}
	return 0
}

func (m *ChargingPileStatus) GetElectricCurrentType() uint32 {
	if m != nil {
		return m.ElectricCurrentType
	}
	return 0
}

func (m *ChargingPileStatus) GetVoltageInput() uint32 {
	if m != nil {
		return m.VoltageInput
	}
	return 0
}

func (m *ChargingPileStatus) GetVoltageOutput() uint32 {
	if m != nil {
		return m.VoltageOutput
	}
	return 0
}

func (m *ChargingPileStatus) GetElectricCurrentOutput() uint32 {
	if m != nil {
		return m.ElectricCurrentOutput
	}
	return 0
}

func (m *ChargingPileStatus) GetGunNumber() uint32 {
	if m != nil {
		return m.GunNumber
	}
	return 0
}

func (m *ChargingPileStatus) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ChargingPileStatus) GetInterfaceType() uint32 {
	if m != nil {
		return m.InterfaceType
	}
	return 0
}

func (m *ChargingPileStatus) GetBaudRate() uint32 {
	if m != nil {
		return m.BaudRate
	}
	return 0
}

func init() {
	proto.RegisterType((*ChargingPileStatus)(nil), "Report.ChargingPileStatus")
}

func init() { proto.RegisterFile("charge_pile_status.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xe5, 0x10, 0xd2, 0x78, 0xc8, 0x3f, 0x6f, 0x93, 0x76, 0x5a, 0x84, 0x64, 0x71, 0xf2,
	0x89, 0x0b, 0x1f, 0x21, 0xf4, 0x90, 0x03, 0xb4, 0x4a, 0x2b, 0xae, 0xd1, 0xc6, 0x3b, 0x98, 0x95,
	0x9c, 0x5d, 0x6b, 0x3d, 0x0b, 0xea, 0x07, 0xe4, 0x7b, 0xa1, 0x9d, 0x24, 0xa8, 0xf4, 0xfa, 0x7b,
	0x6f, 0xe7, 0x3d, 0x3d, 0x1b, 0xb0, 0xfe, 0xa9, 0x43, 0x43, 0xbb, 0xce, 0xb6, 0xb4, 0xeb, 0x59,
	0x73, 0xec, 0x3f, 0x75, 0xc1, 0xb3, 0x57, 0xa3, 0x2d, 0x75, 0x3e, 0xf0, 0xc7, 0x3f, 0x43, 0x50,
	0xeb, 0x64, 0xb2, 0xae, 0x79, 0xb0, 0x2d, 0x3d, 0x8a, 0x49, 0x2d, 0x60, 0x6c, 0x74, 0xbf, 0x8b,
	0xd1, 0x1a, 0xcc, 0xca, 0xac, 0xca, 0xd5, 0x04, 0x86, 0x75, 0x67, 0x0d, 0x0e, 0xca, 0xac, 0x1a,
	0xaa, 0x02, 0xf2, 0x27, 0x7b, 0xa0, 0x9e, 0xf5, 0xa1, 0xc3, 0x37, 0x82, 0x66, 0x30, 0x3a, 0x26,
	0xe0, 0xb0, 0xcc, 0xaa, 0xa9, 0xba, 0x86, 0xf9, 0x96, 0x74, 0x9b, 0x6c, 0xeb, 0x18, 0x02, 0x39,
	0xc6, 0xb7, 0x65, 0x56, 0x0d, 0x5e, 0x0a, 0xdf, 0x7d, 0xcb, 0xba, 0x21, 0x1c, 0x89, 0x70, 0x0b,
	0xaa, 0x3e, 0x3a, 0xef, 0x83, 0xa1, 0xf0, 0x2d, 0x1e, 0xf6, 0x14, 0xf0, 0x42, 0xe2, 0x11, 0x16,
	0xe7, 0x9a, 0x5f, 0x62, 0xd0, 0x6c, 0xbd, 0xc3, 0xb1, 0xe4, 0xbc, 0x50, 0xd6, 0xba, 0xd3, 0xb5,
	0xe5, 0x67, 0xcc, 0xe5, 0xde, 0x12, 0x26, 0xff, 0x14, 0xdf, 0x33, 0x82, 0xd0, 0x02, 0xf2, 0x47,
	0xd6, 0x81, 0x53, 0x3e, 0xbe, 0x93, 0xea, 0x73, 0xb8, 0xb8, 0x73, 0x46, 0xc0, 0x44, 0xc0, 0x0d,
	0x14, 0xe2, 0xf9, 0x4a, 0x4c, 0x61, 0x4b, 0xda, 0x58, 0xd7, 0xe0, 0xf4, 0xdc, 0xfe, 0xce, 0x99,
	0xff, 0x84, 0x99, 0x08, 0x00, 0x03, 0x6b, 0x70, 0x2e, 0x9d, 0x0a, 0xc8, 0xd3, 0x16, 0xd6, 0xbb,
	0x8d, 0xc1, 0x85, 0xa0, 0x2b, 0x98, 0x31, 0x85, 0x83, 0x75, 0xba, 0x7d, 0x7a, 0xee, 0x68, 0x63,
	0xb0, 0x10, 0xae, 0x00, 0x82, 0x66, 0x32, 0x0f, 0xfe, 0x37, 0x05, 0x54, 0x65, 0x56, 0x65, 0xea,
	0x3d, 0x5c, 0x52, 0x4b, 0x35, 0x07, 0x5b, 0x9f, 0xa6, 0x4b, 0x4f, 0xf0, 0x52, 0x1e, 0x2c, 0x61,
	0xf2, 0xeb, 0x38, 0xdb, 0xc6, 0x75, 0x91, 0x71, 0x29, 0x74, 0x05, 0xd3, 0x13, 0xbd, 0x8f, 0x9c,
	0xf0, 0x4a, 0xf0, 0x07, 0x58, 0xbd, 0xba, 0x74, 0x92, 0xaf, 0xce, 0x3d, 0x9b, 0xe8, 0x4e, 0x43,
	0x5f, 0x0b, 0x4a, 0xdf, 0xd9, 0x1b, 0x42, 0x3c, 0x9f, 0xb5, 0x8e, 0x29, 0xfc, 0xd0, 0x35, 0x49,
	0x87, 0x1b, 0xc1, 0x0b, 0x18, 0xef, 0x75, 0x34, 0x5b, 0xcd, 0x84, 0xb7, 0x89, 0xec, 0x47, 0xf2,
	0x5b, 0x7d, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x93, 0x7d, 0xad, 0xbe, 0x72, 0x02, 0x00, 0x00,
}
