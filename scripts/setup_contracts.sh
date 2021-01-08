#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

CONTRACTS_DIR="./eth-solidity"
DEST_DIR="./bindings/ethereum"

set -eux

case $TARGET in
	"build")
    pushd $CONTRACTS_DIR

    make install-deps
    make bindings

    popd

    mkdir -p $DEST_DIR
    cp -r $CONTRACTS_DIR/build/bindings/go/* $DEST_DIR
		;;


esac
