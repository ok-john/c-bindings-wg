//+go:build cgo

package main

/*

#include "wireguard.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#cgo CFLAGS: -Wall

void list_devices(void)
{
	char *device_names, *device_name;
	size_t len;

	device_names = wg_list_device_names();
	if (!device_names) {
		perror("Unable to get device names");
		exit(1);
	}
	wg_for_each_device_name(device_names, device_name, len) {
		wg_device *device;
		wg_peer *peer;
		wg_key_b64_string key;

		if (wg_get_device(&device, device_name) < 0) {
			perror("Unable to get device");
			continue;
		}
		if (device->flags & WGDEVICE_HAS_PUBLIC_KEY) {
			wg_key_to_base64(key, device->public_key);
			printf("%s has public key %s\n", device_name, key);
		} else
			printf("%s has no public key\n", device_name);
		wg_for_each_peer(device, peer) {
			wg_key_to_base64(key, peer->public_key);
			printf(" - peer %s\n", key);
		}
		wg_free_device(device);
	}
	free(device_names);
}
*/
import (
	"C"
)

func main() {
	C.list_devices()
}
