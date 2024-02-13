#include <stdio.h>
#include <stdlib.h>

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

Node * node_init(int value, Node * next, Node * prev){
    Node * node = (Node*)malloc(sizeof(Node));
    node->value = value;
    node->next = next;
    node->prev = prev;
    return node;
}

List * list_init(int first_val){	
    List * list = (List*)malloc(sizeof(List));
    Node * node = node_init(first_val,NULL,NULL);
    list->size = 1;
    list->head = node;
    list->tail = node;
    return list;
}

void free_list(List * list){
    /*TODO go through and free all the nodes*/
    Node * cur = list->head;
    while (cur != NULL) {
       free(cur->prev);
       cur = cur->next;
    } 
    free(list);
}

void print_list(List list){
    printf("(%d)",list.head->value);
    Node * cur = list.head->next;
    while (cur != NULL) {
       printf("<->(%d)",cur->value);
       cur = cur->next;
    } 
    printf("\n");
}
void add_item(List * list,int value){
    Node * new_node = node_init(value,NULL,list->tail);
    list->tail->next = new_node;
    list->tail = new_node;
    list->size ++;
}

