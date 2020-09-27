#include <stdlib.h>
#include "utils.c"
#include <stdio.h>

int		ft_strlen(char *str)
{
	int i;

	i = 0;
	while (str[i] != '\0')
	{
		i += 1;
	}
	return (i);
}

typedef struct s_table_config {
    int max_i;
    int max_j;
} t_table_config;

t_table_config *create_table_config(int max_i, int max_j)
{
    t_table_config *config;

    config = malloc(sizeof(t_table_config));
    if (config == NULL)
        return (NULL);
    config->max_i = max_i;
    config->max_j = max_j;
    return (config);
}

typedef struct s_point {
    int i;
    int j;
} t_point;

t_point *create_point(int i, int j)
{
    t_point *point;

    point = malloc(sizeof(t_point));
    if (point == NULL)
        return (NULL);
    point->i = i;
    point->j = j;
    return (point);
}

void print_point(t_point *point)
{
    printf("point (%d, %d)\n", point->i, point->j);
}

typedef struct s_rectangle {
    t_point *top_left;
    t_point *bottom_right;
    int square;
} t_rectangle;

t_rectangle *create_rectangle(int top_i, int top_j, int bottom_i, int bottom_j)
{
    t_rectangle *rectangle;

    rectangle = malloc(sizeof(t_rectangle));
    if (rectangle == NULL)
        return (NULL);
    
    rectangle->top_left = create_point(top_i, top_j);
    rectangle->bottom_right = create_point(bottom_i, bottom_j);
    rectangle->square = (bottom_i - top_i) * (bottom_j - bottom_i); // todo: overflow checks
    return (rectangle);
}

void ft_print_table(char **table, int size)
{
    int i;

    i = 0;
    while (i < size)
    {
        ft_putstr(table[i]);
        ft_putstr("\n");
        i += 1;
    }
}

void proccess_table(char **table, t_point *top_left, t_point *bottom_right)
{
    int i;
    int j;

    // printf("top left | ");
    // print_point(top_left);
    // printf("bottom right | ");
    // print_point(bottom_right);

    i = top_left->i;
    j = top_left->j;
    while (i <= bottom_right->i)
    {
        j = top_left->j; // ?
        while(j <= bottom_right->j)
        {
            // printf("char = %c (eq to 'o' ? = %d)\n", table[i][j], table[i][j] == 'o');
            if (table[i][j] == 'o') // todo: any symbol
            {
                // printf("\n --- HELL BLOCK (%d, %d) --- \n", i, j);
                proccess_table(table, top_left, create_point(i - 1, bottom_right->j)); // top
                proccess_table(table, top_left, create_point(bottom_right->i, j - 1)); // left
                proccess_table(table, create_point(top_left->i, j + 1), bottom_right); // right
                proccess_table(table, create_point(i + 1, top_left->j) , bottom_right); // bottom
                // printf("\n ### HELL BLOCK ### \n");
                return ; //kek?
            }

            j += 1;
        }

        i += 1;
    }
    i -= 1;
    j -= 1;

    int size = (bottom_right->i - top_left->i + 1) * (bottom_right->j - top_left->j + 1);

    if (size > 10 &&  (bottom_right->i - top_left->i) == (bottom_right->j - top_left->j) )
    {

        printf("top left | ");
        print_point(top_left);
        printf("bottom right | ");
        print_point(bottom_right);
        printf("\nSIZE: %d | (i = %d) (j = %d)\n\n\n\n\n", size, (i), (j)); // lol?
    }
}


#include <fcntl.h>

void parse_file(char *filename)
{
    int fd;
    char buff[10];

    fd = open(filename, O_RDONLY);
    if (fd < 0)
    {
        // handle error
    }

}

int main()
{
    char *table[9];

    table[0] = "...........................";
    table[1] = "....o......................";
    table[2] = "............o..............";
    table[3] = "...........................";
    table[4] = "....o......................";
    table[5] = "...............o...........";
    table[6] = "...........................";
    table[7] = "......o..............o.....";
    table[8] = "..o.......o................";



    // table[0] = "..........";
    // table[1] = "..........";
    // table[2] = ".....o....";  // (2, 12)
    // table[3] = "..........";
    // table[4] = "..........";

    /*   
        top:
            (0,0)
    
     */

    int max_i;
    int max_j;
    t_table_config *table_config;

    max_i = 8;
    max_j = ft_strlen(table[0]) - 1;
    printf("max length: %d\n", max_j);
    table_config = create_table_config(max_i, max_j);

    t_point *top_left = create_point(0, 0);
    t_point *bottom_right = create_point(max_i, max_j);

    proccess_table(table, top_left, bottom_right);
}
