#!/usr/env bash

set -e

function update_menu()
{
  options=("Update" "Exit")
  select OPTIONS in "${options[@]}" do
    case $OPTIONS in
      Update);;
      Exit);;
    esac
  done
}