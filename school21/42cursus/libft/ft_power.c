/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_power.c                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 10:16:27 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 10:16:29 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

int		ft_power(int n, int power)
{
	if (power < 0)
		return (0);
	if (power == 0)
		return (1);
	return (n * ft_power(n, power - 1));
}
