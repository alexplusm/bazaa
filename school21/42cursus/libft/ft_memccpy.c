#include "libft.h"

void *ft_memccpy(void *dst, const void *src, int c, int n)
{
    unsigned char	stop_char;
	unsigned char	*d;
	unsigned char	*s;

	stop_char = (unsigned char)c;
	d = (unsigned char *)dst;
	s = (unsigned char *)src;
	while (n--)
	{
		*d++ = *s++;
		if (*(d - 1) == stop_char)
			return (d);
	}
	return (NULL);
}
