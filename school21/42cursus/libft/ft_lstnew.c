#include "libft.h"

t_list *ft_lstnew(void *content)
{
    t_list *item;

    item = malloc(sizeof(t_list));
    if (item == NULL)
        return NULL;
    item->content = content;
    item->next = NULL;
    return item;
}
