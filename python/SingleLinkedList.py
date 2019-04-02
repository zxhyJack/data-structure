class Node():
    def __init__(self, element):
        self.element = element
        self.next = None

class LinkedList():
    def __init__(self):
        self.head = None
        self.length = 0

    def append(self, element):
        node = Node(element)
        if self.head == None:
            self.head = node
        else:
            current = self.head
            while current.next:
                current = current.next
            current.next = node

    def to_string(self):
        current = self.head
        res = ''
        while current:
            res += str(current.element) + (',' if current.next else '')
            current = current.next
        return '(' + res + ')'

list = LinkedList()
list.append(1)
print(list.to_string())
list.append(2)
print(list.to_string())