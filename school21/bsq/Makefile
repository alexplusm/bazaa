SRC		= src/utils.c src/main.c
NAME	= bsq
CC		= gcc
RM		= rm -f
CFLAGS	= -Wall -Wextra -Werror
INCL	= includes/

${NAME}:
	${CC} ${SRC} -I${INCL} -o ${NAME}

all:	${NAME}
clean:

fclean:		clean
	$(RM) ${NAME}
re:			fclean all
