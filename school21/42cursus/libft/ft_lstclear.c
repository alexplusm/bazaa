/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_lstclear.c                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:10:44 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:10:45 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

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
