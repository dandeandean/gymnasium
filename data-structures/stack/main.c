#include "stack.c"
#include "stack.h"
#include <stdio.h>

int main(int argc, char **argv) {
  // printf("%s", argv[0]);
  if (argc > 2) {
    printf("not impl\n");
  }
  Stack *s = stack_init();
  stack_push(s, 6);
  stack_push(s, 9);
  stack_push(s, 4);
  stack_push(s, 2);
  stack_push(s, 0);
  stack_print(*s);
  stack_pop(s);
  stack_pop(s);
  stack_pop(s);
  stack_print(*s);
  stack_push(s, 0);
  stack_push(s, 0);
  stack_print(*s);
  return 0;
}
