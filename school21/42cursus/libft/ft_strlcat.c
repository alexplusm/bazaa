#include "libft.h"

int ft_strlcat(
    char * restrict dst, const char * restrict src, ft_size_t dstsize
    )
{
    ft_size_t i, j;
    int res;

    if (dstsize == 0)
        return dstsize + ft_strlen(src);
    i = 0;
    while (dst[i] != '\0')
        i += 1;
    res = i + ft_strlen(src);
    if (i >= dstsize - 1)
        return res;
    j = 0;
    while (src[j] != '\0' && i < dstsize - 1)
    {
        dst[i] = src[j];
        i += 1;
        j += 1;
    }
    return res;
}