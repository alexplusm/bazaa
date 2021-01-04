### INFO

exclude /pft

#### Makefile snippet for testing
test: $(NAME)
    $(CC) $(NAME) main.c -o main.out
    ./main.out

"%[flags][width][.precision][length]specifier"
