from datetime import datetime


def get_timestamp(delta=None):
    date = datetime.today()
    if delta is not None:
        date += delta

    return round(date.timestamp())
