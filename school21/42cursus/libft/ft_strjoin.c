#include "libft.h"

char *ft_strjoin(char const *s1, char const *s2)
{
    size_t  s1_size;
    size_t  s2_size;
    size_t  i;
    char    *res;

    s1_size = ft_strlen(s1);
    s2_size = ft_strlen(s2);
    res = malloc(sizeof(char) * (s1_size + s2_size + 1));
    if (res == NULL)
        return NULL;
    i = 0;
    while (i++ < s1_size)
        res[i] = s1[i];
    while (i++ < s1_size + s2_size)
        res[i] = s2[i - s1_size];
    res[i+1] = '\0';
    return res;
}
