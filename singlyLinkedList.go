package linkedList

import (
	"errors"
	"fmt"
)

type SinglyLinkedList struct {
	head, tail *Node
}

type SinglyLinkedListNumber struct {
	head, tail *NumNode
}

type NumNode struct {
	Value int
	Next  *NumNode
}

type Stack struct {
	people []*Person
	count  int
}

func (s *Stack) Push(n *Person) {
	if s.count >= len(s.people) {
		people := make([]*Person, len(s.people)*2)
		copy(people, s.people)
		s.people = people
	}
	s.people[s.count] = n
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	node := s.nodes[s.count-1]
	s.count--
	return node
}

// Retrieve first node
func (list *SinglyLinkedList) First() *Node {
	return list.head
}

func (list *SinglyLinkedListNumber) First() *NumNode {
	return list.head
}

// Append node
func (list *SinglyLinkedList) Push(p Person) {
	node := &Node{Person: p}
	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
}

func (list *SinglyLinkedListNumber) PushNumber(value int) {
	node := &NumNode{Value: value}
	if list.head == nil {
		list.head = node
	} else {
		list.tail.Next = node
	}
	list.tail = node
}

// Find node
func (list *SinglyLinkedList) Find(name string) *Node {
	var result *Node = nil
	for n := list.First(); n != nil; n = n.Next() {
		if n.Person.Name == name {
			result = n
			fmt.Printf("Found : %v\n", name)
		}
	}
	if result == nil {
		fmt.Printf("Not FOund : %v\n", name)
	}
	return result
}

// Delete node
func (list *SinglyLinkedList) Delete(name string) {
	node2del := list.Find(name)
	if node2del == nil { // There is no such name
		fmt.Println("There is nothing to Delete")
		return
	} else if node2del == list.head && node2del == list.tail { // when there is single one node
		list.head = nil
		list.tail = nil
	} else if node2del == list.head { // deleting first node
		list.head = node2del.Next()
	} else {
		previous_node := list.First()
		for node := list.First(); node != node2del; node = node.Next() {
			previous_node = node
		}
		if node2del == list.tail {
			fmt.Println("Deleting tail")
			list.tail = previous_node
			list.tail.next = nil
		} else {
			fmt.Println("Deleting non tail")
			previous_node.next = node2del.Next()
		}
	}
	// deleting next one(even if the last node)
	fmt.Printf("Removed: %v菫ハ", name)
}

// Pop last item from list
func (list *SinglyLinkedList) Pop() (p Person, err error) {
	var errEmpty = errors.New("Error - List is empty")

	if list.tail == nil {
		err = errEmpty
	} else if list.tail == list.head {
		p = list.head.Person
		list.head = nil
		list.tail = nil
	} else {
		previous_node := list.First()
		for node := list.First(); node != list.tail; node = node.Next() {
			previous_node = node
		}
		p = previous_node.next.Person

		list.tail = previous_node
		list.tail.next = nil
	}

	return p, err
}

func (list *SinglyLinkedList) Reverse() {
	currentNode := list.First()
	list.head = list.tail
	list.tail = currentNode

	var previousNode *Node = nil
	for currentNode != nil {
		temp := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = temp
	}
}

func (list *SinglyLinkedList) RemoveDuplicates() {
	currentNode := list.First()
	var previousNode *Node = nil //非初期化(nil)の値は全型に対して共通の値なので、previousNode := nilだとNg

	nameExistMap := make(map[string]bool)

	for currentNode != nil { // tempを使わなくてもfor n:= list.First(); n!=nil; n = n.Next()でもいける
		temp := currentNode.next
		if _, ok := nameExistMap[currentNode.Name]; ok {
			//if nameExistMap[currentNode.Name] { // でもいける
			previousNode.next = currentNode.Next()
		} else {
			nameExistMap[currentNode.Name] = true
			previousNode = currentNode
		}
		currentNode = temp
	}
}

func (list *SinglyLinkedList) FindInReverseOrder(idx int) *Node {
	node1 := list.First()
	node2 := list.First()

	for i := 0; i < idx; i++ {
		if node2 == nil {
			return nil
		}
		node2 = node2.Next()
	}

	for node2.Next() != nil {
		node1 = node1.Next()
		node2 = node2.Next()
	}
	return node1
}

func (list *SinglyLinkedList) DeleteNonHeadNode(node *Node) {
	if node == nil || node.Next() == nil {
		// この場合、最終nodeは削除できない
		return
	}
	node.Person = node.Next().Person
	node.next = node.Next().Next()
}

func (list *SinglyLinkedList) insertNodeAndSortList(x string) *SinglyLinkedList {
	beforeList := new(SinglyLinkedList)
	afterList := new(SinglyLinkedList)

	currentNode := list.First()

	// listの分割
	for currentNode != nil {
		next := currentNode.Next()
		currentNode.next = nil

		if currentNode.Name < x {
			// 前半のリストの最後にnodeを挿入
			if beforeList.head == nil {
				beforeList.head = currentNode
				beforeList.tail = beforeList.First()
			} else {
				beforeList.tail.next = currentNode
				beforeList.tail = currentNode
			}
		} else {
			// 後半のリストの最後にnodeを挿入
			if afterList.head == nil {
				afterList.head = currentNode
				afterList.tail = afterList.First()

			} else {
				afterList.tail.next = currentNode
				afterList.tail = currentNode
			}
		}
		currentNode = next
	}

	if beforeList == nil {
		return afterList
	}

	// beforeListとafterListをマージ
	beforeList.tail.next = afterList.head
	return beforeList
}

func AddNumList(l1 *SinglyLinkedListNumber, l2 *SinglyLinkedListNumber) *SinglyLinkedListNumber {
	list := new(SinglyLinkedListNumber)
	node1 := l1.head
	node2 := l2.head
	val := 0

	for node1 != nil || node2 != nil {
		if &node1.Value != nil {
			val += node1.Value
		}
		if &node2.Value != nil {
			val += node2.Value
		}

		node := &NumNode{Value: val}
		node.Value = val % 10
		node.Next = nil
		fmt.Println(val)

		if list.head == nil {
			list.head = node
			list.tail = node
		} else {
			list.tail.Next = node
			list.tail = list.tail.Next
		}

		node1 = node1.Next
		node2 = node2.Next
		val /= 10
	}

	return list

	//	node := &NumNode{value: value}
	//	if list.head == nil {
	//		list.head = node
	//	} else {
	//		list.tail.next = node
	//	}
	//	list.tail = node
}

func (singlyList *SinglyLinkedList) IsPalindrome() bool {
	fast := singlyList.head
	slow := singlyList.head
	stack := new(Stack)

	for fast != nil && fast.Next() != nil {
		stack.Push(&fast.Person)
		slow = slow.Next()
		fast = fast.Next().Next()
	}

	if fast != nil {
		slow = slow.Next()
	}

	for slow != nil {
		top := stack.Pop()
		if top.Person != slow.Person {
			return false
		}
		slow = slow.Next()
	}
	return true
}
