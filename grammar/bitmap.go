package main

import "fmt"

func (b *BitMap) Set(num int) {
	idx, pos := b.Offset(num)

	if b.size <= idx {
		b.Grow(idx)
	}

	if num == 0 {
		b.bits[idx] &^= 0x01 << pos
	} else {
		fmt.Println()
		b.bits[idx] |= 0x01 << pos
	}

}

func (b *BitMap) Get(num int) bool {
	idx, pos := b.Offset(num)
	if b.size < idx {
		return false
	}
	return (b.bits[idx]>>pos)&0x01 != 0
}

func (b *BitMap) Reset(num int) {
	panic("implement me")
}

func (b *BitMap) Has(num int) bool {
	idx, pos := b.Offset(num)

	return idx < b.size && (b.bits[idx]&(1<<pos)) != 0
}

func (b *BitMap) Offset(num int) (int, int) {
	/**
	idx:表示在数组的下标
	pos:这个数组下标中的byte位中的第几位
	*/
	idx, pos := num/mask, num%mask
	return idx, pos
}

func (b *BitMap) Grow(idx int) bool {
	bytes := make([]byte, idx)
	b.bits = append(b.bits, bytes...)
	b.size++
	return true

}

/**
  bit[0]	 0 0 0 0 0 0 0 0
  bit[1]     0 0 0 0 0 0 0 0

  idx = x/8
  pos = x%8
*/

const (
	mask = 8
)

type BitMapper interface {
	Set(num int)
	Get(num int) bool
	Reset(num int)
	Has(num int) bool
	Offset(num int) (int, int)
	Grow(num int) bool
}
type BitMap struct {
	bits   []byte
	size   int
	maxLen int
}

func NewBitmap() *BitMap {
	return &BitMap{
		bits:   make([]byte, 1),
		size:   1,
		maxLen: 1,
	}
}

func main() {
	bitmap := NewBitmap()
	bitmap.Set(10)
	// fmt.Printf()
	bitmap.Set(1)
	bitmap.Set(10)

	bitmap.Set(10)

	fmt.Printf("%08b", bitmap.bits)
	//bitmap.Set(3)
	//fmt.Printf("%08b", bitmap.bits)
	fmt.Println()
	get := bitmap.Get(3)
	fmt.Println(get)
	get = bitmap.Get(10)
	fmt.Println(get)

}
