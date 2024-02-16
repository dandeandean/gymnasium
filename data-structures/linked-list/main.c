#include <stdio.h>
#include "linked_list.c"

int main(int argc, char *argv[]){
    if (argc > 1) {
        List * plist = list_init(atoi(argv[1]));
        int i = 2;
        while (i < argc) {
            add_item(plist,atoi(argv[i]));
            i ++;
        }
        print_list(*plist);
        return 0;
    }
    List * plist = list_init(69);
    print_list(*plist);
    add_item(plist,420);
    add_item(plist,500);
    add_item(plist,600);
    print_list(*plist);
    remove_item(plist,500);
    remove_item(plist,700);
    print_list(*plist);
    free_list(plist);
    return 0;
}