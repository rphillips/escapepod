#!/bin/bash

START=$(pwd)
cd testdata
curl -L -O "http://scripting.com/misc/userlandSamples.zip"
unzip userlandSamples.zip workspace.userlandSamples.fttb
cd "$START"
