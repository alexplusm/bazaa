#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdio.h>

#include "matrix.h"


//----- FOR TESTING

void print_matrix_config(t_matrix_config *matrix_config)
{
	printf("MATRIX CONFIG ---- \n");
	printf("max_i: %d | max_j: %d\n", matrix_config->max_i, matrix_config->max_j);
	printf("empty: %c | obstacle: %c | fill: %c |\n", matrix_config->empty, matrix_config->obstacle, matrix_config->full);
	printf("END MATRIX CONFIG \n");
}


void	ft_putstr(char *str)
{
	int i;

	i = 0;
	while (str[i] != '\0')
	{
		write(1, (str + i), 1);
		i += 1;
	}
}


void ft_print_table(char **table, int size)
{
    int i;

    i = 0;
    while (i < size)
    {
        ft_putstr(table[i]);
        // ft_putstr("\n");
        i += 1;
    }
}

//-----




void print_error()
{
	write(2, "map error\n", 10);
}

int ft_open(char *filename)
{
	return (open(filename, O_RDONLY));
}

int ft_reopen_file(int fd, char *filename)
{
	close(fd);
	return (ft_open(filename));
}

int decade_power(int power)
{
	int i;
	int result;

	result = 1;
	i = 0;
	while (i < power)
	{
		result *= 10;
		i += 1;
	}
	return result;
}

int atoi_kek(char *str)
{
	int result;
	int len;
	int i;

	result = 0;
	i = 0;
	len = 0;
	while (str[len] != '\0')
		len += 1;
	// todo:  proccess another chars
	while (str[i] != '\0')
	{
		result = result + (str[i] - '0') * decade_power(len - i - 1);
		i += 1;
	}
	return result;
}

char *ft_read(int fd, int length)
{
	char	buff;
	char	*result;
	int		read_bytes;
	int		i;

	i = 0;
	read_bytes = 1;
	result = malloc(sizeof(char) * (length + 1));
	if (result == NULL)
		return (NULL);
	while (i < length && read_bytes)
	{
		read_bytes = read(fd, &buff, 1);
		result[i] = buff;
		i += 1;
	}
	result[i] = '\0';
	return (i == length ? result : NULL);	// error_map
}

void ft_skip_first_line_in_file(int fd)
{
	char	buff;
	int		read_bytes;

	read_bytes = 1;
	while (read_bytes != 0 && buff != '\n')
		read_bytes = read(fd, &buff, 1);
}

int count_lines_size(int fd, int *first_line_size, int *second_line_size)
{
	char	buff;
	int		read_bytes;
	int		char_number_in_first_row;
	int		char_number_in_second_row;
	int		first_line_counted;
	
	first_line_counted = 0;
	read_bytes = 1;
	char_number_in_first_row = 0;
	char_number_in_second_row = 0;
	while (read_bytes != 0)
	{
		read_bytes = read(fd, &buff, 1);
		if (read_bytes == -1)
			return (-1);
		if (!first_line_counted)
			char_number_in_first_row += 1;
		else
			char_number_in_second_row += 1;
		if (buff == '\n')
		{
			if (!first_line_counted)
				first_line_counted = 1;
			else
				break ;
		}
	}
	// проверка валидности посчитанных значений?
	*first_line_size = char_number_in_first_row - 1;
	*second_line_size = char_number_in_second_row - 1;
	return (0);
}

char **read_matrix(int fd, t_matrix_config *config)
{
	char **matrix;
	// char *row;
	int i;
	int j;

	i = 0;
	matrix = malloc(sizeof(char *) * config->max_i);
	ft_skip_first_line_in_file(fd);
	while (i < config->max_i)
	{
		matrix[i] = ft_read(fd, config->max_j + 1);
		// ft_read(fd, 1);
		i += 1;
		// todo: if error -> free
	}
	return matrix;
}

void resolve_map(char *filename)
{
	char	buff;
	int		fd;
	int		char_number_in_first_row;
	int		char_number_in_second_row;

	char_number_in_first_row = 0;
	char_number_in_second_row = 0;
	fd = ft_open(filename);
	if (fd < 0)
	{
		print_error(); // todo: may be выше вызывать это?
		return ;
	}
	count_lines_size(fd, &char_number_in_first_row, &char_number_in_second_row);
	fd = ft_reopen_file(fd, filename);
	char *nums = ft_read(fd, char_number_in_first_row - MATRIX_CONFIG_STR_LENGTH);
	char *str_matrix_config = ft_read(fd, MATRIX_CONFIG_STR_LENGTH);
	t_matrix_config *config = create_matrix_config(atoi_kek(nums), char_number_in_second_row);
	fill_matrix_config(config, str_matrix_config);
	// printf("count char first: %d | count char second: %d\n", char_number_in_first_row, char_number_in_second_row);
	// printf("nums = %d | config = %s \n", atoi_kek(nums), str_matrix_config);
	// print_matrix_config(config);
	// fd = ft_reopen_file(fd, filename);
	char **matrix = read_matrix(fd, config);
	ft_print_table(matrix, config->max_i);


	// ft_matrix_symbol_is_valid()
}

int main (int argc, char *argv[]) {
	int i;

	if (argc == 1) 
	{
		// todo: stdin
	}
	else 
	{
		i = 1;
		while (i < argc)
		{
			// resolve_map(argv[i]);
			i += 1;
		}
		resolve_map("map1.txt");
	}
	return 0;
}
