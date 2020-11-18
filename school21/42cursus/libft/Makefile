SRC = 	ft_memset.c ft_atoi.c

OBJ = 	${SRC:.c=.o}

CC = gcc

CFLAGS = -Wall -Werror -Wextra

RM = rm -rf

NAME = libft.a

# todo: flags (-Wall -Wextra ...)

$(NAME): all

all:
	$(CC) ${CFLAGS} $(SRC) -c
	ar rcs $(NAME) ${OBJ}
	$(MAKE) clean

clean:
	$(RM) $(OBJ)

fclean: clean
	$(RM) $(NAME)

re:	fclean all