/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_strdup.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:14:01 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:14:03 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

char	*ft_strdup(const char *s1)
{
	size_t	size;
	size_t	i;
	char	*res;

	size = 0;
	i = 0;
	while (s1[size] != '\0')
		size += 1;
	res = malloc(sizeof(char) * (size + 1));
	if (res == NULL)
	{
		return (NULL);
	}
	while (i < size)
	{
		res[i] = s1[i];
		i += 1;
	}
	res[i] = '\0';
	return (res);
}
