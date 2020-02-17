#!/bin/bash

ROOTDIR=$PWD

#Run authentication tests
cd $ROOTDIR/authentications/tests
go test -v

#Run users tests
cd $ROOTDIR/users/tests
go test -v