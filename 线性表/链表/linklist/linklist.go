package linklist

import "errors"

// Node 结点
type Node struct {
	elem interface{} //数据项
	next *Node       //指向下一个结点的指针
}

// LinkList 链表
type LinkList struct {
	head   *Node
	length uint //链表长度
}

// InitNode 链表结点初始化
func InitNode(elem interface{}) *Node {
	return &Node{elem, nil}
}

// InitLinkList 链表初始化 头指针 单链表中间的数据结点没有别名，无法直接获取值，只有循环依次读取 读取当前节点数据时 只能获取下个结点信息 无法获取上个结点的信息
func InitLinkList() *LinkList {
	return &LinkList{InitNode(nil), 0}
}

// InsertElem 插入数据(头插法) 头插法是从一个空表开始，重复读入数据，生成新结点，将读入的数据插入到新结点的数据域，新结点插入到当前链表的表头上，直到读入结束
func (l *LinkList) InsertElem(elem interface{}) error {
	if l.head == nil {
		return nil
	}
	newNode := InitNode(elem) //需要插入的结点及数据
	p2 := l.head.next         //把头结点需要指向的结点赋值给一个新的*Node
	l.head.next = newNode     //头结点指向插入的数据结点
	newNode.next = p2         //数据插入数据的结点指向头结点
	l.length++
	return nil
}

// InsertElemTail 尾插法 将新结点插入到当前链表的表尾，需要增设尾指针rear 使新插入的数据结点始终指向链表的尾结点 读取顺序和插入顺序一致
func (l *LinkList) InsertElemTail(elem interface{}) {
	//初始化插入的数据结点
	newNode := InitNode(elem)
	temp := l.head
	for temp.next != nil { //检查头指针指向的next是否为空，为空时插入数据
		temp = temp.next
	}
	temp.next = newNode //插入数据
	newNode.next = nil  //将插入的数据结点指向nil结束
	l.length++          //链表长度+1

}

// InsertPostion 在第i个结点后插入数据
func (l *LinkList) InsertPostion(i int, elem interface{}) {
	newNode := InitNode(elem)
	temp := l.head
	for i > 0 {
		temp = temp.next
		i--
	}
	newNode.next = temp.next
	temp.next = newNode
	l.length++
}

// DeletePostion 删除指定位置结点 删除第i个结点
func (l *LinkList) DeletePostion(i int) (err error) {
	i = i - 1
	if i < 0 || i > int(l.length) {
		return errors.New("插入位置非法！")
	}
	//newNode := InitNode(elem)
	temp := l.head
	for i > 0 {
		temp = temp.next
		i--
	}
	temp.next = temp.next.next
	l.length--
	return nil
}

// QueryPostion 根据位置查询结点数据
func (l *LinkList) QueryPostion(i int) (elem interface{}, err error) {
	i = i - 1
	if i < 0 || i > int(l.length) {
		return nil, errors.New("查询错误，请检查i值")
	}
	//newNode := InitNode(elem)
	temp := l.head
	for i > 0 {
		temp = temp.next
		i--
	}
	elem = temp.next.elem
	return elem, nil
}

// GetElem 根据位置获取元素
func GetElem(l LinkList, elem interface{}, i int) (interface{}, error) {
	var j int
	var p *Node
	p = l.head.next
	j = 1
	for p != nil && j < i {
		p = p.next
		j = j + 1
	}
	if p == nil || j > i {
		err := errors.New("指针为空或者")
		return 0, err
	}
	elem = p.elem
	return elem, nil
}
