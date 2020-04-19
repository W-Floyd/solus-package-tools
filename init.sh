#!/bin/bash

while read -r __program; do
    if ! which "${__program}" &>/dev/null; then
        echo "Please install ${__program}"
        exit 1
    fi
done <<<"cobra
patch
diff"

################################################################################

__package_name="solus-package-tools"
__package_path="github.com/W-Floyd/solus-package-tools"

__package_full="${__package_path}/${__package_name}"

__package_author="William Floyd <william.png2000@gmail.com>"
__package_license="MIT"

__cobra() {
    cobra -a "${__package_author}" -l "${__package_license}" ${@}
}

if [ -d "${__package_name}" ] && [ -d ".${__package_name}" ]; then

    if [ -d '.patches' ]; then
        rm -r '.patches'
    fi

    mkdir '.patches'

    pushd "${__package_name}" 1>/dev/null

    find . -type f | while read -r __file; do
        __patch="$(diff -Nau "../.${__package_name}/${__file}" "${__file}" | sed -e '1,2s/ .*//')"
        if ! [ -z "${__patch}" ]; then
            mkdir -p "$(dirname "../.patches/${__file}")"
            echo "${__patch}" >"../.patches/${__file}"
        fi
    done

    popd 1>/dev/null

    rm -r "${__package_name}"
    rm -r ".${__package_name}"
fi

if ! [ -d "${__package_name}" ]; then
    mkdir "${__package_name}"
fi

################################################################################

pushd "${__package_name}" 1>/dev/null

{

    __cobra init --pkg-name "${__package_full}"

    while read -r __command; do
        __cobra -l "${__package_license}"
    done <<<'build
git
list
graph
info
bootstrap'

} 1>/dev/null

popd 1>/dev/null

################################################################################

cp -r "${__package_name}" ".${__package_name}"

if [ -d '.patches' ]; then

    pushd '.patches' 1>/dev/null

    find . -type f | while read -r __file; do

        if ! [ -d "../${__package_name}/$(dirname "${__file}")" ]; then
            mkdir -p "../${__package_name}/$(dirname "${__file}")"
        fi

        touch "../${__package_name}/${__file}"

        patch --ignore-whitespace "../${__package_name}/${__file}" "${__file}" 1>/dev/null
    done

    popd 1>/dev/null

fi

################################################################################

pushd "${__package_name}" 1>/dev/null

go build -o ~/go/bin/${__package_name}

popd 1>/dev/null

exit
