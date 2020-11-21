#include "libft.h"

void *ft_memcpy(void *restrict dst, const void *restrict src, int len)
{
    unsigned char *ptr_d;
    const unsigned char *ptr_s;

    if (len <= 0)
        return dst;

    ptr_d = dst;
    ptr_s = src;
    while (len > 0)
    {
        *ptr_d++ = *ptr_s++;
        len -= 1;
    }
    return dst;
}
