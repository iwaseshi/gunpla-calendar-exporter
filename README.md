# gunpla-calendar-exporter

確定しているガンプラの発売日 を
任意のカレンダーアプリ等に登録できる `ics` ファイル

![calender](assets/image.png)

# URL一覧

| カテゴリ | URL |
| :-: | - |
| 2024/04 | https://storage.cloud.google.com/gunpla-calendar-exporter/2024April.ics |

# 設定例

## iOS (iPhone/iPad)

* ホーム画面の `設定 (Settings)` を開く
* `パスワードとアカウント (Passwords & Accounts)` -> `アカウントを追加 (Add Account)`
* 一番下の `その他 (Other)`
* 一番下の `照会するカレンダーを追加 (Add Subscribed Calendar)`
* URL を貼り付けて `次へ (Next)`
* そのまま右上の `保存 (Save)`

## Google Calendar (PC)

* icsファイルのリンクをクリックしファイルをダウンロードする。
* `設定` を開く
* 左側メニュー `カレンダーの追加` -> `インポート`
* ダウンロードしたicsを指定し、どのマイカレンダーにインポートするか指定
* `インポート`

もしくは

* `設定` を開く
* 左側メニュー `カレンダーの追加` -> `URL で追加`
* URL を貼り付け
* `カレンダーを追加`

PC から追加後はモバイル版 Google Calendar でも閲覧できます

## appendix

依存ライブラリ

* github.com/chromedp/chromedp
* github.com/arran4/golang-ical
