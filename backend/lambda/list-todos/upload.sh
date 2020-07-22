#! /usr/bin/env bash

PROFILE="$GOFULLSTACKPROFILE"
ROLE="$GOFULLSTACKROLE"

build() {
  go build
  zip list-todos.zip list-todos
}

upload_function() {
  # aws lambda delete-function --function-name list-todos --profile $PROFILE
  aws lambda create-function --function-name list-todos --zip-file fileb://list-todos.zip --handler list-todos --runtime go1.x --role $ROLE --profile $PROFILE
}

update_function() {
  aws lambda update-function-code --function-name list-todos --profile $PROFILE --zip-file fileb://list-todos.zip
}

clean_up() {
	rm list-todos list-todos.zip
}

while getopts ":u" opt; do
  case ${opt} in
    u ) # process option u
			UPDATE=true
      ;;
    \? ) echo "Usage: upload.sh [-u]"
      ;;
  esac
done

build

if [[ $UPDATE = "true" ]]
then
	printf "Updating..."
	update_function
else
	printf "Uploading..."
	upload_function
fi

clean_up
