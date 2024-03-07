#include "array_list.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

ArrayList *alinit(int capacity) {
  int *data = (int *)malloc(sizeof(int) * capacity);
  ArrayList *out = (ArrayList *)malloc(sizeof(ArrayList));
  out->capacity = capacity;
  out->data = data;
  out->used = 0;
  return out;
}

int *double_capacity(ArrayList *al) {
  int *old_data = al->data;
  int old_capacity = al->capacity;
  // move data to new spot
  al->capacity *= 2;
  al->data = (int *)malloc(sizeof(int) * al->capacity);
  // copy over contents
  memcpy(al->data, old_data, sizeof(int) * old_capacity);
  return 0;
}

void add_item(ArrayList *al, int value) {
  if (al->used + 1 > al->capacity) {
    double_capacity(al);
  }
  al->data[al->used] = value;
  al->used++;
}

void print_al(ArrayList al) {
  printf("{");
  for (int i = 0; i < al.capacity; i++) {
    if (i < al.used) {
      printf("%d ", al.data[i]);
    } else {
      printf("_ ");
    }
  }
  printf("}\n");
}
