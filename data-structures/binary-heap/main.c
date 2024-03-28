#include "heap.h"
#include "heap.c"
int main() {
  Heap *h = heap_init(1);
  heap_insert(h, 3);
  heap_insert(h, 2);
  heap_insert(h, 1);
  heap_insert(h, 4);
  heap_insert(h, 5);
  heap_insert(h, 0);
  printf("______\n");
  print_heap(*h);
  int i = heap_pop(h);
  printf("popped: %d\n",i);
  printf("______\n");
  print_heap(*h);
  return 0;
}
