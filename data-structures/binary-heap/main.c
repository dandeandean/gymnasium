#include "array_list.c"
// #include "array_list.h"
#include "heap.c"
// #include "heap.h"

int main() {
  Heap *h = heap_init(1);
  heap_insert(h, 3);

  print_heap(*h);
  printf("______\n");
  heap_insert(h, 2);
  print_heap(*h);
  heap_insert(h, 1);
  printf("______\n");
  print_heap(*h);
  heap_insert(h, 4);
  heap_insert(h, 6);
  heap_insert(h, 0);
  printf("______\n");
  print_heap(*h);
  return 0;
}
