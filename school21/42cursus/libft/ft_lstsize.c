/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_lstsize.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:11:47 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:11:48 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

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
