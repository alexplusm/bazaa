#include "libft.h"

void ft_lstclear(t_list **lst, void (*del)(void*))
{
    // t_list *cursor;

    // if (lst == NULL || *lst == NULL)
    //     return ;
    // cursor = *lst;
    // while (cursor->next)
    // {
    //     del(cursor->content);
    //     cursor->next = cursor;
    // }
    // *lst = NULL;

    if (lst && *lst && del)
	{
		if ((*lst)->next)
			ft_lstclear(&(*lst)->next, del);
        else
            *lst = NULL;
		ft_lstdelone((*lst), del);
	}

}
