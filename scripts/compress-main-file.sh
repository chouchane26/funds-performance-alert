#!/bin/bash

mkdir ../deployment && cd ../deployment

cp ../main . && cp -r ../src/resources .

zip -r main.zip .