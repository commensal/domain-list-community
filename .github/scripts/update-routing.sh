#!/usr/bin/env bash
# Rewrite the two Happ/Incy deeplinks and index.html.
# LastUpdated moves forward only when geosite.dat, geoip.dat or profile settings changed,
# which is the signal Happ uses to refresh routing configuration.
set -euo pipefail

root="${1:?repository root}"
old_geosite="${2:-}"
old_geoip="${3:-}"
# The workflow builds both files in its own directory, not inside the checkout.
new_geosite="${4:-$root/geosite.dat}"
new_geoip="${5:-$root/geoip.dat}"

# The heredoc occupies stdin, so the paths travel through the environment.
ROUTING_ROOT="$root" \
ROUTING_OLD_GEOSITE="$old_geosite" \
ROUTING_OLD_GEOIP="$old_geoip" \
ROUTING_NEW_GEOSITE="$new_geosite" \
ROUTING_NEW_GEOIP="$new_geoip" \
python3 - << 'PY'
import base64, json, os, pathlib, time

root = pathlib.Path(os.environ["ROUTING_ROOT"])
old_geosite = os.environ["ROUTING_OLD_GEOSITE"]
old_geoip = os.environ["ROUTING_OLD_GEOIP"]
new_geosite = os.environ["ROUTING_NEW_GEOSITE"]
new_geoip = os.environ["ROUTING_NEW_GEOIP"]

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
_, orig_default = load(default_path)
_, orig_full = load(full_path)

previous = 0
for profile in (orig_default, orig_full):
    try:
        previous = max(previous, int(profile.get("LastUpdated") or 0))
    except (TypeError, ValueError):
        pass

def make_profile(name, proxy_sites, proxy_ip):
    return {
        "Name": name,
        "LastUpdated": str(previous),
        "GlobalProxy": "false",
        "RouteOrder": "block-proxy-direct",
        "RemoteDNSType": "DoU",
        "RemoteDNSDomain": "dns.google",
        "RemoteDNSIP": "8.8.8.8",
        "DomesticDNSType": "DoU",
        "DomesticDNSDomain": "dns.adguard-dns.com",
        "DomesticDNSIP": "94.140.14.14",
        "Geoipurl": "https://cdn.jsdelivr.net/gh/commensal/domain-list-community@release/geoip.dat",
        "Geositeurl": "https://cdn.jsdelivr.net/gh/commensal/domain-list-community@release/geosite.dat",
        "DnsHosts": {
            "cloudflare-dns.com": "1.1.1.1",
            "dns.google": "8.8.8.8",
            "dns.adguard-dns.com": "94.140.14.14",
        },
        "DirectSites": [],
        "DirectIp": [
            "10.0.0.0/8",
            "172.16.0.0/12",
            "192.168.0.0/16",
            "169.254.0.0/16",
            "224.0.0.0/4",
            "255.255.255.255",
        ],
        "BlockSites": [
            "appcenter.ms",
            "firebase.io",
            "crashlytics.com",
        ],
        "BlockIp": [],
        "DomainStrategy": "AsIs",
        "FakeDNS": "true",
        "UseChunkFiles": "true",
        "ProxySites": proxy_sites,
        "ProxyIp": proxy_ip,
    }

default_profile = make_profile(
    "Black List",
    proxy_sites=[
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
    ],
    proxy_ip=[
        "geoip:telegram",
        "geoip:twitter",
        "geoip:meta",
        "geoip:discord",
    ],
)

full_profile = make_profile(
    "Black List (Full)",
    proxy_sites=[
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
    ],
    proxy_ip=[
        "geoip:telegram",
        "geoip:twitter",
        "geoip:meta",
        "geoip:discord",
        "geoip:roblox",
        "geoip:google-meet",
    ],
)

def fields_differ(p1, p2):
    f1 = {k: v for k, v in p1.items() if k != "LastUpdated"}
    f2 = {k: v for k, v in p2.items() if k != "LastUpdated"}
    return f1 != f2

config_changed = fields_differ(orig_default, default_profile) or fields_differ(orig_full, full_profile)
should_bump = changed or config_changed

# A newer unix time is what makes an already imported profile fetch geo files.
# An unchanged profile must keep the old value, or clients re-download for nothing.
stamp = str(max(int(time.time()), previous + 1)) if should_bump else str(previous)
default_profile["LastUpdated"] = stamp
full_profile["LastUpdated"] = stamp

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
print(f"profile_changed={'true' if config_changed else 'false'}")
print(f"routing_changed={'true' if should_bump else 'false'}")
print(f"last_updated={stamp}")
PY
