#include "libft.h"

void ft_lstiter(t_list *lst, void (*f)(void *))
{
    t_list *cursor;

    if (lst == NULL || f == NULL)
        return ;
    cursor = lst;
    while (cursor->next)
    {
        f(cursor->content);
        cursor = cursor->next;
    }
}
