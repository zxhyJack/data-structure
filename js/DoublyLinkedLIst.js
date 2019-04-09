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

  // achieve by self
  // 这种方式考虑不全面
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
      let current = this.head;
      if (!this.head) {     // 当链表为空时，插入节点既是头结点也是尾节点
        this.head = node;
        this.tail = node;
        this.length++;
        return;
      }
      if (position === 0) {   // 在链表头插入元素
        node.next = current;
        current.prev = node;
        this.head = node;
      } else if (position === this.length) {  // 在链表最后插入元素
        this.tail.next = node;
        node.prev = this.tail;
        this.tail = node;
      } else {                  // 在中间位置插入元素
        let index = 0;
        while (current && index !== position) {
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
}

let list = new DoublyLinkedList();
// list.append(1);
// list.append(2);
// list.append(3);
// console.log(list.findByValue(1));
// list.insert(0, 'hello');
list.insertRef(0, 1);
list.insertRef(1, 3);
list.insertRef(list.length, 2);
list.insertRef(1, 'hello');
console.log(list.toString());
