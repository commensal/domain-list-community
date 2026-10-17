#!/usr/bin/env python3
"""Fail the build when a routing profile names a category the dat files do not contain."""

import base64
import json
import pathlib
import sys


def varint(data, index):
    value = shift = 0
    while True:
        byte = data[index]
        index += 1
        value |= (byte & 0x7F) << shift
        if byte < 128:
            return value, index
        shift += 7


def messages(data, field_number):
    index = 0
    found = []
    while index < len(data):
        key, index = varint(data, index)
        field, wire = key >> 3, key & 7
        if wire == 0:
            _, index = varint(data, index)
        elif wire == 2:
            length, index = varint(data, index)
            if field == field_number:
                found.append(data[index:index + length])
            index += length
        else:
            raise ValueError(f"unsupported wire type {wire}")
    return found


def country_codes(blob):
    # GeoSiteList.entry and GeoIPList.entry are both field 1, and both messages
    # store the category code in field 1.
    return {messages(entry, 1)[0].decode("utf-8") for entry in messages(blob, 1)}


def profile_codes(path, key):
    payload = path.read_text(encoding="utf-8").split("://routing/add/", 1)[1].strip()
    profile = json.loads(base64.b64decode(payload))
    prefix = "geosite:" if "Sites" in key else "geoip:"
    return {
        item[len(prefix):]
        for item in profile.get(key, [])
        if item.startswith(prefix)
    }


def main():
    geosite_path, geoip_path, root = sys.argv[1:]
    root = pathlib.Path(root)
    geosite = country_codes(pathlib.Path(geosite_path).read_bytes())
    geoip = country_codes(pathlib.Path(geoip_path).read_bytes())
    # The profile writes geosite:custom, but Xray uppercases that code before
    # the byte comparison. A lowercase code in the dat is therefore unreachable.
    mixed = sorted(code for code in geosite | geoip if code != code.upper())
    if mixed:
        raise SystemExit("category codes must be uppercase: " + ", ".join(mixed))

    missing = []
    for folder in ("HAPP", "INCY"):
        for name in ("DEFAULT.DEEPLINK", "FULL.DEEPLINK"):
            path = root / folder / name
            if not path.is_file():
                missing.append(f"missing profile {folder}/{name}")
                continue
            for key, target, label in [
                ("ProxySites", geosite, "geosite"),
                ("BlockSites", geosite, "geosite"),
                ("DirectSites", geosite, "geosite"),
                ("ProxyIp", geoip, "geoip"),
                ("BlockIp", geoip, "geoip"),
                ("DirectIp", geoip, "geoip"),
            ]:
                for code in sorted(profile_codes(path, key)):
                    if code.upper() not in target:
                        missing.append(f"{folder}/{name} {label}:{code} (in {key})")
    if missing:
        raise SystemExit("profiles reference missing categories:\n" + "\n".join(missing))
    print(f"verified {len(geosite)} geosite and {len(geoip)} geoip categories")


if __name__ == "__main__":
    main()
