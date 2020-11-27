/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_substr.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:15:32 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:15:33 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

char	*ft_substr(char const *s, unsigned int start, size_t len)
{
	char *substr;
	char	*result;
	size_t	i;

	if (s == NULL)
		return NULL;
	if (ft_strlen(s) <= start)
	{
		substr = malloc(1);
    	if (substr == NULL)
			return NULL;
		substr[0] = '\0';
		return substr;
	}
	i = 0;
	if (!s || start + len > ft_strlen(s))
		return (NULL);
	if ((result = malloc(sizeof(char) * len + 1)) == NULL)
		return NULL;
	while (len)
	{
		result[i++] = s[start++];
		len--;
	}
	result[i] = '\0';
	return (result);
}
