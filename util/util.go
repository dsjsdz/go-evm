package util

const (
	MethodServerTimeGet = "80"
	MethodDeviceOffline = "81"
)

const (
	SuccessStatusCode = "01"
	FailStatusCode    = "0f"
)

// 0D 24 帧头
// 0D 0A 帧尾
// 帧长度= 从0D 24 开始一直到0D 0A结束，总共字节长度，0x28=40个字节
// 1D 00 帧长度
// 80 功能码
const (
	FrameHeaderSize    = 2  // 帧头长度
	FrameLenSize       = 2  //帧长度
	FrameFuncCodeSize  = 1  //功能码
	FrameMachineIdSize = 14 // 设备ID
	FrameProtocolSize  = 1  //协议版本号
	FrameDataSize      = 15 //预留
	FrameCheckCodeSize = 1  // 校验码
	FrameFooterSize    = 2  // 帧尾
)
