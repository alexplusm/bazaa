#include "libft.h"

// TODO
int ft_strlcat(
    char * restrict dst, const char * restrict src, ft_size_t dstsize
    )
{
    int i;
    unsigned char *ptr;

    if (dstsize == 0) {
        return 665; // TODO: lengths
    }

    i = 0;
    ptr = (unsigned char*) src;

    while (dst[i] != '\0')
    {
        if (i == dstsize)
        {

            return 666;
        }
    }
    return 0;
}