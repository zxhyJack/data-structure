class Node {
	constructor(element) {
		this.element = element;
		this.prev = null;
		this.next = null;
	}
}

class DoublyLinkedList {
	constructor() {
		this.head = null;
		this.tail = null;
		this.length = 0;
	}

	append(element) {
		let current = this.head;
		let node = new Node(element);
		if (this.head === null) {
			this.head = node;
			this.tail = node;
			this.length++;
			return;
		}
		while (current.next !== null) {
			current = current.next;
		}
		current.next = node;
		node.prev = current;
		this.tail = node;
		this.length++;
	}

	findByValue(element) {
		let current = this.head;
		if (current === null) {
			return false;
		}
		while (current.element !== element) {
			current = current.next;
		}
		return current === null ? false : current;
	}

	// insert(element, newElement) {
	//   let current = this.findByValue(element);
	//   if (current === -1) {
	//     return -1;
	//   }
	//   let node = new Node(newElement);
	//   current.next = node;
	//   node.prev = current;
	//   this.length++;
	// }

	findByIndex(position) {
		if (position < 0 || position >= this.length) {
			return false
		}
		let current = this.head;
		let index = 0;
		while (current && index !== position) {
			current = current.next;
			index++;
		}
		return current === null ? false : current;
	}

	// realize by self
	// 这种方式考虑不全面,仅考虑了在链表中间插入的情况
	insert(position, element) {
		let current = this.findByIndex(position);
		if (!current) {
			return false;
		}
		let node = new Node(element);
		let tmpNode = current.next;   // 
		node.next = current;
		current.prev = node;
		this.length++;
		if (tmpNode === null) {   // 
			return true
		}
		tmpNode.prev = node;
		return true;
	}

	// according to book
	insertRef(position, element) {
		if (position > -1 && position <= this.length) {
			let node = new Node(element);
			let current;
			if (!this.head) {     // 当链表为空时，插入节点既是头结点也是尾节点
				this.head = node;
				this.tail = node;
				this.length++;
				return;
			}
			if (position === 0) {   // 在链表头插入元素
				current = this.head;
				node.next = current;
				current.prev = node;
				this.head = node;
			} else if (position === this.length) {  // 在链表最后插入元素
				current = this.tail;
				current.next = node;
				node.prev = current;
				this.tail = node;
			} else {                  // 在中间位置插入元素
				let index = 0;
				current = this.head;
				while (index !== position) {
					current = current.next;
					index++;
				}
				current.prev.next = node;
				node.prev = current.prev;
				node.next = current;
				current.prev = node;
			}

			this.length++;
			return true;
		} else {
			return false;
		}

	}

	toString() {
		let current = this.head;
		let str = '';
		while (current !== null) {
			str += current.element + (current.next === null ? '' : ',');
			current = current.next;
		}
		return '(' + str + ')';
	}

	insertBook(position, element) {
		//检查越界值
		if (position >= 0 && position <= this.length) {
			let node = new Node(element),
				current = this.head,
				previous = this.tail,
				index = 0;
			// 在头部插入元素
			if (position === 0) {
				if (!this.head) {
					this.head = node;
					this.tail = node;
				} else {
					node.next = current;
					current.prev = node;
					this.head = node;
				}
			} else if (position === this.length) {	// 在尾部插入元素
				while (current.next) {
					current = current.next;
				}
				current.next = node;
				node.prev = current
				this.tail = node;
			} else {	// 在链表中间插入元素
				while (index++ < position) {
					current = current.next;
				}
				previous = current.prev;	// previous用来暂存current前面的节点，防止current.prev被覆盖
				node.next = current;
				current.prev = node;
				previous.next = node;
				node.prev = previous;
			}
			this.length++;
		} else {
			return false;
		}
	};

	removeAt(position) {
		// 检查越界值
		if (position >= 0 && position < this.length) {
			let current = this.head,
				previous = null,
				index = 0;
			if (position === 0) {
				this.head = current.next;
			} else if (position === this.length - 1) {
				previous = this.tail.prev;
				this.tail = previous;
				previous.next = null;
			} else {
				while (index++ < position) {
					current = current.next;
				}
				//将previous与current的下一项链接起来
				previous = current.prev;
				previous.next = current.next;
				current.next.prev = previous;
			}
			this.length--;
		} else {
			return false
		}
	}
} 

let list = new DoublyLinkedList();
// list.append(1);
// list.append(2);
// list.append(3);
// console.log(list.findByValue(1));
// list.insert(0, 'hello');
// list.insertRef(0, 1);
// list.insertRef(1, 3);
// list.insertRef(list.length, 2);
// list.insertRef(1, 'hello');
// list.insertBook(2, 'world');
// list.insertRef(1, 'hahah');
list.insertBook(0, 1);
list.insertBook(1, 2);
list.insertBook(0, 3);
// list.insertBook(1, 4); // expect (3,4,1,2)
list.removeAt(1);
console.log(list.toString());
