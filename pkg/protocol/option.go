package protocol

// Option type defines the higher 24 bits of the command(32 bits).

// Option for all
const (
	// 离线模式(与成员识别命令一起使用)
	IPMSG_ABSENCEOPT uint64 = 0x00000100
	// 服务器(预订)
	IPMSG_SERVEROPT uint64 = 0x00000200

	// 成员识别命令
	IPMSG_DIALUPOPT uint64 = 0x00010000
	// 附件
	IPMSG_FILEATTACHOPT uint64 = 0x00200000
	// 密码
	IPMSG_ENCRYPTOPT uint64 = 0x00400000
	// 消息使用 UTF-8
	IPMSG_UTF8OPT uint64 = 0x00800000
)

// Option for send command
const (

	// 检查发送
	IPMSG_SENDCHECKOPT uint64 = 0x00000100
	// 加密
	IPMSG_SECRETOPT uint64 = 0x00000200
	// 检查加密(版本8追加)
	// 广播(全局选择)
	IPMSG_BROADCASTOPT uint64 = 0x00000400
	// 多播(多个选择)
	IPMSG_MULTICASTOPT uint64 = 0x00000800

	// 不弹出消息窗口
	IPMSG_NOPOPUPOPT uint64 = 0x00001000
	// 自动应答
	IPMSG_AUTORETOPT uint64 = 0x00002000
	// 重传标志(HOSTLIST 获取时使用)
	IPMSG_RETRYOPT uint64 = 0x00004000
	// 密码
	IPMSG_PASSWORDOPT uint64 = 0x00008000
	// 附件加密请求
	IPMSG_ENCFILEOPT uint64 = 0x00010000

	// 不写入日志
	IPMSG_NOLOGOPT uint64 = 0x00020000
	// 新版本多播
	IPMSG_NEWMULTIOPT uint64 = 0x00040000
	// 启启动时不加入用户列表
	IPMSG_NOADDLISTOPT uint64 = 0x00080000

	// 检查加密
	IPMSG_READCHECKOPT uint64 = 0x00100000
	// 0x00100200
	IPMSG_SECRETEXOPT uint64 = IPMSG_READCHECKOPT | IPMSG_SECRETOPT

	// 能使用UTF-8
	IPMSG_CAPUTF8OPT uint64 = 0x01000000
	// 支持 IPDict 格式
	IPMSG_CAPIPDICTOPT uint64 = 0x02000000
	// 在密文中包含附件信息和目的地信息
	IPMSG_ENCEXTMSGOPT uint64 = 0x04000000
	// 支持嵌入图片的消息
	IPMSG_CLIPBOARDOPT uint64 = 0x08000000

	// 群主
	IPMSG_DIR_MASTER uint64 = 0x10000000

	// 0x00002400
	IPMSG_NO_REPLY_OPTS uint64 = IPMSG_BROADCASTOPT | IPMSG_AUTORETOPT
)
