class Node {
  constructor(element) {
    this.element = element;
    this.pre = null;
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
    let newNode = new Node(element);
    if (this.head === null) {
      this.head = newNode;
      this.length++;
      return;
    }
    while (current.next !== null) {
      current = current.next;
    }
    current.next = newNode;
    newNode.pre = current;
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
  //   let newNode = new Node(newElement);
  //   current.next = newNode;
  //   newNode.pre = current;
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
  insert(position, element) {
    let current = this.findByIndex(position);
    if (!current) {
      return false;
    }
    let newNode = new Node(element);
    let tmpNode = current.next;
    current.next = newNode;
    newNode.pre = current;
    newNode.next = tmpNode;
    this.length++;
    if(tmpNode===null){
      return true
    }
    tmpNode.pre = newNode;
    return true;
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
list.append(1);
list.append(2);
list.append(3);
// console.log(list.findByValue(1));
list.insert(2, 'hello')
console.log(list.toString());
