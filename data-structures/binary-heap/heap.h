#ifndef HEAP_H
#define HEAP_H

#include "array_list.h"

/* Min Heap */
typedef struct {
  ArrayList *list;
  int depth;
  int item_count;
} Heap;

void heap_insert(Heap *h, int val);
int heap_pop(Heap *h);
Heap *heap_init(int capacity);

#endif
