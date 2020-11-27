/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_strnstr.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:15:08 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:15:10 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

// TODO
char *ft_strnstr(const char *haystack, const char *needle, ft_size_t len)
{
    ft_size_t i, j;

    // TODO: костыыыль
    if (needle[0] == '\0' && len == 0)
        return (char *)haystack;

    i = 0;
    while(haystack[i] != '\0' && i < len)
    {
        j = 0;
        while ((i + j) < len && haystack[i + j] == needle[j] && needle[j] != '\0')
        {
            j++;
        }
        if (needle[j] == '\0')
            return (char *)(haystack) + i;

        i++;
    }
    return 0;
}
