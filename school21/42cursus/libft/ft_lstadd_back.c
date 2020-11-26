#include "libft.h"

void ft_lstadd_back(t_list **lst, t_list *new)
{
    t_list *cursor;

    if (lst == NULL)
        return ;
    if (*lst == NULL)
    {
        *lst = new;
        return ;
    }
    cursor = *lst;
    while (cursor->next)
        cursor = cursor->next;
    cursor->next = new;
}
