#!/bin/bash

while read -r __program; do
    if ! which "${__program}" &> /dev/null; then
        echo "Please install ${__program}"
        exit 1
    fi
done <<< "cobra"

################################################################################

__package_name="solus-package-util"
__package_path="github.com/W-Floyd/solus-package-tools"

__package_full="${__package_path}/${__package_name}"

if [ -d "${__package_name}" ] && [ -d ".${__package_name}" ]; then

    if ! [ -d '.patches' ]; then
        rm -r '.patches'
    fi

    mkdir '.patches'

    pushd "${__package_name}"

    find . -type f | while read -r __file; do
        __patch="$(diff -Nau "../.${__package_name}/${__file}" "${__file}" | sed -e '1,2s/ .*//')"
        if ! [ -z "${__patch}" ]; then
            mkdir -p "$(dirname "../.patches/${__file}")"
            echo "${__patch}" > "../.patches/${__file}"
        fi
    done

    popd

    rm -r "${__package_name}"
    rm -r ".${__package_name}"
fi

################################################################################

cobra init "${__package_full}"

while read -r __command; do
    cobra add -t "${__package_full}" "${__command}"
done <<< 'build
git
list
graph
info
bootstrap'

cobra add -t "${__package_full}" -p gitCmd bump
cobra add -t "${__package_full}" -p gitCmd upgrade
cobra add -t "${__package_full}" -p gitCmd rebuild

################################################################################

cp -r "${__package_name}" ".${__package_name}"

if [ -d '.patches' ]; then

    pushd '.patches'

    find . -type f | while read -r __file; do
        patch "../${__package_name}/${__file}" "${__file}"
    done

    popd

fi

pushd "${__package_name}"

go build -o ~/go/bin/${__package_name}

popd

exit
