package sqllist

import (
	"errors"
	"fmt"
)

// DefaultSize 指定默认顺序表的大小
//const DefaultSize = 100

// ElementType 定义顺序表中的元素类型
//type ElementType interface{}

// SqList 定义顺序表的结构体 即对应抽象数据对象
type SqList struct {
	elem    []int //元素数组
	maxsize int   // 当前顺序表的容量
	length  int   //当前顺序表的下标
}

// InitList 初始化
func (s *SqList) InitList(size int, arr []int) *SqList {
	//初始化结构体实体
	list := &SqList{
		elem:    make([]int, 0, size), //长度为0，容量为maxsize
		maxsize: size,
		length:  0}
	//将传入的数据导入初始化的实体
	for i := range arr {
		list.elem = append(list.elem, arr[i])
		list.length++
	}
	return list
}

// GetElem 根据下标获取指定元素
func (s *SqList) GetElem(pos int) (int, error) {
	err := s.validPostion(pos) //检查插入位置有没有越界
	if err != nil {
		return -1, err
	}
	if s.length == 0 { //说明是空表
		fmt.Println("NOT FOUND")
		return -1, nil
	}
	return s.elem[pos], nil //返回查找到的元素
}

// LocateElem 查找指定元素的位置
func (s *SqList) LocateElem(e int) {
	for x := range s.elem {
		if e == s.elem[x] {
			fmt.Printf("查找成功，元素为第%d个\n", x)
			return
		}
	}
	fmt.Println("查找失败")
	return
}

// InsertElem 根据指定位置插入元素
func (s *SqList) InsertElem(i int, e int) {
	if i < 1 || i > s.maxsize { //检查插入位置有没有越界
		fmt.Println("请检查i值")
		return
	}
	if s.maxsize == s.length { //当前存储空间已满
		fmt.Println("存储空间已满")
		return
	}
	var e2 int //为新元素创建“座位”
	s.elem = append(s.elem, e2)
	for x := s.length - 1; x >= i-1; x-- {
		s.elem[x+1] = s.elem[x]
	}
	s.elem[i-1] = e
	s.length++
	fmt.Println("插入成功")
	return
}

// DeleteElem 删除第i个元素
func (s *SqList) DeleteElem(i int) *SqList {
	if i < 1 || i > s.length { //检查删除位置是否越界
		fmt.Println("删除位置i非法")
		return s
	}
	//这里只展示删除点前后元素位置移动操作，直接利用append()函数特性进行“删除”更常用
	for x := i; x < s.length; x++ {
		s.elem[x-1] = s.elem[x]
	}
	s.length--
	//从i到表位整体前移1个单位长度，剔除表尾冗余。
	s.elem = append(s.elem[:s.length])
	fmt.Println("删除成功")
	return s
}

// UpdateElem 更新指定位置的元素
func (s *SqList) UpdateElem(e int, i int) error {
	err := s.validPostion(i) //检查更新位置是否越界
	if err != nil {
		return err
	}
	s.elem[i-1] = e
	return nil
}

// AddElement 顺序表后面追加元素
func (s *SqList) AddElement(x int) error {
	if cap(s.elem) == s.length+1 {
		return errors.New("The SqList is full")
	}
	s.elem[s.length+1] = x
	s.length++
	return nil
}

// PrintList 打印顺序表中的内容
func (s *SqList) PrintList() {
	fmt.Println("The elements is:")
	for i := 0; i <= s.length; i++ {
		tmp, _ := s.GetElem(i)
		fmt.Println(" ", tmp)
	}
}

// insertCheck 校验插入的时候输入的pos是否合法。
func (s *SqList) insertCheck(pos int) error {
	if pos > cap(s.elem) || pos < 0 {
		return errors.New("SqList pos out 0f bounds error")
	}
	if pos > s.length+1 || pos < 0 {
		return errors.New("you set a wrong position")
	}
	return nil
}

// validPostion 校验修改（update Or remove Or get）的时候输入的pos是否合法。
func (s *SqList) validPostion(pos int) error {
	if pos > cap(s.elem) || pos < 0 {
		return errors.New("SqList pos out 0f bounds error")
	}
	if pos > s.length || pos < 0 {
		return errors.New("you set a wrong position")
	}
	return nil
}
