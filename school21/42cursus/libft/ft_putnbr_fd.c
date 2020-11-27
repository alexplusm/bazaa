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

static int ft_digits_cnt(unsigned int n)
{
	int res;

	res = 1;
	while (n /= 10)
		res++;
	return res;
}

static int ft_pow(int n, int pow)
{
	if (pow <= 0)
		return 1;
	return n * ft_pow(n, pow - 1);
}

void ft_putnbr_fd(int n, int fd)
{
	unsigned int	num;
	char ch;
	int pow;
	
	if (n < 0)
	{
		ft_putchar_fd('-', fd);
		num = -n;
	}
	else	
		num = n;
	
	int cnt = ft_digits_cnt(num);
	while (cnt--)
	{
		pow = ft_pow(10, cnt);
		ch = (num / pow) + '0';
		num %= pow;
		ft_putchar_fd(ch, fd);
	}
}
