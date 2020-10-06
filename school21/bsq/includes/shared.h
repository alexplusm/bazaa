#ifndef SHARED_H
# define SHARED_H

# define BUFFER_SIZE 1024
# define STDIN_FILE_DESCRIPTOR 0

typedef enum {
    FALSE,
    TRUE
}   t_bool;

typedef	enum {
	EMPTY,
	OBSTACLE,
	FILL,
}	t_block;

#endif
