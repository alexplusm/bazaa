#include "libft.h"

// TODO
char *ft_strnstr(const char *haystack, const char *needle, ft_size_t len)
{
    ft_size_t i, j;

    i = 0;
    while(haystack[i] != '\0')
    {
        j = 0;
        while (haystack[i + j] == needle[j])
            j++;
        if (needle[j] == '\0' || j == len-1)
            return (char *)(haystack) + i;
        i++;
    }
    return 0;
}
