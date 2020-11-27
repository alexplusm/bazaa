/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_memcpy.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:12:34 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:12:36 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

void	*ft_memcpy(void *dst, const void *src, size_t len)
{
	unsigned char		*ptr_d;
	unsigned char		*ptr_s;

	if (len <= 0 || (dst == NULL && src == NULL))
		return (dst);
	ptr_d = (unsigned char *)dst;
	ptr_s = (unsigned char *)src;
	while (len > 0)
	{
		*ptr_d++ = *ptr_s++;
		len -= 1;
	}
	return (dst);
}
