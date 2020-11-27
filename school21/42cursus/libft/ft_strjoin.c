/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_strjoin.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:14:09 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:14:11 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

char	*ft_strjoin(char const *s1, char const *s2)
{
	size_t	s1_size;
	size_t	s2_size;
	char	*res;

	if (s1 == NULL || s2 == NULL)
		return (NULL);
	s1_size = ft_strlen(s1);
	s2_size = ft_strlen(s2);
	if ((res = malloc(sizeof(char) * (s1_size + s2_size + 1))) == NULL)
		return (NULL);
	ft_strlcpy(res, s1, s1_size + 1);
	ft_strlcpy(res + s1_size, s2, s2_size + 1);
	return (res);
}
