/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_lstnew.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:11:39 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:11:41 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

t_list *ft_lstnew(void *content)
{
    t_list *item;

    item = malloc(sizeof(t_list));
    if (item == NULL)
        return NULL;
    item->content = content;
    item->next = NULL;
    return item;
}
