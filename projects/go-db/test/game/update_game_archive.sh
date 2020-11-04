# TODO: use this vars in curl exec
FPATH1="/Users/a.mogilevskiy/work/5gen/clean-city/archives/low.zip"
FPATH2="/Users/a.mogilevskiy/work/5gen/clean-city/archives/medium.zip"

URL="http://localhost:1234/api/v1/game/some-game-id"

# -F 'archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/medium.zip' \

curl \
    -F 'archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/low1.zip' \
    -X PUT "$URL"
