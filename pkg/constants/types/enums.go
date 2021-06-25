package types

// 注意，每个 const 值务必声明类型

type PullPolicy string

const (
	PullAlways       PullPolicy = "Always"       // 总是
	PullNever        PullPolicy = "Never"        // 从不
	PullIfNotPresent PullPolicy = "IfNotPresent" // 如果存在
)

//go:generate tools gen enum ExampleType
type ExampleType uint8 // 定义枚举为uint8的别名

// 通过const定义枚举常量，iota表明下面的常量值从0开始依次递增。枚举第一个固定为UNKNOWN
const (
	EXAMPLE_TYPE_UNKNOWN ExampleType = iota
	EXAMPLE_TYPE__A                  // 类型A描述
	EXAMPLE_TYPE__B                  // 类型B描述
)
