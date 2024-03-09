#include "array_list.c"
#include "array_list.h"
#include "heap.c"

int main() {
  ArrayList *al = alinit(4);
  add_item(al, 74);
  add_item(al, 42);
  print_al(*al);

  add_item(al, 49);
  add_item(al, 49);

  print_al(*al);

  add_item(al, 70);
  print_al(*al);
  swap_items(al, 0, 2);
  print_al(*al);

  return 0;
}
