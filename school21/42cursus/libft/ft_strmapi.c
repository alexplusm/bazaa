#include "libft.h"

char *ft_strmapi(char const *s, char (*f)(unsigned int, char))
{
    char	*str;
    size_t	i;

	if (s == NULL || f == NULL)
		return (NULL);
    str = ft_strdup(s);
    if (str == NULL)
        return (NULL);
    i = 0;
    while (str[i])
    {
        str[i] = f(i, str[i]);
        i++;
    }
    return str;
}
