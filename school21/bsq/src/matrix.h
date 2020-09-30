#ifndef MATRIX_H

# include <stdlib.h>

# define MATRIX_H

# define MATRIX_CONFIG_STR_LENGTH 3

typedef struct s_matrix_config {
	unsigned int	max_i;
	unsigned int	max_j;
	char			empty;
	char			obstacle;
	char			full;
}					t_matrix_config;

t_matrix_config *create_matrix_config(int max_i, int max_j);

void fill_matrix_config(t_matrix_config *matrix_config, char *config);

int ft_matrix_symbol_is_valid(char c, t_matrix_config *matrix_config);

#endif
