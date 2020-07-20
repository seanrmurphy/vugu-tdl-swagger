#! /usr/bin/env bash

update()  {
  echo "Updating create-todo function"
  cd create-todo
  ./upload.sh -u

  echo
  echo
  echo "Updating delete-todo function"
  cd ../delete-todo
  ./upload.sh -u

  echo
  echo
  echo "Updating get-todo function"
  cd ../get-todo
  ./upload.sh -u

  echo
  echo
  echo "Updating list-todos function"
  cd ../list-todos
  ./upload.sh -u

  echo
  echo
  echo "Updating update-todo function"
  cd ../update-todo
  ./upload.sh -u
}


upload()  {
  echo "Uploading create-todo function"
  cd create-todo
  ./upload.sh

  echo
  echo
  echo "Uploading delete-todo function"
  cd ../delete-todo
  ./upload.sh

  echo
  echo
  echo "Uploading get-todo function"
  cd ../get-todo
  ./upload.sh

  echo
  echo
  echo "Uploading list-todos function"
  cd ../list-todos
  ./upload.sh

  echo
  echo
  echo "Uploading update-todo function"
  cd ../update-todo
  ./upload.sh
}

check_env_vars() {
	if [ -z "$GOFULLSTACKPROFILE" ]
	then
		echo "Environment variable GOFULLSTACKPROFILE not defined...exiting..."
		exit
	fi

	if [ -z "$GOFULLSTACKROLE" ]
	then
		echo "Environment variable GOFULLSTACKROLE not defined...exiting..."
		exit
	fi
}

while getopts ":u" opt; do
  case ${opt} in
    u ) # process option u
			UPDATE=true
      ;;
    \? ) echo "Usage: upload-all-functions [-u]"
      ;;
  esac
done

check_env_vars

if [[ $UPDATE = "true" ]]
then
	printf "Updating..."
	update
else
	printf "Uploading..."
	upload
fi



