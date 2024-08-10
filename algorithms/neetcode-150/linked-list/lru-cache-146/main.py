class Node:
    def __init__(self, key:int, val: int, next=None, prev=None):
        self.key = key
        self.val = val
        self.next = next
        self.prev = prev
    def __repr__(self):
        return f"({self.key}, {self.val})"

class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.size = 0
        self.MRU : Node | None = None
        self.LRU : Node | None = None
        self.mapping : dict[int, Node] = {}

    def __update(self, node: Node) -> None:
        if node.prev:
            node.prev.next = node.next
        if node.next:
            node.next.prev = node.prev
        node.prev = None
        node.next = self.MRU
        if self.MRU:
            self.MRU.prev = node
        self.MRU = node

    def __evict(self) -> None:
        if self.LRU:
            self.mapping.pop(self.LRU.key)
            self.LRU = self.LRU.prev
            self.capacity -= 1
        if self.LRU:
            self.LRU.next = None

    def get(self, key: int) -> int:
        if key in self.mapping:
            entry = self.mapping[key]
            self.__update(entry)
            return entry.val
        return -1

    def put(self, key: int, value: int) -> None:
        if self.size >= self.capacity:
            self.__evict()
        new_node: Node = Node(key, value)
        if self.size == 0:
            self.LRU = new_node
        self.size += 1 
        self.mapping[key] = new_node
        self.__update(new_node)



if __name__ == "__main__":
    c = LRUCache(2)
    c.put(1,1)
    c.put(2,2)
    print(c.mapping)
    c.put(3,3)
    print(c.mapping)
    cur = c.MRU
    while cur:
        print(cur)
        cur = cur.next

