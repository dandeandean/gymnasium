#include "array_list.c"
#include "array_list.h"
#include "heap.c"
#include "heap.h"

int main() {
  Heap *h = heap_init(1);
  ArrayList *al = h->list;
  add_item(al, 74);
  add_item(al, 42);

  add_item(al, 49);
  add_item(al, 49);

  add_item(al, 70);
  swap_items(al, 0, 2);

  add_item(al, 75);
  add_item(al, 80);
  add_item(al, 55);
  // print_al(*al);
  print_heap(*h);

  return 0;
}
