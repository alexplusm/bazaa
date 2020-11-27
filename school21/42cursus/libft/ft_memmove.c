/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_memmove.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:12:42 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:12:44 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

void	*ft_memmove(void *dst, const void *src, int len)
{
	int				i;
	unsigned char	*ptr_s;
	unsigned char	*ptr_d;

	if (dst == NULL && src == NULL)
		return (NULL);
	i = 0;
	ptr_d = (unsigned char *)dst;
	ptr_s = (unsigned char *)src;
	if (ptr_d > ptr_s)
	{
		while (len--)
			ptr_d[len] = ptr_s[len];
	}
	else
	{
		while (i < len)
		{
			ptr_d[i] = ptr_s[i];
			i += 1;
		}
	}
	return (dst);
}
