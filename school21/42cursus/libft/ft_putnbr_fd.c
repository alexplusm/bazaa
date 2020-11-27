/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_putnbr_fd.c                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:13:26 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:13:27 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

static int	ft_digits_cnt(unsigned int n)
{
	int	res;

	res = 1;
	while (n /= 10)
		res++;
	return (res);
}

void		ft_putnbr_fd(int n, int fd)
{
	unsigned int	num;
	char			ch;
	int				pow;
	int				cnt;

	if (n < 0)
	{
		ft_putchar_fd('-', fd);
		num = -n;
	}
	else
		num = n;
	cnt = ft_digits_cnt(num);
	while (cnt--)
	{
		pow = ft_power_bonus(10, cnt);
		ch = (num / pow) + '0';
		num %= pow;
		ft_putchar_fd(ch, fd);
	}
}
