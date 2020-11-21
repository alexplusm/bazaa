#include "libft.h"

void *ft_memccpy(void *restrict dst, const void *restrict src, int c, int n)
{
    int i;
    const unsigned char *ptr_s;
   
    ptr_s = src;
    i = 0;
    while(ptr_s[i] != c && i < n)
        i+=1;
    
    return ft_memcpy(dst, src, i);
}
