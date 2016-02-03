#!/bin/sh

set -e

if [ -z "${MOUNT_DIR}" ]; then
  MOUNT_DIR="/var/lib/openstorage/btrfs"
fi
if [ -z "${MOUNT_IMAGE}" ]; then
  MOUNT_IMAGE="/var/lib/openstorage/btrfs.img"
fi
if [ -z "${IMAGE_SIZE}" ]; then
  IMAGE_SIZE="1G"
fi
if [ -z "${NO_RM}" ]; then
  echo "***WARNING***: this will remove ${MOUNT_DIR} and ${MOUNT_IMAGE}, sleeping for 5 seconds so you can ctrl+c..." >&2
  sleep 5
fi

apt-get install -yq btrfs-tools

if [ -z "${NO_RM}" ]; then
  umount ${MOUNT_DIR} || true
  rm -f ${MOUNT_IMAGE}
  rm -rf ${MOUNT_DIR}
  truncate ${MOUNT_IMAGE} -s ${IMAGE_SIZE}
  mkfs.btrfs ${MOUNT_IMAGE}
  mkdir -p ${MOUNT_DIR}
fi

mount ${MOUNT_IMAGE} ${MOUNT_DIR} || true
