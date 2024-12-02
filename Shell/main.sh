#!/usr/env bash

set -e

USER_NAME=$(whoami)
echo "Whelcome to ${USER_NAME} !"

function main()
{
  options=("Test" "Setting" "Exit")
  select OPTIONS in "${options[@]}" do
    case $OPTIONS in
      Test);;
      Setting);;
      Exit);;
    esac
  done
}
main