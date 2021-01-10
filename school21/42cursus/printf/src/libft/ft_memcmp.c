/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_memcmp.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:12:27 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:12:29 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

int		ft_memcmp(const void *s1, const void *s2, size_t n)
{
	size_t			i;
	unsigned char	*ptr_1;
	unsigned char	*ptr_2;

	i = 0;
	ptr_1 = (unsigned char*)s1;
	ptr_2 = (unsigned char*)s2;
	while (i < n && ptr_1[i] == ptr_2[i])
		i += 1;
	return ((i == n) ? 0 : ptr_1[i] - ptr_2[i]);
}
