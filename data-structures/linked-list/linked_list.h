#ifndef LINKED_LIST_H
#define LINKED_LIST_H

#define TRUE 1
#define FALSE 0

typedef struct Node{
    int value;
    struct Node * prev;
    struct Node * next;
} Node;

typedef struct{
    Node * head;
    Node * tail;
    int size;
}List;


#endif