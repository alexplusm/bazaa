#include "libft.h"

// // t_list *ft_lstmap(t_list *lst, void *(*f)(void *),void (*del)(void *))
// // {
// //     // t_list *cursor;
    

// //     // cursor = lst;
// //     // while(cursor->next)
// //     // {
        
// //     //     cursor = cursor->next;
// //     // }
// // }

// static void	ft_del(void *content, size_t content_size)
// {
// 	(void)content_size;
// 	free(content);
// }

// t_list		*ft_lstmap(t_list *lst, t_list *(*f)(t_list *elem))
// {
// 	t_list *elem;
// 	t_list *prev;
// 	t_list *head;

// 	prev = NULL;
// 	head = NULL;
// 	if (f)
// 		while (lst)
// 		{
// 			if (!(elem = f(lst)))
// 			{
// 				if (head)
// 					ft_lstdel(&head, &ft_del);
// 				return (NULL);
// 			}
// 			if (prev)
// 				prev->next = elem;
// 			else
// 				head = elem;
// 			lst = lst->next;
// 			prev = elem;
// 		}
// 	return (head);
// }



t_list
	*ft_lstmap(t_list *lst, void *(*f)(void*), void (*del)(void *))
{
	t_list	*first;
	t_list	*new;

	if (!f || !del)
		return (NULL);
	first = NULL;
	while (lst)
	{
		if (!(new = ft_lstnew((*f)(lst->content))))
		{
			while (first)
			{
				new = first->next;
				(*del)(first->content);
				free(first);
				first = new;
			}
			lst = NULL;
			return (NULL);
		}
		ft_lstadd_back(&first, new);
		lst = lst->next;
	}
	return (first);
}