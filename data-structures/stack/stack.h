#ifndef STACK_H
#define STACK_H
#include "../array-list/array_list.c"

typedef struct Stack
{
    /* The bottom will be at index zero */
    int index_top;
    ArrayList * data;
}Stack;

int stack_push(Stack *s, int val);
int stack_pop(Stack *s);
void stack_print(Stack s);
#endif