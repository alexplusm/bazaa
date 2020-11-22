#include "libft.h"

int ft_strlcpy(
    char *restrict dst, const char *restrict src, ft_size_t dstsize
    )
{
    ft_size_t i;

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
