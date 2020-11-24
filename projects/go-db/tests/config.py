import os


def get_config():
    server_port = os.environ["SERVER_PORT_OUTER"]
    host = os.environ["SERVER_HOST"]
    # host = "http://127.0.0.1"

    return {"server_port": server_port, "host": host}


def get_url_start():
    c = get_config()
    return c["host"] + ":" + c["server_port"]
