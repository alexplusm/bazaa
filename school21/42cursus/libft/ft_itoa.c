/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_itoa.c                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:10:10 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:10:12 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

static int	ft_digits_cnt(unsigned int n)
{
	int res;

	res = 1;
	while (n /= 10)
		res++;
	return (res);
}

char		*ft_itoa(int n)
{
	unsigned int	num;
	int				negative;
	int				cnt;
	size_t			mem_cnt;
	char			*res;

	negative = (n < 0) ? 1 : 0;
	num = (n < 0) ? -n : n;
	cnt = ft_digits_cnt(num);
	mem_cnt = (negative) ? cnt + 2 : cnt + 1;
	if ((res = malloc(sizeof(char) * mem_cnt)) == NULL)
		return (NULL);
	if (negative)
		res[0] = '-';
	res[cnt + negative] = '\0';
	if (num == 0)
		res[0] = '0';
	while (num)
	{
		res[cnt + negative - 1] = (num % 10) + '0';
		num /= 10;
		cnt--;
	}
	return (res);
}
