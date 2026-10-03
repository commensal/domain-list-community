# ⚡ Compact GeoSite & GeoIP for Happ / Incy (Xray & Sing-box)

[![Build & Release](https://github.com/commensal/domain-list-community/actions/workflows/build.yml/badge.svg)](https://github.com/commensal/domain-list-community/actions/workflows/build.yml)
[![GitHub Release](https://img.shields.io/github/v/release/commensal/domain-list-community?color=blue&label=Latest%20Release)](https://github.com/commensal/domain-list-community/releases/latest)
[![Daily Update](https://img.shields.io/badge/Auto--Update-Daily%2002%3A00%20UTC-success)](https://github.com/commensal/domain-list-community/actions)
[![Clients](https://img.shields.io/badge/Clients-Happ%20%7C%20Incy%20%7C%20Xray%20%7C%20Sing--box-orange)](https://github.com/commensal/domain-list-community)

Оптимизированный генератор легковесных баз правил **`geosite.dat`** и **`geoip.dat`** для выборочной маршрутизации трафика (Split Tunneling). 

Базы автоматически компилируются каждый день в **02:00 UTC** через GitHub Actions, очищаются от мусора, нормализуются и объединяются со списками сообщества **[itdoginfo/allow-domains](https://github.com/itdoginfo/allow-domains)** и вашими персональными правилами.

---

## 🚀 Быстрая установка правил в 1 клик

Нажмите на ссылку ниже **с телефона** — профиль автоматически добавится в приложение.

| Клиент | 📦 Дефолтный профиль (Рекомендуемый / Легкий) | 🔥 Полный профиль (Все категории + RU-BLOCK) |
| :--- | :--- | :--- |
| 🟢 **Happ** (iOS / Android) | [📲 **Добавить в Happ**](https://raw.githubusercontent.com/commensal/domain-list-community/refs/heads/master/HAPP/DEFAULT.DEEPLINK) | [📲 **Добавить в Happ**](https://raw.githubusercontent.com/commensal/domain-list-community/refs/heads/master/HAPP/FULL.DEEPLINK) |
| 🔵 **Incy** (iOS / Android) | [📲 **Добавить в Incy**](https://raw.githubusercontent.com/commensal/domain-list-community/refs/heads/master/INCY/DEFAULT.DEEPLINK) | [📲 **Добавить в Incy**](https://raw.githubusercontent.com/commensal/domain-list-community/refs/heads/master/INCY/FULL.DEEPLINK) |

> 💡 **В чем разница?** 
> * **Дефолтный профиль (Рекомендуемый):** Включает только заблокированные зарубежные сервисы и исключает подсети хостингов (`cloudflare`, `cloudfront`, `digitalocean`, `hetzner`). Благодаря этому **Авиасейлс, Госуслуги и все российские банки (Сбер, Т-Банк и др.) работают напрямую** и не блокируются. Потребляет минимум памяти.
> * **Полный профиль:** Дополнительно включает масштабные базы `RU-BLOCK` (весь реестр РКН) и `GEOBLOCK`.

---

### 📋 Ссылки для ручного копирования

Если вы настраиваете приложение вручную, скопируйте нужную ссылку ниже:

#### 📦 1. Дефолтный профиль (Рекомендуемый / Легкий)
Оптимальный набор правил без тяжелых баз. Не ломает Авиасейлс и российские банки.

* **Включает домены (`ProxySites`):** `TELEGRAM`, `META`, `CUSTOM`, `DISCORD`, `PORN`, `YOUTUBE`, `GOOGLE-PLAY`, `GOOGLE-MEET`, `HDREZKA`, `NEWS`, `ROBLOX`, `TIKTOK`, `TWITTER`.
* **Включает подсети (`ProxyIp`):** `geoip:custom`, `geoip:telegram`, `geoip:twitter`, `geoip:meta`, `geoip:discord`.

* **Для приложения Happ:**
```text
happ://routing/add/ewogICAgIkJsb2NrSXAiOiBbCiAgICBdLAogICAgIkJsb2NrU2l0ZXMiOiBbCiAgICAgICAgImFwcGNlbnRlci5tcyIsCiAgICAgICAgImZpcmViYXNlLmlvIiwKICAgICAgICAiY3Jhc2hseXRpY3MuY29tIgogICAgXSwKICAgICJEaXJlY3RJcCI6IFsKICAgICAgICAiMTAuMC4wLjAvOCIsCiAgICAgICAgIjE3Mi4xNi4wLjAvMTIiLAogICAgICAgICIxOTIuMTY4LjAuMC8xNiIsCiAgICAgICAgIjE2OS4yNTQuMC4wLzE2IiwKICAgICAgICAiMjI0LjAuMC4wLzQiLAogICAgICAgICIyNTUuMjU1LjI1NS4yNTUiCiAgICBdLAogICAgIkRpcmVjdFNpdGVzIjogWwogICAgXSwKICAgICJEbnNIb3N0cyI6IHsKICAgICAgICAiY2xvdWRmbGFyZS1kbnMuY29tIjogIjEuMS4xLjEiLAogICAgICAgICJkbnMuZ29vZ2xlIjogIjguOC44LjgiCiAgICB9LAogICAgIkRvbWFpblN0cmF0ZWd5IjogIklQSWZOb25NYXRjaCIsCiAgICAiRG9tZXN0aWNETlNEb21haW4iOiAiaHR0cHM6Ly9nZW9oaWRlLnJ1L2Rucy1xdWVyeSIsCiAgICAiRG9tZXN0aWNETlNJUCI6ICI0NS4xNTUuMjA0LjE5MCIsCiAgICAiRG9tZXN0aWNETlNUeXBlIjogIkRvSCIsCiAgICAiRmFrZUROUyI6ICJ0cnVlIiwKICAgICJHZW9pcHVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb2lwLmRhdCIsCiAgICAiR2Vvc2l0ZXVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb3NpdGUuZGF0IiwKICAgICJHbG9iYWxQcm94eSI6ICJmYWxzZSIsCiAgICAiTGFzdFVwZGF0ZWQiOiAxNzg4MzM3MjYyLAogICAgIk5hbWUiOiAiQmxhY2sgTGlzdCIsCiAgICAiUHJveHlJcCI6IFsKICAgICAgICAiZ2VvaXA6Y3VzdG9tIiwKICAgICAgICAiZ2VvaXA6dGVsZWdyYW0iLAogICAgICAgICJnZW9pcDp0d2l0dGVyIiwKICAgICAgICAiZ2VvaXA6bWV0YSIsCiAgICAgICAgImdlb2lwOmRpc2NvcmQiCiAgICBdLAogICAgIlByb3h5U2l0ZXMiOiBbCiAgICAgICAgImdlb3NpdGU6VEVMRUdSQU0iLAogICAgICAgICJnZW9zaXRlOk1FVEEiLAogICAgICAgICJnZW9zaXRlOkNVU1RPTSIsCiAgICAgICAgImdlb3NpdGU6RElTQ09SRCIsCiAgICAgICAgImdlb3NpdGU6UE9STiIsCiAgICAgICAgImdlb3NpdGU6WU9VVFVCRSIsCiAgICAgICAgImdlb3NpdGU6R09PR0xFLVBMQVkiLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1NRUVUIiwKICAgICAgICAiZ2Vvc2l0ZTpIRFJFWktBIiwKICAgICAgICAiZ2Vvc2l0ZTpORVdTIiwKICAgICAgICAiZ2Vvc2l0ZTpST0JMT1giLAogICAgICAgICJnZW9zaXRlOlRJS1RPSyIsCiAgICAgICAgImdlb3NpdGU6VFdJVFRFUiIKICAgIF0sCiAgICAiUmVtb3RlRE5TRG9tYWluIjogImh0dHBzOi8vZXUuZ2VvaGlkZS5ydS9kbnMtcXVlcnkiLAogICAgIlJlbW90ZUROU0lQIjogIjIxNy42MC4yNDUuMjE5IiwKICAgICJSZW1vdGVETlNUeXBlIjogIkRvSCIsCiAgICAiUm91dGVPcmRlciI6ICJibG9jay1wcm94eS1kaXJlY3QiCn0K
```

* **Для приложения Incy:**
```text
incy://routing/add/ewogICAgIkJsb2NrSXAiOiBbCiAgICBdLAogICAgIkJsb2NrU2l0ZXMiOiBbCiAgICAgICAgImFwcGNlbnRlci5tcyIsCiAgICAgICAgImZpcmViYXNlLmlvIiwKICAgICAgICAiY3Jhc2hseXRpY3MuY29tIgogICAgXSwKICAgICJEaXJlY3RJcCI6IFsKICAgICAgICAiMTAuMC4wLjAvOCIsCiAgICAgICAgIjE3Mi4xNi4wLjAvMTIiLAogICAgICAgICIxOTIuMTY4LjAuMC8xNiIsCiAgICAgICAgIjE2OS4yNTQuMC4wLzE2IiwKICAgICAgICAiMjI0LjAuMC4wLzQiLAogICAgICAgICIyNTUuMjU1LjI1NS4yNTUiCiAgICBdLAogICAgIkRpcmVjdFNpdGVzIjogWwogICAgXSwKICAgICJEbnNIb3N0cyI6IHsKICAgICAgICAiY2xvdWRmbGFyZS1kbnMuY29tIjogIjEuMS4xLjEiLAogICAgICAgICJkbnMuZ29vZ2xlIjogIjguOC44LjgiCiAgICB9LAogICAgIkRvbWFpblN0cmF0ZWd5IjogIklQSWZOb25NYXRjaCIsCiAgICAiRG9tZXN0aWNETlNEb21haW4iOiAiaHR0cHM6Ly9nZW9oaWRlLnJ1L2Rucy1xdWVyeSIsCiAgICAiRG9tZXN0aWNETlNJUCI6ICI0NS4xNTUuMjA0LjE5MCIsCiAgICAiRG9tZXN0aWNETlNUeXBlIjogIkRvSCIsCiAgICAiRmFrZUROUyI6ICJ0cnVlIiwKICAgICJHZW9pcHVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb2lwLmRhdCIsCiAgICAiR2Vvc2l0ZXVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb3NpdGUuZGF0IiwKICAgICJHbG9iYWxQcm94eSI6ICJmYWxzZSIsCiAgICAiTGFzdFVwZGF0ZWQiOiAxNzg4MzM3MjYyLAogICAgIk5hbWUiOiAiQmxhY2sgTGlzdCIsCiAgICAiUHJveHlJcCI6IFsKICAgICAgICAiZ2VvaXA6Y3VzdG9tIiwKICAgICAgICAiZ2VvaXA6dGVsZWdyYW0iLAogICAgICAgICJnZW9pcDp0d2l0dGVyIiwKICAgICAgICAiZ2VvaXA6bWV0YSIsCiAgICAgICAgImdlb2lwOmRpc2NvcmQiCiAgICBdLAogICAgIlByb3h5U2l0ZXMiOiBbCiAgICAgICAgImdlb3NpdGU6VEVMRUdSQU0iLAogICAgICAgICJnZW9zaXRlOk1FVEEiLAogICAgICAgICJnZW9zaXRlOkNVU1RPTSIsCiAgICAgICAgImdlb3NpdGU6RElTQ09SRCIsCiAgICAgICAgImdlb3NpdGU6UE9STiIsCiAgICAgICAgImdlb3NpdGU6WU9VVFVCRSIsCiAgICAgICAgImdlb3NpdGU6R09PR0xFLVBMQVkiLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1NRUVUIiwKICAgICAgICAiZ2Vvc2l0ZTpIRFJFWktBIiwKICAgICAgICAiZ2Vvc2l0ZTpORVdTIiwKICAgICAgICAiZ2Vvc2l0ZTpST0JMT1giLAogICAgICAgICJnZW9zaXRlOlRJS1RPSyIsCiAgICAgICAgImdlb3NpdGU6VFdJVFRFUiIKICAgIF0sCiAgICAiUmVtb3RlRE5TRG9tYWluIjogImh0dHBzOi8vZXUuZ2VvaGlkZS5ydS9kbnMtcXVlcnkiLAogICAgIlJlbW90ZUROU0lQIjogIjIxNy42MC4yNDUuMjE5IiwKICAgICJSZW1vdGVETlNUeXBlIjogIkRvSCIsCiAgICAiUm91dGVPcmRlciI6ICJibG9jay1wcm94eS1kaXJlY3QiCn0K
```

#### 🔥 2. Полный профиль (Со всеми категориями)
Включает **абсолютно всё** из Дефолтного профиля, а также гигантские базы `RU-BLOCK` и `GEOBLOCK`.

* **Для приложения Happ:**
```text
happ://routing/add/ewogICAgIkJsb2NrSXAiOiBbXSwKICAgICJCbG9ja1NpdGVzIjogWwogICAgICAgICJhcHBjZW50ZXIubXMiLAogICAgICAgICJmaXJlYmFzZS5pbyIsCiAgICAgICAgImNyYXNobHl0aWNzLmNvbSIKICAgIF0sCiAgICAiRGlyZWN0SXAiOiBbCiAgICAgICAgIjEwLjAuMC4wLzgiLAogICAgICAgICIxNzIuMTYuMC4wLzEyIiwKICAgICAgICAiMTkyLjE2OC4wLjAvMTYiLAogICAgICAgICIxNjkuMjU0LjAuMC8xNiIsCiAgICAgICAgIjIyNC4wLjAuMC80IiwKICAgICAgICAiMjU1LjI1NS4yNTUuMjU1IgogICAgXSwKICAgICJEaXJlY3RTaXRlcyI6IFtdLAogICAgIkRuc0hvc3RzIjogewogICAgICAgICJjbG91ZGZsYXJlLWRucy5jb20iOiAiMS4xLjEuMSIsCiAgICAgICAgImRucy5nb29nbGUiOiAiOC44LjguOCIKICAgIH0sCiAgICAiRG9tYWluU3RyYXRlZ3kiOiAiSVBJZk5vbk1hdGNoIiwKICAgICJEb21lc3RpY0ROU0RvbWFpbiI6ICJodHRwczovL2dlb2hpZGUucnUvZG5zLXF1ZXJ5IiwKICAgICJEb21lc3RpY0ROU0lQIjogIjQ1LjE1NS4yMDQuMTkwIiwKICAgICJEb21lc3RpY0ROU1R5cGUiOiAiRG9IIiwKICAgICJGYWtlRE5TIjogInRydWUiLAogICAgIkdlb3NpdGV1cmwiOiAiaHR0cHM6Ly9jZG4uanNkZWxpdnIubmV0L2doL2NvbW1lbnNhbC9kb21haW4tbGlzdC1jb21tdW5pdHlAcmVsZWFzZS9nZW9pcC5kYXQiLAogICAgIkdlb3NpdGV1cmwiOiAiaHR0cHM6Ly9jZG4uanNkZWxpdnIubmV0L2doL2NvbW1lbnNhbC9kb21haW4tbGlzdC1jb21tdW5pdHlAcmVsZWFzZS9nZW9zaXRlLmRhdCIsCiAgICAiR2xvYmFsUHJveHkiOiAiZmFsc2UiLAogICAgIkxhc3RVcGRhdGVkIjogMTc4ODMzNzI2MiwKICAgICJOYW1lIjogIkJsYWNrIExpc3QgKEZ1bGwpIiwKICAgICJQcm94eUlwIjogWwogICAgICAgICJnZW9pcDpjdXN0b20iLAogICAgICAgICJnZW9pcDp0ZWxlZ3JhbSIsCiAgICAgICAgImdlb2lwOnR3aXR0ZXIiLAogICAgICAgICJnZW9pcDptZXRhIiwKICAgICAgICAiZ2VvaXA6ZGlzY29yZCIsCiAgICAgICAgImdlb2lwOnJvYmxveCIsCiAgICAgICAgImdlb2lwOmdvb2dsZS1tZWV0IgogICAgXSwKICAgICJQcm94eVNpdGVzIjogWwogICAgICAgICJnZW9zaXRlOkNVU1RPTSIsCiAgICAgICAgImdlb3NpdGU6UlUtQkxPQ0siLAogICAgICAgICJnZW9zaXRlOkdFT0JMT0NLIiwKICAgICAgICAiZ2Vvc2l0ZTpURUxFR1JBTSIsCiAgICAgICAgImdlb3NpdGU6TUVUQSIsCiAgICAgICAgImdlb3NpdGU6WU9VVFVCRSIsCiAgICAgICAgImdlb3NpdGU6RElTQ09SRCIsCiAgICAgICAgImdlb3NpdGU6VFdJVFRFUiIsCiAgICAgICAgImdlb3NpdGU6VElLVE9LIiwKICAgICAgICAiZ2Vvc2l0ZTpST0JMT1giLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1QTEFZIiwKICAgICAgICAiZ2Vvc2l0ZTpHT09HTEUtTUVFVCIsCiAgICAgICAgImdlb3NpdGU6R09PR0xFLUFJIiwKICAgICAgICAiZ2Vvc2l0ZTpDTE9VREZMQVJFIiwKICAgICAgICAiZ2Vvc2l0ZTpDTE9VREZST05UIiwKICAgICAgICAiZ2Vvc2l0ZTpESUdJVEFMT0NFQU4iLAogICAgICAgICJnZW9zaXRlOkhFVFpORVIiLAogICAgICAgICJnZW9zaXRlOk9WSCIsCiAgICAgICAgImdlb3NpdGU6SERSRVpLQSIsCiAgICAgICAgImdlb3NpdGU6SE9EQ0EiLAogICAgICAgICJnZW9zaXRlOk5FV1MiLAogICAgICAgICJnZW9zaXRlOlBPUk4iLAogICAgICAgICJnZW9zaXRlOkFOSU1FIgogICAgXSwKICAgICJSZW1vdGVETlNEb21haW4iOiAiaHR0cHM6Ly9ldS5nZW9oaWRlLnJ1L2Rucy1xdWVyeSIsCiAgICAiUmVtb3RlRE5TSVAiOiAiMjE3LjYwLjI0NS4yMTkiLAogICAgIlJlbW90ZUROU1R5cGUiOiAiRG9IIiwKICAgICJSb3V0ZU9yZGVyIjogImJsb2NrLXByb3h5LWRpcmVjdCIKfQo=
```

* **Для приложения Incy:**
```text
incy://routing/add/ewogICAgIkJsb2NrSXAiOiBbXSwKICAgICJCbG9ja1NpdGVzIjogWwogICAgICAgICJhcHBjZW50ZXIubXMiLAogICAgICAgICJmaXJlYmFzZS5pbyIsCiAgICAgICAgImNyYXNobHl0aWNzLmNvbSIKICAgIF0sCiAgICAiRGlyZWN0SXAiOiBbCiAgICAgICAgIjEwLjAuMC4wLzgiLAogICAgICAgICIxNzIuMTYuMC4wLzEyIiwKICAgICAgICAiMTkyLjE2OC4wLjAvMTYiLAogICAgICAgICIxNjkuMjU0LjAuMC8xNiIsCiAgICAgICAgIjIyNC4wLjAuMC80IiwKICAgICAgICIyNTUuMjU1LjI1NS4yNTUiCiAgICBdLAogICAgIkRpcmVjdFNpdGVzIjogWwogICAgXSwKICAgICJEbnNIb3N0cyI6IHsKICAgICAgICAiY2xvdWRmbGFyZS1kbnMuY29tIjogIjEuMS4xLjEiLAogICAgICAgICJkbnMuZ29vZ2xlIjogIjguOC44LjgiCiAgICB9LAogICAgIkRvbWFpblN0cmF0ZWd5IjogIklQSWZOb25NYXRjaCIsCiAgICAiRG9tZXN0aWNETlNEb21haW4iOiAiaHR0cHM6Ly9nZW9oaWRlLnJ1L2Rucy1xdWVyeSIsCiAgICAiRG9tZXN0aWNETlNJUCI6ICI0NS4xNTUuMjA0LjE5MCIsCiAgICAiRG9tZXN0aWNETlNUeXBlIjogIkRvSCIsCiAgICAiRmFrZUROUyI6ICJ0cnVlIiwKICAgICJHZW9pcHVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb2lwLmRhdCIsCiAgICAiR2Vvc2l0ZXVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb3NpdGUuZGF0IiwKICAgICJHbG9iYWxQcm94eSI6ICJmYWxzZSIsCiAgICAiTGFzdFVwZGF0ZWQiOiAxNzg4MzM3MjYyLAogICAgIk5hbWUiOiAiQmxhY2sgTGlzdCAoRnVsbCkiLAogICAgIlByb3h5SXAiOiBbCiAgICAgICAgImdlb2lwOmN1c3RvbSIsCiAgICAgICAgImdlb2lwOnRlbGVncmFtIiwKICAgICAgICAiZ2VvaXA6dHdpdHRlciIsCiAgICAgICAgImdlb2lwOm1ldGEiLAogICAgICAgICJnZW9pcDpkaXNjb3JkIiwKICAgICAgICAiZ2VvaXA6cm9ibG94IiwKICAgICAgICAiZ2VvaXA6Z29vZ2xlLW1lZXQiCiAgICBdLAogICAgIlByb3h5U2l0ZXMiOiBbCiAgICAgICAgImdlb3NpdGU6Q1VTVE9NIiwKICAgICAgICAiZ2Vvc2l0ZTpSVS1CTE9DSyIsCiAgICAgICAgImdlb3NpdGU6R0VPQkxPQ0siLAogICAgICAgICJnZW9zaXRlOlRFTEVHUkFNIiwKICAgICAgICAiZ2Vvc2l0ZTpNRVRBIiwKICAgICAgICAiZ2Vvc2l0ZTpZT1VUVUJFIiwKICAgICAgICAiZ2Vvc2l0ZTpESVNDT1JEIiwKICAgICAgICAiZ2Vvc2l0ZTpUV0lUVEVSIiwKICAgICAgICAiZ2Vvc2l0ZTpUSUtUT0siLAogICAgICAgICJnZW9zaXRlOlJPQkxPWCIsCiAgICAgICAgImdlb3NpdGU6R09PR0xFLVBMQVkiLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1NRUVUIiwKICAgICAgICAiZ2Vvc2l0ZTpHT09HTEUtQUkiLAogICAgICAgICJnZW9zaXRlOkNMT1VERkxBUkUiLAogICAgICAgICJnZW9zaXRlOkNMT1VERlJPTlQiLAogICAgICAgICJnZW9zaXRlOkRJR0lUQUxPQ0VBTiIsCiAgICAgICAgImdlb3NpdGU6SEVUWk5FUiIsCiAgICAgICAgImdlb3NpdGU6T1ZIIiwKICAgICAgICAiZ2Vvc2l0ZTpIRFJFWktBIiwKICAgICAgICAiZ2Vvc2l0ZTpIT0RDQSIsCiAgICAgICAgImdlb3NpdGU6TkVXUyIsCiAgICAgICAgImdlb3NpdGU6UE9STiIsCiAgICAgICAgImdlb3NpdGU6QU5JTUUiCiAgICBdLAogICAgIlJlbW90ZUROU0RvbWFpbiI6ICJodHRwczovL2V1Lmdlb2hpZGUucnUvZG5zLXF1ZXJ5IiwKICAgICJSZW1vdGVETlNJUCI6ICIyMTcuNjAuMjQ1LjIxOSIsCiAgICAiUmVtb3RlRE5TVHlwZSI6ICJEb0giLAogICAgIlJvdXRlT3JkZXIiOiAiYmxvY2stcHJveHktZGlyZWN0Igp9
```

---

## 🧠 Архитектура DNS: как работают ИИ без прокси и почему именно GeoHide

В этом профиле маршрутизации используется продуманная гибридная схема DNS, решающая одну из главных проблем пользователей современных VPN:

### ⚠️ Проблема: Google блокирует доступ к Gemini через VPS
Большинство пользователей пускают весь трафик или заблокированные ресурсы через зарубежные VPS (Hetzner, DigitalOcean, OVH и др.). 
Однако **Google активно отслеживает IP-адреса хостинг-провайдеров и дата-центров**. 
При попытке войти в **Google Gemini**, **Google AI Studio** или пользоваться расширенным поиском через обычный VPS вы сталкиваетесь с ошибками:
* `403 Forbidden` / `Location not supported`
* Бесконечные проверки Cloudflare / Google ReCAPTCHA
* Внезапные сбросы сессии

### ✅ Решение: Smart DoH через GeoHide (`DomesticDNSDomain`)
В конфиге в качестве внутреннего DNS задан:
* **`DomesticDNSDomain`**: `https://geohide.ru/dns-query` (IP: `45.155.204.190`, тип `DoH`)

**GeoHide** работает как **Smart DNS**:
1. Он не пускает тяжелый трафик через медленные туннели, а точечно модифицирует DNS-ответы для регионально заблокированных ИИ-сервисов (Google Gemini, ChatGPT/OpenAI, Claude/Anthropic).
2. Запросы к этим сервисам идут **напрямую (Direct)** через чистые распределенные SNI-шлюзы.
3. **Результат:** Gemini, ChatGPT и Claude работают **без прокси на максимальной скорости вашего интернет-провайдера**, а Google не видит «подозрительный» дата-центровый IP вашего сервера и не блокирует аккаунт!

### 🛡 Внешний безопасный DoH (`RemoteDNSDomain`)
* **`RemoteDNSDomain`**: `https://eu.geohide.ru/dns-query` (IP: `217.60.245.219`, тип `DoH`)
Европейская нода GeoHide используется для разрешения доменов, которые направляются в прокси. Она гарантирует:
- Полное шифрование DNS-запросов через HTTPS (DoH).
- Защиту от прослушивания и подмены DNS со стороны местного интернет-провайдера (Anti-DNS Poisoning).
- Отсутствие утечек реального DNS (No DNS Leaks).

### ⚡ FakeDNS (`"FakeDNS": "true"`)
* **Включен (`true`)**: Позволяет клиенту мгновенно отвечать на DNS-запросы фиктивными IP-адресами из пула `198.18.0.0/15`, сокращая задержку резолвинга до 0 мс (0 ms RTT) и предотвращая утечки DNS.
* В современных клиентах (включая **Incy** на iOS и клиенты на Android) звонки и WebRTC работают штатно и без сбоев.

### 📌 Bootstrap DNS (`DnsHosts`)
```json
"DnsHosts": {
    "cloudflare-dns.com": "1.1.1.1",
    "dns.google": "8.8.8.8"
}
```
Статическая привязка IP-адресов исключает «замкнутый круг» (deadlock), когда для подключения к DoH-серверу клиенту сначала нужно было бы определить его IP через обычный незашифрованный DNS.

---

## ⚙️ Параметры конфигурации профиля

Полный разбор параметров, настроенных в правиле:

| Параметр | Значение | Зачем это нужно |
| :--- | :--- | :--- |
| **`DomainStrategy`** | `IPIfNonMatch` | Умная проверка: сначала сверяются доменные списки (`geosite`). Если совпадений нет, домен резолвится в IP и проверяются IP-правила (`geoip`). Дает максимальное быстродействие. |
| **`RouteOrder`** | `block-proxy-direct` | Порядок обработки: **1. Block** (реклама/слежка) ➔ **2. Proxy** (заблокированные сайты) ➔ **3. Direct** (весь остальной Рунет и обычные сайты). |
| **`GlobalProxy`** | `false` | Выборочный режим (Split Tunneling). Трафик российских банков, Госуслуг и локальных сайтов идет напрямую и не ломается. |
| **`FakeDNS`** | `true` | Мгновенный DNS-ответ из пула 198.18.0.0/15 (0 мс задержки) и защита от утечек. |
| **`BlockSites`** | `appcenter.ms`, `firebase.io`, `crashlytics.com` | Блокировка мобильной телеметрии, трекеров и аналитики крашей. Экономит заряд батареи и трафик. |
| **`DirectIp`** | `10.0.0.0/8`, `192.168.0.0/16`, `172.16.0.0/12`... | Исключение локальных сетей (RFC 1918). Доступ к роутеру, сетевым принтерам и умному дому всегда работает напрямую. |
| **`ProxySites`** | Категории `geosite:*` | Список категорий доменов, направляемых в прокси. |
| **`ProxyIp`** | Категории `geoip:*` | Список подсетей сервисов, трафик к которым отправляется через прокси. |

---

## 📦 Доступные категории в базах

### 🌐 Домены (`geosite.dat`)
Скомпилированы из репозитория `itdoginfo/allow-domains` + персональные правила:

* `geosite:CUSTOM` — ваши персональные домены из файла `my-domains.txt`.
* `geosite:TELEGRAM` — веб-версии, CDN и инфраструктура Telegram.
* `geosite:META` — Instagram, Facebook, Threads, WhatsApp web.
* `geosite:YOUTUBE` — видеохостинг YouTube и CDN `googlevideo`.
* `geosite:TWITTER` — социальная сеть X (Twitter) и медиа-сервера.
* `geosite:DISCORD` — голосовой и текстовый мессенджер Discord.
* `geosite:TIKTOK` — платформа TikTok и стриминговые сервера.
* `geosite:ROBLOX` — игровая платформа Roblox.
* `geosite:GOOGLE-PLAY` / `geosite:GOOGLE-MEET` — сервисы Google.
* `geosite:GOOGLE-AI` — Google Gemini, AI Studio и связанные API.
* `geosite:HDREZKA` — онлайн-кинотеатр HDRezka.
* `geosite:NEWS` — независимые и зарубежные новостные издания.
* `geosite:PORN` / `geosite:HODCA` / `geosite:ANIME` — тематические категории.
* `geosite:CLOUDFLARE` / `geosite:CLOUDFRONT` — зарубежные CDN.
* `geosite:RU-BLOCK` / `geosite:GEOBLOCK` — реестры блокировок РКН и сервисы с гео-ограничениями.

### 📍 IP-адреса и подсети (`geoip.dat`)
В дефолтном профиле маршрутизации используются только специализированные подсети:
* `geoip:custom` — ваши персональные IP из `my-ips.txt`.
* `geoip:telegram`
* `geoip:twitter`
* `geoip:meta`
* `geoip:discord`

> 🛡 **Обратите внимание:** Подсети хостингов (`geoip:cloudflare`, `geoip:cloudfront`, `geoip:digitalocean`, `geoip:hetzner`, `geoip:ovh`) исключены из автоматической маршрутизации `ProxyIp`. Это предотвращает ошибочное попадание в VPN белых российских сайтов (Авиасейлс, Госуслуги, банки), защищенных через Cloudflare или размещенных на зарубежных серверах.

---

## 🏆 Чем этот репозиторий отличается от других

1. **Минимальный вес баз (~28 Кб против 30–60 Мб в стандартных V2Fly):**
   Обычные сборки geosite содержат сотни тысяч доменов со всего мира, из-за чего мобильные клиенты долго загружают конфигурацию, греют процессор и потребляют батарею. Здесь содержатся **только актуальные списки**, необходимые для обхода блокировок.
2. **Идеальная совместимость имен правил:**
   Строгий компилятор `domain-list-community` падает при наличии подчеркиваний `_` в названиях списков. В нашем пайплайне все имена автоматически валидируются и конвертируются в kebab-case (`google_play` ➔ `google-play`), предотвращая падение клиентов.
3. **Защита от Windows переносов (`\r\n`):**
   При редактировании `my-domains.txt` на Windows невидимый символ возврата каретки `\r` очищается пайплайном на лету.
4. **Автоматическое слияние IPv4 + IPv6:**
   Списки подсетей из разных каталогов автоматически склеиваются в единые категории `geoip` без дубликатов.

---

## 📥 Прямые ссылки на файлы баз

Если вам нужно указать ссылки на файлы вручную в других клиентах (v2rayN, Nekoray, Clash, Sing-box):

### Вариант 1: GitHub Releases (Рекомендуется — свежие файлы без задержек)
* **geosite.dat:** `https://github.com/commensal/domain-list-community/releases/latest/download/geosite.dat`
* **geoip.dat:** `https://github.com/commensal/domain-list-community/releases/latest/download/geoip.dat`

### Вариант 2: Raw GitHub (Кэш обновления до 5 минут)
* **geosite.dat:** `https://raw.githubusercontent.com/commensal/domain-list-community/release/geosite.dat`
* **geoip.dat:** `https://raw.githubusercontent.com/commensal/domain-list-community/release/geoip.dat`

### Вариант 3: jsDelivr CDN
* **geosite.dat:** `https://cdn.jsdelivr.net/gh/commensal/domain-list-community@release/geosite.dat`
* **geoip.dat:** `https://cdn.jsdelivr.net/gh/commensal/domain-list-community@release/geoip.dat`

---

## 🛠 Добавление собственных доменов и IP

В корне репозитория есть два файла:
* **`my-domains.txt`** — добавьте сюда ваши домены (по одному на строку). Они автоматически попадут в категорию `geosite:CUSTOM`.
* **`my-ips.txt`** — добавьте сюда ваши IP-адреса или подсети в формате CIDR (например, `1.2.3.4/32`). Они попадут в категорию `geoip:custom`.

После коммита изменений GitHub Actions автоматически пересоберет базы и опубликует свежий релиз!
