#include "libft.h"

#include <stdio.h>

int ft_includes(char c, char const *str)
{
    size_t  l;
    
    l = ft_strlen(str);
    while (l--)
    {
        if (str[l] == c)
            return 1;
    }
    return 0;
}

size_t ft_right_idx(char const *s1, char const *set)
{
    size_t s_len;

    s_len = ft_strlen(s1);
    while (s_len--)
    {
        if (!ft_includes(s1[s_len], set))
            return s_len;
    }
    return s_len;
}

size_t ft_left_idx(char const *s1, char const *set)
{
    size_t i;
    size_t s_len;

    s_len = ft_strlen(s1);
    i = 0;
    while (i < s_len)
    {
        if (!ft_includes(s1[i], set))
            return i;
        i++;
    }
    return i;
}

char *ft_strtrim(char const *s1, char const *set)
{
    size_t set_len;
    size_t s_len;
    size_t i;
    size_t mem_cnt;
    char *res;
    unsigned char *ptr_s;
    size_t right_idx;
    size_t left_idx;

    ptr_s = (unsigned char *)s1;

    s_len = ft_strlen(s1);
    set_len = ft_strlen(set);

    right_idx = ft_right_idx(s1, set);
    left_idx = ft_left_idx(s1, set);
    
    // printf("len: %zu | left_idx: %zu | right_idx: %zu\n", s_len, left_idx, right_idx);
    
    mem_cnt = s_len + 1 - left_idx - (s_len - right_idx);
    // printf("mem_cnt: %zu\n", mem_cnt);
    i = 0;
    res = malloc(sizeof(char) * mem_cnt);
    if (res == NULL)
        return (NULL);
    ptr_s += left_idx;
    while (i < mem_cnt)
    {
        res[i] = ptr_s[i];
        i++;
    }
    res[i] = 0;
    return res;
}

