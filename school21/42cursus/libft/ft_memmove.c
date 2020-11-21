#include "libft.h"

void *ft_memmove(void *dst, const void *src, int len)
{
    int i;
    const unsigned char *ptr_s;
    unsigned char *ptr_d;

    i = 0;
    ptr_d = dst;
    ptr_s = src;
    while(i < len)
    {
        ptr_d[i] = ptr_s[i];
        i += 1;
    } 
    
    return dst;
}