#include "libft.h"

void ft_lstclear(t_list **lst, void (*del)(void*))
{
    t_list *cursor;

    if (lst == NULL || *lst == NULL)
        return ;
    cursor = *lst;
    while (cursor)
    {
        del(cursor);
        cursor = cursor->next;
    }
    *lst = NULL;
}
