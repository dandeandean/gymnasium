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

int double_capacity_plus_one(ArrayList *al) {
  // realloc
  al->capacity = 2 * al->capacity + 1;
  int *temp = (int *)realloc(al->data, sizeof(int) * al->capacity);
  if (temp == NULL) {
    // something horrible happened
    return EXIT_FAILURE;
  }
  al->data = temp;
  return EXIT_SUCCESS;
}

int add_item(ArrayList *al, int value) {
  int doubled = 0;
  if (al->used + 1 > al->capacity) {
    double_capacity_plus_one(al);
    doubled = 1;
  }
  al->data[al->used] = value;
  al->used++;
  return doubled;
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

int swap_items(ArrayList *al, int i1, int i2) {
  if (i1 == i2 || i1 < 0 || i2 < 0) {
    return EXIT_FAILURE;
  }
  if (i1 > al->used || i2 > al->used) {
    return EXIT_FAILURE;
  }
  int i1_temp = al->data[i1];
  al->data[i1] = al->data[i2];
  al->data[i2] = i1_temp;
  return EXIT_SUCCESS;
}
