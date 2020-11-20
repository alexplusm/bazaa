# INFO: in (venv) can't use "source .local.env"
export $(grep -v '^#' .local.env | xargs)