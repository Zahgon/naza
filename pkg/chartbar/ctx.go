// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package chartbar

type Ctx struct {
	option Option
}

// WithItems
//
// @param items: 注意，内部不会修改切片底层数据的值以及顺序
func (ctx *Ctx) WithItems(items []Item) string {
	_ = "STUB: not implemented"
	// 拷贝一份，避免修改外部切片的原始顺序
	return ""
}

// 排序

// noop

// 选取需要的元素

// count柱状最长画多长
// num字段多长

// 如果都是整数，且实际最大值最小值的差值小于柱状最大长度限制

// 都是正整数,按原始值绘制

// 最小的负值画1

// 都是正数的情况，最大的画满柱状条，其他的按与最大占比画
// round四舍五入

// 有负数的情况，最小的负数画1，最大的画满

// 最小可能和最大的比太小了

// -3是因为整数不需要小数点和小数点的后两位

func (ctx *Ctx) WithAnySlice(a interface{}, iterateTransFn func(originItem interface{}) Item, modOptions ...ModOption) string {
	_ = "STUB: not implemented"
	return ""
}

func (ctx *Ctx) WithMap(m map[string]int) string { _ = "STUB: not implemented"; return "" }

func (ctx *Ctx) WithMapFloat(m map[string]float64) string { _ = "STUB: not implemented"; return "" }

func (ctx *Ctx) WithCsv(filename string) (string, error) {
	_ = "STUB: not implemented"
	// 读取
	return "", nil
}
