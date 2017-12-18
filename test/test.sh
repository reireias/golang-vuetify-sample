#!/bin/bash

expect -c "
set timeout 10
spawn vue init reireias/golang-vuetify-sample test
expect \"name\"
send \"\n\"
expect \"author\"
send \"\n\"
interact
"
