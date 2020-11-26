#include "libft.h"

int ft_lstsize(t_list *lst)
{
    int cnt;
    t_list *cursor;
    
    cnt = 0;
    if (lst == NULL)
        return cnt;
    cursor = lst;
    while (cursor)
    {
        cursor = cursor->next;
        cnt++;
    }
    return cnt;
}
