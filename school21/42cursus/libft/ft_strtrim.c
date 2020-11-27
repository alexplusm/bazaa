/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_strtrim.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:15:24 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:15:26 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

int static ft_includes(char c, char const *str, size_t len)
{
    while (len--)
        if (str[len] == c)
            return 1;
    return 0;
}

size_t static ft_right_idx(
    char const *s1, char const *set, size_t s_len, size_t set_len
)
{
    while (s_len--)
        if (!ft_includes(s1[s_len], set, set_len))
            return s_len;
    return s_len + 1;
}

size_t static ft_left_idx(
    char const *s1, char const *set, size_t s_len, size_t set_len
)
{
    size_t i;

    i = 0;
    while (i < s_len)
    {
        if (!ft_includes(s1[i], set, set_len))
            return i;
        i++;
    }
    return i;
}

size_t static ft_get_memcnt_and_set_left_idx(
    char const *s1, char const *set, size_t *left_idx
)
{
    size_t set_len;
    size_t s_len;
    size_t right_i;
    size_t left_i;

    s_len = ft_strlen(s1);
    set_len = ft_strlen(set);
    right_i = ft_right_idx(s1, set, s_len, set_len);
    left_i = ft_left_idx(s1, set, s_len, set_len);
    *left_idx = left_i;
    return (left_i < right_i) ? s_len + 2 - left_i - (s_len - right_i) : 1;
}

char *ft_strtrim(char const *s1, char const *set)
{
    size_t	mem_cnt;
    size_t	left_i;
    char	*res;

	if (s1 == NULL || set == NULL)
		return (NULL);
    mem_cnt = ft_get_memcnt_and_set_left_idx(s1, set, &left_i);
    res = malloc(sizeof(char) * (mem_cnt));
    if (res == NULL)
        return (NULL);
    ft_strlcpy(res, s1 + left_i, mem_cnt);
    return (res);
}

