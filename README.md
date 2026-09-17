# ⚡ Compact GeoSite & GeoIP for Happ / Incy (Xray & Sing-box)

[![Build & Release](https://github.com/commensal/domain-list-community/actions/workflows/build.yml/badge.svg)](https://github.com/commensal/domain-list-community/actions/workflows/build.yml)
[![GitHub Release](https://img.shields.io/github/v/release/commensal/domain-list-community?color=blue&label=Latest%20Release)](https://github.com/commensal/domain-list-community/releases/latest)
[![Daily Update](https://img.shields.io/badge/Auto--Update-Daily%2002%3A00%20UTC-success)](https://github.com/commensal/domain-list-community/actions)
[![Clients](https://img.shields.io/badge/Clients-Happ%20%7C%20Incy%20%7C%20Xray%20%7C%20Sing--box-orange)](https://github.com/commensal/domain-list-community)

Оптимизированный генератор легковесных баз правил **`geosite.dat`** и **`geoip.dat`** для выборочной маршрутизации трафика (Split Tunneling). 

Базы автоматически компилируются каждый день в **02:00 UTC** через GitHub Actions, очищаются от мусора, нормализуются и объединяются со списками сообщества **[itdoginfo/allow-domains](https://github.com/itdoginfo/allow-domains)** и вашими персональными правилами.

---

## 🚀 Готовые Deeplink ссылки для импорта правил

> ⚠️ **Важная информация по импорту:**
> Веб-интерфейс GitHub в целях безопасности блокирует переход по сторонним протоколам (`happ://` и `incy://`) при обычном клике мышкой.
> 
> **Как мгновенно добавить правило:**
> 1. Нажмите **иконку копирования (📋)** в правом верхнем углу нужного блока с кодом ниже.
> 2. **В мобильном браузере:** Вставьте скопированную ссылку в адресную строку Safari или Chrome на телефоне и нажмите **Перейти (Go)** — система предложит открыть Happ или Incy и сразу сохранит правило.
> 3. **Внутри приложения:** В Happ или Incy перейдите во вкладку **Маршрутизация (Routing)** ➔ нажмите **+** (Добавить) ➔ выберите **Импорт из буфера**.

---

### 🔥 1. Полный профиль (Рекомендуется — все категории репозитория)
Включает полный набор: `RU-BLOCK`, `GEOBLOCK`, `DISCORD`, `GOOGLE-AI`, `TELEGRAM`, `META`, `YOUTUBE`, `TWITTER`, `TIKTOK`, `ROBLOX`, `CLOUDFLARE`, `CLOUDFRONT`, `DIGITALOCEAN`, `HETZNER`, `OVH`, `HDREZKA`, `NEWS`, `HODCA`, `PORN`, `ANIME` + все подсети IP (`geoip`).

#### 📱 Для приложения Happ (iOS / Android):
```text
happ://routing/add/ewogICAgIkJsb2NrSXAiOiBbXSwKICAgICJCbG9ja1NpdGVzIjogWwogICAgICAgICJhcHBjZW50ZXIubXMiLAogICAgICAgICJmaXJlYmFzZS5pbyIsCiAgICAgICAgImNyYXNobHl0aWNzLmNvbSIKICAgIF0sCiAgICAiRGlyZWN0SXAiOiBbCiAgICAgICAgIjEwLjAuMC4wLzgiLAogICAgICAgICIxNzIuMTYuMC4wLzEyIiwKICAgICAgICAiMTkyLjE2OC4wLjAvMTYiLAogICAgICAgICIxNjkuMjU0LjAuMC8xNiIsCiAgICAgICAgIjIyNC4wLjAuMC80IiwKICAgICAgICAiMjU1LjI1NS4yNTUuMjU1IgogICAgXSwKICAgICJEaXJlY3RTaXRlcyI6IFtdLAogICAgIkRuc0hvc3RzIjogewogICAgICAgICJjbG91ZGZsYXJlLWRucy5jb20iOiAiMS4xLjEuMSIsCiAgICAgICAgImRucy5nb29nbGUiOiAiOC44LjguOCIKICAgIH0sCiAgICAiRG9tYWluU3RyYXRlZ3kiOiAiSVBJZk5vbk1hdGNoIiwKICAgICJEb21lc3RpY0ROU0RvbWFpbiI6ICJodHRwczovL2dlb2hpZGUucnUvZG5zLXF1ZXJ5IiwKICAgICJEb21lc3RpY0ROU0lQIjogIjQ1LjE1NS4yMDQuMTkwIiwKICAgICJEb21lc3RpY0ROU1R5cGUiOiAiRG9IIiwKICAgICJGYWtlRE5TIjogInRydWUiLAogICAgIkdlb2lwdXJsIjogImh0dHBzOi8vY2RuLmpzZGVsaXZyLm5ldC9naC9jb21tZW5zYWwvZG9tYWluLWxpc3QtY29tbXVuaXR5QHJlbGVhc2UvZ2VvaXAuZGF0IiwKICAgICJHZW9zaXRldXJsIjogImh0dHBzOi8vY2RuLmpzZGVsaXZyLm5ldC9naC9jb21tZW5zYWwvZG9tYWluLWxpc3QtY29tbXVuaXR5QHJlbGVhc2UvZ2Vvc2l0ZS5kYXQiLAogICAgIkdsb2JhbFByb3h5IjogImZhbHNlIiwKICAgICJMYXN0VXBkYXRlZCI6IDE3ODgzMzcyNjIsCiAgICAiTmFtZSI6ICJCbGFjayBMaXN0IChGdWxsKSIsCiAgICAiUHJveHlJcCI6IFsKICAgICAgICAiZ2VvaXA6Y3VzdG9tIiwKICAgICAgICAiZ2VvaXA6dGVsZWdyYW0iLAogICAgICAgICJnZW9pcDp0d2l0dGVyIiwKICAgICAgICAiZ2VvaXA6bWV0YSIsCiAgICAgICAgImdlb2lwOmRpc2NvcmQiLAogICAgICAgICJnZW9pcDpjbG91ZGZyb250IiwKICAgICAgICAiZ2VvaXA6Y2xvdWRmbGFyZSIsCiAgICAgICAgImdlb2lwOmRpZ2l0YWxvY2VhbiIsCiAgICAgICAgImdlb2lwOmhldHpuZXIiLAogICAgICAgICJnZW9pcDpvdmgiLAogICAgICAgICJnZW9pcDpyb2Jsb3giLAogICAgICAgICJnZW9pcDpnb29nbGUtbWVldCIKICAgIF0sCiAgICAiUHJveHlTaXRlcyI6IFsKICAgICAgICAiZ2Vvc2l0ZTpDVVNUT00iLAogICAgICAgICJnZW9zaXRlOlJVLUJMT0NLIiwKICAgICAgICAiZ2Vvc2l0ZTpHRU9CTE9DSyIsCiAgICAgICAgImdlb3NpdGU6VEVMRUdSQU0iLAogICAgICAgICJnZW9zaXRlOk1FVEEiLAogICAgICAgICJnZW9zaXRlOllPVVRVQkUiLAogICAgICAgICJnZW9zaXRlOkRJU0NPUkQiLAogICAgICAgICJnZW9zaXRlOlRXSVRURVIiLAogICAgICAgICJnZW9zaXRlOlRJS1RPSyIsCiAgICAgICAgImdlb3NpdGU6Uk9CTE9YIiwKICAgICAgICAiZ2Vvc2l0ZTpHT09HTEUtUExBWSIsCiAgICAgICAgImdlb3NpdGU6R09PR0xFLU1FRVQiLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1BSSIsCiAgICAgICAgImdlb3NpdGU6Q0xPVURGTEFSRSIsCiAgICAgICAgImdlb3NpdGU6Q0xPVURGUk9OVCIsCiAgICAgICAgImdlb3NpdGU6RElHSVRBTE9DRUFOIiwKICAgICAgICAiZ2Vvc2l0ZTpIRVRaTkVSIiwKICAgICAgICAiZ2Vvc2l0ZTpPVkgiLAogICAgICAgICJnZW9zaXRlOkhEUkVaS0EiLAogICAgICAgICJnZW9zaXRlOkhPRENBIiwKICAgICAgICAiZ2Vvc2l0ZTpORVdTIiwKICAgICAgICAiZ2Vvc2l0ZTpQT1JOIiwKICAgICAgICAiZ2Vvc2l0ZTpBTklNRSIKICAgIF0sCiAgICAiUmVtb3RlRE5TRG9tYWluIjogImh0dHBzOi8vZXUuZ2VvaGlkZS5ydS9kbnMtcXVlcnkiLAogICAgIlJlbW90ZUROU0lQIjogIjIxNy42MC4yNDUuMjE5IiwKICAgICJSZW1vdGVETlNUeXBlIjogIkRvSCIsCiAgICAiUm91dGVPcmRlciI6ICJibG9jay1wcm94eS1kaXJlY3QiCn0=
```

#### 📱 Для приложения Incy (iOS / Android):
```text
incy://routing/add/ewogICAgIkJsb2NrSXAiOiBbXSwKICAgICJCbG9ja1NpdGVzIjogWwogICAgICAgICJhcHBjZW50ZXIubXMiLAogICAgICAgICJmaXJlYmFzZS5pbyIsCiAgICAgICAgImNyYXNobHl0aWNzLmNvbSIKICAgIF0sCiAgICAiRGlyZWN0SXAiOiBbCiAgICAgICAgIjEwLjAuMC4wLzgiLAogICAgICAgICIxNzIuMTYuMC4wLzEyIiwKICAgICAgICAiMTkyLjE2OC4wLjAvMTYiLAogICAgICAgICIxNjkuMjU0LjAuMC8xNiIsCiAgICAgICAgIjIyNC4wLjAuMC80IiwKICAgICAgICAiMjU1LjI1NS4yNTUuMjU1IgogICAgXSwKICAgICJEaXJlY3RTaXRlcyI6IFtdLAogICAgIkRuc0hvc3RzIjogewogICAgICAgICJjbG91ZGZsYXJlLWRucy5jb20iOiAiMS4xLjEuMSIsCiAgICAgICAgImRucy5nb29nbGUiOiAiOC44LjguOCIKICAgIH0sCiAgICAiRG9tYWluU3RyYXRlZ3kiOiAiSVBJZk5vbk1hdGNoIiwKICAgICJEb21lc3RpY0ROU0RvbWFpbiI6ICJodHRwczovL2dlb2hpZGUucnUvZG5zLXF1ZXJ5IiwKICAgICJEb21lc3RpY0ROU0lQIjogIjQ1LjE1NS4yMDQuMTkwIiwKICAgICJEb21lc3RpY0ROU1R5cGUiOiAiRG9IIiwKICAgICJGYWtlRE5TIjogInRydWUiLAogICAgIkdlb2lwdXJsIjogImh0dHBzOi8vY2RuLmpzZGVsaXZyLm5ldC9naC9jb21tZW5zYWwvZG9tYWluLWxpc3QtY29tbXVuaXR5QHJlbGVhc2UvZ2VvaXAuZGF0IiwKICAgICJHZW9zaXRldXJsIjogImh0dHBzOi8vY2RuLmpzZGVsaXZyLm5ldC9naC9jb21tZW5zYWwvZG9tYWluLWxpc3QtY29tbXVuaXR5QHJlbGVhc2UvZ2Vvc2l0ZS5kYXQiLAogICAgIkdsb2JhbFByb3h5IjogImZhbHNlIiwKICAgICJMYXN0VXBkYXRlZCI6IDE3ODgzMzcyNjIsCiAgICAiTmFtZSI6ICJCbGFjayBMaXN0IChGdWxsKSIsCiAgICAiUHJveHlJcCI6IFsKICAgICAgICAiZ2VvaXA6Y3VzdG9tIiwKICAgICAgICAiZ2VvaXA6dGVsZWdyYW0iLAogICAgICAgICJnZW9pcDp0d2l0dGVyIiwKICAgICAgICAiZ2VvaXA6bWV0YSIsCiAgICAgICAgImdlb3NpdGU6ZGlzY29yZCIsCiAgICAgICAgImdlb2lwOmNsb3VkZnJvbnQiLAogICAgICAgICJnZW9pcDpjbG91ZGZsYXJlIgogICAgXSwKICAgICJQcm94eVNpdGVzIjogWwogICAgICAgICJnZW9zaXRlOkNVU1RPTSIsCiAgICAgICAgImdlb3NpdGU6UlUtQkxPQ0siLAogICAgICAgICJnZW9zaXRlOkdFT0JMT0NLIiwKICAgICAgICAiZ2Vvc2l0ZTpURUxFR1JBTSIsCiAgICAgICAgImdlb3NpdGU6TUVUQSIsCiAgICAgICAgImdlb3NpdGU6WU9VVFVCRSIsCiAgICAgICAgImdlb3NpdGU6RElTQ09SRCIsCiAgICAgICAgImdlb3NpdGU6VFdJVFRFUiIsCiAgICAgICAgImdlb3NpdGU6VElLVE9LIiwKICAgICAgICAiZ2Vvc2l0ZTpST0JMT1giLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1QTEFZIiwKICAgICAgICAiZ2Vvc2l0ZTpHT09HTEUtTUVFVCIsCiAgICAgICAgImdlb3NpdGU6R09PR0xFLU1FRVQiLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1BSSIsCiAgICAgICAgImdlb3NpdGU6Q0xPVURGTEFSRSIsCiAgICAgICAgImdlb3NpdGU6Q0xPVURGUk9OVCIsCiAgICAgICAgImdlb3NpdGU6RElHSVRBTE9DRUFOIiwKICAgICAgICAiZ2Vvc2l0ZTpIRVRaTkVSIiwKICAgICAgICAiZ2Vvc2l0ZTpPVkgiLAogICAgICAgICJnZW9zaXRlOkhEUkVaS0EiLAogICAgICAgICJnZW9zaXRlOkhPRENBIiwKICAgICAgICAiZ2Vvc2l0ZTpORVdTIiwKICAgICAgICAiZ2Vvc2l0ZTpQT1JOIiwKICAgICAgICAiZ2Vvc2l0ZTpBTklNRSIKICAgIF0sCiAgICAiUmVtb3RlRE5TRG9tYWluIjogImh0dHBzOi8vZXUuZ2VvaGlkZS5ydS9kbnMtcXVlcnkiLAogICAgIlJlbW90ZUROU0lQIjogIjIxNy42MC4yNDUuMjE5IiwKICAgICJSZW1vdGVETlNUeXBlIjogIkRvSCIsCiAgICAiUm91dGVPcmRlciI6ICJibG9jay1wcm94eS1kaXJlY3QiCn0=
```

---

### 📦 2. Базовый профиль (Исходный шаблон без RU-BLOCK и GEOBLOCK)
Если вам нужен облегченный список без реестра запрещенных сайтов РФ.

#### 📱 Для приложения Happ:
```text
happ://routing/add/ewogICAgIkJsb2NrSXAiOiBbCiAgICBdLAogICAgIkJsb2NrU2l0ZXMiOiBbCiAgICAgICAgImFwcGNlbnRlci5tcyIsCiAgICAgICAgImZpcmViYXNlLmlvIiwKICAgICAgICAiY3Jhc2hseXRpY3MuY29tIgogICAgXSwKICAgICJEaXJlY3RJcCI6IFsKICAgICAgICAiMTAuMC4wLjAvOCIsCiAgICAgICAgIjE3Mi4xNi4wLjAvMTIiLAogICAgICAgICIxOTIuMTY4LjAuMC8xNiIsCiAgICAgICAgIjE2OS4yNTQuMC4wLzE2IiwKICAgICAgICAiMjI0LjAuMC4wLzQiLAogICAgICAgICIyNTUuMjU1LjI1NS4yNTUiCiAgICBdLAogICAgIkRpcmVjdFNpdGVzIjogWwogICAgXSwKICAgICJEbnNIb3N0cyI6IHsKICAgICAgICAiY2xvdWRmbGFyZS1kbnMuY29tIjogIjEuMS4xLjEiLAogICAgICAgICJkbnMuZ29vZ2xlIjogIjguOC44LjgiCiAgICB9LAogICAgIkRvbWFpblN0cmF0ZWd5IjogIklQSWZOb25NYXRjaCIsCiAgICAiRG9tZXN0aWNETlNEb21haW4iOiAiaHR0cHM6Ly9nZW9oaWRlLnJ1L2Rucy1xdWVyeSIsCiAgICAiRG9tZXN0aWNETlNJUCI6ICI0NS4xNTUuMjA0LjE5MCIsCiAgICAiRG9tZXN0aWNETlNUeXBlIjogIkRvSCIsCiAgICAiRmFrZUROUyI6ICJ0cnVlIiwKICAgICJHZW9pcHVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb2lwLmRhdCIsCiAgICAiR2Vvc2l0ZXVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb3NpdGUuZGF0IiwKICAgICJHbG9iYWxQcm94eSI6ICJmYWxzZSIsCiAgICAiTGFzdFVwZGF0ZWQiOiAxNzg4MzM3MjYyLAogICAgIk5hbWUiOiAiQmxhY2sgTGlzdCIsCiAgICAiUHJveHlJcCI6IFsKICAgICAgICAiZ2VvaXA6Y3VzdG9tIiwKICAgICAgICAiZ2VvaXA6dGVsZWdyYW0iLAogICAgICAgICJnZW9pcDp0d2l0dGVyIiwKICAgICAgICAiZ2VvaXA6bWV0YSIsCiAgICAgICAgImdlb3NpdGU6ZGlzY29yZCIsCiAgICAgICAgImdlb2lwOmNsb3VkZnJvbnQiLAogICAgICAgICJnZW9pcDpjbG91ZGZsYXJlIgogICAgXSwKICAgICJQcm94eVNpdGVzIjogWwogICAgICAgICJnZW9zaXRlOlRFTEVHUkFNIiwKICAgICAgICAiZ2Vvc2l0ZTpNRVRBIiwKICAgICAgICAiZ2Vvc2l0ZTpDVVNUT00iLAogICAgICAgICJnZW9zaXRlOlBPUk4iLAogICAgICAgICJnZW9zaXRlOllPVVRVQkUiLAogICAgICAgICJnZW9zaXRlOkNMT1VERkxBUkUiLAogICAgICAgICJnZW9zaXRlOkNMT1VERFJPTlQiLAogICAgICAgICJnZW9zaXRlOkRJR0lUQUxPQ0VBTiIsCiAgICAgICAgImdlb3NpdGU6R09PR0xFLVBMQVkiLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1NRUVUIiwKICAgICAgICAiZ2Vvc2l0ZTpIRFJFWktBIiwKICAgICAgICAiZ2Vvc2l0ZTpIRVRaTkVSIiwKICAgICAgICAiZ2Vvc2l0ZTpIT0RDQSIsCiAgICAgICAgImdlb3NpdGU6TkVXUyIsCiAgICAgICAgImdlb3NpdGU6T1ZIIiwKICAgICAgICAiZ2Vvc2l0ZTpST0JMT1giLAogICAgICAgICJnZW9zaXRlOlRJS1RPSyIsCiAgICAgICAgImdlb3NpdGU6VFdJVFRFUiIKICAgIF0sCiAgICAiUmVtb3RlRE5TRG9tYWluIjogImh0dHBzOi8vZXUuZ2VvaGlkZS5ydS9kbnMtcXVlcnkiLAogICAgIlJlbW90ZUROU0lQIjogIjIxNy42MC4yNDUuMjE5IiwKICAgICJSZW1vdGVETlNUeXBlIjogIkRvSCIsCiAgICAiUm91dGVPcmRlciI6ICJibG9jay1wcm94eS1kaXJlY3QiCn0=
```

#### 📱 Для приложения Incy:
```text
incy://routing/add/ewogICAgIkJsb2NrSXAiOiBbCiAgICBdLAogICAgIkJsb2NrU2l0ZXMiOiBbCiAgICAgICAgImFwcGNlbnRlci5tcyIsCiAgICAgICAgImZpcmViYXNlLmlvIiwKICAgICAgICAiY3Jhc2hseXRpY3MuY29tIgogICAgXSwKICAgICJEaXJlY3RJcCI6IFsKICAgICAgICAiMTAuMC4wLjAvOCIsCiAgICAgICAgIjE3Mi4xNi4wLjAvMTIiLAogICAgICAgICIxOTIuMTY4LjAuMC8xNiIsCiAgICAgICAgIjE2OS4yNTQuMC4wLzE2IiwKICAgICAgICAiMjI0LjAuMC4wLzQiLAogICAgICAgICIyNTUuMjU1LjI1NS4yNTUiCiAgICBdLAogICAgIkRpcmVjdFNpdGVzIjogWwogICAgXSwKICAgICJEbnNIb3N0cyI6IHsKICAgICAgICAiY2xvdWRmbGFyZS1kbnMuY29tIjogIjEuMS4xLjEiLAogICAgICAgICJkbnMuZ29vZ2xlIjogIjguOC44LjgiCiAgICB9LAogICAgIkRvbWFpblN0cmF0ZWd5IjogIklQSWZOb25NYXRjaCIsCiAgICAiRG9tZXN0aWNETlNEb21haW4iOiAiaHR0cHM6Ly9nZW9oaWRlLnJ1L2Rucy1xdWVyeSIsCiAgICAiRG9tZXN0aWNETlNJUCI6ICI0NS4xNTUuMjA0LjE5MCIsCiAgICAiRG9tZXN0aWNETlNUeXBlIjogIkRvSCIsCiAgICAiRmFrZUROUyI6ICJ0cnVlIiwKICAgICJHZW9pcHVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb2lwLmRhdCIsCiAgICAiR2Vvc2l0ZXVybCI6ICJodHRwczovL2Nkbi5qc2RlbGl2ci5uZXQvZ2gvY29tbWVuc2FsL2RvbWFpbi1saXN0LWNvbW11bml0eUByZWxlYXNlL2dlb3NpdGUuZGF0IiwKICAgICJHbG9iYWxQcm94eSI6ICJmYWxzZSIsCiAgICAiTGFzdFVwZGF0ZWQiOiAxNzg4MzM3MjYyLAogICAgIk5hbWUiOiAiQmxhY2sgTGlzdCIsCiAgICAiUHJveHlJcCI6IFsKICAgICAgICAiZ2VvaXA6Y3VzdG9tIiwKICAgICAgICAiZ2VvaXA6dGVsZWdyYW0iLAogICAgICAgICJnZW9pcDp0d2l0dGVyIiwKICAgICAgICAiZ2VvaXA6bWV0YSIsCiAgICAgICAgImdlb3NpdGU6ZGlzY29yZCIsCiAgICAgICAgImdlb2lwOmNsb3VkZnJvbnQiLAogICAgICAgICJnZW9pcDpjbG91ZGZsYXJlIgogICAgXSwKICAgICJQcm94eVNpdGVzIjogWwogICAgICAgICJnZW9zaXRlOlRFTEVHUkFNIiwKICAgICAgICAiZ2Vvc2l0ZTpNRVRBIiwKICAgICAgICAiZ2Vvc2l0ZTpDVVNUT00iLAogICAgICAgICJnZW9zaXRlOlBPUk4iLAogICAgICAgICJnZW9zaXRlOllPVVRVQkUiLAogICAgICAgICJnZW9zaXRlOkNMT1VERkxBUkUiLAogICAgICAgICJnZW9zaXRlOkNMT1VERFJPTlQiLAogICAgICAgICJnZW9zaXRlOkRJR0lUQUxPQ0VBTiIsCiAgICAgICAgImdlb3NpdGU6R09PR0xFLVBMQVkiLAogICAgICAgICJnZW9zaXRlOkdPT0dMRS1NRUVUIiwKICAgICAgICAiZ2Vvc2l0ZTpIRFJFWktBIiwKICAgICAgICAiZ2Vvc2l0ZTpIRVRaTkVSIiwKICAgICAgICAiZ2Vvc2l0ZTpIT0RDQSIsCiAgICAgICAgImdlb3NpdGU6TkVXUyIsCiAgICAgICAgImdlb3NpdGU6T1ZIIiwKICAgICAgICAiZ2Vvc2l0ZTpST0JMT1giLAogICAgICAgICJnZW9zaXRlOlRJS1RPSyIsCiAgICAgICAgImdlb3NpdGU6VFdJVFRFUiIKICAgIF0sCiAgICAiUmVtb3RlRE5TRG9tYWluIjogImh0dHBzOi8vZXUuZ2VvaGlkZS5ydS9kbnMtcXVlcnkiLAogICAgIlJlbW90ZUROU0lQIjogIjIxNy42MC4yNDUuMjE5IiwKICAgICJSZW1vdGVETlNUeXBlIjogIkRvSCIsCiAgICAiUm91dGVPcmRlciI6ICJibG9jay1wcm94eS1kaXJlY3QiCn0=
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
* Для всех запросов клиент мгновенно генерирует фиктивный локальный IP-адрес из диапазона `198.18.0.0/15` и отдает его операционной системе без ожидания ответа от удаленного DNS-сервера.
* Это снижает задержку (RTT) до **0 мс** при открытии сайтов на мобильном интернете и полностью исключает DNS-утечки.

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
| **`BlockSites`** | `appcenter.ms`, `firebase.io`, `crashlytics.com` | Блокировка мобильной телеметрии, трекеров и аналитики крашей. Экономит заряд батареи и трафик. |
| **`DirectIp`** | `10.0.0.0/8`, `192.168.0.0/16`, `172.16.0.0/12`... | Исключение локальных сетей (RFC 1918). Доступ к роутеру, сетевым принтерам и умному дому всегда работает напрямую. |
| **`ProxySites`** | Категории `geosite:*` | Список категорий, направляемых в прокси (см. таблицу ниже). |
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
* `geosite:HDREZKA` — онлайн-кинотеатр HDRezka.
* `geosite:CLOUDFLARE` / `geosite:CLOUDFRONT` — зарубежные CDN.
* `geosite:HETZNER` / `geosite:DIGITALOCEAN` / `geosite:OVH` — зарубежные хостинг-провайдеры.
* `geosite:RU-BLOCK` — реестр сайтов, заблокированных на территории РФ (РКН).
* `geosite:GEOBLOCK` — сервисы с географическими ограничениями.
* `geosite:NEWS` / `geosite:HODCA` / `geosite:PORN` / `geosite:ANIME` — тематические категории.

### 📍 IP-адреса и подсети (`geoip.dat`)
Собраны с автоматическим объединением диапазонов IPv4 и IPv6:
* `geoip:custom` — ваши персональные IP из `my-ips.txt`.
* `geoip:telegram`
* `geoip:twitter`
* `geoip:meta`
* `geoip:discord`
* `geoip:cloudfront`
* `geoip:cloudflare`

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
