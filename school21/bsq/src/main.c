#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdlib.h>

#include "shared.h"
#include "utils.h"

typedef struct	s_grid {
	unsigned int	max_j;
	unsigned int	max_i;
	char			*source;
	unsigned char	**map;
	unsigned char	translate[3];
}				t_grid;

//структура для решения
typedef struct	s_solution {
	unsigned int	size;
	unsigned int	x;
	unsigned int	y;
}				t_solution;

t_bool			ft_load_grid(char *path, t_grid *grid);
void			ft_free_grid(t_grid *grid);
void			ft_process_grid(t_grid *grid, t_solution *sol);
t_bool			ft_parse_header_line(t_grid *grid, char *line, unsigned int length);
t_bool			ft_parse_normal_line(t_grid *grid, char *line, unsigned int length);
t_bool			ft_parse_grid(int fd, t_grid *grid);
t_bool			ft_read_full(int fd, char **file_content, unsigned int *total);
t_bool			ft_has_width_changed(t_grid *grid, unsigned int *curr_w);
t_bool			ft_process_lines(t_grid *grid, unsigned int index, char *file_content, unsigned int total);

t_bool			find_solution(t_solution *sol, t_grid *grid);
void			ft_print_solution(t_grid *grid, t_solution *solution);
void			actualize_sol(t_solution *sol, t_solution *intent, t_grid *grid);
t_bool			fits(t_solution *sol, t_grid *grid);
t_bool			fits_succ(t_solution *intent, t_grid *grid);
t_bool			dont_fit_basic(t_solution *intent, t_grid *grid);

void ft_print_error()
{
	write(2, "map error\n", 10);
}

t_bool	ft_read_full(int fd, char **file_content, unsigned int *total)
{
	char				buffer[BUFFER_SIZE];
	unsigned int		byte_read;

	if (read(fd, 0, 0) == -1)
		return (FALSE);
	*total = 0;
	*file_content = NULL;
	while ((byte_read = read(fd, buffer, BUFFER_SIZE)) > 0)
	{
		if (!(*file_content = ft_extend_array(*file_content,
				buffer, *total, *total + byte_read)))
			return (FALSE);
		(*total) += byte_read;
	}
	if (byte_read == 0)
		return (TRUE);
	return (FALSE);
}

t_bool	ft_parse_header_line(t_grid *grid, char *line, unsigned int length)
{
	t_bool	result;

	if (length < 4)
		return (FALSE);


	grid->translate[FILL] = line[length - 1];
	grid->translate[OBSTACLE] = line[length - 2];
	grid->translate[EMPTY] = line[length - 3];
	result = ft_atoi(line, length - 3, &(grid->max_i));
	return (result);
}

t_bool	ft_parse_normal_line(t_grid *grid, char *line, unsigned int length)
{
	unsigned int	index;
	char	current;

	if (length < 1)
		return (FALSE);
	grid->max_j = length;
	index = 0;
	while (index < length)
	{
		current = line[index];
		if (current != grid->translate[EMPTY]
				&& current != grid->translate[OBSTACLE]
				&& current != grid->translate[FILL])
			return (FALSE);
		index++;
	}
	return (TRUE);
}

t_bool	ft_has_width_changed(t_grid *grid, unsigned int *curr_w)
{
	if (*curr_w == (unsigned int)-1)
		*curr_w = grid->max_j;
	return (*curr_w != grid->max_j);
}

t_bool	ft_process_lines(t_grid *grid, unsigned int index,
							char *file_content, unsigned int total)
{
	unsigned int	jndex;
	unsigned int	y;
	unsigned int	curr_w;

	y = 0;
	curr_w = -1;
	while (index < total && y < grid->max_i)
	{
		jndex = index;
		while (file_content[jndex] != '\n')
			if (jndex++ + 1 >= total)
				return (FALSE);
		if (ft_parse_normal_line(grid, file_content + index, jndex - index))
		{
			grid->map[y++] = (unsigned char *)(file_content + index);
			if (ft_has_width_changed(grid, &curr_w))
				return (FALSE);
		}
		else
			return (FALSE);
		index = jndex + 1;
	}
	return (y == grid->max_i);
}

t_bool	ft_parse_grid(int fd, t_grid *grid)
{
	char			*file_content;
	unsigned int	total;
	unsigned int	index;

	grid->map = 0;
	if (!ft_read_full(fd, &file_content, &total))
		return (FALSE);
	grid->source = file_content;
	index = 0;
	while (index < total)
	{
		if (file_content[index] == '\n')
		{
			if (!ft_parse_header_line(grid, file_content, index))
				return (FALSE);
			if (!(grid->map = malloc(sizeof(char *) * grid->max_i)))
				return (FALSE);
			return (ft_process_lines(grid, index + 1, file_content, total));
		}
		index++;
	}
	return (FALSE);
}
 
