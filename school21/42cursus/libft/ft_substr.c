#include "libft.h"

char	*ft_substr(char const *s, unsigned int start, size_t len)
{
	size_t			i;
	unsigned char	*ptr;
	char 			*substr;

	i = 0;
	ptr = (unsigned char *)s;
	if (ft_strlen(s) <= start)
	{
		substr = malloc(1);
    	if (substr == NULL)
			return NULL;
		substr[0] = '\0';
		return substr;
	}
	// size_t str_len = ft_strlen(s)
	while (i < len && *(ptr + start + i) != '\0')
		i++;
	substr = malloc(i + 1);
    if (substr == NULL)
		return NULL;
	substr[i + 1] = '\0';
    while (i--)
		substr[i] = ptr[start + i];
	return substr;
}
