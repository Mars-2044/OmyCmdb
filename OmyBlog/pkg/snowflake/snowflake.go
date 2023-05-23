package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

func GetID() (sid uint64) {
	var st sonyflake.Settings
	// 此处没有实现MachineID, CheckID方法, 默认通过IP计算得出MachineID, 而不进行CheckID
	// 然而好多人用了 SnowFlake 算法，却返回一个 Number 类型的 ID 给前端……孰不知 JS 的整数最大只支持 53 位，溢出了都……
	st.StartTime = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	sf := sonyflake.NewSonyflake(st)

	id, err := sf.NextID()
	if err != nil {
		fmt.Printf("获取随机ID失败: %s", err)
	}

	return id
	// fmt.Println("Generated id:", id)
}
