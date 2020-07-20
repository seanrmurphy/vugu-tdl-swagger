#! /usr/bin/env bash

PROFILE="$GOFULLSTACKPROFILE"
ROLE="$GOFULLSTACKROLE"

build() {
  go build
  zip update-todo.zip update-todo
}

upload_function() {
  # aws lambda delete-function --function-name update-todo --profile $PROFILE
  aws lambda create-function --function-name update-todo --zip-file fileb://update-todo.zip --handler update-todo --runtime go1.x --role $ROLE --profile $PROFILE
}

update_function() {
  aws lambda update-function-code --function-name update-todo --profile $PROFILE --zip-file fileb://update-todo.zip
}

clean_up() {
	rm update-todo update-todo.zip
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

