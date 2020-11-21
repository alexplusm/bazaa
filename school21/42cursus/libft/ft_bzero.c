#include "libft.h"

void ft_bzero(void *s, int len)
{
    unsigned char *ptr;
    if (len <= 0)
        return;

    ptr=s;
    // TODO: SegFault
    while(len > 0)
    {
        *(ptr++) = '\0';
        len -= 1;
    }
}
