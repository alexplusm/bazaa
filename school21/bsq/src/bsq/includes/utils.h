
#ifndef UTILS_H
# define UTILS_H
# include <stdlib.h>
# include "shared.h"

char	*ft_extend_array(char *orig, char *n_cont, unsigned int old_len, unsigned int len);

char	*ft_str_n_copy(char *dest, char *src, int n);

t_bool	ft_atoi(char *str, unsigned int n, unsigned int *result);

#endif
