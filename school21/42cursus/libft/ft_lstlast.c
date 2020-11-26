#include "libft.h"

t_list *ft_lstlast(t_list *lst)
{
    t_list *cursor;

    if (lst == NULL)
        return (NULL);
    cursor = lst;
    while (cursor->next)
        cursor = cursor->next;
    return cursor;
}
