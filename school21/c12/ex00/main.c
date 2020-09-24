#include "ft_list.h"
#include "ft_create_elem.c"
#include <stdio.h>

int main()
{
    char *str = "HALLO";
    t_list *list_item;

    list_item = ft_create_elem(str);

    printf("%s\n", list_item->data);
    printf("%p\n", list_item->next);
}