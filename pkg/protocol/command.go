package protocol

// Command types defines the lower 8 bits of the command(32 bits) in the packet.
const (
	// 无操作 No operation
	IPMSG_NOOPERATION uint64 = 0x00000000

	// 开启服务(启动时广播) Start service (broadcast at startup)
	IPMSG_BR_ENTRY uint64 = 0x00000001

	// 关闭服务(结束时广播) Stop service (broadcast at end)
	IPMSG_BR_EXIT uint64 = 0x00000002

	// 在线通知 Online notification
	IPMSG_ANSENTRY uint64 = 0x00000003

	// 更改为离线模式 Change to offline mode
	IPMSG_BR_ABSENCE uint64 = 0x00000004

	///////////////////////////////////////////
	// 搜索可通信的用户列表 Search for a list of users who can communicate
	IPMSG_BR_ISGETLIST uint64 = 0x00000010

	// 上面命令的确认命名，把自己标记为可通信 Confirm the above command and mark yourself as communicable
	IPMSG_OKGETLIST uint64 = 0x00000011

	// 用户列表请求 User list request
	IPMSG_GETLIST uint64 = 0x00000012

	// 用户列表传输 User list transfer
	IPMSG_ANSLIST uint64 = 0x00000013

	///////////////////////////////////////////
	// 发送消息 Send message
	IPMSG_SENDMSG uint64 = 0x00000020

	// 确认消息接收 Confirm message reception
	IPMSG_RECVMSG uint64 = 0x00000021

	///////////////////////////////////////////
	// 加密消息阅读通知 Encrypted message read notification
	// IPMSG_READMSG           0x00000030
	IPMSG_READMSG uint64 = 0x00000030

	// 加密消息废弃通知 Encrypted message discard notification
	IPMSG_DELMSG uint64 = 0x00000031

	// 加密消息阅读确认(版本8追加) Encrypted message read confirmation (added in version 8)
	IPMSG_ANSREADMSG uint64 = 0x00000032

	///////////////////////////////////////////
	// 获取 IPMSG 版本 Get IPMSG version
	IPMSG_GETINFO uint64 = 0x00000040

	// 发送 IPMSG 版本
	IPMSG_SENDINFO uint64 = 0x00000041

	///////////////////////////////////////////
	// 获取离线信息 Get offline information
	IPMSG_GETABSENCEINFO uint64 = 0x00000050

	// 发送离线信息
	IPMSG_SENDABSENCEINFO uint64 = 0x00000051

	///////////////////////////////////////////
	// 附件请求 Attachment request
	IPMSG_GETFILEDATA uint64 = 0x00000060

	// 附件请求放弃 Attachment request abandoned
	IPMSG_RELEASEFILES uint64 = 0x00000061

	// 文件夹附件请求 Folder attachment request
	IPMSG_GETDIRFILES uint64 = 0x00000062

	///////////////////////////////////////////
	// 获取 RSA 公匙 RSA public key acquisition
	// IPMSG_GETPUBKEY         0x00000072
	IPMSG_GETPUBKEY uint64 = 0x00000072

	// RSA 公匙响应
	
	///////////////////////////////////////////
	// 群成员在线通知群主 Group member online notification group owner
	IPMSG_DIR_POLL uint64 = 0x000000b0

	// 群管理员任命 Agent appointment
	IPMSG_DIR_POLLAGENT uint64 = 0x000000b1

	// 群主任命管理员广播请求 Group owner appoints administrator broadcast request
	// IPMSG_DIR_BROADCAST     0x000000b2
	IPMSG_DIR_BROADCAST uint64 = 0x000000b2

	// 群主任命管理员广播响应 Group owner appoints administrator broadcast response
	IPMSG_DIR_ANSBROAD uint64 = 0x000000b3

	// 群主更改群成员名单通知 Group owner changes group member list notification
	// IPMSG_DIR_PACKET        0x000000b4
	IPMSG_DIR_PACKET uint64 = 0x000000b4

	// 群主解除管理员请求 Group owner relieves administrator request
	IPMSG_DIR_REQUEST uint64 = 0x000000b5
	
	// 管理员解除自我管理 Group owner relieves self management
	IPMSG_DIR_AGENTPACKET uint64 = 0x000000b6
)
