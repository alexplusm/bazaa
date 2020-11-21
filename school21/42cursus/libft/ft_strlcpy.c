#include "libft.h"

//  TODO!
//  TODO: протестить, 
// если придет src без терминального нуля (на библитечной функции)
int ft_strlcpy(char *restrict dst, const char *restrict src, int dstsize)
{
    int i;
    unsigned char *ptr;

    if (dstsize == 0)
        return 0;

    i = 0;
    ptr = (unsigned char*)src;

    while (i < dstsize-1 && ptr[i] != '\0')
    {
        dst[i] = ptr[i];
        i += 1;
    }
    dst[i] = '\0';
    return ft_strlen(src); // TODO: нужно ли добавить + 1?
    // посмотреть man | и библиотечную функцию
}
