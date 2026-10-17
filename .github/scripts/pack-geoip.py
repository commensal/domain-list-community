#!/usr/bin/env python3
"""Pack CIDR text files into a v2ray geoip.dat, preserving filename case."""

import ipaddress
import pathlib
import sys


def varint(value):
    out = bytearray()
    while value > 0x7F:
        out.append((value & 0x7F) | 0x80)
        value >>= 7
    out.append(value)
    return bytes(out)


def key(field, wire):
    return varint((field << 3) | wire)


def length_delimited(field, data):
    return key(field, 2) + varint(len(data)) + data


def cidr_message(network):
    body = length_delimited(1, network.network_address.packed)
    if network.prefixlen:
        body += key(2, 0) + varint(network.prefixlen)
    return body


def main():
    source, output = sys.argv[1:]
    entries = bytearray()
    for path in sorted(pathlib.Path(source).glob("*.txt")):
        cidrs = []
        for line in path.read_text(encoding="utf-8").splitlines():
            line = line.strip()
            if line and not line.startswith("#"):
                cidrs.append(ipaddress.ip_network(line, strict=False))
        if not cidrs:
            raise SystemExit(f"{path.name} has no CIDR")
        message = length_delimited(1, path.stem.encode()) + b"".join(
            length_delimited(2, cidr_message(cidr)) for cidr in cidrs
        )
        entries += length_delimited(1, message)
        print(f"{path.stem}: {len(cidrs)} CIDR")
    pathlib.Path(output).write_bytes(entries)


if __name__ == "__main__":
    main()
