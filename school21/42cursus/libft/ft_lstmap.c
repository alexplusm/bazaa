/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_lstmap.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:11:16 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:11:18 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

t_list *ft_lstmap(t_list *lst, void *(*f)(void *),void (*del)(void *))
{
	t_list *new_list;
	t_list *cursor;
	t_list *item;

	if (lst == NULL || f == NULL)
		return (NULL);
	cursor = lst;
	if ((item = ft_lstnew(f(cursor->content))) == NULL)
		return (NULL);
	new_list = item;
	cursor = cursor->next;
	while (cursor)
	{
		if ((item = ft_lstnew(f(cursor->content))) == NULL)
		{
			if (del != NULL)
				ft_lstclear(&new_list, del);
			return (NULL);
		}
		ft_lstadd_back(&new_list, item);
		cursor = cursor->next;
	}		
	return new_list;
}
