ARCHIVE_LOW1="archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/low1.zip"
ARCHIVE_MEDIUM="archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/medium.zip"

GAME_ID1="be359a94-d2cf-4ec4-8dc2-fbd8350c89ec" # past
GAME_ID2="2bc4c2aa-aa87-406c-b922-334a95fed451" # future

URL="http://localhost:1234/api/v1/game/${GAME_ID2}"

#-F "$ARCHIVE_LOW1" \

curl \
  -F "$ARCHIVE_LOW1" \
  -F "$ARCHIVE_MEDIUM" \
  -X PUT "$URL"
