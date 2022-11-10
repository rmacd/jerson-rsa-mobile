#!/usr/bin/env sh
#######################################################
#                                                     #
# Usage:                                              #
# ./go.sh 1.14.15 mylibraryname                       #
#                                                     #
# By: @jerson                                         #
#######################################################

GO_VERSION="$1"
LIBRARY_NAME="$2"

if [[ -z $1 || -z $2 ]] ; then
    echo "version and library_name must be provided, eg $0 1.14.15 rsamobile"
    exit 1
fi

set -e
set -x

# check valid suffix
if [[ ! "${LIBRARY_NAME}" =~ ^[a-z_]+$ ]] ; then
    echo "invalid suffix, alpha and/or underscores only"
    exit 1
fi

patch_file() {
    if [ -z $1 ] ; then exit 1 ; fi
    sed -i "
        s/crosscall1/crosscall1_${SUFFIX}/g;
        s/_cgo_panic/_cgo_panic_${SUFFIX}/g;
        s/_cgo_topofstack/_cgo_topofstack_${SUFFIX}/g;
        s/crosscall2/crosscall2_${SUFFIX}/g;
        s/_cgo_release_context/_cgo_release_context_${SUFFIX}/g;
        s/_cgo_sys_thread_start/_cgo_sys_thread_start_${SUFFIX}/g;
        s/x_cgo_init/x_cgo_init_${SUFFIX}/g;
        s/_cgo_get_context_function/_cgo_get_context_function_${SUFFIX}/g;
        s/_cgo_try_pthread_create/_cgo_try_pthread_create_${SUFFIX}/g;
        s/_cgo_wait_runtime_init_done/_cgo_wait_runtime_init_done_${SUFFIX}/g;
        s/x_cgo_notify_runtime_init_done/x_cgo_notify_runtime_init_done_${SUFFIX}/g;
        s/x_cgo_set_context_function/x_cgo_set_context_function_${SUFFIX}/g;
        s/x_cgo_sys_thread_create/x_cgo_sys_thread_create_${SUFFIX}/g;
        s/x_cgo_setenv/x_cgo_setenv_${SUFFIX}/g;
        s/x_cgo_unsetenv/x_cgo_unsetenv_${SUFFIX}/g;
        s/x_cgo_callers/x_cgo_callers_${SUFFIX}/g;
        s/_cgo_yield/_cgo_yield_${SUFFIX}/g;
        s/x_cgo_thread_start/x_cgo_thread_start_${SUFFIX}/g;
        s/crosscall_amd64/crosscall_amd64_${SUFFIX}/g;
        s/crosscall_arm64/crosscall_arm64_${SUFFIX}/g;
        s/crosscall_arm/crosscall_arm_${SUFFIX}/g;
        s/crosscall_arm7/crosscall_arm7_${SUFFIX}/g;
        s/crosscall_arm1/crosscall_arm1_${SUFFIX}/g;
        s/xx_cgo_panicmem/xx_cgo_panicmem_${SUFFIX}/g;
        s/darwin_arm_init_thread_exception_port/darwin_arm_init_thread_exception_port_${SUFFIX}/g;
        s/darwin_arm_init_mach_exception_handler/darwin_arm_init_mach_exception_handler_${SUFFIX}/g" "$1"
    echo -n "."
}

patch_go_cgo() {
    FIND_DIRECTORY="$1"
    SUFFIX="$2"

    IFS=$'\n'
    files=($(LC_ALL=C find "${FIND_DIRECTORY}" -type f))
    unset IFS

    set +x
    echo -n "Patching files"
    for file in "${files[@]}" ; do
        patch_file "$file"
    done
    set -x
}

FILE=go${GO_VERSION}.tar.gz
BASE_DIRECTORY="$(mktemp -d -t go-$(date +%s)-XXXX)"
GO_DIRECTORY="${BASE_DIRECTORY}/go"
GO_SOURCE_FILE="${BASE_DIRECTORY}/${FILE}"

echo "(1/5) Downloading go source"
curl -L -o "${GO_SOURCE_FILE}" https://go.dev/dl/go${GO_VERSION}.src.tar.gz
echo "(2/5) Extraction..."
tar -C "${BASE_DIRECTORY}" -xzf "${GO_SOURCE_FILE}"
echo "(3/5) Patch source"
patch_go_cgo "${GO_DIRECTORY}" "${LIBRARY_NAME}"
echo "(4/5) Build"
cd "${GO_DIRECTORY}/src"
./make.bash
echo "(5/5) Clean"
rm -rf "${GO_SOURCE_FILE}"

