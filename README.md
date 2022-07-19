## コツ

## 1章: 字句解析器

- 60分あれば写経はいける！ 👍
- 細かいアルゴリズムやGo的なところはともかく！
- **Lexerのテストケースは1トークンずつ足していくとマジで分かりやすい**。最大でも1行。Lexerは意味の解釈はしないから1トークンずつでいいのだ。一気にやると、わかりにくいしやる気なくなる。
- JSONのレキサー(Notパーサー！)も、同じパターンで結構いけるんじゃね？
- 電卓のレキサーなら、登場するトークンが少ないからもっといい感じにいけるんじゃね？ お！？！？？
  - INT ( ) + - * / と ILLEGAL
  - これだけじゃん！！
  - **2文字トークンが存在しないから**、先読みもいらない

## 2章

- 経験が無いから、難しく感じるだけって話、なるほど
  - というわけで、プログラミング言語の構文解析器をシリアライズ形式（JSON、YAML、TOML、 INIなど） のパーサーと同様に身近に感じ、直感的に理解できるようになるためには、プログラミング
    言語の構文解析器が生成するデータ構造について理解する必要がある。

### let文を正しく構文解析するというのは一体どういうことだろうか？

それは、元のlet文に含まれる 情報を正確に表現したASTを構文解析器が生成することを意味する。これはもっともらしく聞こえる。 とは言うものの、まだASTは実装していないし、それがどういう形であるべきかすらわかっていない。
そこで、まずはMonkeyのソースコードをよく観察し、どのように構造化されているかを見ることにし よう。そうすれば、let文を正確に表現するにあたって、ASTに必要となる部品を定義できるようにな る。

## 各Nodeが何で、どういう型なのか、どういう属性なのかをいちいち確認したほうがいい!

っていうか、そうしないと無理だね。無の写経になる。