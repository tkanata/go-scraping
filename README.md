# go-scraping
goqueryを使って特定の文字列が含まれる箇所をスクレイピングする。
index.htmlをダブルクリックしたらなんか画面出てくるので、スクレイピングしたいページのURLをフォームに入力してください。
送信ボタンか何か押したら結果画面に遷移します。
現在の仕様だと対象の文字列が存在しない場合は真っ白な画面に遷移し、対象が存在する場合は対象の文字列と文字列と同じ行にあるhtmlタグが抽出されます。
