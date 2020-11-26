#include "libft.h"

#include "stdio.h"

void ft_lstclear(t_list **lst, void (*del)(void*))
{
    // t_list *cursor;

    // if (lst == NULL || *lst == NULL)
    //     return ;
    // cursor = *lst;
    // while (cursor)
    // {
    //     // printf("ppp: %p | %s\n", cursor->content, (char*)cursor->content);
    //     del(cursor->content);
    //     cursor = cursor->next;
    // }
    // free(*lst);
    // // cursor = NULL;
    // *lst = NULL;



    t_list	*tmp;

	if (!del || !lst || !*lst)
		return ;
	while (lst && *lst)
	{
		tmp = (*lst)->next;
		ft_lstdelone(*lst, del);
		*lst = tmp;
	}
    // ft_lstdelone(*lst, del);

    // if (lst && *lst && del)
	// {
	// 	if ((*lst)->next)
	// 		ft_lstclear(&(*lst)->next, del);
    //     // else
    //         // free(*lst);
	// 	ft_lstdelone((*lst), del);
	// }
    

}


// void	ft_lstdel(t_list **alst, void (*del)(void *, size_t))
// {
// 	if (alst && *alst && del)
// 	{
// 		if ((*alst)->next)
// 			ft_lstdel(&(*alst)->next, del);
// 		ft_lstdelone(&(*alst), del);
// 	}
// }
