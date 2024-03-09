#include "heap.h"
#include "array_list.h"

Heap *heap_init(int capacity) {
  Heap *out = (Heap *)malloc(sizeof(Heap));
  out->list = alinit(capacity);
  out->depth = 0;
  out->item_count = 0;
  return out;
}

void heap_insert(Heap *h, int val) {}

int heap_pop(Heap *h) { return h->list->data[0]; }

void print_heap(Heap heap) {
  int twos_cap = 1;
  int twos_counter = 1;
  for (int i = 0; i < heap.list->capacity; i++) {
    if (i < heap.list->used) {
      printf("%d ", heap.list->data[i]);
    } else {
      printf("_ ");
    }
    twos_counter++;
    if (twos_counter >= twos_cap) {
      /* Should always end on a print \n */
      printf("\n");
      twos_counter = 0;
      twos_cap *= 2;
    }
  }
}
