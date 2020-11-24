#include "libft.h"

char *ft_strdup(const char *s1) 
{
    ft_size_t size;
    ft_size_t i;
    char *res;

    size = 0;
    i = 0;
    while (s1[size] != '\0')
        size += 1;
    res = malloc(sizeof(char) * size);
    if (res == NULL)
    {    
        errno = ENOMEM;
        return NULL;
    }
    while (i < size)
    {
        res[i] = s1[i];
        i += 1;
    }
    res[i] = '\0';
    return res;
}
