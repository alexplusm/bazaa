#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>

void print_file(char *filename);
void fill_array_null(char array[5]);

int main(int argc, char *argv[])
{
    char *str;
    str = argv[1];
        
    print_file(str);
    return (0);
}

void print_file(char *filename)
{
    int fd;
    char buff[5];
    int read_bytes;

    fd = open(filename, 0);

    while (read_bytes != 0)
    {
        read_bytes = read(fd, &buff, 5);
        printf("%s", buff);
        fill_array_null(buff);
    }
}

void fill_array_null(char array[5])
{
    int i;

    i = 0;
    while (i < 5)
    {
        array[i] = 0;
        i += 1;
    }
}