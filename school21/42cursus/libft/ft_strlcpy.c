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
