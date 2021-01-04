# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: cdeon <cdeon@student.21-school.ru>         +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2020/12/27 09:16:13 by cdeon             #+#    #+#              #
#    Updated: 2020/12/27 09:16:15 by cdeon            ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

NAME 		= libftprintf.a

LIB_DIR		= libft
LIB_NAME 	= libft.a

SRCS 		= ft_printf.c
OBJS 		= ${SRCS:.c=.o}

CFLAGS		=	-Wall -Wextra -Werror

CC			= gcc
RM			= rm -f

INCLUDE		=	-I ft_printf.h -I $(LIB_DIR)

all: $(NAME)

.c.o:
	${CC} ${CFLAGS} $(INCLUDE) -g -c $< -o ${<:.c=.o}

$(NAME): $(OBJS)
	make -C $(LIB_DIR)
	cp $(LIB_DIR)/$(LIB_NAME) $(NAME)
	ar rcs ${NAME} ${OBJS}

clean:
	make clean -C $(LIB_DIR)
	rm -f $(OBJS)

fclean: clean
	make fclean -C $(LIB_DIR)
	rm -f $(NAME)

re: fclean all

# TODO: for test
test: $(NAME)
	$(CC) $(NAME) main.c -o main.out
	./main.out
