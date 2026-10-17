#!/usr/bin/env bash
# Rewrite the two Happ/Incy deeplinks and index.html.
# LastUpdated moves forward only when geosite.dat or geoip.dat actually changed,
# which is the signal Happ uses to download the geo files again.
set -euo pipefail

root="${1:?repository root}"
old_geosite="${2:-}"
old_geoip="${3:-}"
# The workflow builds both files in its own directory, not inside the checkout.
new_geosite="${4:-$root/geosite.dat}"
new_geoip="${5:-$root/geoip.dat}"

python3 - "$root" "$old_geosite" "$old_geoip" "$new_geosite" "$new_geoip" << 'PY'
import base64, json, pathlib, sys, time

root, old_geosite, old_geoip, new_geosite, new_geoip = sys.argv[1:]
root = pathlib.Path(root)

def same(a, b):
    if not a or not pathlib.Path(a).is_file() or not pathlib.Path(b).is_file():
        return False
    return pathlib.Path(a).read_bytes() == pathlib.Path(b).read_bytes()

# A missing previous file is not evidence of a change: publishing would
# otherwise move LastUpdated on every run that failed to fetch release.
changed = (
    pathlib.Path(old_geosite).is_file()
    and pathlib.Path(old_geoip).is_file()
    and not (same(old_geosite, new_geosite) and same(old_geoip, new_geoip))
)

def load(path):
    text = path.read_text(encoding="utf-8").strip()
    scheme, payload = text.split("://routing/add/", 1)
    raw = base64.b64decode(payload)
    # Happ accepts standard base64; keep the decoded JSON strict.
    profile = json.loads(raw)
    return scheme, profile

def dump(scheme, profile):
    body = json.dumps(profile, ensure_ascii=False, indent=4) + "\n"
    return f"{scheme}://routing/add/{base64.b64encode(body.encode()).decode()}"

default_path = root / "HAPP" / "DEFAULT.DEEPLINK"
full_path = root / "HAPP" / "FULL.DEEPLINK"
_, default_profile = load(default_path)
_, full_profile = load(full_path)

previous = 0
for profile in (default_profile, full_profile):
    try:
        previous = max(previous, int(profile.get("LastUpdated") or 0))
    except (TypeError, ValueError):
        pass

# A newer unix time is what makes an already imported profile fetch geo files.
# An unchanged pair must keep the old value, or clients re-download for nothing.
stamp = str(max(int(time.time()), previous + 1)) if changed else str(previous)

def apply(profile, name):
    profile["Name"] = name
    profile["LastUpdated"] = stamp
    profile["GlobalProxy"] = "false"
    profile["RouteOrder"] = "block-proxy-direct"
    profile["RemoteDNSType"] = "DoU"
    profile["RemoteDNSDomain"] = "https://dns.dns-ai.ru/dns-query"
    profile["RemoteDNSIP"] = "8.8.8.8"
    profile["DomesticDNSType"] = "DoU"
    profile["DomesticDNSDomain"] = "https://dns.google/dns-query"
    profile["DomesticDNSIP"] = "77.88.8.8"
    profile["Geoipurl"] = "https://cdn.jsdelivr.net/gh/commensal/domain-list-community@release/geoip.dat"
    profile["Geositeurl"] = "https://cdn.jsdelivr.net/gh/commensal/domain-list-community@release/geosite.dat"
    profile["DnsHosts"] = {
        "cloudflare-dns.com": "1.1.1.1",
        "dns.google": "8.8.8.8",
    }
    profile["DirectSites"] = []
    profile["DirectIp"] = [
        "10.0.0.0/8",
        "172.16.0.0/12",
        "192.168.0.0/16",
        "169.254.0.0/16",
        "224.0.0.0/4",
        "255.255.255.255",
    ]
    profile["BlockSites"] = [
        "geosite:category-ads-all",
        "appcenter.ms",
        "firebase.io",
        "crashlytics.com",
    ]
    profile["BlockIp"] = []
    profile["DomainStrategy"] = "AsIs"
    profile["FakeDNS"] = "true"
    profile["UseChunkFiles"] = "true"

apply(default_profile, "Black List")
default_profile["ProxySites"] = [
    "geosite:telegram",
    "geosite:meta",
    "geosite:custom",
    "geosite:discord",
    "geosite:porn",
    "geosite:youtube",
    "geosite:google-play",
    "geosite:google-meet",
    "geosite:hdrezka",
    "geosite:news",
    "geosite:roblox",
    "geosite:tiktok",
    "geosite:twitter",
    "geosite:google-ai",
    "geosite:geoblock",
]
default_profile["ProxyIp"] = [
    "geoip:telegram",
    "geoip:twitter",
    "geoip:meta",
    "geoip:discord",
]

apply(full_profile, "Black List (Full)")
full_profile["ProxySites"] = [
    "geosite:custom",
    "geosite:ru-block",
    "geosite:geoblock",
    "geosite:telegram",
    "geosite:meta",
    "geosite:youtube",
    "geosite:discord",
    "geosite:twitter",
    "geosite:tiktok",
    "geosite:roblox",
    "geosite:google-play",
    "geosite:google-meet",
    "geosite:google-ai",
    "geosite:cloudflare",
    "geosite:cloudfront",
    "geosite:digitalocean",
    "geosite:hetzner",
    "geosite:ovh",
    "geosite:hdrezka",
    "geosite:hodca",
    "geosite:news",
    "geosite:porn",
    "geosite:anime",
]
full_profile["ProxyIp"] = [
    "geoip:telegram",
    "geoip:twitter",
    "geoip:meta",
    "geoip:discord",
    "geoip:roblox",
    "geoip:google-meet",
]

default_path.write_text(dump("happ", default_profile) + "\n", encoding="utf-8", newline="\n")
full_path.write_text(dump("happ", full_profile) + "\n", encoding="utf-8", newline="\n")
(root / "INCY" / "DEFAULT.DEEPLINK").write_text(dump("incy", default_profile) + "\n", encoding="utf-8", newline="\n")
(root / "INCY" / "FULL.DEEPLINK").write_text(dump("incy", full_profile) + "\n", encoding="utf-8", newline="\n")

page = root / "index.html"
html = page.read_text(encoding="utf-8")
marker = 'const fullB64 = "'
start = html.find(marker)
end = html.find('";', start)
if start < 0 or end < 0:
    raise SystemExit("index.html has no fullB64 payload")
payload = dump("happ", full_profile).split("://routing/add/", 1)[1]
html = html[:start] + marker + payload + html[end:]
page.write_text(html, encoding="utf-8", newline="\n")

print(f"geo_changed={'true' if changed else 'false'}")
print(f"last_updated={stamp}")
print(
    "compared "
    f"{pathlib.Path(new_geosite).stat().st_size if pathlib.Path(new_geosite).is_file() else 'missing'} "
    "geosite bytes with "
    f"{pathlib.Path(old_geosite).stat().st_size if pathlib.Path(old_geosite).is_file() else 'missing'}"
)
PY

# LIGHT was a byte-for-byte copy of DEFAULT and is not a third profile.
rm -f "$root/HAPP/LIGHT.DEEPLINK" "$root/INCY/LIGHT.DEEPLINK"
