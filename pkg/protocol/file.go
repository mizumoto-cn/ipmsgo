package protocol

// lower 8 bits
const (
	IPMSG_FILE_REGULAR   uint64 = 0x00000001 // 普通文件
	IPMSG_FILE_DIR       uint64 = 0x00000002 // 文件夹
	IPMSG_FILE_RETPARENT uint64 = 0x00000003 // 返回到父级目录
	IPMSG_FILE_SYMLINK   uint64 = 0x00000004 // 软链接
	IPMSG_FILE_CDEV      uint64 = 0x00000005 // for UNIX 字符设备
	IPMSG_FILE_BDEV      uint64 = 0x00000006 // for UNIX 伪文件系统
	IPMSG_FILE_FIFO      uint64 = 0x00000007 // for UNIX 管道
	IPMSG_FILE_RESFORK   uint64 = 0x00000010 // for Mac
	IPMSG_FILE_CLIPBOARD uint64 = 0x00000020 // for Windows Clipboard
)
