#! /usr/bin/env bash

PROFILE="$GOFULLSTACKPROFILE"
ROLE="$GOFULLSTACKROLE"

build() {
  go build
  zip delete-todo.zip delete-todo
}

upload_function() {
  # aws lambda delete-function --function-name delete-todo --profile $PROFILE
  aws lambda create-function --function-name delete-todo --zip-file fileb://delete-todo.zip --handler delete-todo --runtime go1.x --role $ROLE --profile $PROFILE
}

update_function() {
  aws lambda update-function-code --function-name delete-todo --profile $PROFILE --zip-file fileb://delete-todo.zip
}

clean_up() {
	rm delete-todo delete-todo.zip
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
