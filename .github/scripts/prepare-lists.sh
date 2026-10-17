#!/usr/bin/env bash
# Turn itdoginfo/allow-domains plus the personal lists into compiler input.
# Category file names stay lowercase: Xray looks up geosite:/geoip: codes
# byte-for-byte after uppercasing the rule, so the dat must store that case.
set -euo pipefail

root="${1:?repository root}"
itdog="${2:?allow-domains checkout}"

data="$root/data"
geoip_source="$root/geoip-source"
mkdir -p "$data" "$geoip_source"
find "$data" -mindepth 1 -delete
find "$geoip_source" -mindepth 1 -delete

append_rules() {
  local src="$1" dest="$2"
  [ -f "$src" ] || return 0
  # A bare "domain:" prefix is redundant. A leading dot is a public-suffix zone,
  # which this compiler rejects as a domain, so it becomes an anchored regexp.
  sed -e 's/\r$//' -e 's/^domain://' -e 's/^\.\(.*\)$/regexp:\\.\1$/' "$src" \
    | grep -v '^#' \
    | grep -v '^package:' \
    | grep -v '^[[:space:]]*$' >> "$dest" || true
}

category_name() {
  local name
  name=$(basename "$1" .lst | tr '[:upper:]' '[:lower:]' | tr '_' '-')
  if [ "$name" = "block" ]; then
    name="ru-block"
  fi
  printf '%s\n' "$name"
}

for f in "$itdog"/Categories/*.lst "$itdog"/Services/*.lst; do
  [ -e "$f" ] || continue
  append_rules "$f" "$data/$(category_name "$f")"
done

# Personal domains that are not covered by an allow-domains category.
if [ -f "$root/my-domains.txt" ]; then
  append_rules "$root/my-domains.txt" "$data/custom"
fi

for f in "$data"/*; do
  [ -f "$f" ] || continue
  sort -u "$f" | sed '/^$/d' > "$f.tmp"
  mv "$f.tmp" "$f"
  if [ ! -s "$f" ]; then
    rm -f "$f"
  fi
done

if [ ! -s "$data/custom" ]; then
  echo "example.com" > "$data/custom"
fi

for family in IPv4 IPv6; do
  for f in "$itdog/Subnets/$family"/*.lst; do
    [ -e "$f" ] || continue
    append_rules "$f" "$geoip_source/$(category_name "$f").txt"
  done
done

if [ -f "$root/my-ips.txt" ]; then
  append_rules "$root/my-ips.txt" "$geoip_source/custom.txt"
fi

for f in "$geoip_source"/*.txt; do
  [ -e "$f" ] || continue
  sort -u "$f" | sed '/^$/d' > "$f.tmp"
  # An empty file is not a category. A placeholder CIDR would publish a group
  # that no profile references, and an empty my-ips.txt must not do that either.
  if [ ! -s "$f.tmp" ]; then
    rm -f "$f" "$f.tmp"
  else
    mv "$f.tmp" "$f"
  fi
done
