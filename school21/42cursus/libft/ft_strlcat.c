/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_strlcat.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:14:21 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:14:23 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

size_t ft_strlcat(char *dst, const char *src, size_t dstsize)
{
    size_t i, j;
    size_t res;
    
    if (dstsize == 0)
        return dstsize + ft_strlen(src);
    i = 0;
    
    while (i < dstsize)
    {
        if (dst[i] == '\0')
            break;
        i += 1;
    }
    res = i + ft_strlen(src);
    if (i >= dstsize - 1)
        return res;
    j = 0;
    while (src[j] != '\0' && i < dstsize - 1)
    {
        dst[i] = src[j];
        i += 1;
        j += 1;
    }
    dst[i] = '\0';
    return res;
}
