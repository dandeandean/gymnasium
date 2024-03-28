#ifndef ARRAY_LIST_H
#define ARRAY_LIST_H

typedef struct {
  int *data;
  int used;
  int capacity;
} ArrayList;

ArrayList *alinit(int capacity);
int *double_capacity_plus_one(ArrayList *al);
int add_item(ArrayList *al, int value);
void print_al(ArrayList al);
int swap_items(ArrayList *al, int i1, int i2);
#endif
