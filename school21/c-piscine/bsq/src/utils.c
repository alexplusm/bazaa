#include "utils.h"

char	*ft_str_n_copy(char *dest, char *src, int n)
{
	int		index;

	index = 0;
	while (index < n && src[index] != '\0')
	{
		dest[index] = src[index];
		index++;
	}
	while (index < n)
	{
		dest[index] = '\0';
		index++;
	}
	return (dest);
}

char	*ft_extend_array(char *orig, char *n_cont, unsigned int old_len, unsigned int len)
{
	char *dest;

	if (!(dest = malloc((len + 1) * sizeof(char))))
		return (NULL);
	if (orig != NULL)
		ft_str_n_copy(dest, orig, old_len);
	ft_str_n_copy(dest + old_len, n_cont, (unsigned int)(len - old_len));
	if (orig != NULL)
		free(orig);
	return (dest);
}

t_bool	ft_atoi(char *str, unsigned int n, unsigned int *result)
{
	unsigned int	index;

	index = 0;
	*result = 0;
	while (index < n)
	{
		if ('0' <= str[index] && str[index] <= '9')
			*result = *result * 10 + str[index] - '0';
		else
			return (FALSE);
		index++;
	}
	return (TRUE);
}
