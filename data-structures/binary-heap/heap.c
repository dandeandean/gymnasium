#include "array_list.h"

typedef struct Node {
    Node * r_child;
    Node * l_child;
    int value;
} Node;

typedef struct {
    Node * root;
    int size;
    int depth;
} Heap;

void add_item(){}

void remove_item(){}