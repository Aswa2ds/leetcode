package codetop

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
}

// func (this *LRUCache) Print() {
// 	fmt.Println("capacity: ", this.capacity)
// 	fmt.Println("cache: ")
// 	for k, v := range this.cache {
// 		fmt.Println("key: ", k, "value: ", v.value)
// 	}
// 	fmt.Println("list: ")
// 	for p := this.head.next; p != this.head; p = p.next {
// 		fmt.Println("key: ", p.key, "value: ", p.value)
// 	}
// }

type Node struct {
	key       int
	value     int
	pre, next *Node
}

func (this *LRUCache) removeNode(node *Node) int {
	rmKey := node.key
	node.pre.next = node.next
	node.next.pre = node.pre
	return rmKey
}

func (this *LRUCache) addToHead(node *Node) {
	node.next = this.head.next
	node.pre = this.head
	node.next.pre = node
	node.pre.next = node
}
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() int {
	return this.removeNode(this.head.pre)
}

func Constructor(capacity int) LRUCache {
	head := &Node{key: -1, value: -1}
	head.next = head
	head.pre = head
	return LRUCache{
		capacity: capacity,
		cache:    map[int]*Node{},
		head:     head,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.moveToHead(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		v.value = value
		this.moveToHead(v)
		return
	}

	if this.capacity == len(this.cache) {
		rmKey := this.removeTail()
		delete(this.cache, rmKey)
	}

	newNode := &Node{key: key, value: value}
	this.addToHead(newNode)
	this.cache[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
