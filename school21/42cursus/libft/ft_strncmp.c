#include "libft.h"

int ft_strncmp(const char *s1, const char *s2, ft_size_t n)
{
    ft_size_t i;

    i = 0;
    while (s1[i] == s2[i] && (i + 1) < n)
        i++;
    return s1[i] - s2[i];
}
