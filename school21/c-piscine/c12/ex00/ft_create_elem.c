#include "ft_list.h"

t_list *ft_create_elem(void *data)
{
    t_list *list_item;

    list_item = malloc(sizeof(t_list));
    if (list_item == NULL)
        return NULL;
    list_item->data = data;
    list_item->next = NULL;
    return list_item;
}