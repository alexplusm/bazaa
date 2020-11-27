/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_memset.c                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: cdeon <marvin@42.fr>                       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/11/27 09:12:55 by cdeon             #+#    #+#             */
/*   Updated: 2020/11/27 09:12:57 by cdeon            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

void *tf_memset(void *b, int c, int len)
{
    // int i;

    // // todo: while
    // for (i = 0; i < len; i++) {
    //     ((int *)b)[i] = c;
    //     printf("i: %d | value: %c | char: %c \n", i, ((char *)b)[i], c);
    // }
    // return b;

    unsigned char *ptr = b;
    while (len-- > 0)
        *ptr++ = c;
    return b;
}
