#include "heap.h"
#include "array_list.h"
#include <stdlib.h>

Heap *heap_init(int capacity) {
  Heap *out = (Heap *)malloc(sizeof(Heap));
  out->list = alinit(capacity);
  out->depth = 0;
  out->item_count = 0;
  return out;
}

void heap_insert(Heap *h, int val) {}

int heap_pop(Heap *h) { return h->list->data[0]; }
