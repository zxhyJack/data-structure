class Node {
  constructor(element) {
    this.element = element;
    this.next = null;
  }
}

class LinkedList {
  constructor() {
    this.head = null;
    this.length = 0;
  }

  append(element) {
    let node = new Node(element);
    if (!this.head) {
      this.head = node;
      this.length++;
    } else {
      let current = this.head;
      while (current.next) {    // current.next 可以使拿到的节点是最后一个节点，而不是空节点，便于下一步添加节点
        current = current.next;
      }
      current.next = node;
      this.length++;
    }
  }

  findByIndex(index) {
    let current = this.head;
    let position = 0;
    if (index < -1 || index > this.length) {
      return -1;
    }
    while (position !== index && current !== null) {
      current = current.next;
      position++;
    }
    return current === null ? -1 : current;
  }

  findByValue(element) {
    let current = this.head;
    while (current !== null && current.element !== element) {
      current = current.next;
    }
    return current === null ? '-1' : current;
  }

  insert(element, newElement) {
    let current = this.findByValue(element);
    if (current === -1) {
      return -1;
    }
    let newNode = new Node(newElement);
    newNode.next = current.next;
    current.next = newNode;
  }

  findPre(element) {
    let current = this.head
    if (current.element === element) {
      return -1;
    }
    while (current.next.element !== element) {
      current = current.next;
    }
    return current === null ? -1 : current;
  }

  remove(element) {
    let current = this.findPre(element);
    if (current === -1) {
      return -1;
    }
    current.next = current.next.next;
  }

  size() {
    return this.length;
  }

  toString() {
    let str = '';
    let current = this.head;
    while (current) {
      str += current.element + (current.next ? ',' : '');
      current = current.next;
    }
    return '(' + str + ')';
  }
}

let list = new LinkedList();
list.append(1);
list.append(2);
console.log(list.toString());
list.append('hello');
console.log(list.toString());
console.log(list.findByValue(2));
console.log(list.size());
console.log(list.findByIndex(2));
list.insert('hello', 'world');
console.log(list.toString());

console.log(list.findPre('hello'));
list.remove(2)
console.log(list.toString());


