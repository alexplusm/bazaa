/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_lstadd_back.c                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:10:22 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:10:24 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

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
