/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_strlcpy.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:14:29 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:14:30 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

int ft_strlcpy(char *dst, const char *src, ft_size_t dstsize)
{
    ft_size_t i;

    if (dst == NULL || src == NULL)
        return 0;

    if (dstsize == 0)
        return ft_strlen(src);
    i = 0;
    while (i < dstsize-1 && src[i] != '\0')
    {
        dst[i] = src[i];
        i += 1;
    }
    dst[i] = '\0';
    return ft_strlen(src);
}
