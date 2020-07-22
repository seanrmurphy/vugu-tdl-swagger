#! /usr/bin/env bash

PROFILE="$GOFULLSTACKPROFILE"
ROLE="$GOFULLSTACKROLE"

build() {
  go build
  zip create-todo.zip create-todo
}

upload_function() {
  # aws lambda delete-function --function-name create-todo --profile $PROFILE
  aws lambda create-function --function-name create-todo --zip-file fileb://create-todo.zip --handler create-todo --runtime go1.x --role $ROLE --profile $PROFILE
}

update_function() {
  aws lambda update-function-code --function-name create-todo --profile $PROFILE --zip-file fileb://create-todo.zip
}

clean_up() {
	rm create-todo create-todo.zip
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
