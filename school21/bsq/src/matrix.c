#include "matrix.h"

t_matrix_config *create_matrix_config(int max_i, int max_j)
{
    t_matrix_config *matrix_config;
    matrix_config = malloc(sizeof(t_matrix_config));
    if (matrix_config == NULL)
        return (NULL);
    matrix_config->max_i = max_i;
    matrix_config->max_j = max_j;
    return (matrix_config);
}

void fill_matrix_config(t_matrix_config *matrix_config, char *config)
{
    matrix_config->empty = config[0];
    matrix_config->obstacle = config[1];
    matrix_config->full = config[2];
}

int ft_matrix_symbol_is_valid(char c, t_matrix_config *matrix_config)
{
    return (c == matrix_config->empty || c == matrix_config->obstacle);
}
