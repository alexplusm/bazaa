#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdlib.h>

# define BUFFER_SIZE 1024
# define STDIN_FILE_DESCRIPTOR 0

typedef struct	s_grid {
	unsigned int	w;
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

typedef	enum {
	false,
	true,
}	t_bool;

typedef	enum {
	empty,
	obstacle,
	fill,
}	t_block;

t_bool			ft_load_grid(char *path, t_grid *grid);
void			ft_free_grid(t_grid *grid);
void			ft_process_grid(t_grid *grid, t_solution *sol);
t_bool			ft_parse_header_line(t_grid *grid, char *line, unsigned int length);
t_bool			ft_parse_normal_line(t_grid *grid, char *line, unsigned int length);
t_bool			ft_read_full(int fd, char **file_content, unsigned int *total);
t_bool			ft_has_width_changed(t_grid *grid, unsigned int *curr_w);
t_bool			ft_process_lines(t_grid *grid, unsigned int index, char *file_content, unsigned int total);
t_bool			ft_parse_grid(int fd, t_grid *grid);
t_bool			find_solution(t_solution *sol, t_grid *grid);
void			ft_print_solution(t_grid *grid, t_solution *solution);
void		actualize_sol(t_solution *sol, t_solution *intent, t_grid *grid);
t_bool	fits(t_solution *sol, t_grid *grid);
t_bool	fits_succ(t_solution *intent, t_grid *grid);
t_bool	dont_fit_basic(t_solution *intent, t_grid *grid);
char	*ft_extend_array(char *orig, char *n_cont, unsigned int old_len, unsigned int len);


t_bool	ft_atoi_n_strict(char *str, unsigned int n, unsigned int *result);

void ft_print_error()
{
	write(2, "map error\n", 10);
}

t_bool	ft_read_full(int fd, char **file_content, unsigned int *total)
{
	char				buffer[BUFFER_SIZE];
	unsigned int		byte_read;

	if (read(fd, 0, 0) == -1)
		return (false);
	*total = 0;
	*file_content = NULL;
	while ((byte_read = read(fd, buffer, BUFFER_SIZE)) > 0)
	{
		if (!(*file_content = ft_extend_array(*file_content,
				buffer, *total, *total + byte_read)))
			return (false);
		(*total) += byte_read;
	}
	if (byte_read == 0)
		return (true);
	return (false);
}

t_bool	ft_parse_header_line(t_grid *grid, char *line, unsigned int length)
{
	t_bool	result;

	if (length < 4)
		return (false);


	grid->translate[fill] = line[length - 1];
	grid->translate[obstacle] = line[length - 2];
	grid->translate[empty] = line[length - 3];
	result = ft_atoi_n_strict(line, length - 3, &(grid->max_i));
	return (result);
}


t_bool	ft_parse_normal_line(t_grid *grid, char *line, unsigned int length)
{
	unsigned int	index;
	char	current;

	if (length < 1)
		return (false);
	grid->w = length;
	index = 0;
	while (index < length)
	{
		current = line[index];
		if (current != grid->translate[empty]
				&& current != grid->translate[obstacle]
				&& current != grid->translate[fill])
			return (false);
		index++;
	}
	return (true);
}

t_bool	ft_has_width_changed(t_grid *grid, unsigned int *curr_w)
{
	if (*curr_w == (unsigned int)-1)
		*curr_w = grid->w;
	return (*curr_w != grid->w);
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
				return (false);
		if (ft_parse_normal_line(grid, file_content + index, jndex - index))
		{
			grid->map[y++] = (unsigned char *)(file_content + index);
			if (ft_has_width_changed(grid, &curr_w))
				return (false);
		}
		else
			return (false);
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
		return (false);
	grid->source = file_content;
	index = 0;
	while (index < total)
	{
		if (file_content[index] == '\n')
		{
			if (!ft_parse_header_line(grid, file_content, index))
				return (false);
			if (!(grid->map = malloc(sizeof(char *) * grid->max_i)))
				return (false);
			return (ft_process_lines(grid, index + 1, file_content, total));
		}
		index++;
	}
	return (false);
}
 
t_bool	ft_load_grid(char *filename, t_grid *grid)
{
	int		fd;
	t_bool	result;

	fd = open(filename, O_RDONLY);
	if (fd < 0)
		return (false);
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
				grid->map[y][x] = grid->translate[fill];
				y++;
			}
			x++;
		}
		y = 0;
		while (y < grid->max_i)
		{
			write(1, grid->map[y], grid->w);
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
		while (j < grid->w - sol->size)
		{
			intent.x = j;
			intent.y = i;
			actualize_sol(sol, &intent, grid);
			j++;
		}
		i++;
	}
	return (sol->size ? true : false);
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
// static t_bool	fits(const t_solution *intent, const t_grid *grid)
{
	int i;
	int i_max;
	int j;
	int j_max;

	if (dont_fit_basic(intent, grid))
		return (false);
	i = intent->y;
	i_max = i + intent->size;
	j_max = intent->x + intent->size;
	while (i < i_max)
	{
		j = intent->x;
		while (j < j_max)
		{
			if (grid->map[i][j] == grid->translate[obstacle])
				return (false);
			j++;
		}
		i++;
	}
	return (true);
}

// static t_bool	fits_succ(const t_solution *intent, const t_grid *grid)
t_bool	fits_succ(t_solution *intent, t_grid *grid)
{
	int i;
	int j;
	int i_max;

	if (dont_fit_basic(intent, grid))
		return (false);
	i = intent->x;
	i_max = i + intent->size;
	j = intent->y + intent->size - 1;
	while (i < i_max)
	{
		if (grid->map[j][i] == grid->translate[obstacle])
			return (false);
		i++;
	}
	i = intent->y;
	i_max = i + intent->size - 1;
	j = intent->x + intent->size - 1;
	while (i < i_max)
	{
		if (grid->map[i][j] == grid->translate[obstacle])
			return (false);
		i++;
	}
	return (true);
}

t_bool	dont_fit_basic(t_solution *intent, t_grid *grid)
// static t_bool	dont_fit_basic(const t_solution *intent, const t_grid *grid)
{
	return (intent->x + intent->size > grid->w
			|| intent->y + intent->size > grid->max_i);
}

// str utils
char	*ft_str_n_copy(char *dest, char *src, int n)
{
	int		index;

	index = 0;
	while (index < n && src[index] != '\0')
	{
		dest[index] = src[index];
		index++;
	}
	while (index < n)
	{
		dest[index] = '\0';
		index++;
	}
	return (dest);
}

char	*ft_extend_array(char *orig, char *n_cont, unsigned int old_len, unsigned int len)
{
	char *dest;

	if (!(dest = malloc((len + 1) * sizeof(char))))
		return (NULL);
	if (orig != NULL)
		ft_str_n_copy(dest, orig, old_len);
	ft_str_n_copy(dest + old_len, n_cont, (unsigned int)(len - old_len));
	if (orig != NULL)
		free(orig);
	return (dest);
}

// in utils
t_bool	ft_atoi_n_strict(char *str, unsigned int n, unsigned int *result)
{
	unsigned int	index;

	index = 0;
	*result = 0;
	while (index < n)
	{
		if ('0' <= str[index] && str[index] <= '9')
			*result = *result * 10 + str[index] - '0';
		else
			return (false);
		index++;
	}
	return (true);
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
