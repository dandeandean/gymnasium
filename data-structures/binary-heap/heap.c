#include "heap.h"
#include "array_list.h"
#include <stdio.h>

Heap *heap_init(int capacity) {
  Heap *out = (Heap *)malloc(sizeof(Heap));
  out->list = alinit(capacity);
  out->depth = 0;
  out->item_count = 0;
  return out;
}

/* Returns 0 if all is well, 1 if it couldn't perform */
void min_heapify(Heap *h, int i) {
  int smallest = i;
  int left_child = 2 * i + 1;
  int right_child = left_child + 1;
  if (h->list->data[right_child] < h->list->data[smallest] &&
      right_child < h->item_count) {
    swap_items(h->list, smallest, right_child);
  }
  if (h->list->data[left_child] < h->list->data[smallest] &&
      left_child < h->item_count) {
    swap_items(h->list, smallest, left_child);
  }
  printf("smallest: %d, i: %d\n", smallest, i);
  if (smallest != i) {
    min_heapify(h, smallest);
  }
}
void heap_insert(Heap *h, int val) {
  // h->list->data[h->item_count] = val;
  add_item(h->list, val);
  h->item_count++;
  for (int i = h->item_count / 2; i >= 0; i--) {
    printf("calling heap on %d...\n", i);
    min_heapify(h, i);
  }
}

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
