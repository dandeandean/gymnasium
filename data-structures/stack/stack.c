#include "stack.h"
#include <stdlib.h>

Stack *stack_init(void) {
  Stack *s = (Stack *)malloc(sizeof(Stack));
  s->data = alinit(1);
  s->index_top = 0;
  return s;
}

int stack_push(Stack *s, int val) {
  add_item(s->data, val);
  s->index_top++;
  return 0;
}

int stack_pop(Stack *s) {
  int out = s->data->data[s->data->used];
  s->data->data[s->data->used] = 0;
  s->data->used--;
  s->index_top--;
  return out;
}

void stack_print(Stack s) {
  printf("|");
  print_al(*s.data);
}
