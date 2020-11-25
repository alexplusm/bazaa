#include "libft.h"

int ft_strncmp(const char *s1, const char *s2, ft_size_t n)
{
    ft_size_t       i;
    unsigned char   *ptr_1;
    unsigned char   *ptr_2;

    i = 0;
    ptr_1 = (unsigned char *) s1;
    ptr_2 = (unsigned char *) s2;
    while ((ptr_1[i] != '\0' || ptr_2[i] != '\0') && (i) < n)
    {
        if (ptr_1[i] != ptr_2[i])
            return (ptr_1[i] - ptr_2[i]);
        i++;
    }
    return (0);
}