t_bool	ft_load_grid(char *filename, t_grid *grid)
{
	int		fd;
	t_bool	result;

	fd = open(filename, O_RDONLY);
	if (fd < 0)
		return (FALSE);
	result = ft_parse_grid(fd, grid);
	close(fd);
	return (result);
}

void	ft_free_grid(t_grid *grid)
{
	free(grid->map);
	free(grid->source);
}

void	ft_process_grid(t_grid *grid, t_solution *sol)
{
	unsigned int x;
	unsigned int y;

	if (find_solution(sol, grid))
	{
		x = sol->x;
		while (x < sol->x + sol->size)
		{
			y = sol->y;
			while (y < sol->y + sol->size)
			{
				grid->map[y][x] = grid->translate[FILL];
				y++;
			}
			x++;
		}
		y = 0;
		while (y < grid->max_i)
		{
			write(1, grid->map[y], grid->max_j);
			write(1, "\n", 1);
			y++;
		}
	}
}




t_bool			find_solution(t_solution *sol, t_grid *grid)
{
	unsigned int		i;
	unsigned int		j;
	t_solution	intent;

	i = 0;
	sol->size = 0;
	intent.size = 1;
	while (i < grid->max_i - sol->size)
	{
		j = 0;
		while (j < grid->max_j - sol->size)
		{
			intent.x = j;
			intent.y = i;
			actualize_sol(sol, &intent, grid);
			j++;
		}
		i++;
	}
	return (sol->size ? TRUE : FALSE);
}

void		actualize_sol(t_solution *sol, t_solution *intent,
// static void		actualize_sol(t_solution *sol, t_solution *intent,
					t_grid *grid)
{
	if (!fits(intent, grid))
		return ;
	sol->x = intent->x;
	sol->y = intent->y;
	sol->size++;
	intent->size++;
	while (fits_succ(intent, grid))
	{
		sol->size++;
		intent->size++;
	}
}

t_bool	fits(t_solution *intent, t_grid *grid)
{
	int i;
	int i_max;
	int j;
	int j_max;

	if (dont_fit_basic(intent, grid))
		return (FALSE);
	i = intent->y;
	i_max = i + intent->size;
	j_max = intent->x + intent->size;
	while (i < i_max)
	{
		j = intent->x;
		while (j < j_max)
		{
			if (grid->map[i][j] == grid->translate[OBSTACLE])
				return (FALSE);
			j++;
		}
		i++;
	}
	return (TRUE);
}

// static t_bool	fits_succ(const t_solution *intent, const t_grid *grid)
t_bool	fits_succ(t_solution *intent, t_grid *grid)
{
	int i;
	int j;
	int i_max;

	if (dont_fit_basic(intent, grid))
		return (FALSE);
	i = intent->x;
	i_max = i + intent->size;
	j = intent->y + intent->size - 1;
	while (i < i_max)
	{
		if (grid->map[j][i] == grid->translate[OBSTACLE])
			return (FALSE);
		i++;
	}
	i = intent->y;
	i_max = i + intent->size - 1;
	j = intent->x + intent->size - 1;
	while (i < i_max)
	{
		if (grid->map[i][j] == grid->translate[OBSTACLE])
			return (FALSE);
		i++;
	}
	return (TRUE);
}

t_bool	dont_fit_basic(t_solution *intent, t_grid *grid)
{
	return (intent->x + intent->size > grid->max_j
			|| intent->y + intent->size > grid->max_i);
}

void		process_args(t_grid *grid, t_solution *sol, int argc, char **argv)
{
	int			i;

	i = 1;
	while (i < argc)
	{
		if (ft_load_grid(argv[i], grid)) // если файлик загружен и все ОК
			ft_process_grid(grid, sol);
		else
			ft_print_error();
		i++;
		if (i != argc)
			write(1, "\n", 1);
		ft_free_grid(grid);
	}
}

int			main(int argc, char **argv)
{
	t_grid		grid;
	t_solution	solution;

	if (argc < 2)
	{
		if (ft_parse_grid(STDIN_FILE_DESCRIPTOR, &grid))
			ft_process_grid(&grid, &solution);
		else
			ft_print_error();
		ft_free_grid(&grid);
	}
	else
		process_args(&grid, &solution, argc, argv);
}
