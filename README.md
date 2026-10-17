# Compact GeoSite & GeoIP for Happ / Incy

[![Build & Release](https://github.com/commensal/domain-list-community/actions/workflows/build.yml/badge.svg)](https://github.com/commensal/domain-list-community/actions/workflows/build.yml)
[![GitHub Release](https://img.shields.io/github/v/release/commensal/domain-list-community?color=blue&label=Latest%20Release)](https://github.com/commensal/domain-list-community/releases/latest)

Компактные `geosite.dat` и `geoip.dat` для раздельной маршрутизации в Happ, Incy, Xray и sing-box.

Списки берутся из [itdoginfo/allow-domains](https://github.com/itdoginfo/allow-domains) и из `my-domains.txt` / `my-ips.txt`. Сборка запускается каждый день в 02:00 UTC, по пушу в `master` и вручную. Новый релиз и новое значение `LastUpdated` появляются только когда содержимое `geosite.dat` или `geoip.dat` изменилось: Happ скачивает геофайлы заново только в этом случае.

## Установка

Откройте ссылку с телефона. Профиль добавится в приложение.

| Клиент | Default | Full |
| :--- | :--- | :--- |
| Happ | [Добавить Default](https://raw.githubusercontent.com/commensal/domain-list-community/refs/heads/master/HAPP/DEFAULT.DEEPLINK) | [Добавить Full](https://raw.githubusercontent.com/commensal/domain-list-community/refs/heads/master/HAPP/FULL.DEEPLINK) |
| Incy | [Добавить Default](https://raw.githubusercontent.com/commensal/domain-list-community/refs/heads/master/INCY/DEFAULT.DEEPLINK) | [Добавить Full](https://raw.githubusercontent.com/commensal/domain-list-community/refs/heads/master/INCY/FULL.DEEPLINK) |

Страница с кнопками полного профиля: [index.html](index.html).

Профиля три не существует. `LIGHT` был копией `DEFAULT` и удалён.

### Что проксируется

**Default** — заблокированные зарубежные сервисы, без подсетей хостингов:

`geosite:telegram`, `meta`, `custom`, `discord`, `porn`, `youtube`, `google-play`, `google-meet`, `hdrezka`, `news`, `roblox`, `tiktok`, `twitter`, `google-ai`, `geoblock`

`geoip:telegram`, `twitter`, `meta`, `discord`

**Full** добавляет реестр блокировок, хостинги и оставшиеся категории:

`geosite:ru-block`, `cloudflare`, `cloudfront`, `digitalocean`, `hetzner`, `ovh`, `hodca`, `anime`

`geoip:roblox`, `google-meet`

Российские банки и Aviasales в Default идут напрямую, потому что их домены не входят в эти списки, а подсети Cloudflare, CloudFront, DigitalOcean, Hetzner и OVH подключены только в Full.

Оба профиля блокируют `appcenter.ms`, `firebase.io` и `crashlytics.com`.

## DNS

Оба профиля используют одинаковые DNS. FakeDNS включён: сервер перехватывает UDP/53, поэтому отдельный DoH-резолвер не нужен.

| Назначение | Тип | Адрес | IP |
| :--- | :--- | :--- | :--- |
| Remote DNS | DoU | `https://dns.dns-ai.ru/dns-query` | `8.8.8.8` |
| Domestic DNS | DoU | `https://dns.adguard-dns.com/dns-query` | `94.140.14.14` |

Фиксированные соответствия: `cloudflare-dns.com` → `1.1.1.1`, `dns.google` → `8.8.8.8`, `dns.adguard-dns.com` → `94.140.14.14`.

`DomainStrategy` — `AsIs`, `RouteOrder` — `block-proxy-direct`, `GlobalProxy` выключен, `UseChunkFiles` включён.

## Имена категорий

В профилях имена строчные: `geosite:custom`, `geoip:telegram`. Happ и Xray перед поиском переводят код в верхний регистр и сравнивают байты как есть, поэтому внутри `geosite.dat` и `geoip.dat` те же категории записаны как `CUSTOM` и `TELEGRAM`. `block` из allow-domains публикуется как `ru-block`, `google_ai` — как `google-ai`.

## Свои домены и IP

- `my-domains.txt` — домены, которых нет в категориях itdog. Комментарии отбрасываются, всё содержимое попадает в один список `geosite:custom`.
- `my-ips.txt` — CIDR для `geoip:custom`. Пока файл пуст, категория не создаётся и в профили не включается. После добавления адресов следующая сборка создаст категорию; ссылку `geoip:custom` нужно будет вернуть в профиль.

Пустая категория не заменяется на `example.com` или `127.0.0.1/32`, кроме `geosite:custom`: компилятору нужен хотя бы один домен, поэтому при пустом `my-domains.txt` туда записывается `example.com`.

## Файлы

Релиз публикуется на ветку `release` и в GitHub Release только при изменении баз.

- `geosite.dat`: `https://cdn.jsdelivr.net/gh/commensal/domain-list-community@release/geosite.dat`
- `geoip.dat`: `https://cdn.jsdelivr.net/gh/commensal/domain-list-community@release/geoip.dat`
- Release: `https://github.com/commensal/domain-list-community/releases/latest/download/geosite.dat`
- Raw: `https://raw.githubusercontent.com/commensal/domain-list-community/release/geosite.dat`
