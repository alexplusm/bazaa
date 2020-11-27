/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_memcpy.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:12:34 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:12:36 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

void *ft_memcpy(void *dst, const void *src, int len)
{
    char *ptr_d;
    const char *ptr_s;

    if (len <= 0 || (dst == NULL && src == NULL))
        return dst;

    ptr_d = dst;
    ptr_s = src;
    while (len > 0)
    {
        *ptr_d++ = *ptr_s++;
        len -= 1;
    }
    return dst;
}
