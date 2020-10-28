#include "ft_list.h"
#include "ft_create_elem.c"
#include <stdio.h>


#include "ft_list_push_front.c"

void print_list(t_list **begin_list)
{
    int id = 0;
    t_list *cursor;

    cursor = *begin_list;

    while (cursor->next)
    {
        printf("id: %d | data: %s | next: %p \n", id, cursor->data, cursor->next);
        cursor = cursor->next;
        id += 1;
    }
    printf("id: %d | data: %s | next: %p \n", id, cursor->data, cursor->next);
    printf("---\n");
}

int main()
{
    t_list *list_item;

    list_item = ft_create_elem("KEKUS");
    list_item->next = ft_create_elem("AZAZA");

    print_list(&list_item);
    ft_list_push_front(&list_item, "WTF");
    print_list(&list_item);
}

