#include "libft.h"

void *ft_calloc(ft_size_t count, ft_size_t size)
{
    char *ptr;
    ft_size_t bytes_count;

    bytes_count = count * size;
    ptr = malloc(bytes_count);
    if (ptr == NULL)
        return NULL;
    while (bytes_count > 0)
    {
        bytes_count -= 1;
        ptr[bytes_count] = 0;
    }
    return ptr;
}
