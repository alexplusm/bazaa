/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_atoi.c                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:08:22 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:08:27 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

#include <stdio.h>

int		ft_atoi(const char *str)
{
	unsigned long	result;
	unsigned long	border;
	size_t			i;
	int				sign;

	printf("1: %zu\n", ~0l);
	printf("2: %zu\n", (~0l) >> 1);
	printf("3: %lu\n", (unsigned long)((long)(((unsigned long)(~0l)) >> 1) / 10));

	result = 0;
	border = (unsigned long)(FT_LONG_MAX / 10);

	printf("4: %lu\n", border);

	i = 0;
	while (ft_isspace(str[i]))
		i++;
	sign = (str[i] == '-') ? -1 : 1;
	if (str[i] == '-' || str[i] == '+')
		i++;
	while (ft_isdigit(str[i]))
	{
		if ((result > border || (result == border && (str[i] - '0') > 7))
															&& sign == 1)
			return (-1);
		else if ((result > border || (result == border && (str[i] - '0') > 8))
																&& sign == -1)
			return (0);
		result = result * 10 + (str[i++] - '0');
	}
	return ((int)(result * sign));
}
