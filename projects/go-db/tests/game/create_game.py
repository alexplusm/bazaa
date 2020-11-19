import sys
sys.path.append('../..')

from tests.config import conn_config

game = 666

if __name__ == '__main__':
    print(conn_config['host'], conn_config['port'])
