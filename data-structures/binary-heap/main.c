#include "array_list.c"
// #include "array_list.h"
#include "heap.c"
#include "heap.h"
// #include "heap.h"

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
  heap_pop(h);
  printf("______\n");
  print_heap(*h);
  return 0;
}
