/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_calloc.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:09:00 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:09:01 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "libft.h"

void	*ft_calloc(size_t count, size_t size)
{
	void	*ptr;
	size_t	bytes_count;

	bytes_count = count * size;
	ptr = malloc(bytes_count);
	if (ptr == NULL)
		return (NULL);
	ft_bzero(ptr, bytes_count);
	return (ptr);
}
