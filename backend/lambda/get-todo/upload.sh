#! /usr/bin/env bash

PROFILE="$GOFULLSTACKPROFILE"
ROLE="$GOFULLSTACKROLE"

build() {
  go build
  zip get-todo.zip get-todo
}

upload_function() {
  # aws lambda delete-function --function-name get-todo --profile $PROFILE
  aws lambda create-function --function-name get-todo --zip-file fileb://get-todo.zip --handler get-todo --runtime go1.x --role $ROLE --profile $PROFILE
}

update_function() {
  aws lambda update-function-code --function-name get-todo --profile $PROFILE --zip-file fileb://get-todo.zip
}

clean_up() {
	rm get-todo get-todo.zip
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
