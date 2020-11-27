/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_lstiter.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:10:59 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:11:02 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

// void ft_lstiter(t_list *lst, void (*f)(void *))
// {
//     t_list *cursor;

//     if (lst == NULL || f == NULL)
//         return ;
//     cursor = lst;
//     while (cursor->next)
//     {
//         f(cursor->content);
//         cursor = cursor->next;
//     }
// }

void	ft_lstiter(t_list *lst, void (*f)(void *))
{
	if (f)
		while (lst)
		{
			f(lst->content);
			lst = lst->next;
		}
}
