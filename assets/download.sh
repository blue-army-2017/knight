#!/bin/bash

PICO_VERSION="2.0.6"

mkdir -p assets/css
# pico.css
wget "https://raw.githubusercontent.com/picocss/pico/v${PICO_VERSION}/css/pico.blue.min.css" -O assets/css/pico.css

mkdir -p assets/font
# material-symbols.woff2
wget "https://github.com/google/material-design-icons/raw/master/variablefont/MaterialSymbolsOutlined%5BFILL,GRAD,opsz,wght%5D.woff2" -O assets/font/material-symbols.woff2
