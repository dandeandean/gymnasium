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
        self.MRU: Node = Node(-1,-1)
        self.LRU: Node = Node(-2,-1, next=self.MRU)
        self.MRU.prev = self.LRU
        self.mapping : dict[int, Node] = {}

    def __push_right(self, node: Node) -> None:
        old_prev = self.MRU.prev
        if old_prev:
            old_prev.next = node
        node.prev = self.MRU.prev
        node.next = self.MRU
        self.size +=1


    def __pop(self, node: Node) -> None:
        old_next = node.next
        old_prev = node.prev
        if old_next:
            old_next.prev = old_prev
        if old_prev:
            old_prev.next = old_next

    def get(self, key: int) -> int:
        if key in self.mapping:
            entry = self.mapping[key]
            self.__pop(entry)
            self.__push_right(entry)
            return entry.val
        return -1

    def put(self, key: int, value: int) -> None:
        new_node: Node = Node(key, value, next=self.MRU)
        if key in self.mapping:
            self.__pop(self.mapping[key])
            self.size -=1
        self.size += 1 
        if self.size >= self.capacity and self.LRU.prev:
            self.mapping.pop(self.LRU.prev)
            self.__pop(self.LRU.prev)
        self.mapping[key] = new_node
        self.__push_right(new_node)

if __name__ == "__main__":
    def print_cache(c: LRUCache):
        print(f"---------<<{c.MRU=}, {c.LRU=}>>")
        print(c.mapping)
        cur = c.MRU
        while cur:
            print(cur)
            cur = cur.next
        print("---------")
    #fs=["get","put","get","put","put","get","get"]
    #argv = [[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
    c = LRUCache(2)
    c.put(2,6)
    print_cache(c)
    c.get(2)
    c.put(1,5)
    print_cache(c)
    c.put(1,2)
    print(c.get(1))
    print(c.get(2))

