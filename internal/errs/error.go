package errs

import (
	"fmt"
	"time"
)

func NewErrIndexOutOfRange(length, index int) error {
	return fmt.Errorf("common-kit: 下标超出范围，长度 %d, 下标 %d", length, index)
}

func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("common-kit: 类型转换失败，预期类型: %s，实际值: %#v", want, got)
}

func NewErrInvalidIntervalValue(internalVal time.Duration) error {
	return fmt.Errorf("common-kit: 无效的间隔时间 %d，预期值应大于0", internalVal)
}

func NewErrInvalidMaxInternalValue(maxInterVal, initialInterVal time.Duration) error {
	return fmt.Errorf("common-kit: 最大重试间隔的时间 [%d] 应大于等于初始重试的间隔时间 [%d] ", maxInterVal, initialInterVal)
}
